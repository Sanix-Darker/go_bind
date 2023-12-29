## GO-BIND

Manipulate objects and use functions from a golang codebase to python, using
C-shared lib with an .so.

There is multiple methods i found for this:
- Using C params as args (see the code inside ./c_params/).
- Using protobuf (by defining some .proto files that later generate appropriate
  code, see the code inside ./protobuf/.
- Using Gopy (the most stable solution so far), even go data-struct are
  transposed into an "exported way" (some glitch but nothing unfixable, check
  (./gopy/)).

All solutions had to load and define typing for methods comming from the
shared libs(except for the gopy solution).

Note: some .c/.h/.o files are going to be generated, don't worry, that's normal
for the communication to work properly between the two.
