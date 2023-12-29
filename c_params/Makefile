BASE=$(shell basename "$(shell pwd)")

build-so:
	rm -rf ${BASE}.so go_bind_py.c go_bind_py.o go_bind_py.so
	@echo "> Building ${BASE}..."
	go build -buildmode=c-shared -o ${BASE}.so .
	@echo "> List exported funcs from ${BASE}..."
	nm -gC ${BASE}.so

build-py: build-so
	python3 ./go_bind_builder.py

run: build-py
	python3 ./test_go_bind.py
