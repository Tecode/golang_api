## beego version & go version

要运行项目不要配置`GOPATH`,需要`GO`版本大于`1.16`，这样才支持使用`go mod` 管理依赖

```bash
| ___ \
| |_/ /  ___   ___
| ___ \ / _ \ / _ \
| |_/ /|  __/|  __/
\____/  \___| \___| v2.0.4

├── GoVersion : go1.20.5
├── GOOS      : windows
├── GOARCH    : amd64
├── NumCPU    : 12
├── GOPATH    : C:\Users\Administrator\go
├── GOROOT    : D:/Go
├── Compiler  : gc
└── Date      : Friday, 16 Jun 2023

go version go1.20.5 windows/amd64
```

## 安装依赖 & [运行](http://localhost:8080/)

```bash
# 安装bee工具 https://beego.gocn.vip/beego/zh/v2.0.x/bee/#bee-%E5%B7%A5%E5%85%B7%E7%9A%84%E5%AE%89%E8%A3%85
#http://localhost:8080/swagger/
#http://localhost:8080

go get golang_apiv2

# 包含swagger文档
bee generate docs
bee run -gendoc=true -downdoc=true

```

## 快捷命令

```bash
# 快速创建一个controller
bee generate controller haouxuan

# 快速创建一个model
bee generate model modelName -fields="name:string,age:int"

# 使用了命名路由需要执行这个命令重新生成commentRouter.go文件
bee generate routers
```