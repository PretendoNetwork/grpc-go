build: pull account miiverse friends go_mod

pull:
	git submodule init
	git submodule update --remote --merge

account:
	mkdir -p account

	protoc \
	--proto_path=../grpc-protobufs/account \
	--go_out=./account \
	--go_opt=paths=source_relative \
	--go-grpc_out=./account \
	--go-grpc_opt=paths=source_relative \
	../grpc-protobufs/account/*.proto

miiverse:
	mkdir -p miiverse

	protoc \
	--proto_path=../grpc-protobufs/miiverse \
	--go_out=./miiverse \
	--go_opt=paths=source_relative \
	--go-grpc_out=./miiverse \
	--go-grpc_opt=paths=source_relative \
	../grpc-protobufs/miiverse/*.proto

friends:
	mkdir -p friends

	protoc \
	--proto_path=../grpc-protobufs/friends \
	--go_out=./friends \
	--go_opt=paths=source_relative \
	--go-grpc_out=./friends \
	--go-grpc_opt=paths=source_relative \
	../grpc-protobufs/friends/*.proto

go_mod:
	test -f go.mod || (echo "go module not found. Creating.." && go mod init github.com/PretendoNetwork/grpc-go)
	go mod tidy

.PHONY: pull account miiverse friends go_mod