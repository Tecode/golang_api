
# API接口

## 安装依赖项

```bash
go get -u github.com/beego/bee
go get -u github.com/astaxie/beego
go get -u github.com/dgrijalva/jwt-go
go get github.com/astaxie/beego/orm
go get -u github.com/go-sql-driver/mysql
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
nohup ./golang_api &
```

> Server on:[ http://127.0.0.1:8080/swagger]( http://127.0.0.1:8080/swagger).

## 错误信息

