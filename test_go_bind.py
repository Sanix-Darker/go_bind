from go_bind_py import lib
from cffi import FFI
import ctypes

ffi = FFI()

def py_to_c_str(_str: str) -> bytes:
    return _str.encode()

# host "spicedb.toucan.local:50051"
# bearer_token "toucan-local-env-dev-spicedb"
# authzed_Client* client = InitClient(host, token);

sta = lib.InitAuthZClient(
   py_to_c_str("spicedb.toucan.local:50051"),
   py_to_c_str("toucan-local-env-dev-spicedb")
)

# mem_addr = ctypes.addressof(0x55850c7a4570)
# cdhex = int(str(sta).split(' ')[-1].split('>')[0], 16)
# cdata_pointer = ffi.cast("void*", cdhex)

__import__('ipdb').set_trace()
print(f">> authZ initialized {sta=}")
print(f">> go_bind_py :: export {dir(lib)}")
