
# Golang API(Beego)

## Creating the Beego API aplication

```bash
go get -u github.com/beego/bee
go get -u github.com/astaxie/beego
go get -u github.com/dgrijalva/jwt-go
```

## Run


```bash
bee run -downdoc=true -gendoc=true
```

## Run product

`Linux`

```bash
runmode = prod

bee pack
nohup ./beeapi &
```

> Server on:[ http://127.0.0.1:8080/swagger]( http://127.0.0.1:8080/swagger).

## Error

> Change file name 'golang_api' to 'beeapi'.

```bash
main.go:4:2: cannot find package "beeapi/routers" in any of:
/usr/local/go/src/beeapi/routers (from $GOROOT)
/Users/ming/go/src/beeapi/routers (from $GOPATH)
```
