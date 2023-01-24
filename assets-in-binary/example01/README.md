# Building a single binary containing templates

This is a complete example to create a single binary with the
[gin-gonic/gin][gin] Web Server with HTML templates.

[gin]: https://github.com/gin-gonic/gin

## How to use

### Prepare Packages

```sh
go get github.com/gin-gonic/gin
go get github.com/jessevdk/go-assets-builder
go install github.com/jessevdk/go-assets-builder@latest
```

### Generate assets.go

```sh
go-assets-builder html -o assets.go
```

### Build the server

```sh
go build -o assets-in-binary
```

### Run

```sh
./assets-in-binary
```

### 参考文档

- [Go实战]gin整合[go-assets][vfsgen][docker]完整打包方案：https://blog.csdn.net/aaaadong/article/details/118155836

