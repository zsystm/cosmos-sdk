.PHONY: proto-gen update-proto-pin

proto-gen:
	go install ./cmd/protoc-gen-go-cosmos-orm
	(cd internal; buf generate)

update-proto-pin:
	buf build ../proto --path ../proto/cosmos/orm --exclude-imports -o proto_image.bin.gz