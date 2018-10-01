# 語言/Web框架 效能比較

幾個比較popular的框架大致說明  
https://blog.csdn.net/csdnnews/article/details/80750455

目前採用最普遍使用的Gin  
https://github.com/gin-gonic/gin

## Build Image
編譯go的binary主要需要指定幾個參數  
GOOS跟GOARCH  
指定目標的OS跟平台架構  

另外有debug跟release的區別  
如果使用go build -gcflags "-l -N"出來會是debug  
如果直接go build出來的就會是release  
```bash
.\build_windows.bat
```
```bash
.\build_linux.bat
```
```bash
.\build_darwin.bat
```
bat檔中指定了上述的參數  
另外會將編譯好的binary copy到archive資料夾中
## Deploy
直接將編譯好之binary置於目標主機上即可執行  
centOS必須執行以下command 調整權限
```bash
chmod u+x filename
```
將filename代換為binary實際的檔名即可  