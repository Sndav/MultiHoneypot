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