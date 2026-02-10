### createWebHistory与createWebHashHistory的区别
```
在Vue 3中，createWebHistory和createWebHashHistory是Vue Router提供的两种路由模式。这两种模式的主要区别在于它们如何与浏览器的历史记录API和URL结构交互。

createWebHistory

createWebHistory基于HTML5的history API，它允许在不重新加载页面的情况下操作浏览器历史记录。使用createWebHistory时，URL看起来很简洁，没有额外的字符，例如/home或/about。这种模式需要服务器配置以处理前端路由，否则直接访问或刷新页面可能会导致404错误。此外，createWebHistory只支持支持HTML5 history API的现代浏览器。

import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import About from '@/views/About.vue'

const router = createRouter({
history: createWebHistory(),
routes: [
{ path: '/home', name: 'home', component: Home },
{ path: '/about', name: 'about', component: About }
]
})
复制
createWebHashHistory

createWebHashHistory使用URL的hash部分来管理路由状态。这种模式在URL中添加了#，例如/#/home或/#/about。它不需要服务器端配置，因为hash值不会发送到服务器，因此可以避免404错误。createWebHashHistory在所有浏览器中都受支持，包括旧版浏览器，但可能对搜索引擎优化（SEO）不利。

import { createRouter, createWebHashHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import About from '@/views/About.vue'

const router = createRouter({
history: createWebHashHistory(),
routes: [
{ path: '/home', name: 'home', component: Home },
{ path: '/about', name: 'about', component: About }
]
})
复制
选择哪种模式

选择哪种路由模式取决于你的具体需求和服务器配置。如果你的服务器已经配置了URL重写并且你希望拥有干净的URL，那么createWebHistory可能是更好的选择。如果你的应用需要支持旧版浏览器或者你不想在服务器上进行额外配置，createWebHashHistory可能更适合。无论选择哪种模式，都要确保它符合你的开发需求和用户体验。
```