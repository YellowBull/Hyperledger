# blockchain -explorer
Hyperledger Explorer是一个简单，功能强大，易于使用，高度维护的开源浏览器，用于查看底层区块链网络上的活动。<br/>
## 克隆存储库
### 使用以下命令克隆此存储库以获取最新版本。

`git clone https://github.com/hyperledger/blockchain-explorer.git`<br/>
`cd blockchain-explorer`<br/>

## 数据库设置

## mysql数据库设置
centos下安装mysql<br/>
`yum list installed mysql*`<br/>
首先需要安装mariadb-server<br/>
`yum install -y mariadb-server`<br/>
启动<br/>
`systemctl start mariadb.service`<br/>
设置开机启动<br/>
`systemctl enable mariadb.service`<br/>
进行一些安全设置，以及修改数据库管理员密码<br/>
`mysql_sceure_installation`<br/>
测试<br/> 
`mysql -u root -p`<br/>

### 运行位于下的数据库设置脚本 db/fabricexplorer.sql

`mysql -u<username> -p < db/fabricexplorer.sql`<br/>

## 结构网络设置
此存储库附带一个示例网络配置，以开始<br/>

`cd first-network`<br/>
下载必要的二进制文件和hhyperledger docker images<br/>
`chmod u+x bootstrap-1.0.2.sh`<br/>
`./bootstrap-1.0.2.sh`<br/>
`mkdir -p ./channel-artifacts`<br/>
`./byfn.sh -m generate -c mychannel`<br/>
`./byfn.sh -m up -c mychannel`<br/>
这带来了一个2渠道名称的组织网络mychannel。<br/>

或者，您可以使用Fabric 构建网络教程来设置自己的网络。一旦你设置网络，请相应地修改值config.json。<br/>

## 运行区块链 - 资源管理器
在另一个终端上，<br/>

`cd blockchain-explorer`<br/>
修改config.json来更新一个频道<br/>
mysql主机，用户名，密码的详细信息<br/>
```javascript
"channel": "mychannel",
 "mysql":{
      "host":"127.0.0.1",
      "database":"fabricexplorer",
      "username":"root",
      "passwd":"123456"
   }
```
如果要连接到非TLS结构对等体，请修改对等URL中的协议（grpcs->grpc）和端口（9051-> 9050），然后删除tls_cacerts。根据这个键，应用程序决定是去TLS还是非TLS路线。<br/>

`npm install`<br/>
`./start.sh`<br/>
在浏览器上启动URL `http://localhost:8080`<br/>