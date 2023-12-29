install-py: ## install python libs/deps
	pip install -r requirements.txt

install-go: ## install go libs/deps
	go install github.com/golang/protobuf/protoc-gen-go@latest

proto: install-py install-go ## build the .proto
	protoc --go_out=. go_bind.proto
	protoc --python_out=. go_bind.proto
	python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. go_bind.proto

build-so: proto ## build the .so shared lib
	rm -rf go_bind.so go_bind_py.c go_bind_py.o go_bind_py.so
	@echo "> Building go_bind..."
	sed -i 's/package __/package main/g' ./go_bind.pb.go
	go build -buildmode=c-shared -o go_bind.so .
	@echo "> List exported funcs from go_bind..."
	nm -gC go_bind.so

run: build-so ## run the test python code
	python3 ./test_go_bind.py
