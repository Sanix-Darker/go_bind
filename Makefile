BASE=$(shell basename "$(shell pwd)")

build-so:
	rm -rf ${BASE}.so go_bind_py.c go_bind_py.o go_bind_py.so
	@echo "> Building ${BASE}..."
	/usr/local/go1.20/go/bin/go build -buildmode=c-shared -o ${BASE}.so .
	nm -gC ${BASE}.so

build-py:
	python3.11 ./go_bind_builder.py

run-py:
	python3.11 ./test-go-bind.py
