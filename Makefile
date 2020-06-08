START=$(pwd)

protoc:
	protoc -I proto/ \
		--go_out=plugins=grpc:proto \
		proto/*.proto
