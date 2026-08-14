## 创建项目
```
npm create vue@latest VueCom
```
## Nexus创建npm-hosted仓库
```
选择仓库类型：在点击 Create repository 后，你会看到一个长长的仓库类型列表。因为我们需要一个能存储和发布私有包的仓库，所以在列表中找到并选择 npm (hosted)。
Name（仓库名称）：这是最重要的，给它起一个你记得住的名字，比如 npm-private或 npm-hosted。请记下这个名字，后续发布包时会用到。
Online（是否启用）：确保这个选项是勾选状态（默认为勾选），表示仓库处于可用状态。
Blob store（存储位置）：这是存放包文件的地方。你可以使用默认的 default，也可以点击左侧的Blob Stores提前创建一个专用的存储空间，方便管理。
关键设置：部署策略（Deployment policy）
在页面下方，找到 Deployment policy 下拉菜单，这个选项决定了你能否覆盖已发布的包版本：
Allow redeploy（允许重新部署）：推荐选择此项。它允许你覆盖已发布的相同版本号的包，在开发和测试阶段非常灵活。
Disable redeploy（禁止重新部署）：更严格，一旦发布就不能更改，适合生产环境。
Read-only（只读）：禁止任何发布操作。
完成创建：检查所有配置无误后，点击页面最下方的 Create repository 蓝色按钮，你的私有 npm 仓库就创建好了。
```
## 修改package.json
```
  "private": false,
  "type": "module",
    "publishConfig": {
    "registry": "http://nexus.dennis.com.cn:8081/repository/npm-hosted/"
  },
  "main": "dist/index.js",         
  "types": "dist/index.d.ts",      
  "files": [                    
    "dist"
  ],
  //这一部分需要修改或添加
```
## 终端执行登陆和发布
```
npm login --username "admin" --password "xiechuan110&" --registry=http://nexus.dennis.com.cn:8081/repository/npm-hosted/
npm publish --registry=http://nexus.dennis.com.cn:8081/repository/npm-hosted/
```
## 另一个项目如何引用
```
npm install my-private-package --registry=http://nexus.dennis.com.cn:8081/repository/npm-group/
```
