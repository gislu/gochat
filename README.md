# GoChat
A open-source WeChat Official Account server framework based on Beego

VERSION = "1.0"
## 获取安装

执行以下命令，就能够在你的`GOPATH/src` 目录下发现gochat
```bash
$ go get github.com/gislu/gochat
```
##初次使用

###配置文件
找到根目录下/conf/id_relative.yaml 这个文件，程序会自动读取配置文件中需要的数据

公共号：请把配置文件的pubid 、pubsecret 、 token一栏后面对应的值换成自己公共号上的AppID/应用ID 、 AppSecret/应用密钥以及开发者配置的token

企业号：请把配置文件中的torpid 、corpsecret 、 key 、token 一栏换成自己企业号对应的CorpID 、 Secret 、企业号应用的AesKey 、以及对应应用开发者配置中的token

###编译及运行应用
进入项目根目录，直接go build编译项目即可
```
$ go build -o gochat main.go
```

创建成功以后，你能得到一个名叫gochat的应用程序,在有该程序的目录下输入
```
$nohup ./gochat &
```
即可运行该程序，然后在公共号/企业号的开发者配置中启用”服务器配置/回调模式“

token随意设置,将url 设置为: 域名/pub(企业号则是:域名/corp) 即可：
```bash
example: www.myhost.com/pub
```
启动”服务器配置/回调模式“，即可成功运行一个基础版的微信公共号/企业号gochat框架

