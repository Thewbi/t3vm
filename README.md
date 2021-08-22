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