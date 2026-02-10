### ollama部署本地AI
#### 下载docker
```
# sudo su
# apt install docker.io
```
#### 下载ollama
```
# docker pull swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/ollama/ollama:0.11.4
```
#### 启动容器使用ollama镜像
```
# docker run -d -v ollama:/root/.ollama -p 11434:11434 --name ollama swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/ollama/ollama:0.11.4
# docker exec -it  ollama bash
```
#### 第二次使用该容器时候
```
# docker ps -a
# docker start ollama
# docker ps
# docker exec -it ollama bash
```
### ollama容器里面的操作
```
root@jasper:/home/jasper# docker exec -it ollama bash
root@f544a4ca7c0c:/# ollama list
NAME               ID              SIZE      MODIFIED          
deepseek-r1:14b    c333b7232bdb    9.0 GB    About an hour ago    
qwen3:14b          bdbd181c33f2    9.3 GB    2 hours ago          
root@f544a4ca7c0c:/# ollama run deepseek-r1:14b
>>> 你是谁
您好！我是由中国的深度求索（DeepSeek）公司开发的智能助手DeepSeek-R1。如您有任何任何问题，我会尽我所能为您提供帮助。

>>> 用python帮我设计判断是否是三角形的代码
Thinking...
好的，我现在需要帮用户设计一个用Python判断三个数是否能构成三角形的代码。首先，我要理解用户的需求，他们可能正在学习编程或者在做相关的项目。

用户的问题很明确，就是编写一个程序来判断输入的三个数是否满足三角形的条件。那我得先回忆一下三角形成立的条件是什么。记得三角形必须满足任意两边之和大于第三
边，所以这三个条件都需要满足：a + b > c，b + c > a，c + a > b。
>>> /bye
```
