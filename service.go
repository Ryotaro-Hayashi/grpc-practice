package grpc_practice

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	// validation
	switch req.Query {
	case "Life", "Universe", "Everything":
	default:
		// gRPC は共通で使われるエラーコードを定めているので、基本は定義済みのコードを使う
		// https://grpc.github.io/grpc/core/md_doc_statuscodes.html
		return nil, status.Error(codes.InvalidArgument, "Contemplate your query")
	}

	// クライアントがタイムアウトを指定しているかチェック
	deadline, ok := ctx.Deadline()
	// 指定されていない、もしくは十分な時間があれば回答
	if !ok || time.Until(deadline) > 750*time.Millisecond {
		time.Sleep(750 * time.Millisecond)
		return &pb.InferResponse{
			Answer:      42,
			Description: []string{"I checked it"},
		}, nil
	}

	// 時間が足りなければ DEADLINE_EXCEEDED (code 4) エラーを返す
	return nil, status.Error(codes.DeadlineExceeded, "It would take longer")
}
