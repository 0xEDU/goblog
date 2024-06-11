NAME = goblog

PROTO_DIR = ./pkg/proto/

UI = ui
STORAGE_API = storage-api

GENERATED_FILES = $(addprefix ${PROTO_DIR}, \
				  article.pb.go \
				  article_grpc.pb.go)

all: build

compile-proto: ${GENERATED_FILES}

${PROTO_DIR}%.pb.go: ${PROTO_DIR}%.proto
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$<

build: build-ui build-storage-api

build-ui:
	go build -o ${UI} ./cmd/${UI}/main.go

build-storage-api:
	go build -o ${STORAGE_API} ./cmd/${STORAGE_API}/main.go
