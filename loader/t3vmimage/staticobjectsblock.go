package t3vmimage

import (
	"fmt"
)

func StaticObjectsBlock(data []byte, size int, flags int) {
	fmt.Println("StaticObjectsBlock() size flags", size, flags)

	//fmt.Printf("%s", hex.Dump(data))

}
