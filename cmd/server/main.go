package main

import (
	"fmt"
	"net"
	"os"

	grpc_practice "github.com/Ryotaro-Hayashi/grpc-practice"

	"github.com/Ryotaro-Hayashi/grpc-practice/pb"

	"google.golang.org/grpc"
)

const portNumber = 13333

func main() {
	serv := grpc.NewServer()

	// 実装した Server を登録
	pb.RegisterComputeServer(serv, &grpc_practice.Server{})

	// 待ち受けソケットを作成
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", portNumber))
	if err != nil {
		fmt.Println("failed to listen:", err)
		os.Exit(1)
	}

	// gRPC サーバーでリクエストの受付を開始
	// l は Close されてから戻るので、main 関数での Close は不要
	if err := serv.Serve(l); err != nil {
		fmt.Println("failed to serve:", err)
		os.Exit(1)
	}
}
