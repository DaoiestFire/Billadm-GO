# 设置输出编码为 UTF-8（防止中文乱码）
[Console]::OutputEncoding = [System.Text.Encoding]::UTF8

# 获取脚本所在目录
$scriptDir = $PSScriptRoot

# 获取项目根目录（上一级）
$projectRoot = Split-Path -Parent $scriptDir

# 定义路径
$kernelDir = Join-Path $projectRoot "kernel"
$outputExeName = "Billadm-Kernel.exe"
$outputPath = Join-Path $kernelDir $outputExeName

# 输出路径信息
Write-Host "📌 脚本所在目录: $scriptDir" -ForegroundColor Green
Write-Host "📁 项目根目录: $projectRoot" -ForegroundColor Green
Write-Host "🔧 Go 项目目录: $kernelDir" -ForegroundColor Green
Write-Host "💾 输出文件: $outputPath" -ForegroundColor Green

# 检查 kernel 目录是否存在
if (-not (Test-Path $kernelDir))
{
    Write-Error "❌ 错误：Go 项目目录不存在: $kernelDir"
    exit 1
}

# 检查是否包含 go.mod（判断是否为有效 Go 项目）
$goMod = Join-Path $kernelDir "go.mod"
if (-not (Test-Path $goMod))
{
    Write-Error "❌ 错误：未找到 go.mod，确认 '$kernelDir' 是有效的 Go 项目"
    exit 1
}

# 删除 Go 项目下的 可执行文件
if (Test-Path $outputPath)
{
    Write-Host "🗑️  正在删除旧的 编译文件..." -ForegroundColor Yellow
    try
    {
        Remove-Item $outputPath -Recurse -Force -ErrorAction Stop
        Write-Host "✅ 成功删除 $outputPath" -ForegroundColor Green
    }
    catch
    {
        Write-Error "❌ 删除 $outputPath 失败: $( $_.Exception.Message )"
        exit 1
    }
}

# 记录当前目录（脚本所在目录），用于最后返回
$initialLocation = Get-Location

# 进入 Go 项目目录
Set-Location $kernelDir
Write-Host "`n🔨 正在编译 Go 项目..." -ForegroundColor Magenta

# 设置编译环境变量（Windows 32位）
$env:GOOS = "windows"
$env:GOARCH = "amd64"
$env:CGO_ENABLED = "1"

# 执行编译命令
& go build -ldflags '-s -w -extldflags "-static"' -o $outputPath

# 检查编译是否成功
if ($LASTEXITCODE -ne 0)
{
    Write-Error "❌ Go 编译失败，退出码: $LASTEXITCODE"
    Set-Location $initialLocation
    exit $LASTEXITCODE
}

# 验证输出文件是否生成
if (Test-Path $outputPath)
{
    Write-Host "✅ 编译成功！生成文件: $outputPath" -ForegroundColor Green
}
else
{
    Write-Error "❌ 编译完成但未生成预期文件: $outputPath"
    Set-Location $initialLocation
    exit 1
}

# 返回脚本所在目录
Set-Location $initialLocation
Write-Host "↩️  已返回脚本所在目录: $scriptDir" -ForegroundColor DarkCyan