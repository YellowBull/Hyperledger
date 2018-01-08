# Hyperledger Sawtooth 环境安装

## 安装相关镜像镜像

新建一个工作目录 <br/>
mkdir sawtooth <br/>
cd sawtooth <br/>
vim sawtooth-default.yaml <br/>
配置docker <br/>
sudo mkdir -p /etc/systemd/system/docker.service.d <br/>
cd /etc/systemd/system/docker.service.d/ <br/>
vim http-proxy.conf <br/>

`[Service]
Environment="HTTP_PROXY=http://proxy.example.com:80/"` <br/>

打开一个终端窗口 <br/>
cd sawtooth <br/>
docker-compose -f sawtooth-default.yaml up <br/>
