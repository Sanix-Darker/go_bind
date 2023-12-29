import pathlib
from cffi import FFI
ffibuilder = FFI()

libpath = pathlib.Path().absolute() / "go_bind.so"
ffibuilder.set_source("go_bind_py",
    """ //passed to the real C compiler
        #include "go_bind.h"
    """,
    extra_objects=[libpath])

ffibuilder.cdef("""
    typedef struct {
        int xxx;
        char yyy[100];
    } GoBindStruct;

    extern int PrintThis(int a, char* b);
    """)

if __name__ == "__main__":
    ffibuilder.compile(
        verbose=True,
        target="go_bind_py.so",
        debug=True
    )
