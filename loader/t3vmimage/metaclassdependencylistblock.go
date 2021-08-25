package t3vmimage

import (
	"fmt"
)

func MetaClassDependencylistBlock(data []byte, size int, flags int) {
	fmt.Println("MetaClassDependencylistBlock() size flags", size, flags)

	//fmt.Printf("%s", hex.Dump(data))

}
