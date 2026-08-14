## 在本地先设置一遍
```
Jasper@DESKTOP-1JMSUMK MINGW64 /c/workspace/VueCom/VueCom (main)
$ npm config set registry=http://nexus.dennis.com.cn:8081/repository/npm-group/

Jasper@DESKTOP-1JMSUMK MINGW64 /c/workspace/VueCom/VueCom (main)
$ npm login
npm notice Log in on http://nexus.dennis.com.cn:8081/repository/npm-group/
Username: admin
Password: 
Logged in on http://nexus.dennis.com.cn:8081/repository/npm-group/.

Jasper@DESKTOP-1JMSUMK MINGW64 /c/workspace/VueCom/VueCom (main)
$ cat ~/.npmrc
registry=http://nexus.dennis.com.cn:8081/repository/npm-group/
//nexus.dennis.com.cn:8081/repository/npm-hosted/:_authToken=NpmToken.<你的token>
//nexus.dennis.com.cn:8081/repository/npm-group/:_authToken=NpmToken.<你的token>
```
## 将~/.npmrc拷贝进项目.npmrc文件
```
registry=http://nexus.dennis.com.cn:8081/repository/npm-group/
//nexus.dennis.com.cn:8081/repository/npm-hosted/:_authToken=NpmToken.<你的token>
//nexus.dennis.com.cn:8081/repository/npm-group/:_authToken=NpmToken.<你的token>
```
## gitlab-ci.yml中进行处理
```
stages:
  - build
  - publish
variables:
  NexusUrl: "http://nexus.dennis.com.cn:8081/repository/npm-hosted/"
build-job:
  stage: build
  image: node:24.13
  artifacts:
    paths:
      - pig-ui/
    expire_in: 1 week
  script:
    # - pwd
    - npm install
    # - sleep 1000
    - npm run package
    - ls
    # - sleep 1000
publish-job:
  stage: publish
  image: node:24.13
  dependencies:
    - build-job
  script:
  - ls
  - cp ./.npmrc ~/.npmrc
  - ls ~/.npmrc
  - npm publish
```