# 語言/Web框架 效能比較

本專案以測試 Web API Server 的效能為主軸
目前正在進行的項目

1. Node.js - Koa + HAProxy
2. Golang - Gin
3. Python - Flask + waitress

## 評測方式

實作 2 個 API

1. Get Redis Value by Key
2. Set Redis Key & Value

使用 wrk2 進行壓力測試
測試命令如下

```bash
wrk -c50 -t4 -RXXXXX 'URL'
```

## 主要評比項目

1. 效能：使用壓測軟體測試每秒可以承受的最大 Request 數量
2. 開發速度：是否有足夠且高品質的第三方 Module？語言+框架本身是否便於使用？
3. 部署：封裝檔案大小；部署時是否需要安裝其他 run-time？是否能將 Module 完全打包成數個檔案？

## 效能測試 1st

| 項目 | Request / Second | Avg Latency |
|:- |:- |:- |
| Go + Gin | 20000 | 16.63 ms |
|          | 30000 | 5.42 s |
| Python + Flask | 2000 | 6.32 s |
| Node.js + Koa | 10000 | 3.87 ms |
|               | 15000 | 2.03 s |
|               | 20000 | 6.36 s |