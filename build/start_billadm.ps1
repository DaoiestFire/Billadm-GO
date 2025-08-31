# 设置输出编码为 UTF-8（防止中文乱码）
[Console]::OutputEncoding = [System.Text.Encoding]::UTF8

# 获取脚本所在目录
$scriptDir = $PSScriptRoot

# 获取项目根目录（上一级）
$projectRoot = Split-Path -Parent $scriptDir

# 定义路径
$targetDir = Join-Path $scriptDir "target"
$appDistDir = Join-Path $projectRoot "app\dist"
$kernelExe = Join-Path $projectRoot "kernel\Billadm-Kernel.exe"

# 输出开始信息
Write-Host "🚀 开始执行一键构建与启动流程" -ForegroundColor Cyan

# 1. 执行 build_vue.ps1
Write-Host "`n📦 正在构建前端 Vue 项目..." -ForegroundColor Magenta
$buildVueScript = Join-Path $scriptDir "build_vue.ps1"
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
Write-Host "`n🔧 正在构建后端 Go 项目..." -ForegroundColor Magenta
$buildKernelScript = Join-Path $scriptDir "build_kernel.ps1"
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

# 3. 删除并重建脚本目录下的 target
Write-Host "`n🗑️  清理旧的 target 目录..." -ForegroundColor Yellow
if (Test-Path $targetDir)
{
    try
    {
        Remove-Item $targetDir -Recurse -Force -ErrorAction Stop
        Write-Host "✅ target 目录已删除" -ForegroundColor Green
    }
    catch
    {
        Write-Error "❌ 删除 target 失败: $( $_.Exception.Message )"
        exit 1
    }
}

try
{
    New-Item -ItemType Directory -Path $targetDir -ErrorAction Stop | Out-Null
    Write-Host "🆕 已创建新的 target 目录" -ForegroundColor Cyan
}
catch
{
    Write-Error "❌ 创建 target 目录失败: $( $_.Exception.Message )"
    exit 1
}

# 4. 拷贝 dist 目录
Write-Host "`n📂 正在拷贝前端 dist 目录..." -ForegroundColor Yellow
if (-not (Test-Path $appDistDir))
{
    Write-Error "❌ 找不到 dist 目录: $appDistDir"
    exit 1
}

try
{
    Copy-Item -Path $appDistDir -Destination $targetDir -Recurse -Force -ErrorAction Stop
    Write-Host "✅ 前端 dist 已拷贝到 $targetDir\dist" -ForegroundColor Green
}
catch
{
    Write-Error "❌ 拷贝 dist 失败: $( $_.Exception.Message )"
    exit 1
}

# 5. 拷贝 Billadm-Kernel.exe
Write-Host "`n📄 正在拷贝后端可执行文件..." -ForegroundColor Yellow
if (-not (Test-Path $kernelExe))
{
    Write-Error "❌ 找不到后端可执行文件: $kernelExe"
    exit 1
}

try
{
    Copy-Item -Path $kernelExe -Destination $targetDir -Force -ErrorAction Stop
    Write-Host "✅ 已拷贝 $kernelExe → $targetDir" -ForegroundColor Green
}
catch
{
    Write-Error "❌ 拷贝可执行文件失败: $( $_.Exception.Message )"
    exit 1
}

# 7. 进入 target 目录并启动服务
Set-Location $targetDir
Write-Host "`n🎮 正在启动 Billadm-Kernel 服务..." -ForegroundColor Green
Write-Host "   执行命令: .\Billadm-Kernel.exe -log_file billadm.log -mode release`n"

# 启动程序（阻塞运行）
& .\Billadm-Kernel.exe -log_file billadm.log -mode release

# 检查启动是否异常退出
if ($LASTEXITCODE -ne 0)
{
    Write-Error "❌ 程序异常退出，退出码: $LASTEXITCODE"
}

# 8. 程序退出后，返回脚本所在目录（即使程序退出也确保返回）
Set-Location $scriptDir
Write-Host "`n↩️  已返回脚本所在目录: $scriptDir" -ForegroundColor DarkCyan