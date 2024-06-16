NAME = goblog
UI_NAME = ui
STORAGE_API_NAME = storage-api

PROTO_DIR = ./pkg/proto/
BUILD_DIR = ./bin/

UI_PATH = ${BUILD_DIR}${UI_NAME}
STORAGE_API_PATH = ${BUILD_DIR}${STORAGE_API_NAME}

GENERATED_FILES = $(addprefix ${PROTO_DIR}, \
				  article.pb.go \
				  article_grpc.pb.go)

all: build

compile-proto: ${GENERATED_FILES}

${PROTO_DIR}%.pb.go: ${PROTO_DIR}%.proto
	@echo "Compiling $<"
	@protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$<

build: build-dir build-ui build-storage-api

build-dir:
	@mkdir -p ${BUILD_DIR}

build-ui: compile-proto
	@go build -o ${UI_PATH} ./cmd/${UI_NAME}/main.go
	@cp -r ./cmd/ui/static ${BUILD_DIR}

run-ui: build-ui
	@./${UI_PATH}

build-storage-api: compile-proto
	@go build -o ${STORAGE_API_PATH} ./cmd/${STORAGE_API_NAME}/main.go

run-storage-api: build-storage-api
	@./${STORAGE_API_PATH}
