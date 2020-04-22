## TCP 并发多连接服务器
适合安卓APP、下位机硬件数据交互、云服务存储。上位机esp8266 上传数据至本tcp服务器，服务器接收并存储至mysql中，同时向在线的APP推送数据；APP可以往下位机发送指令、查询服务器上保存的数据。

#### 一、 预备条件

-   mysql 8.0
-   windows/linux
-   git
-   NetAssit

#### 二、 配置说明
配置文件在config/server.conf内，可修改数据库类型、名称。

#### 三、测试例子
App、Node1、Node2连接至tcp服务器

##### App
认证身份
```
{"did":1234,"role":1}
```
获取在线node列表
```
{ "to_did":0,"cmd":"getNodes"}
```
向下位机发送指令
```
{ "to_did":2345,"cmd":"send:hello"}
```
获取node的历史数据
```
{ "to_did":2345,"cmd":"getData"}
```
##### Node1、2
认证身份
```
{"did":2345,"role":2}
```
上传数据
```
{"light":100,"mq2":10.1,"mq135":10.3,"temp":30.0,"wet":100.0}
```
