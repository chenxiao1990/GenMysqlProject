# GenMysqlProject

#### 介绍
通过mysql数据表初始化各个表的curd功能，包含api、service、dao、model模块 。 

api使用gin框架提供对外http接口

mysql数据库访问使用gorm框架


#### 安装教程

+ 下载编译

```
git clone git@gitee.com:290746987/GenMysqlProject.git
进入GenMysqlProject
go build
```

+ go get 直接安装

```
go get -u gitee.com/290746987/GenMysqlProject
还是需要把模板文件夹下载下来, 模板文件夹是ProjectTemplate
在ProjectTemplate的同目录执行 GenMysqlProject 命令
```


#### 使用说明

+ 使用 GenMysqlProject -h 获取说明
+ 命令 GenMysqlProject --ipport "192.168.0.86:3306" --dbname "cx" --user "root" --pass "12345678" --outname "myproject" 



#### 软件架构

+ config

配置相关的信息在config.json中，项目运行时config.json放在同目录下./config/config.json

+ api

对外提供http接口服务
api解析service层需要的参数，调用service服务

+ service

对api层提供服务
定义服务的参数struct，返回struct
service可以调用dao层的接口对数据库进行操作，但是不可以直接调用操作数据库
多数情况service层是直接调用一个dao层进行返回，也可以进行多个表数据之间的计算整合等操作。

+ dao

对service层提供数据库操作接口
可以提供基础的增删改查，如果联合查询，那么接口应该定义在主表的dao文件中

+ model

数据库表结构
提供gorm的自动创建表功能


#### 注意事项

+ 表中如果无id字段，那么要删掉dao层中相关使用id的操作
+ 如果有两个表，表名是 something, somethings 那么需要手动处理一下，会有两个Something结构生成
+ 主流的字段格式已解析，可能会有漏掉的，后边会根据反馈进行修改

 