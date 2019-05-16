
# API接口

## 安装依赖项

```bash
go get -u github.com/beego/bee
go get -u github.com/astaxie/beego
go get -u github.com/dgrijalva/jwt-go
```

## 运行


```bash
bee run -downdoc=true -gendoc=true
```

## 发布

`Linux`

```bash
runmode = prod

bee pack
nohup ./beeapi &
```

> Server on:[ http://127.0.0.1:8080/swagger]( http://127.0.0.1:8080/swagger).

## 错误信息

> 修改文件名'golang_api' 为 'beeapi'

```bash
main.go:4:2: cannot find package "beeapi/routers" in any of:
/usr/local/go/src/beeapi/routers (from $GOROOT)
/Users/ming/go/src/beeapi/routers (from $GOPATH)
```
