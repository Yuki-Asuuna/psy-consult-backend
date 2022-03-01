# psy-consult-backend项目（心院咨询热线后端）

服务器ip&port：8.130.13.233:8000

系统信息：Alibaba Cloud Linux 3.2104 64位

系统配置：2 vCPU 2 GiB （I/O优化）ecs.t6-c1m1.large 1Mbps

## 访问示例

没有申请域名，目前只能通过ip+端口号访问

Request：

```
http://8.130.13.233:8000/ping
```

Response：

```json
{
    "code": 0,
    "message": "OK",
    "result": "pong"
}
```

## Redis

由于安全问题，Redis不对公网开放 （docker部署）

服务器ip：8.130.13.233（与服务器ip相同）

Port：6379

用户名：/

密码：/

## Chevereto（图床服务）

服务器ip：8.130.13.233（与服务器ip相同）

Port：80

使用教程：https://chevereto-free.github.io/api/#api-key

