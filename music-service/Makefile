gen-api:
	protoc  --proto_path=./api \
            --proto_path=./third_party \
            --validate_out=paths=source_relative,lang=go:./api \
			--go_out=paths=source_relative:./api \
         	--go-grpc_out=paths=source_relative:./api \
       		 api/v1/proto/music-service.proto