## 示例错误
```
import App from './App.vue';
// 报错: 找不到模块 './App.vue' 或其相应的类型声明
```
## 解决方法
```
在项目根目录下找到或创建 env.d.ts 文件，并添加以下代码：
declare module '*.vue' {
 import type { DefineComponent } from 'vue';
 const vueComponent: DefineComponent<{}, {}, any>;
 export default vueComponent;
}
```
## vue2 optionsAPI
```
<template>
    <div class="person">
        <h2>姓名:{{ name }}</h2>
        <h2>年龄:{{ age }}</h2>
        <button @click="changeName">修改名字</button>
        <button @click="changeAge">修改年龄</button>
        <button @click="showTel">查看联系方式</button>
    </div>
</template>

<script>

    export default {
        name:'Person',
        data(){
            return {
                name:'张三',
                age:18,
                tel:'1388888888888'
            }
        },
        methods:{
            changeName(){
                this.name='zhang-san'
            },
            changeAge(){
                this.age+=1
            },
            showTel(){
                alert(this.tel)
            }
        }

}
</script>

<style lang="scss" scoped>
.person {
    background-color: skyblue;
    box-shadow: 0 0 10px;
    border-radius: 10px;
    padding: 20px;
}
button{
    margin: 0 5px
}
</style>
```
## vue2和vue3可以共存但是vue2可以调用vue3 vue3不能影响vue2
```
<script lang="ts" setup></script>相当于写了一个setup函数并且自动return
```
## 下载插件
```
(base) PS C:\workspace\VueCom\demo>  npm i vite-plugin-vue-setup-extend -D
npm warn deprecated sourcemap-codec@1.4.8: Please use @jridgewell/sourcemap-codec instead

added 3 packages, and audited 201 packages in 4s

49 packages are looking for funding
  run `npm fund` for details

found 0 vulnerabilities
```
## vite.config.ts引入插件
```
import VueSetup from 'vite-plugin-vue-setup-extend'
  plugins: [
    vue(),
    vueDevTools(),
    VueSetup()
  ],
```