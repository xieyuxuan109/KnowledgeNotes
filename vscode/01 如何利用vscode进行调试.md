### 如何利用vscode进行调试
```
    go install github.com/go-devel/devel/cmd/dlv@latest
```
### 需要添加launch.json
```
{
    "version":"0.2.0",
    "configurations":{
        "name":"Launch Package"，
        "type":"go",
        "request":"launch",
        "mode":"auto",
        "program":"${workspaceFolder}"//根据实际情况修改，这个是指main.go所在位置，默认是和.vscode同一级
    }
}
```