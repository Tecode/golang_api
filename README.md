
# API接口

## 安装依赖项

```bash
go get -u github.com/beego/bee
go get -u github.com/astaxie/beego
go get -u github.com/dgrijalva/jwt-go
go get github.com/astaxie/beego/orm
go get -u github.com/go-sql-driver/mysql
go get github.com/nfnt/resize
```

## 运行


```bash
bee run -downdoc=true -gendoc=true
```

## 版本信息

```bash
______                                           ______
| ___ \                                          | ___ \
| |_/ /  ___   ___                               | |_/ /  ___   ___
| ___ \ / _ \ / _ \                              | ___ \ / _ \ / _ \
| |_/ /|  __/|  __/                              | |_/ /|  __/|  __/
\____/  \___| \___| v1.10.0                      \____/  \___| \___| v1.10.0

├── Beego     : 1.11.2                           ├── Beego     : 1.11.2
├── GoVersion : go1.11.2                         ├── GoVersion : go1.10.4
├── GOOS      : windows                          ├── GOOS      : linux
├── GOARCH    : amd64                            ├── GOARCH    : amd64
├── NumCPU    : 12                               ├── NumCPU    : 6
├── GOPATH    : C:\Users\Administrator\go        ├── GOPATH    : /home/xm/go/src/golang_api:/home/xm/go
├── GOROOT    : D:/Golang                        ├── GOROOT    : /usr/lib/go-1.10
├── Compiler  : gc                               ├── Compiler  : gc
└── Date      : Sunday, 19 May 2019              └── Date      : Monday, 20 May 2019
```

## 发布

`Linux`

```bash
runmode = prod

bee pack
nohup ./golang_api &
```

> Server on:[ http://127.0.0.1:8080/swagger]( http://127.0.0.1:8080/swagger).

## 错误信息

