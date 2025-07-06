# 设置项目名称和输出目录
$projectName = "billadm"
$outputDir = Join-Path $PSScriptRoot "target\bin"

# 如果 target 目录存在则先删除
if (Test-Path $outputDir) {
    Remove-Item -Recurse -Force $outputDir
}

# 创建输出目录
New-Item -ItemType Directory -Path $outputDir | Out-Null

# 定义目标平台数组：支持 amd64 和 arm64
$targets = @(
    @{ GOOS = "windows"; GOARCH = "amd64" },
    @{ GOOS = "windows"; GOARCH = "arm64" }
)

# 获取脚本所在目录的上级目录作为项目根目录
$projectRoot = Split-Path $PSScriptRoot -Parent

# 进入 kernel/cmd 子目录
$buildDir = Join-Path $projectRoot "kernel\cmd"
Set-Location $buildDir

Write-Host "当前工作目录: $(Get-Location)"

# 遍历并构建每个平台
foreach ($target in $targets) {
    $goos = $target.GOOS
    $goarch = $target.GOARCH

    # 设置环境变量
    $env:GOOS = $goos
    $env:GOARCH = $goarch

    # 编译参数（可选）
    $ldflags = "-s -w -X 'constant.Mode=release'"

    # 构建输出文件名
    $filename = "$projectName-$goos-$goarch"
    if ($goos -eq "windows") {
        $filename += ".exe"
    }

    $outputFile = Join-Path $outputDir $filename

    Write-Host "Building for $goos/$goarch..."

    # 执行 go build 命令
    go build -o $outputFile -ldflags $ldflags .

    if ($LASTEXITCODE -ne 0) {
        Write-Error "Build failed for $goos/$goarch."
        exit 1
    }

    Write-Host "Built successfully: $outputFile"
}

# 返回原始工作目录
Set-Location $PSScriptRoot

# 清理环境变量
Remove-Item Env:\GOOS -ErrorAction SilentlyContinue
Remove-Item Env:\GOARCH -ErrorAction SilentlyContinue

Write-Host "All builds completed successfully!"