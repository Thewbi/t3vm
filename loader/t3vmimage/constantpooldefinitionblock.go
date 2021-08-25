package t3vmimage

import (
	"fmt"
	"log"

	"github.com/Thewbi/t3vm/loader"
)

// Constant Pool Definition Block
// Identifier: "CPDF"
// The Constant Pool Definition block specifies the overall structure of a constant pool.
// This block contains the number of pages in the pool, and the size of each page.
//
// An image file must have exactly one Constant Pool Definition block for each constant pool,
// and the Constant Pool Definition block must precede the first Constant Pool Page block for that pool
// (i.e., the Definition block must be located at an earlier byte position in the image file than any
// Page block for the same pool).
//
// A Constant Pool Definition block's "mandatory" flag should be set to 1 if the pool identifier is 1 or 2.
// (In other words, all pools currently defined are mandatory. However, it is possible that new pool types
// added in future versions of the VM would not be mandatory, hence we refrain here from specifying that all
// constant pool definition blocks are mandatory.)

func ConstantPoolDefinitionBlock(data []byte, size int, flags int) {
	fmt.Println("ConstantPoolDefinitionBlock() size flags", size, flags)

	//fmt.Printf("%s", hex.Dump(data))

	offset := 0

	// UINT2 (pool identifier)
	// 1 - Byte-code pool (Virtual machine byte code)
	// 2 - Constant data pool
	poolIdentifier := loader.ReadU2_Buffer_LE(data, offset)
	offset += 2
	log.Print("poolIdentifier:", poolIdentifier)
	if poolIdentifier == 1 {
		log.Println("Byte-code pool")
	} else if poolIdentifier == 2 {
		log.Println("Constant data pool")
	} else {
		log.Println("Unknown")
	}

	// UINT4 (number of pages in the pool)
	numberOfPages := loader.ReadU4_Buffer_LE(data, offset)
	offset += 4
	log.Println("numberOfPages:", numberOfPages)

	// UINT4 (size in bytes of each page in the pool)
	sizeOfEachPageInBytes := loader.ReadU4_Buffer_LE(data, offset)
	offset += 4
	log.Println("sizeOfEachPageInBytes:", sizeOfEachPageInBytes)

}
