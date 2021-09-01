package main

import (
	"context"
	grpc_practice "github.com/Ryotaro-Hayashi/grpc-practice"
)

// Server interfaceであるComputeServerを実装する型
type Server struct {
	// 将来protoファイルにRPCが追加されてインタフェースが拡張された際にビルドエラーになるのを防止する仕組み
	grpc_practice.UnimplementedComputeServer
}

// インタフェースが実装できていることをコンパイル時に確認
var _ grpc_practice.ComputeServer = &Server{}


func (s *Server) Boot(req *grpc_practice.BootRequest, stream grpc_practice.Compute_BootServer) error {
	panic("not implemented") // TODO: Implement
}

func (s *Server) Infer(ctx context.Context, req *grpc_practice.InferRequest) (*grpc_practice.InferResponse, error) {
	panic("not implemented") // TODO: Implement
}
