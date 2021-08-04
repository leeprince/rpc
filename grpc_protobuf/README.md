# 安装 protobuf
- 到 https://github.com/protocolbuffers/protobuf/releases/ 下载相应系统的 protoc 文件
> 本人系统是：macos, 下载的是：protoc-3.17.3-osx-x86_64.zip

- 解压后，将 protoc（protoc-3.17.3-osx-x86_64/bin/protoc） 复制到 $GOPATH/bin 路径下

# 安装 protobuf、grpc 的 go 包
```
go get google.golang.org/protobuf
go get google.golang.org/grpc
```

# 安装命令到 $GOPATH/bin
```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

> 注意：go get google.golang.org/grpc可能下载失败;用下面这个方式代替
```
git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc 
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto cd $GOPATH/src/
go install google.golang.org/grpc
```


# 一、编写 proto 文件 


# 二、根据 proto 文件生成 protobuf 的 go 文件
## （一）【推荐】protoc 生成 protobuf go 文件到远程仓库（github）

protoc 生成的 grpc protobuf go 文件外部 github 项目(单独存放 proto 文件也保存在外部项目中)。 服务端和客户端很好共用同一份 protobuf go 文件

> protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/*.proto
### 1. 生成：greeter.pb.go

```
protoc -I=. --go_out=../../grpc/ ./greeter.proto
```

### 2. 生成：greeter.pb.go、greeter_grpc.pb.go
> 方式一
1. cd proto/greeter # 进入greeter.proto所在目录
2. 执行：
    ```
    protoc -I=. --go_out=../../grpc --go-grpc_out=../../grpc/ ./greeter.proto
    ```
3. 在 grpc/ 目录下自动创建greeter 目录并生成的文件：greeter.pb.go、greeter_grpc.pb.go

> 方式二
1. cd . # 在当前项目根目录下 protobuf/
2. 执行：
```
protoc -I=proto/greeter --go_out=./grpc --go-grpc_out=./grpc/ proto/greeter/greeter.proto
```
3. 在 grpc/ 目录下自动创建 greeter 目录并生成的文件：greeter.pb.go、greeter_grpc.pb.go


## （二）protoc 生成 protobuf go 文件到 GOROOT
```
protoc --go_out=$GOROOT/src/rpc/greeter/ --go_opt=paths=source_relative \
    --go-grpc_out=$GOROOT/src/rpc/greeter/ --go-grpc_opt=paths=source_relative \
    greeter.proto
```



## （三）protoc 生成 protobuf go 文件保存到当前项目
> 服务端和客户端需同时存在同一份 protobuf go 文件
```server
protoc --go_out=../../server/protobuf/hello --go_opt=paths=source_relative \
    --go-grpc_out=../../server/protobuf/hello --go-grpc_opt=paths=source_relative \
    hello.proto
```
```client
protoc --go_out=../../client/protobuf/hello --go_opt=paths=source_relative \
    --go-grpc_out=../../client/protobuf/hello --go-grpc_opt=paths=source_relative \
    hello.proto
```
