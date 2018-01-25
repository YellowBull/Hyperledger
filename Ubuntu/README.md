# Ubuntu 下部署 fabric

## 环境搭建

* ubuntu 16.0+ <br/>
* docker 17.0+ <br/>
### ubuntu 升级
`sudo do-release-upgrade`<br/> 
### Docker 环境搭建

#### Docker环境搭建[原始方式](https://yeasy.gitbooks.io/docker_practice/content/install/ubuntu.html)这里就不展示了<br/>
#### 在测试或开发环境中 Docker 官方为了简化安装流程，提供了一套便捷的安装脚本，Ubuntu 系统上可以使用这套脚本安装：<br/>
`curl -fsSL get.docker.com -o get-docker.sh` <br/>
`sh get-docker.sh --mirror Aliyun` <br/>

#### 启动 Docker CE
`systemctl enable docker` <br/>
`systemctl start docker` <br/>

#### 建立 Docker 组 <br/>
`groupadd docker` <br/>
`usermod -aG docker $USER` <br/>

#### 测试 Docker 是否安装正确(一般不会错) <br/>

#### 镜像加速 (不加速保证你后悔)<br/>

#### 在 `/etc/docker/daemon.json` 中写入以下信息 <br/>
```javascript
{
  "registry-mirrors": [
    "https://registry.docker-cn.com"
  ]
}
```
#### 之后重新启动服务 <br/>
`systemctl daemon-reload` <br/>
`systemctl restart docker` <br/>
* 注意：如果您之前查看旧教程，修改了 docker.service 文件内容，请去掉您添加的内容（--registry-mirror=https://registry.docker-cn.com），这里不再赘述。 <br/>

#### 检查加速器是否生效 <br/>

`docker info` <br/>
输出 `Registry Mirrors: `<br/>
` https://registry.docker-cn.com/` <br/>

### Docker-Compose 环境搭建

#### Docker-compose是支持通过模板脚本批量创建Docker容器的一个组件。在安装Docker-Compose之前，需要安装Python-pip，运行脚本：
`apt-get install python-pip` <br/>

#### 安装完成后，接下来从DaoClound安装Docker-compose，运行脚本：
`curl -L http://get.daocloud.io/docker/compose/releases/download/1.10.1/docker-compose-`\`uname -s\`-\`uname -m\` > ~/docker-compose <br/>
`sudo mv ~/docker-compose /usr/local/bin/docker-compose ` <br/>
`chmod +x /usr/local/bin/docker-compose` <br/>

#### Fabric源码下载

`go get github.com/hyperledger/fabric` <br/>

#### 这个可能等的时间比较久，等完成后，我们可以在~/go/src/github.com/hyperledger/fabric中找到所有的最新的源代码。
#### 由于Fabric一直在更新，所有我们并不需要最新最新的源码，需要切换到v1.0.0版本的源码即可：

`cd ~/go/src/github.com/hyperledger/fabric` <br/>
`git checkout v1.0.0` <br/>

#### Fabric Docker镜像的下载

`cd ~/go/src/github.com/hyperledger/fabric/examples/e2e_cli/` <br/>
`source download-dockerimages.sh -c x86_64-1.0.0 -f x86_64-1.0.0` <br/>

### docker-compose 配置文件准备

在Fabric的源码中，提供了单机部署4Peer+1Orderer的示例，在Example/e2e_cli文件夹中。我们可以在其中一台机器上运行单机的Fabric实例，确认无误后，在该机器上，生成公私钥，修改该机器中的Docker-compose配置文件，然后把这些文件分发给另外4台机器。我们就以orderer.example.com这台机器为例<br/>

#### 单机运行4+1 Fabric实例，确保脚本和镜像正常
`./network_setup.sh up`<br/>
这个命令可以在本机启动4+1的Fabric网络并且进行测试，跑Example02这个ChainCode。我们可以看到每一步的操作，最后确认单机没有问题。确认我们的镜像和脚本都是正常的，我们就可以关闭Fabric网络，继续我们的多机Fabric网络设置工作。关闭Fabric命令：<br/>
`./network_setup.sh down`<br/>

#### 生成公私钥、证书、创世区块等
公私钥和证书是用于Server和Server之间的安全通信，另外要创建Channel并让其他节点加入Channel就需要创世区块，这些必备文件都可以一个命令生成，官方已经给出了脚本：<br/>
`./generateArtifacts.sh mychannel`<br/>
运行这个命令后，系统会创建channel-artifacts文件夹，里面包含了mychannel这个通道相关的文件，另外还有一个crypto-config文件夹，里面包含了各个节点的公私钥和证书的信息。<br/>

#### 设置peer节点的docker-compose文件
e2e_cli中提供了多个yaml文件，我们可以基于docker-compose-cli.yaml文件创建：<br/>
cp docker-compose-cli.yaml docker-compose-peer.yaml <br/>

然后修改docker-compose-peer.yaml，去掉orderer的配置，只保留一个peer和cli，因为我们要多级部署，节点与节点之前又是通过主机名通讯，所以需要修改容器中的host文件，也就是extra_hosts设置，修改后的peer配置如下：<br/>

```javascript
peer0.org1.example.com:
  container_name: peer0.org1.example.com
  extends:
    file:  base/docker-compose-base.yaml
    service: peer0.org1.example.com
  extra_hosts:
   - "orderer.example.com:10.174.13.185"
```
同样，cli也需要能够和各个节点通讯，所以cli下面也需要添加extra_hosts设置，去掉无效的依赖，并且去掉command这一行，因为我们是每个peer都会有个对应的客户端，也就是cli，所以我只需要去手动执行一次命令，而不是自动运行。修改后的cli配置如下：<br/>

```javascript
cli: 
  container_name: cli 
  image: hyperledger/fabric-tools 
  tty: true 
  environment: 
    - GOPATH=/opt/gopath 
    - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock 
    - CORE_LOGGING_LEVEL=DEBUG 
    - CORE_PEER_ID=cli 
    - CORE_PEER_ADDRESS=peer0.org1.example.com:7051 
    - CORE_PEER_LOCALMSPID=Org1MSP 
    - CORE_PEER_TLS_ENABLED=true 
    - CORE_PEER_TLS_CERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/server.crt 
     - CORE_PEER_TLS_KEY_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/server.key 
    - CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt 
    - CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp 
  working_dir: /opt/gopath/src/github.com/hyperledger/fabric/peer 
  volumes: 
      - /var/run/:/host/var/run/ 
      - ../chaincode/go/:/opt/gopath/src/github.com/hyperledger/fabric/examples/chaincode/go 
      - ./crypto-config:/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ 
      - ./scripts:/opt/gopath/src/github.com/hyperledger/fabric/peer/scripts/ 
      - ./channel-artifacts:/opt/gopath/src/github.com/hyperledger/fabric/peer/channel-artifacts 
  depends_on: 
    - peer0.org1.example.com 
  extra_hosts: 
   - "orderer.example.com:10.174.13.185" 
   - "peer0.org1.example.com:10.51.120.220" 
   - "peer1.org1.example.com:10.51.126.19" 
   - "peer0.org2.example.com:10.51.116.133" 
   - "peer1.org2.example.com:10.51.126.5"
```

在单击模式下，4个peer会映射主机不同的端口，但是我们在多机部署的时候是不需要映射不同端口的，所以需要修改base/docker-compose-base.yaml文件，将所有peer的端口映射都改为相同的：<br/>

```javascript
ports: 
  - 7051:7051 
  - 7052:7052 
  - 7053:7053
```
#### 设置orderer节点的docker-compose文件
与创建peer的配置文件类似，我们也复制一个yaml文件出来进行修改：<br/>
`cp docker-compose-cli.yaml docker-compose-orderer.yaml`<br/>
orderer服务器上我们只需要保留order设置，其他peer和cli设置都可以删除。orderer可以不设置extra_hosts。<br/>

#### 分发配置文件
前面4步的操作，我们都是在orderer.example.com上完成的，接下来我们需要将这些文件分发到另外4台服务器上。Linux之间的文件传输，我们可以使用scp命令。<br/>

我先登录peer0.org1.example.com，将本地的e2e_cli文件夹删除：<br/>
`rm e2e_cli –R`<br/>
然后再登录到orderer服务器上，退回到examples文件夹，因为这样可以方便的把其下的e2e_cli文件夹整个传到peer0服务器上。<br/>

`scp -r e2e_cli fabric@10.51.120.220:/home/fabric/go/src/github.com/hyperledger/fabric/examples/`<br/>
我们在前面配置的就是peer0.org1.example.com上的节点，所以复制过来后不需要做任何修改。<br/>

再次运行scp命令，复制到peer1.org1.example.com上，然后我们需要对docker-compose-peer.yaml做一个小小的修改，将启动的容器改为peer1.org1.example.com，并且添加peer0.org1.example.com的IP映射，对应的cli中也改成对peer1.org1.example.com的依赖。这是修改后的peer1.org1.example.com上的配置文件：<br/>

```javascript
version: '2'

services:

  peer1.org1.example.com: 
    container_name: peer1.org1.example.com 
    extends: 
      file:  base/docker-compose-base.yaml 
      service: peer1.org1.example.com 
    extra_hosts: 
     - "orderer.example.com:10.174.13.185" 
     - "peer0.org1.example.com:10.51.120.220"

  cli: 
    container_name: cli 
    image: hyperledger/fabric-tools 
    tty: true 
    environment: 
      - GOPATH=/opt/gopath 
      - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock 
      - CORE_LOGGING_LEVEL=DEBUG 
      - CORE_PEER_ID=cli 
      - CORE_PEER_ADDRESS=peer1.org1.example.com:7051 
      - CORE_PEER_LOCALMSPID=Org1MSP 
      - CORE_PEER_TLS_ENABLED=true 
      - CORE_PEER_TLS_CERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/server.crt 
      - CORE_PEER_TLS_KEY_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/server.key 
      - CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/ca.crt 
      - CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp 
    working_dir: /opt/gopath/src/github.com/hyperledger/fabric/peer 
    volumes: 
        - /var/run/:/host/var/run/ 
        - ../chaincode/go/:/opt/gopath/src/github.com/hyperledger/fabric/examples/chaincode/go 
         - ./crypto-config:/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ 
        - ./scripts:/opt/gopath/src/github.com/hyperledger/fabric/peer/scripts/ 
        - ./channel-artifacts:/opt/gopath/src/github.com/hyperledger/fabric/peer/channel-artifacts 
    depends_on: 
      - peer1.org1.example.com 
    extra_hosts: 
     - "orderer.example.com:10.174.13.185" 
     - "peer0.org1.example.com:10.51.120.220" 
     - "peer1.org1.example.com:10.51.126.19" 
     - "peer0.org2.example.com:10.51.116.133" 
     - "peer1.org2.example.com:10.51.126.5"
```
接下来继续使用scp命令将orderer上的文件夹传送给peer0.org2.example.com和peer1.org2.example.com，然后也是修改一下docker-compose-peer.yaml文件，使得其启动对应的peer节点。<br/>

### 启动Fabric

#### 启动orderer
让我们首先来启动orderer节点，在orderer服务器上运行：<br/>
`docker-compose -f docker-compose-orderer.yaml up –d` <br/>
运行完毕后我们可以使用docker ps看到运行了一个名字为orderer.example.com的节点。<br/>

#### 启动peer
然后我们切换到peer0.org1.example.com服务器，启动本服务器的peer节点和cli，命令为：<br/>
`docker-compose -f docker-compose-peer.yaml up –d`<br/>
运行完毕后我们使用docker ps应该可以看到2个正在运行的容器。<br/>
接下来依次在另外3台服务器运行启动peer节点容器的命令：<br/>
`docker-compose -f docker-compose-peer.yaml up –d`<br/>
现在我们整个Fabric4+1服务器网络已经成型，接下来是创建channel和运行ChainCode。<br/>

### 创建Channel测试ChainCode
切换到peer0.org1.example.com服务器上，使用该服务器上的cli来运行创建Channel和运行ChainCode的操作。首先进入cli容器：<br/>
`docker exec -it cli bash`<br/>
官方已经提供了完整的创建Channel和测试ChainCode的脚本，并且已经映射到cli容器内部，所以我们只需要在cli内运行如下命令：<br/>
`./scripts/script.sh mychannel`<br/>
那么该脚本就可以一步一步的完成创建通道，将其他节点加入通道，更新锚节点，创建ChainCode，初始化账户，查询，转账，再次查询等链上代码的各个操作都可以自动化实现。直到最后，系统提示：<br/>
```javascript
===================== All GOOD, End-2-End execution completed =====================
```
```javascript
说明我们的4+1的Fabric多级部署成功了。我们现在是在peer0.org1.example.com的cli容器内，我们也可以切换到peer0.org2.example.com服务器，运行docker ps命令，可以看到本来是2个容器的，现在已经变成了3个容器，因为ChainCode会创建一个容器：

docker ps 
CONTAINER ID        IMAGE                                 COMMAND                  CREATED             STATUS              PORTS                              NAMES 
add457f79d57        dev-peer0.org2.example.com-mycc-1.0   "chaincode -peer.a..."   11 minutes ago      Up 11 minutes                                          dev-peer0.org2.example.com-mycc-1.0 
0c06fb8e8f20        hyperledger/fabric-tools              "/bin/bash"              13 minutes ago      Up 13 minutes                                          cli 
632c3e5d3a5e        hyperledger/fabric-peer               "peer node start"        13 minutes ago      Up 13 minutes       0.0.0.0:7051-7053->7051-7053/tcp   peer0.org2.example.com
```
### 总结
其中与单机部署最大的不同的地方就是在单机部署的时候，我们是在同一个docker网络中，所以相互之间通过主机名通讯很容易，而在多机环境中，就需要额外设置DNS或者就是通过extra_hosts参数，设置容器中的hosts文件。而且不能一股脑的就跟cli一样把5台机器的域名IP配置到peer中，那样会报错的，所以只需要设置需要的即可。<br/>
 









