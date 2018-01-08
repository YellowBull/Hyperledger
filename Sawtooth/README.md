# Hyperledger Sawtooth 环境安装

## 安装相关镜像镜像

新建一个工作目录
mkdir sawtooth
cd sawtooth
vim sawtooth-default.yaml
配置docker
sudo mkdir -p /etc/systemd/system/docker.service.d
cd /etc/systemd/system/docker.service.d/
vim http-proxy.conf

`[Service]
Environment="HTTP_PROXY=http://proxy.example.com:80/"`

打开一个终端窗口
cd sawtooth
docker-compose -f sawtooth-default.yaml up
