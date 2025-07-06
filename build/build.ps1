# 切换到脚本目录
Set-Location $PSScriptRoot

# 删除并重建 target 目录
if (Test-Path "target") {
    Remove-Item "target" -Recurse -Force
}
New-Item "target" -ItemType Directory | Out-Null

# 创建 target/backend/bin 子目录
New-Item "target/backend/bin" -ItemType Directory | Out-Null
# 创建 target/backend/log 子目录
New-Item "target/backend/log" -ItemType Directory | Out-Null

# 执行 build_kernel.ps1 脚本
Write-Host "执行 build_kernel.ps1..."
if (Test-Path "build_kernel.ps1") {
    & ".\build_kernel.ps1"
    if ($LASTEXITCODE -ne 0) {
        Write-Error "build_kernel.ps1 执行失败。构建中止。"
        exit 1
    }
} else {
    Write-Error "找不到 build_kernel.ps1 文件。"
    exit 1
}

# 检查系统架构
$arch = Get-CimInstance -Query "SELECT * FROM Win32_Processor" | Select-Object -ExpandProperty Architecture
switch ($arch) {
    9 { $archName = "amd64" }
    12 { $archName = "arm64" }
    default { $archName = "unknown" }
}

if ($archName -notin @("amd64", "arm64")) {
    Write-Error "不支持的系统架构: $arch"
    exit 1
}

Write-Host "检测到系统架构为: $archName"

# 定义源文件路径和目标路径
$sourceFile = "target/bin/billadm-windows-$archName.exe"
$destFile = "target/backend/bin/billadm.exe"

# 拷贝可执行文件
if (Test-Path $sourceFile) {
    Write-Host "复制 $sourceFile 到 $destFile"
    Copy-Item -Path $sourceFile -Destination $destFile -Force
} else {
    Write-Error "找不到源文件: $sourceFile"
    exit 1
}

# 拷贝conf目录到目标目录
# 获取脚本所在目录的上级目录作为项目根目录
$projectRoot = Split-Path $PSScriptRoot -Parent

# 进入 kernel 子目录
$kernelDir = Join-Path $projectRoot "kernel"
$confDir = Join-Path $kernelDir "conf"
$destDir = Join-Path $PSScriptRoot "target" "backend"
if (Test-Path $confDir) {
    Write-Host "复制 $confDir 到 $destDir"
    Copy-Item -Path "$confDir" -Destination "$destDir" -Recurse
} else {
    Write-Error "找不到源目录: $confDir"
    exit 1
}

Write-Host "构建流程已完成。"