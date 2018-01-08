# Hyperledger Sawtooth 环境安装

## 安装相关镜像镜像

* 新建一个工作目录 <br/>
mkdir sawtooth <br/>
cd sawtooth <br/>
vim sawtooth-default.yaml <br/>
* 配置docker <br/>
sudo mkdir -p /etc/systemd/system/docker.service.d <br/>
cd /etc/systemd/system/docker.service.d/ <br/>
vim http-proxy.conf <br/>
`[Service]` <br/>
`Environment="HTTP_PROXY=http://proxy.example.com:80/"` <br/>

* 打开一个终端窗口 <br/>
cd sawtooth <br/>
docker-compose -f sawtooth-default.yaml up <br/>

## 登录到客户端容器

* 指定客户端容器名称,登录到客户端容器
docker exec -it sawtooth-shell-default bash

* 确认连接
curl http://rest-api:8080/blocks

## 使用锯齿和CLI命令

* 使用intkey创建和提交事务
intkey create_batch --count 10 --key-count 5 
intkey load -f batches.intkey -U http：// rest-api：8080

* 用锯齿批处理提交交易
sawtooth batch submit -f batches.intkey --url http://rest-api:8080

* 查看块列表
sawtooth block list --url http://rest-api:8080

* 查看特定块
sawtooth block show --url http://rest-api:8080 {BLOCK_ID}

## 查看全局状态

* 查看节点列表（地址）
sawtooth state list --url http://rest-api:8080

* 在地址查看数据
sawtooth state show --url http://rest-api:8080 {STATE_ADDRESS}

## 连接到REST API

* 从客户端容器
curl http://rest-api:8080/blocks

* 从主机操作系统
curl http://localhost:8080/blocks

## 连接到每个容器

* 客户端容器
docker exec -it sawtooth-shell-default bash

* 验证器容器
docker exec -it sawtooth-validator-default bash

* REST API容器
docker exec -it sawtooth-rest-api-default bash

* 设置交易处理器容器
docker exec -it sawtooth-settings-tp-default bash

* IntegerKey交易处理器容器
docker exec -it sawtooth-intkey-tp-python-default bash

* XO交易处理器容器
docker exec -it sawtooth-xo-tp-python-default bash

* 查看日志文件
docker logs {CONTAINER}

* 配置事务系列的列表
docker exec -it sawtooth-validator-default bash











