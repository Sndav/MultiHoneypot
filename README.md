# Multi-Honey

一个Go语言编写的可扩展的蜜罐框架

## 架构

```           
端口监听  <----------->  代理 <----------> 后端
```
在代理中我们可以记录所有攻击流量

## 运行
配置 `config/config.ini` 以及 `./cmd/backend/main.go`
```shell script
go run ./cmd/backend/main.go
```


## 感谢
1. `https://github.com/src-d/go-mysql-server`
2. `https://github.com/40t/go-sniffer`