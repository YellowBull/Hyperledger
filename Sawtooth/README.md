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

* 指定客户端容器名称,登录到客户端容器 <br/>
docker exec -it sawtooth-shell-default bash <br/>

* 确认连接 <br/>
curl http://rest-api:8080/blocks <br/>

## 使用锯齿和CLI命令

* 使用intkey创建和提交事务 <br/>
intkey create_batch --count 10 --key-count 5  <br/>
intkey load -f batches.intkey -U http://rest-api:8080 <br/>

* 用锯齿批处理提交交易 <br/>
sawtooth batch submit -f batches.intkey --url http://rest-api:8080 <br/>

* 查看块列表 <br/>
sawtooth block list --url http://rest-api:8080 <br/>

* 查看特定块 <br/>
sawtooth block show --url http://rest-api:8080 {BLOCK_ID} <br/>

## 查看全局状态

* 查看节点列表（地址） <br/>
sawtooth state list --url http://rest-api:8080 <br/>

* 在地址查看数据 <br/>
sawtooth state show --url http://rest-api:8080 {STATE_ADDRESS} <br/>

## 连接到REST API

* 从客户端容器 <br/>
curl http://rest-api:8080/blocks <br/>

* 从主机操作系统 <br/>
curl http://localhost:8080/blocks <br/>

## 连接到每个容器

* 客户端容器 <br/>
docker exec -it sawtooth-shell-default bash <br/>

* 验证器容器 <br/>
docker exec -it sawtooth-validator-default bash <br/>

* REST API容器 <br/>
docker exec -it sawtooth-rest-api-default bash <br/>

* 设置交易处理器容器 <br/>
docker exec -it sawtooth-settings-tp-default bash <br/>

* IntegerKey交易处理器容器 <br/>
docker exec -it sawtooth-intkey-tp-python-default bash <br/>

* XO交易处理器容器 <br/>
docker exec -it sawtooth-xo-tp-python-default bash <br/>

* 查看日志文件 <br/>
docker logs {CONTAINER} <br/>

* 配置事务系列的列表 <br/>
docker exec -it sawtooth-validator-default bash <br/>











