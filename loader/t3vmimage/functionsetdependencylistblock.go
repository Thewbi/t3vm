package t3vmimage

import (
	"fmt"
)

func FunctionSetDependencyListBlock(data []byte, size int, flags int) {
	fmt.Println("FunctionSetDependencyListBlock() size flags", size, flags)

	//fmt.Printf("%s", hex.Dump(data))

}
