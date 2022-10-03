
# go-psmp(golang public-seat-manager-platform)

一款go 语言写的轻量级后台管理监控平台, 集成了 redis mysql cron email等基础封装
## 功能

### 1. 短链服务
### 2. 邮件告警服务
### 3. 钉钉告警(开发中...)
### 4.[Agent项目](https://github.com/fengzijk/psmp-agent/tree/master)
    1. cpu监控 (1.0.0 已经完成)
    2. 磁盘监控(1.0.1 计划中)
    3. web服务监控(1.0.0 已经完成)



## 使用示例

### 1. application.yml 配置文件
```yml
server:
  port: 8080
mysql:
  username: root
  password: 123456
  host: 192.168.2.11
  port: 3306
  database: go_psmp

redis:
  address: 192.168.2.11:6379
  database: 0
  password: 123456
  
short:
  prefix: http://localhost:8080
  length: 8

email:
  user: 
  password: 
  host : smtp.163.com
  port: 465
  toUser:
 
# 定时任务
task-cron:
  # 邮件告警 每5秒执行一次
  send-alarm-email: "*/5 * * * * *"
  # agent 心跳检查 五秒一次
  agent-heartbeat-alarm: "*/5 * * * * *"
```



## 打包命令

```shell
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64

go build -o go-psmp main.go

```