# 设置输出编码为 UTF-8（防止中文乱码）
[Console]::OutputEncoding = [System.Text.Encoding]::UTF8

# 获取脚本所在目录
$scriptDir = $PSScriptRoot

# 获取项目根目录（上一级）
$projectRoot = Split-Path -Parent $scriptDir

# 定义路径
$vueDir = Join-Path $projectRoot "app"
$distDir = Join-Path $vueDir "dist"
$packageJson = Join-Path $vueDir "package.json"

# 输出路径信息
Write-Host "📌 脚本所在目录: $scriptDir" -ForegroundColor Green
Write-Host "📁 项目根目录: $projectRoot" -ForegroundColor Green
Write-Host "🚀 Vue 项目目录: $vueDir" -ForegroundColor Green
Write-Host "📦 构建输出目录: $distDir" -ForegroundColor Green

# 检查 Vue 项目目录
if (-not (Test-Path $vueDir))
{
    Write-Error "❌ 错误：Vue 项目目录不存在: $vueDir"
    exit 1
}

# 检查 package.json
if (-not (Test-Path $packageJson))
{
    Write-Error "❌ 错误：未找到 package.json，确认 '$vueDir' 是 Vue 项目目录"
    exit 1
}

# 删除 dist 目录（如果存在）
if (Test-Path $distDir)
{
    Write-Host "🗑️  正在删除旧的 dist 目录..." -ForegroundColor Yellow
    try
    {
        Remove-Item $distDir -Recurse -Force -ErrorAction Stop
        Write-Host "✅ 成功删除 dist 目录: $distDir" -ForegroundColor Green
    }
    catch
    {
        Write-Error "❌ 删除 dist 目录失败: $( $_.Exception.Message )"
        exit 1
    }
}
else
{
    Write-Host "🔍 dist 目录不存在，跳过删除。" -ForegroundColor Cyan
}

# 记录当前目录（脚本所在目录），后续要返回
$initialLocation = Get-Location

# 进入 Vue 项目目录并执行构建
Set-Location $vueDir
Write-Host "`n🔨 正在构建 Vue 项目..." -ForegroundColor Magenta
Write-Host "   执行命令: npm run build`n"

# 执行构建命令
& npm run build

# 检查构建是否成功
if ($LASTEXITCODE -ne 0)
{
    Write-Error "❌ 构建失败，退出码: $LASTEXITCODE"
    # 即使失败也尝试返回原目录
    Set-Location $initialLocation
    exit $LASTEXITCODE
}

# 构建成功
Write-Host "`n🎉 构建成功！输出位于: $distDir" -ForegroundColor Green

# 返回脚本所在目录
Set-Location $initialLocation
Write-Host "↩️  已返回脚本所在目录: $scriptDir" -ForegroundColor DarkCyan