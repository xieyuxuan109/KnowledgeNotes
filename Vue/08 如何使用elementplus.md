## element plus安装与使用
```
npm create vue@latest
npm install element-plus --save
```
## 修改main.ts文件
```
import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import ElementPlus from 'element-plus' //新增
import 'element-plus/dist/index.css' //新增
const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(ElementPlus) //新增
app.mount('#app')
```
## 配置tsconfig.node.json文件
```
"types": ["node","element-plus/global"], //修改这一行
```
## 修改helloworld.vue组件
```
  <div class="greetings">
    <h1 class="green">{{ msg }}</h1>
    <el-empty description="空空如也"></el-empty> //新增
    <h3>
      You’ve successfully created a project with
      <a href="https://vite.dev/" target="_blank" rel="noopener">Vite</a> +
      <a href="https://vuejs.org/" target="_blank" rel="noopener">Vue 3</a>. What's next?
    </h3>
  </div>
```
```
其中el-empty组件是一个空态页组件 用来展示无数据时的页面占位符
```
## 按钮组件
```
element-plus中提供了el-button组件来创建按钮 el-button自建提供了很多属性来对按钮的样式进行定制
```
## size属性
```
size 属性 主要用来设置按钮大小尺寸 default（默认值） large small三种值
<template>
  <div class="about">
    <h1>Button组件的用法</h1>
    <div>
      <el-button>Default</el-button>
      <el-button size="large">Large</el-button>
      <el-button size="small">Large</el-button>
    </div>
  </div>
</template>
```
## type 属性
```
type属性'default' | 'primary' | 'success' | 'warning' | 'danger' | 'info' 
primary:常规风格
success:成功风格
warning:警告风格
danger:危险风格
info:详情风格
text : 文本风格
```
## plain属性
```
控制是填充风格还是描边风格
<el-button type="success" plain>Large</el-button>
<el-button type="warning" plain>Large</el-button>
```
## circle和round
```
circle表示是不是圆角
round表示是不是圆形
```
## 换行
```
<br></br>
```
## loading和disabled
```
可以使用loading插槽或loadingIcon属性自定图标，代码如下所示：
<br>
  <br>
  <!--loading属性-->
  <el-button>默认按钮</el-button>
  <el-button loading="true" type="primary">常规按钮</el-button>
  <el-button type="success">成功按钮</el-button>
  <el-button loading type="info">信息按钮</el-button>
  <el-button loading :loading-icon="Eleme" type="warning">警告按钮</el-button>
  <el-button loading type="danger">
    <template #loading>
      <el-icon><Eleme/></el-icon>
    </template>
  危险按钮</el-button>
```
## tag和icon属性
```
tag属性用来自定义属性标签
      <!-- tag属性 -->
       <br><br>
      <el-button>默认按钮</el-button>
      <el-button tag="div" role="button" tabindex="0">div</el-button>
      <el-button tag="a" href="http://www.baidu.com" target="_blank" type="success">你好</el-button>
```
## icon引入
```
安装图标库
首先，确认你是否已经安装了 @element-plus/icons-vue 包。如果没有，在项目根目录运行以下命令之一进行安装：

bash
npm install @element-plus/icons-vue
# 或
yarn add @element-plus/icons-vue
# 或
pnpm install @element-plus/icons-vue
注册图标组件
安装后，你需要在 Vue 应用中将所有图标进行全局注册，这样它们才能在模板中使用。
在你的 main.js 或 main.ts 文件中添加以下代码：

javascript
import { createApp } from 'vue'
import App from './App.vue'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

const app = createApp(App)

// 全局注册所有图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

app.mount('#app')
使用 :icon 绑定图标组件
完成以上两步后，你的代码就应该能正常工作了。

html
<template>
  <!-- 直接使用 :icon 绑定导入的图标组件 -->
  <el-button :icon="Plus" tag="div" role="button" tabindex="0">div</el-button>
  <el-button :icon="EditPen" tag="a" href="http://www.baidu.com" target="_blank" type="success">你好</el-button>
</template>

<script setup>
// 确保在 <script setup> 中导入了图标组件
import { Plus, EditPen } from '@element-plus/icons-vue'
</script>
```
## color属性
```
自定义按钮颜色 并自动计算hover和active触发后的颜色，代码如下所示
<el-button color="#d52">谢宇轩</el-button>
```
## 标签组件 el-tag
## type 属性
```
type标签类型
success 成功风格
info 详情风格
warning 警告风格
danger 危险风格
  <el-tag type="success">成功风格</el-tag>
  <el-tag type="info">成功风格</el-tag>
  <el-tag type="warning">成功风格</el-tag>
  <el-tag type="danger">成功风格</el-tag>
```
## size属性 round属性
```
large small default
```
## hit属性表示是否描边
```
:hit="true" "false"
```
## color属性 给el-tag设置背景颜色
## effect属性
```
用来设置主题 dark light plain
```
## 空态图用法
## count属性控制渲染骨架模板个数 骨架屏额外渲染行数为rows控制 loading控制是否显示出来
```
  <!-- 空态图用法 -->
   <el-skeleton :animated="true"></el-skeleton>
   <el-skeleton :animated="false"></el-skeleton>
```
## el-avatar shape用来设置形状 circle 和 square size控制大小
```
   <el-avatar shape = "circle" :src = "url" />
    </div>
  </div>

</template>
<script setup>
// 确保在 <script setup> 中导入了图标组件
import { Plus, EditPen } from '@element-plus/icons-vue'
import {reactive,toRefs} from 'vue'
const state = reactive({
  url:'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png'
})
const {url}=toRefs(state)
```