codegen:
	go install ./cmd/protoc-gen-go-cosmos-orm
	(cd internal; buf generate)
	(cd client/graphql/internal; buf generate)
