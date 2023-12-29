
install:
	pip install pybindgen
	pip install --upgrade setuptools wheel
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/go-python/gopy@latest

build-so: install
	# WORKS with python3.10 only(as far as i tested it)
	gopy build -output=build -vm=python3 github.com/sanix-darker/gopy
	sed -i 's/bind_TEstData(/TEstData(/g' ./build/go_bind.py

run: build-so
	python3.10 ./test_go_bind.py
