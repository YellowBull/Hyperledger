# Ubuntu 下部署 fabric

## 环境搭建

* ubuntu 16.0+ <br/>
* docker 17.0+ <br/>

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
`curl -L http://get.daocloud.io/docker/compose/releases/download/1.10.1/docker-compose-`uname -s`-`uname -m` > ~/docker-compose ` <br/>
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

在Fabric的源码中，提供了单机部署4Peer+1Orderer的示例，在Example/e2e_cli文件夹中。我们可以在其中一台机器上运行单机的Fabric实例，确认无误后，在该机器上，生成公私钥，修改该机器中的Docker-compose配置文件，然后把这些文件分发给另外4台机器。我们就以orderer.example.com这台机器为例

#### 单机运行4+1 Fabric实例，确保脚本和镜像正常