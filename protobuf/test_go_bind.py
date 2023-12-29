from ctypes import c_char_p, CDLL
import pathlib
from go_bind_pb2 import Pack as PackMessage

# Load the shared library
libpath = pathlib.Path().absolute() / "go_bind.so"
lib = CDLL(str(libpath))

get_data_info = lib.GetDataInfo
get_data_info.argtypes = [c_char_p]
get_data_info.restype = c_char_p

# pack stuffs
pk = PackMessage(data="ok doki", len=101)
data = pk.SerializeToString()

# call the exported method
result = get_data_info(data)

# unpack stuffs...
pko = PackMessage()
pko.ParseFromString(result)
print(f">> {pko.data=}")
print(f">> {pko.len=}")
