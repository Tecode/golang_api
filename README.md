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
# 全局安装beego工具
go get -u github.com/beego/bee/v2
# 注意设置 GO111MODULE="auto"，项目中有go.mod会提示错误
go env -w GO111MODULE=auto


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

## 部署

```shell
bee version
bee pack
tar -zxvf golang_apiv2.tar.gz
nohup ./golang_apiv2 &
rm -rf golang_apiv2.tar.gz
```

## Go Env 配置

```cgo
GO111MODULE="on"
GOARCH="amd64"
GOBIN=""
GOCACHE="/root/.cache/go-build"
GOENV="/root/.config/go/env"
GOEXE=""
GOEXPERIMENT=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOINSECURE=""
GOMODCACHE="/root/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="linux"
GOPATH="/root/go"
GOPRIVATE=""
GOPROXY="https://goproxy.cn,direct"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GOVCS=""
GOVERSION="go1.20.5"
GCCGO="gccgo"
GOAMD64="v1"
AR="ar"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD="/dev/null"
GOWORK=""
CGO_CFLAGS="-O2 -g"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-O2 -g"
CGO_FFLAGS="-O2 -g"
CGO_LDFLAGS="-O2 -g"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build903738740=/tmp/go-build -gno-record-gcc-switches"
```