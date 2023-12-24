from cffi import FFI
ffibuilder = FFI()

ffibuilder.set_source("go_bind_py",
    """ //The generated header for go-bind
        #include "go_bind.h"
    """,
    extra_objects=["go_bind.so"])

ffibuilder.cdef("""
    extern void* InitAuthZClient(char* host, char* bearerToken);
    """)

if __name__ == "__main__":
    ffibuilder.compile(
        verbose=True,
        target="go_bind_py.so",
        debug=True
    )
