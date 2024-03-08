# Change this to your own Go module
ROOT_FOLDER := github.com
GO_MODULE := github.com/espitman/hexaGRPC/protos


.PHONY: tidy
tidy:
	go mod tidy


.PHONY: clean
clean:
ifeq ($(OS), Windows_NT)
	if exist "protogen" rd /s /q protogen	
else
	rm -fR ./protogen
endif


.PHONY: protoc
# (un)comment or add folder that contains proto as required
protoc:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
	--go-grpc_opt=module=${GO_MODULE} --go-grpc_out=. \
	-I . \
  -I ./validate/ \
  --validate_out="lang=go:." \
	./proto/greet/*.proto \

.PHONY: build
build: clean protoc move-validate tidy


.PHONY: run
run:
	go run main.go


.PHONY: execute
execute: build run


.PHONY: move-validate
move-validate: 
	mv ${GO_MODULE}/protogen/greet/* ./protogen/greet
	rm -fR  ./${ROOT_FOLDER}
