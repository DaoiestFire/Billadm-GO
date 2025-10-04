# 设置输出编码为 UTF-8（防止中文乱码）
[Console]::OutputEncoding = [System.Text.Encoding]::UTF8

# 获取脚本所在目录
$scriptDir = $PSScriptRoot

# 获取项目根目录（上一级）
$projectRoot = Split-Path -Parent $scriptDir

# 定义路径
$buildDir = Join-Path $projectRoot "build"
$electronDir = Join-Path $projectRoot "electron"
$appDistDir = Join-Path $projectRoot "app\dist"
$kernelExe = Join-Path $projectRoot "kernel\Billadm-Kernel.exe"
$sqlPath = Join-Path $projectRoot "kernel\billadm.sql"

# 输出开始信息
Write-Host "📦 开始执行 Electron 应用打包流程" -ForegroundColor Cyan

# 1. 执行 build_vue.ps1
Write-Host "`n🛠️  正在构建前端 Vue 项目..." -ForegroundColor Magenta
$buildVueScript = Join-Path $buildDir "build_vue.ps1"
if (-not (Test-Path $buildVueScript))
{
    Write-Error "❌ 找不到前端构建脚本: $buildVueScript"
    exit 1
}

& $buildVueScript
if ($LASTEXITCODE -ne 0)
{
    Write-Error "❌ 前端构建失败，终止流程。"
    exit $LASTEXITCODE
}
Write-Host "✅ 前端构建成功" -ForegroundColor Green


# 2. 执行 build_kernel.ps1
Write-Host "`n⚙️  正在构建后端 Go 项目..." -ForegroundColor Magenta
$buildKernelScript = Join-Path $buildDir "build_kernel.ps1"
if (-not (Test-Path $buildKernelScript))
{
    Write-Error "❌ 找不到后端构建脚本: $buildKernelScript"
    exit 1
}

& $buildKernelScript
if ($LASTEXITCODE -ne 0)
{
    Write-Error "❌ 后端构建失败，终止流程。"
    exit $LASTEXITCODE
}
Write-Host "✅ 后端构建成功" -ForegroundColor Green


# 3. 进入 electron 目录前的准备：清理旧文件
Write-Host "`n🧹 正在清理 electron 目录旧文件..." -ForegroundColor Yellow
$targetDistDir = Join-Path $electronDir "dist"
$targetKernelExe = Join-Path $electronDir "Billadm-Kernel.exe"
$targetSqlPath = Join-Path $electronDir "billadm.sql"

# 删除 dist 目录（如果存在）
if (Test-Path $targetDistDir)
{
    try
    {
        Remove-Item $targetDistDir -Recurse -Force -ErrorAction Stop
        Write-Host "🗑️  已删除旧的 dist 目录" -ForegroundColor DarkGray
    }
    catch
    {
        Write-Error "❌ 删除 dist 目录失败: $( $_.Exception.Message )"
        exit 1
    }
}

# 删除旧的 Billadm-Kernel.exe（如果存在）
if (Test-Path $targetKernelExe)
{
    try
    {
        Remove-Item $targetKernelExe -Force -ErrorAction Stop
        Write-Host "🗑️  已删除旧的 Billadm-Kernel.exe" -ForegroundColor DarkGray
    }
    catch
    {
        Write-Error "❌ 删除 Billadm-Kernel.exe 失败: $( $_.Exception.Message )"
        exit 1
    }
}

# 删除旧的 billadm.sql（如果存在）
if (Test-Path $targetSqlPath)
{
    try
    {
        Remove-Item $targetSqlPath -Force -ErrorAction Stop
        Write-Host "🗑️  已删除旧的 billadm.sql" -ForegroundColor DarkGray
    }
    catch
    {
        Write-Error "❌ 删除 billadm.sql 失败: $( $_.Exception.Message )"
        exit 1
    }
}


# 4. 拷贝新的 dist 目录到 electron
Write-Host "`n📂 正在拷贝前端 dist 到 electron..." -ForegroundColor Yellow
if (-not (Test-Path $appDistDir))
{
    Write-Error "❌ 找不到前端 dist 目录: $appDistDir"
    exit 1
}

try
{
    Copy-Item -Path $appDistDir -Destination $electronDir -Recurse -Force -ErrorAction Stop
    Write-Host "✅ 前端 dist 已拷贝至 $electronDir\dist" -ForegroundColor Green
}
catch
{
    Write-Error "❌ 拷贝 dist 失败: $( $_.Exception.Message )"
    exit 1
}


# 5. 拷贝后端可执行文件到 electron
Write-Host "`n📄 正在拷贝后端可执行文件到 electron..." -ForegroundColor Yellow
if (-not (Test-Path $kernelExe))
{
    Write-Error "❌ 找不到后端可执行文件: $kernelExe"
    exit 1
}

try
{
    Copy-Item -Path $kernelExe -Destination $electronDir -Force -ErrorAction Stop
    Write-Host "✅ 已拷贝 Billadm-Kernel.exe 到 $electronDir" -ForegroundColor Green
}
catch
{
    Write-Error "❌ 拷贝可执行文件失败: $( $_.Exception.Message )"
    exit 1
}

# 拷贝sql文件到 electron
Write-Host "`n📄 正在拷贝billadm.sql到 electron..." -ForegroundColor Yellow
if (-not (Test-Path $sqlPath))
{
    Write-Error "❌ 找不到sql文件: $sqlPath"
    exit 1
}

try
{
    Copy-Item -Path $sqlPath -Destination $electronDir -Force -ErrorAction Stop
    Write-Host "✅ 已拷贝 billadm.sql 到 $electronDir" -ForegroundColor Green
}
catch
{
    Write-Error "❌ 拷贝sql文件失败: $( $_.Exception.Message )"
    exit 1
}


# 6. 进入 electron 目录并执行打包命令
Write-Host "`n🚀 正在进入 electron 目录并执行打包命令..." -ForegroundColor Green
Set-Location $electronDir

Write-Host "   执行命令: npm run package" -ForegroundColor Yellow

# 执行打包命令
npm run package

# 检查打包是否成功
if ($LASTEXITCODE -ne 0)
{
    Write-Error "❌ Electron 打包失败，退出码: $LASTEXITCODE"
    # 打包失败后也返回脚本目录
    Set-Location $scriptDir
    exit $LASTEXITCODE
}

Write-Host "🎉 Electron 应用打包成功！" -ForegroundColor Green


# 7. 返回脚本所在目录
Set-Location $scriptDir
Write-Host "`n↩️  已返回脚本目录: $scriptDir" -ForegroundColor DarkCyan