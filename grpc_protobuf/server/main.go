package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	// pb "rpc/grpc_protobuf/proto/hello" // protoc 生成 protobuf go 文件到GOROOT
	// pb "rpc/grpc_protobuf/client/protobuf/hello" // protoc 生成 protobuf go 文件到当前项目
	pb "github.com/leeprince/protobuf/grpc/hello" // protoc 生成 protobuf go 文件到远程仓库（github）【推荐】
)

const LISTENADDRESS = "127.0.0.1:12345"

type GreeterServer struct {
	pb.UnsafeGreeterServer
}

func main() {
	l, e := net.Listen("tcp", LISTENADDRESS)
	log.Println("Listen at ", l.Addr())
	if e != nil {
		log.Fatal("Listen error:", e)
	}
	defer l.Close()
	
	s := grpc.NewServer()
	
	/** go 服务需要注册是因为 go 是强语言，没办法像 php 一样字符串加双括号就可以作为方法「字符串()」 */
	// 注册服务
	pb.RegisterGreeterServer(s, &GreeterServer{})
	
	if err := s.Serve(l); err != nil {
		fmt.Printf("Failed to server: %v", err)
	}
	
}

func (s *GreeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "hello " + in.GetName()}, nil
}

