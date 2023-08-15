## beego version & go version

要运行项目不要配置`GOPATH`,需要`GO`版本大于`1.16`，这样才支持使用`go mod` 管理依赖

## 新增配置文件conf/app.conf

```text
appname = golang_apiv2
httpport = 8080
runmode = prod
CopyRequestBody = true

# 是否需要创建数据表
createtable = false
mysqlurl = localhost:3306
mysqlaccount = root
mysqlpassword = 123456
msqldatabase = golang_apiv2

# 发送邮件的邮箱配置
emailhost = smtp.qq.com #邮件发送的地址
emailport = 587 #端口号
emailaccount= 邮箱账号
emailpassword = 邮箱密码（开启POP3/IMAP/SMTP/Exchange/CardDAV/CalDAV服务后的授权码）
```
## Nginx配置

```sh
server {
  listen [::]:443 ssl ipv6only=on http2;
  listen       443 ssl http2;
  server_name  soscoon.com www.soscoon.com;

  ssl_certificate       /home/xm/certificate/soscoon.com_nginx/soscoon.com_bundle.pem;
  ssl_certificate_key   /home/xm/certificate/soscoon.com_nginx/soscoon.com.key;


  # ssl验证相关配置
  ssl_session_timeout  5m;    #缓存有效期
  ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE:ECDH:AES:HIGH:!NULL:!aNULL:!MD5:!ADH:!RC4;    
  #加密算法
  ssl_protocols TLSv1 TLSv1.1 TLSv1.2;    #安全链接可选的加密协议
  ssl_prefer_server_ciphers on;   #使用服务器端的首选算法

  #开启gzip
  gzip  on;
  #低于1kb的资源不压缩
  gzip_min_length 1k;
  #压缩级别1-9，越大压缩率越高，同时消耗cpu资源也越多，建议设置在5左右。
  gzip_comp_level 5;
  #需要压缩哪些响应类型的资源，多个空格隔开。不建议压缩图片.
  gzip_types text/plain application/javascript application/x-javascript text/javascript text/xml text/css;
  #配置禁用gzip条件，支持正则。此处表示ie6及以下不启用gzip（因为ie低版本不支持）
  gzip_disable "MSIE [1-6]\.";
  #是否添加“Vary: Accept-Encoding”响应头
  gzip_vary on;

  charset utf-8;
    #access_log  /var/log/nginx/admin-nginx-access.log  main;
    access_log off;

    #charset koi8-r;
    #access_log  /var/log/nginx/host.access.log  main;

    location /ws {
        proxy_pass http://localhost:8080/ws;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "Upgrade";
    }

    location / {
        proxy_buffers 8 1024k;
        proxy_buffer_size 1024k;
        proxy_pass      http://localhost:8080;
    }

    # redirect server error pages to the static page /50x.html
    #
    #error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/html;
    }

    # proxy the PHP scripts to Apache listening on 127.0.0.1:80
    #
    #location ~ \.php$ {
    #    proxy_pass   http://127.0.0.1;
    #}
}

# 将http请求重定向到https
server{
  listen 80 default_server;
  server_name _;
  return 301 https://$host$request_uri;
}

```

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

## 使用到的第三方插件

```sh
go get gopkg.in/gomail.v2
go get github.com/patrickmn/go-cache
go get github.com/beego/beego/v2/core/validation
go get github.com/gorilla/websocket
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

### 自动化脚本
```shell
#!/bin/bash

export BUILD_ID=dontKillMe
export PATH=$PATH:/usr/local/go/bin
# 获取程序的进程ID（PID）
PID=$(pgrep golang_apiv2)

if [ -n "$PID" ]; then
  # 存在进程ID，说明程序正在运行
  echo "程序正在运行，进程ID为: $PID"
  
  # 终止程序
  kill $PID
  
  echo "程序已终止"
else
  echo "程序未在后台运行"
fi
# 从这里开始记录错误，如果解压或者运行失败会捕获到做自动化程序会自动终止
# set +e 表示不记录错误
set -e
/usr/local/go/bin/bee pack
tar -zxvf golang_apiv2.tar.gz
nohup ./golang_apiv2 > /home/xm/logs/golang_api2.txt 2>&1 &
sleep 10
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