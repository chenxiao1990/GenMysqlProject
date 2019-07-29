# GenMysqlProject

#### 介绍
通过mysql数据表初始化 curd项目 。  使用 gin  gorm

#### 安装教程

1. xxxx
2. xxxx
3. xxxx

#### 使用说明

1. xxxx
2. xxxx
3. xxxx


#### 软件架构
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


 