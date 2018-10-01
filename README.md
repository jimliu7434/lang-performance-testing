# 語言/Web框架 效能比較 - Node.js

Node.js 版本因為語言特性關係
所以搭配 HAProxy & Docker-Compose 進行部署

本次使用 **Node.js \* 4 + HAProxy \* 1** 來進行評測

## Build Image

```bash
# build app image
docker build -t nodekoa-redis-test:1.0.0 .
# build haproxy image
cd haproxy && docker build -t haproxy-dev:1.8 . && cd ..
```

## Deploy

```bash
docker-compose up -d --scale koa=4
```