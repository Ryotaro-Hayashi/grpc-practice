protoc-doc:
	protoc -I. --doc_out=. --doc_opt=markdown,deepthought.md deepthought.proto
