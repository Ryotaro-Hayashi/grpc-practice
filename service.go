package grpc_practice

import (
	"context"
	"time"

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
	// 無限ループ
	for {
		select {
		// クライアントがリクエストをキャンセルしたら終わり
		case <-stream.Context().Done():
			return nil
		// 1秒待機してデータを送信
		case <-time.After(1 * time.Second):
		}

		if err := stream.Send(&pb.BootResponse{
			Message: "I THINK THEREFORE I AM.",
		}); err != nil {
			return err
		}
	}
}

func (s *Server) Infer(ctx context.Context, req *pb.InferRequest) (*pb.InferResponse, error) {
	panic("not implemented") // TODO: Implement
}
