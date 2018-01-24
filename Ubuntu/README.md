# Ubuntu 下部署 fabric

## 环境搭建

* ubuntu 16.0+ <br/>
* docker 17.0+ <br/>

### docker 环境搭建

#### docker环境搭建[原始方式](https://yeasy.gitbooks.io/docker_practice/content/install/ubuntu.html)这里就不展示了<br/>
#### 在测试或开发环境中 Docker 官方为了简化安装流程，提供了一套便捷的安装脚本，Ubuntu 系统上可以使用这套脚本安装：<br/>
`curl -fsSL get.docker.com -o get-docker.sh` <br/>
`sh get-docker.sh --mirror Aliyun` <br/>

#### 启动 Docker CE
`systemctl enable docker` <br/>
`systemctl start docker` <br/>

#### 建立 docker 组 <br/>
`groupadd docker` <br/>
`usermod -aG docker $USER` <br/>

#### 测试 Docker 是否安装正确(一般不会错) <br/>

#### 镜像加速 (不加速保证你后悔)<br/>

#### 在 `/etc/docker/daemon.json` 中写入以下信息 <br/>
`{ `<br/>
`  "registry-mirrors": `[ <br/>
`    "https://registry.docker-cn.com"` <br/>
`  ] `<br/>
`}` <br/>
#### 之后重新启动服务 <br/>
`systemctl daemon-reload` <br/>
`systemctl restart docker` <br/>
* 注意：如果您之前查看旧教程，修改了 docker.service 文件内容，请去掉您添加的内容（--registry-mirror=https://registry.docker-cn.com），这里不再赘述。 <br/>

#### 检查加速器是否生效 <br/>

`docker info` <br/>
输出 `Registry Mirrors: `<br/>
` https://registry.docker-cn.com/` <br/>


