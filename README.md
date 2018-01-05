# Hyperledger composer 多节点部署

## 准备工作

要运行Hyperledger Composer和Hyperledger Fabric，建议您至少拥有4Gb的内存。

### 以下是安装所需开发工具的先决条件：

操作系统：Centos7.1</br>
Docker：版本17.03或更高</br>
Docker-Compose：版本1.8或更高</br>
node：8.9或更高（注意版本9不支持） </br>
npm：v5.x</br>
git：2.9.x或更高版本</br>
Python：2.7.x</br>

### 开发工具的具体安装方式

#### 操作系统

centos7的操作系统我就不说了，这个是基础

#### docker 环境搭建

添加docker资源库的配置文件 <br/>
yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo <br/>
通过docker资源文件下载docker <br/>
yum install docker-ce <br/>
查看dockers运行状态 <br/>
systemctl status docker <br/>
启动docker <br/>
systemctl start docker <br/>
查看安装生成的文件 <br/>
rpm -ql docker-ce <br/>
修改用户权限，把需要操作docker的用户加入docker用户组 <br/>
usermod -G docker,wheel username <br/>

#### docker-compose 环境搭建

curl -L https://get.daocloud.io/docker/compose/releases/download/1.9.0/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose <br/>
chmod +x /usr/local/bin/docker-compose <br/>

#### node 环境搭建

cd /usr/local/src/ <br/>
sudo wget https://nodejs.org/dist/latest-v8.x/node-v8.9.3-linux-x64.tar.gz   <br/>
sudo tar zxvf node-v8.9.3-linux-x64.tar.gz  <br/>
cd node-v8.9.3-linux-x64/bin <br/>
sudo ln -s /usr/local/src/node-v8.9.3-linux-x64/bin/node /usr/local/bin/node <br/>
sudo ln -s /usr/local/src/node-v8.9.3-linux-x64/bin/npm /usr/local/bin/npm <br/>

#### cmpm 环境搭建（用于国内替代npm）

cnpm环境搭建
npm install -g cnpm --registry=https://registry.npm.taobao.org












