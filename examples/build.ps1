# File: build.ps1
# Project: examples
# Created Date: 16/06/2022
# Author: Shun Suzuki
# -----
# Last Modified: 14/08/2022
# Modified By: Shun Suzuki (suzuki@hapis.k.u-tokyo.ac.jp)
# -----
# Copyright (c) 2022 Shun Suzuki. All rights reserved.
# 

Param(
    [string]$version = "2.3.1",
    [switch]$force
)

function ColorEcho($color, $PREFIX, $message) {
    Write-Host $PREFIX -ForegroundColor $color -NoNewline
    Write-Host ":", $message
}

$script_path = Split-Path $MyInvocation.MyCommand.Path -Parent

if (-not $force) {
    if (Test-Path "${script_path}/VERSION") { 
        $file = New-Object System.IO.StreamReader("${script_path}/VERSION", [System.Text.Encoding]::GetEncoding("utf-8"))
        $line = $file.ReadLine()
        if ($line -eq $version) {
            return;
        }
    }
}

ColorEcho "Green" "INFO" "Download ", $version, " libraries..."
$url = "https://github.com/shinolab/autd3/releases/download/v" + $version + "/autd3-v" + $version + "-win-x64.zip"

curl.exe -Lo "${script_path}/tmp.zip" $url

Expand-Archive "${script_path}/tmp.zip" -Force

Move-Item "${script_path}/tmp/bin/*" "${script_path}/" -Force
Move-Item "${script_path}/tmp/LICENSE" "${script_path}/" -Force
Move-Item "${script_path}/tmp/NOTICE" "${script_path}/" -Force

Remove-Item "${script_path}/tmp.zip" -Force
Remove-Item "${script_path}/tmp" -Recurse -Force

$file = New-Object System.IO.StreamWriter("${script_path}/VERSION", $false, [System.Text.Encoding]::GetEncoding("utf-8"))
$file.Write($version)
$file.Close()

ColorEcho "Green" "INFO" "Done"
