## 一、开发流程
### 1.创建maven项目
```
//在pom.xml里面引入配置
<!--所有springboot项目都继承这个父项目-->
    <parent>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-parent</artifactId>
        <version>3.0.5</version>
    </parent>
```
### 2.导入场景
```
    <dependencies>
<!--        web开发场景启动器-->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
```
### 3.主程序
```
@SpringBootApplication //这是一个springboot应用
public class MainApplication {
    public static void main(String[] args) {
        SpringApplication.run(MainApplication.class, args);
    }
}
```
### 4.业务书写
```
//@Controller
//@ResponseBody
@RestController
public class HelloController {
    @GetMapping("/hello")
    public String hello(){
        return "hello world";
    }
}
```
### 5.测试
```
springboot 应用默认启动localhost:8080
```
### 6.打包插件
```
<!--        SpringBoot应用打包插件-->
    </dependencies>
    <build>
        <plugins>
            <plugin>
                <groupId>org.springframework.boot</groupId>
                <artifactId>spring-boot-maven-plugin</artifactId>
            </plugin>
        </plugins>
    </build>
```
```
使用mvn clean package把应用打成jar包
使用java -jar demo.jar启动项目
```
## 二、优点
### 1.简化整合
```
导入相关的场景，拥有相关的功能。场景启动器
默认支持的所有场景：https://docs.spring.io/spring-boot/docs/current/reference/html/using.html#using.build-systems.starters
* 官方提供的场景：命名为：spring-boot-starter-*
* 第三方提供场景：命名为：*-spring-boot-starter
场景一导入，万物皆就绪
```
### 2. 简化开发
```
无需编写任何配置，直接开发业务
```
### 3. 简化配置
```
application.properties：
* 集中式管理配置。只需要修改这个文件就行。
* 配置基本都有默认值
* 能写的所有配置都在：https://docs.spring.io/spring-boot/docs/current/reference/html/application-properties.html#appendix.application-properties
```
### 4. 简化部署
```
打包为可执行的jar包。
linux服务器上有java环境即可
```
### 5. 简化运维
```
修改配置（外部放一个）、监控、健康检查。
```
## 三、spring initializer
```
一键创建spring项目
```
## 四、spring boot好处
### 依赖管理机制
#### 1、为什么导入 starter-web 所有相关依赖都导入进来？
```
开发什么场景，导入什么场景启动器。
maven依赖传递原则。A-B-C：A就拥有B和C
导入场景启动器。场景启动器 自动把这个场景的所有核心依赖全部导入进来
```
#### 2、为什么版本号都不用写？
```
每个boot项目都有一个父项目 spring-boot-starter-parent
parent的父项目是 spring-boot-dependencies
父项目 版本仲裁中心，把所有常见的jar的依赖版本都声明好了。
比如：mysql-connector-j
```
#### 3、自定义版本号
```
利用maven的就近原则
直接在当前项目 properties 标签中声明父项目用的版本属性的key
直接在导入依赖的时候声明版本
```
#### 4、导入第三方包要自己写版本
```
<!-- https://mvnrepository.com/artifact/com.alibaba/druid -->
<dependency>
    <groupId>com.alibaba</groupId>
    <artifactId>druid</artifactId>
    <version>1.2.16</version>
</dependency>
```
### 默认的包扫描规则
```
@SpringBootApplication 标注的类就是主程序类
SpringBoot只会扫描主程序所在的包及其下面的子包，自动的component-scan功能
```
#### 自定义扫描路径
```
@SpringBootApplication(scanBasePackages = "com.atguigu")
@ComponentScan("com.atguigu") 直接指定扫描的路径
```
#### 配置默认值
```
配置文件的所有配置项是和某个类的对象值进行一一绑定的。
绑定了配置文件中每一项值的类：配置属性类。
比如：
  ServerProperties 绑定了所有Tomcat服务器有关的配置
  MultipartProperties 绑定了所有文件上传相关的配置
  ...参照官方文档：
```
#### 按需加载自动配置
```
导入场景 spring-boot-starter-web
场景启动器除了会导入相关功能依赖，导入一个 spring-boot-starter，是所有 starter 的 starter，基础核心starter
spring-boot-starter 导入了一个包 spring-boot-autoconfigure。包里面都是各种场景的 AutoConfiguration 自动配置类
虽然全场景的自动配置都在 spring-boot-autoconfigure 这个包，但不是全都开启的。
导入哪个场景就开启哪个自动配置
```
