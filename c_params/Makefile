install-py:
	pip install -r requirements.txt

build-so: install-py
	rm -rf go_bind.so go_bind_py.c go_bind_py.o go_bind_py.so
	@echo "> Building go_bind..."
	go build -buildmode=c-shared -o go_bind.so .
	@echo "> List exported funcs from go_bind..."
	nm -gC go_bind.so

build-py: build-so
	python3 ./go_bind_builder.py

run: build-py
	python3 ./test_go_bind.py
