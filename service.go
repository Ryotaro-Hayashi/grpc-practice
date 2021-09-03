package grpc_practice

import (
	"context"

	"github.com/Ryotaro-Hayashi/grpc-practice/pb"
)

// Server interfaceであるComputeServerを実装する型
type Server struct {
	// 将来protoファイルにRPCが追加されてインタフェースが拡張された際にビルドエラーになるのを防止する仕組み
	pb.UnimplementedComputeServer
}

// インタフェースが実装できていることをコンパイル時に確認
var _ pb.ComputeServer = &Server{}

func (s *Server) Boot(req *pb.BootRequest, stream pb.Compute_BootServer) error {
	panic("not implemented") // TODO: Implement
}

func (s *Server) Infer(ctx context.Context, req *pb.InferRequest) (*pb.InferResponse, error) {
	panic("not implemented") // TODO: Implement
}
