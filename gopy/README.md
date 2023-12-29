## GO-BIND-GOPY

Using ![gopy](https://github.com/go-python/gopy), this tool with python3.10
(the 3.11 gave me headaches with some weird incompatibilities on
libpython-3.11.a).

## REQUIREMENTS

- make
- golang (>1.18 recommended)
- python (3.10)
- gopy (will be installed on make install)


## HOW TO BUILD AND RUN

```console
$ make run

#
# python3.10 ./test_go_bind.py
#  >> ex.Data='test' :: 0
#  >> pp.Data="Hello there from GO code 'test';" :: 32
```
