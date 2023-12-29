## GO-BIND-PROTOBUF

Using protobuf, set a dataStruct you can manipulate between golang code and python code.

## REQUIREMENTS

- make
- golang (>1.18 recommended)
- python (>=3.11 recommended)
- cffi (from pypi)

## HOW TO BUILD AND RUN

```console
$ make run

#
# python3 ./test_go_bind.py
# Data: ok doki, Length: 101
# >> pko.data="Hi from GO, your input was 'ok doki'"
# >> pko.len=36
```
