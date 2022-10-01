# go-psmp(golang public-seat-manager-platform)

一款go 语言写的轻量级后台管理平台, 集成了 redis mysql cron email等基础封装
## 功能

### 1. 短链服务
### 2. 邮件告警服务
### 3. 钉钉告警(开发中...)

## 使用示例

### 1. application.yml 配置文件
~~~yml
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
  user: xxx@163.com
  password: 1111
  host : smtp.163.com:25
~~~

