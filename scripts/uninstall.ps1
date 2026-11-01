# Casbin Gateway clean uninstaller for Windows PowerShell.
#
# Usage:
#   irm https://raw.githubusercontent.com/tonnychiulab/casbin-gateway-tw/master/scripts/uninstall.ps1 | iex

$ErrorActionPreference = 'SilentlyContinue'

Write-Host "正在為您完整解除安裝 Casbin Gateway..." -ForegroundColor Cyan

# 1. 停止任何可能殘留的行程
Get-Process -Name "casbin-gateway", "casbin-gateway-desktop" -ErrorAction SilentlyContinue | Stop-Process -Force -ErrorAction SilentlyContinue

# 2. 移除桌面與開始功能表捷徑
$DesktopLnk = Join-Path ([System.Environment]::GetFolderPath('Desktop')) 'Casbin Gateway.lnk'
$StartLnk = Join-Path ([System.Environment]::GetFolderPath('Programs')) 'Casbin Gateway.lnk'
$StartupLnk = Join-Path ([System.Environment]::GetFolderPath('Startup')) 'Casbin Gateway.lnk'
Remove-Item -Path $DesktopLnk, $StartLnk, $StartupLnk -Force -ErrorAction SilentlyContinue
Write-Host "  - 捷徑與啟動項目已移除"

# 3. 移除開機自動啟動機碼與自訂網址協定 (ccswitch://)
Remove-ItemProperty -Path 'HKCU:\Software\Microsoft\Windows\CurrentVersion\Run' -Name 'CasbinGateway' -ErrorAction SilentlyContinue
Remove-Item -Path 'HKCU:\Software\Classes\ccswitch' -Recurse -Force -ErrorAction SilentlyContinue
Write-Host "  - 註冊表開機啟動與 ccswitch:// 協定已清除"

# 4. 從使用者 PATH 環境變數中移除 bin 目錄
foreach ($binDirName in @('casbin-gateway-tw\bin', 'casbin-gateway\bin')) {
    $targetBin = Join-Path $env:LOCALAPPDATA $binDirName
    $UserPath = [System.Environment]::GetEnvironmentVariable('PATH', 'User')
    if ($UserPath -like "*$targetBin*") {
        $NewPath = ($UserPath -split ';' | Where-Object { $_ -ne $targetBin -and $_ -ne '' }) -join ';'
        [System.Environment]::SetEnvironmentVariable('PATH', $NewPath, 'User')
        Write-Host "  - 已從使用者 PATH 環境變數中移除: $targetBin"
    }
}

# 5. 刪除整個安裝目錄
$dirsToClean = if ($env:INSTALL_DIR) { @($env:INSTALL_DIR) } else { @("$env:LOCALAPPDATA\casbin-gateway-tw", "$env:LOCALAPPDATA\casbin-gateway") }
foreach ($dir in $dirsToClean) {
    if (Test-Path $dir) {
        Remove-Item -Path $dir -Recurse -Force -ErrorAction SilentlyContinue
        Write-Host "  - 已清理安裝目錄: $dir"
    }
}

Write-Host "`n✅ Casbin Gateway 已成功完整解除安裝乾淨！" -ForegroundColor Green
