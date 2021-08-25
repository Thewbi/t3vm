# t3vm
TADS 3 Virtual Machine. TADS 3 is the 3rd version of the Text Adventure Development System.


## Running

```
go run cmd/t3vm/main.go
```



## Compile a helloworld application

http://www.tads.org/t3doc/doc/sysman/hello.htm

1. Install FrobTADS
Download FrobTADS from here: http://www.tads.org/tads3.htm#
Install it
Test it: In the console the command t3make should be available!

2. Compile a hello world app
Create a file called helloworld.t

```
#include "tads.h"

main(args)
{
  "Hello from TADS 3!!!\b";
}
```

Compile it using t3make:

```
t3make helloworld.t
```

You get a file called helloworld.t3

3. Run that file in the t3vm
In main.go, change the file that is loaded to the helloworld.t3 created in the preceding step.
Run the app.

```
go run cmd/t3vm/main.go
```


## Data holders

DataHolders are used in the SYMD block (Symbolic Names Block)

DataHolders are described here: http://www.tads.org/t3spec/bincode.htm

DataHolders are stored using bytes

byte 0 - type code
byte 1 - data (meaning is determined individually for each type)
byte 2 - data (meaning is determined individually for each type)
byte 3 - arbitrary
byte 4 - arbitrary

Type Codes:

VM_NIL          1 - Boolean: False or NULL-Pointer or nil
VM_TRUE         2 - Boolean: True
VM_STACK        3 -	Reserved for implementation use for storing native machine pointers to stack frames (see note below)	none
VM_CODEPTR      4 - Reserved for implementation use for storing native machine pointers to code (see note below)	none
VM_OBJ          5 - object reference as a 32-bit unsigned object ID number	UINT4
VM_PROP	        6 - property ID as a 16-bit unsigned number	UINT2
VM_INT          7 - integer as a 32-bit signed number	INT4
VM_SSTRING      8 - single-quoted string; 32-bit unsigned constant pool offset	UINT4
VM_DSTRING      9 - double-quoted string; 32-bit unsigned constant pool offset	UINT4
VM_LIST	       10 - list constant; 32-bit unsigned constant pool offset	UINT4
VM_CODEOFS     11 - code offset; 32-bit unsigned code pool offset	UINT4
VM_FUNCPTR     12 - function pointer; 32-bit unsigned code pool offset	UINT4
VM_EMPTY       13 - no value (this is useful in some cases to represent an explicitly unused data slot, such as a slot that has never been initialized)	none
VM_NATIVE_CODE 14 - Reserved for implementation use for storing native machine pointers to native code (see note below)	none
VM_ENUM	       15 - enumerated constant; 32-bit integer	UINT4



Example:

06 6b 00 00 00 - DataHolder
0b             - Length of symbolic name (0x0b = 11d)
6f 62 6a 54 6f 53 74 72 69 6e 67


# TODO
if the tads image file contains a SINI block, that block must be executed before the first instruction from the
ENTP block.