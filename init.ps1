# File: init.ps1
# Project: gomod-test
# Created Date: 16/06/2022
# Author: Shun Suzuki
# -----
# Last Modified: 16/06/2022
# Modified By: Shun Suzuki (suzuki@hapis.k.u-tokyo.ac.jp)
# -----
# Copyright (c) 2022 Shun Suzuki. All rights reserved.
# 


Param(
    [string]$version = "2.2.0"
)

function ColorEcho($color, $PREFIX, $message) {
    Write-Host $PREFIX -ForegroundColor $color -NoNewline
    Write-Host ":", $message
}

$script_path = Split-Path $MyInvocation.MyCommand.Path -Parent

Remove-Item "${script_path}/*.dll" -Force
Remove-Item "${script_path}/*.so" -Force
Remove-Item "${script_path}/*.dylib" -Force

ColorEcho "Green" "INFO" "Download Windows ", $version, " libraries..."
$url = "https://github.com/shinolab/autd3/releases/download/v" + $version + "/autd3-v" + $version + "-win-x64.zip"
curl.exe -Lo "${script_path}/tmp.zip" $url
Expand-Archive "${script_path}/tmp.zip" -Force
Move-Item "${script_path}/tmp/bin/*" "${script_path}/" -Force
Remove-Item "${script_path}/tmp.zip" -Force
Remove-Item "${script_path}/tmp" -Recurse -Force

ColorEcho "Green" "INFO" "Download Linux", $version, "libraries..."
$url = "https://github.com/shinolab/autd3/releases/download/v" + $version + "/autd3-v" + $version + "-linux-x64.tar.gz"
curl.exe -Lo "${script_path}/tmp.tar.gz" $url
mkdir tmp
tar.exe -xzf "${script_path}/tmp.tar.gz" -C tmp
Move-Item "${script_path}/tmp/bin/*" "${script_path}/" -Force
Remove-Item "${script_path}/tmp.tar.gz" -Force
Remove-Item "${script_path}/tmp" -Recurse -Force

ColorEcho "Green" "INFO" "Download macOS", $version, "libraries..."
$url = "https://github.com/shinolab/autd3/releases/download/v" + $version + "/autd3-v" + $version + "-macos-universal.tar.gz"
curl.exe -Lo "${script_path}/tmp.tar.gz" $url
mkdir tmp
tar.exe -xzf "${script_path}/tmp.tar.gz" -C tmp
Move-Item "${script_path}/tmp/bin/*" "${script_path}/" -Force
Remove-Item "${script_path}/tmp.tar.gz" -Force
Remove-Item "${script_path}/tmp" -Recurse -Force

ColorEcho "Green" "INFO" "Done"
