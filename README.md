# 基于Fabric的农产品溯源与交易系统 



### 模块和技术

这个demo只写了fabric区块链网络和gin框架的后端，前端不太熟悉就没写。

```
network： fabric区块链网络

server： 使用gin写的后端

区块链：Fabric 2.2.10

链码： Go

环境： Ubuntu20.04，  Docker 20.10.21，  Docker-compose 1.25.0,  Go 1.18.2
```

##### 区块链部分

只用了最简单的双peer节点和一个order节点，其实应该给每一个组织节点都有一个peer

分为7个链码，将各部分的数据分别上链，然后可以溯源查询。本想将溯源过程和商城结合起来使用，但最后还是做的有些简陋。

```
farmer：负责记录作物和作物的生长过程，可以查询作物信息和作物生长信息

driver：负责作物运输环节数据，记录运输信息，根据运输ID或作物ID查询运输信息

material：记录和查询加工信息

sell：记录和查询交易信息

seller：虽然叫seller，但其实记录的是商品信息，上架、修改、撤销商品等，查询商品信息

capital：记录资产信息，查询ID对应资产，修改资产等

users：记录账户信息，可以新建账户，更改密码等
```

##### gin后端

接收GET或POST请求，调用区块链中相应的链码方法



### 启动demo

##### 拉取docker镜像

```sh
docker pull hyperledger/fabric-peer:2.2.10
docker pull hyperledger/fabric-orderer:2.2.10
docker pull hyperledger/fabric-ca:1.5.5
docker pull hyperledger/fabric-tools:2.2.10
docker pull hyperledger/fabric-ccenv:2.2.10
docker pull hyperledger/fabric-baseos:2.2.10
docker pull couchdb:3.1.1
```

##### 启动network下的start.sh

```sh
chmod +x ./start.sh
sudo ./start.sh
```

##### 启动server下的start.sh

```sh
chmod +x ./start.sh
sudo ./start.sh
```

##### 访问

因为没有前端，所以这里只能用浏览器或者别的软件直接发送相应的GET或POST请求

访问本机的8282端口即可

查询数据测试

![image-20230417180450888](https://files.cnblogs.com/files/blogs/753861/image-20230417180450888.gif?t=1681736591)

将Data中的数据base64解码得到如下数据

![image-20230417180538016](https://files.cnblogs.com/files/blogs/753861/image-20230417180538016.gif?t=1681736597)

这些操作本该在前端完成并展示的，此处只能这样展示了

##### 停止

停止server，按下Ctrl+C停止运行main.go，然后执行stop.sh

停止network，执行stop.sh脚本即可

### 一些错误

```
Error: failed to create deliver client for orderer: orderer client failed to connect to orderer.example.com:7050: failed to create new connection: connection error: desc = "transport: error while dialing: dial tcp: lookup orderer.example.com: no such host"
```

需要关闭防火墙
