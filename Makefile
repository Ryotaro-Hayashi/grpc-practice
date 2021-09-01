protoc-doc:
	protoc --doc_out=. --doc_opt=markdown,deepthought.md deepthought.proto

protoc-go:
	protoc --go_out=module=github.com/Ryotaro-Hayashi/grpc-practice:. deepthought.proto

protoc-go-grpc:
	protoc --go-grpc_out=module=github.com/Ryotaro-Hayashi/grpc-practice:. deepthought.proto
