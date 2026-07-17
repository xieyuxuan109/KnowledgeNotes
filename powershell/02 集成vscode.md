## 下载powershell插件
## 安装powershell最新版
## 编写test.ps1
```
param(
    [Parameter(Mandatory)]
    [int]$age,
    [string]$username = "xieyuxuan"
)
if ($age -eq 10)
{
    Write-Host "Hello,$age=$username"
} else {
    Write-Host "Hello,$age!=$username"<# Action when all if and elseif conditions are false #>
}
```
## 编写launch.json
```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Hello Script",
            "type": "PowerShell",
            "request": "launch",
            "script": "${file}",
            "cwd": "${workspaceFolder}",
            "args": [
                "-age 10",
                "-username xieyuxuan" //注意变量写法
            ]
        },
        
    ]
}
```