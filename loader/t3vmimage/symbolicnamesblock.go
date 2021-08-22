package t3vmimage

import (
	"encoding/hex"
	"fmt"
)

func SymbolicNamesBlock(data []byte, size int, flags int) {
	fmt.Println("SymbolicNamesBlock() size flags", size, flags)

	fmt.Printf("%s", hex.Dump(data))

}
