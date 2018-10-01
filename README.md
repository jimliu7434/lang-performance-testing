# 語言/Web框架 效能比較 - Python

## Build Image

```bash
docker build -t py-test:1.0.0 .
```

## Deploy

```bash
docker run -d -p 5000:5000 py-test:1.0.0
```