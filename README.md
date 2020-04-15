# ui

![Image1](https://raw.githubusercontent.com/chenxiao1990/GenMysqlProject/master/images/1.1.png)

![Image2](https://raw.githubusercontent.com/chenxiao1990/GenMysqlProject/master/images/2.png)

![Image3](https://raw.githubusercontent.com/chenxiao1990/GenMysqlProject/master/images/3.png)

# GenMysqlProject

#### 介绍
json 转 go struct
mysql 数据库表 转 go struct

生成mysql数据表的 gorm查询go代码

通过mysql数据表初始化各个表的curd功能，包含api、service、dao、model模块 。 

api使用gin框架提供对外http接口

mysql数据库访问使用gorm框架



#### 安装教程
+ go get
```
go get -u -v github.com/chenxiao1990/GenMysqlProject

(go get 最好在go.mod的工程下使用， 好走配置的proxy加速， 不然就得科学上网了)
```
+ 下载源码编译

```
git clone git@github.com:chenxiao1990/GenMysqlProject.git
cd GenMysqlProject
go build 
./GenMysqlProject
```

 
#### 使用说明

编译成可执行程序后直接运行，然后浏览器打开 http://localhost:8008/vue

+ 生成工程    
  输入工程名，点击生成，会在运行目录下生成工程
+ 生成查询代码  
  进行一些点击操作生成查询数据库的代码


#### 注意事项

+ 表中如果无id字段，那么要删掉dao层中相关使用id的操作
+ 如果有两个表，表名是 something, somethings 那么需要手动处理一下，会有两个Something结构生成
+ 主流的字段格式已解析，可能会有漏掉的，后边会根据反馈进行修改

 

# web端工程的 静态资源打包 go-bindata工具

```
本工程中已打包，所以你不想做更改的话 可以不用管这个

go-bindata -fs -prefix "webui/dist/" webui/dist/...

-fs 表示生成 func AssetFile() http.FileSystem

-prefix  去掉生成包的访问前缀 webui/dist

成功生成bindata.go文件 里面就是一堆静态文件的字符串
GRouter = gin.Default()
GRouter.StaticFS("/vue", AssetFile())
这样 gin就添加了 dist里的静态文件访问  
例如 webui/dist/css/app.xxx.css 访问路径为 http://ip:port/vue/css/app.xxx.css
```
