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

# 编写 proto 文件 


# 根据 proto 文件生成 protobuf 的 go 文件
## 1. cd 到 *.proto 文件所在目录

## 2. 执行如下命令 
### 2.1 protoc 生成 protobuf go 文件格式
```
protoc --go_out={保存*.pb.go文件路径} --go_opt=paths=source_relative \
    --go-grpc_out={保存*_grpc.pb.go文件路径} --go-grpc_opt=paths=source_relative \
    {proto 文件}
```

#### 2.1.1 protoc 生成 protobuf go 文件保存到当前项目
protoc 生成的 grpc protobuf go 文件保存到当前项目（同一个 go.mod 项目）中
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

#### 2.1.2 protoc 生成 protobuf go 文件到GOROOT
protoc 生成的 grpc protobuf go 文件保存到 GOROOT
> 服务端和客户端很好共用同一份 protobuf go 文件
```
protoc --go_out=$GOROOT/src/rpc/grpc_protobuf/proto/hello/ --go_opt=paths=source_relative \
    --go-grpc_out=$GOROOT/src/rpc/grpc_protobuf/proto/hello/ --go-grpc_opt=paths=source_relative \
    hello.proto
```

#### 2.1.3 protoc 生成 protobuf go 文件到远程仓库（github）【推荐】
protoc 生成的 grpc protobuf go 文件外部 github 项目(单独存放 proto 文件也保存在外部项目中)

```

```
