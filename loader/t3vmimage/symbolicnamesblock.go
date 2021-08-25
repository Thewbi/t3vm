package t3vmimage

import (
	"fmt"
	"log"

	"github.com/Thewbi/t3vm/loader"
)

func SymbolicNamesBlock(data []byte, size int, flags int) {
	fmt.Println("SymbolicNamesBlock() size flags", size, flags)

	//fmt.Printf("%s", hex.Dump(data))

	// number of entries in the table
	offset := 0
	numberOfEntries := loader.ReadU2_Buffer_LE(data, offset)
	offset += 2
	log.Println("numberOfEntries:", numberOfEntries)

	for i := 0; i < int(numberOfEntries); i++ {

		// data holder
		//numberOfEntries := loader.ReadU2_Buffer_LE(data, 0)
		offset += 5

		// length of identifier
		lengthIdentifier := loader.ReadU1_Buffer_LE(data, offset)
		log.Println("lengthIdentifier:", lengthIdentifier)
		offset += 1

		// identifier
		identifier := string(loader.Read_Buffer_LE(data, offset, int(lengthIdentifier)))
		offset += int(lengthIdentifier)
		log.Println("identifier:", identifier)

	}

}
