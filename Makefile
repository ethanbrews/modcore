.PHONY: proto build build-core build-cli clean deps test sync

# Detect OS
ifeq ($(OS),Windows_NT)
    MKDIR = if not exist "$1" mkdir "$1"
    RM = if exist "$1" rmdir /s /q "$1"
    EXE_EXT = .exe
else
    MKDIR = mkdir -p $1
    RM = rm -rf $1
    EXE_EXT =
endif

# Directories
BIN_DIR := ./bin
PROTO_GEN_DIR := ./proto/gen

# Generate protobuf files
proto:
	$(call MKDIR,$(PROTO_GEN_DIR))
	protoc --go_out=$(PROTO_GEN_DIR) --go_opt=paths=source_relative --go-grpc_out=$(PROTO_GEN_DIR) --go-grpc_opt=paths=source_relative --proto_path=./proto ./proto/modcore.proto

# Build executables
build-cli: proto
	$(call MKDIR,$(BIN_DIR))
	cd cli && go build -o ../$(BIN_DIR)/cli$(EXE_EXT)

build-core: proto
	$(call MKDIR,$(BIN_DIR))
	cd core && go build -o ../$(BIN_DIR)/core$(EXE_EXT)

build: build-cli build-core

sync:
	go work sync

clean:
	$(call RM,$(PROTO_GEN_DIR))
	$(call RM,$(BIN_DIR))