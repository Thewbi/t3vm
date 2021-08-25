package t3vmimage

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/Thewbi/t3vm/loader"
)

// Constant Pool Page Block
// Identifier: "CPPG"
//
// A Constant Pool Page block provides the data for one page of a constant pool.
//
// The block contains a small header that specifies the pool with which the page
// is associated and the page's index within the pool, followed by the data contained in the page.
//
// The first Constant Pool Page block associated with a particular constant pool must be preceded
// by a Constant Pool Definition block for that pool. (Other blocks may intervene;
// the requirement is merely that the Definition block be earlier in the file than the first
// associated Page block.)
//
// Constant Pool Page blocks for a particular pool need not be in any particular order in the image file,
// and need not be contiguous.
//
// A Constant Pool Page block's mandatory flag should be set to the same value as the mandatory flag for
// the associated Constant Pool Definition block.
//
// The loaded image has one Constant Pool Page block for each page of each constant pool.

func ConstantPoolPageBlock(data []byte, size int, flags int) {
	fmt.Println("ConstantPoolPageBlock() size flags", size, flags)

	//fmt.Printf("%s", hex.Dump(data))

	offset := 0

	// "CPPG" Header

	// UINT2 (pool identifier)
	// The "pool identifier" value has the same meaning as for a Constant Pool Definition block.
	// 1 - Byte-code pool (Virtual machine byte code)
	// 2 - Constant data pool
	poolIdentifier := loader.ReadU2_Buffer_LE(data, offset)
	offset += 2
	log.Println("poolIdentifier:", poolIdentifier)

	// UINT4 (page index)
	// The "page index" value specifies the index of the page within the pool.
	// The first page in the pool has index 0, the second page has index 1, and so on.
	// The starting byte offset of the data in the page is given by multiplying the pool's page size
	// (specified in the Constant Pool Definition block for the pool) by the page index.
	pageIndex := loader.ReadU4_Buffer_LE(data, offset)
	offset += 4
	log.Println("pageIndex:", pageIndex)

	// UBYTE (xor mask)
	// The "xor mask" is a byte value that must be XOR'ed with each byte of the page data bytes whenever
	// the page data bytes are loaded. If this is zero, it has no effect and can be ignored
	// (since x xor 0 equals x for any value of x).
	//
	// Otherwise, whenever the page data bytes are loaded, each byte must be XOR'ed with this value.
	xorMask := loader.ReadU1_Buffer_LE(data, offset)
	offset += 1
	log.Println("xorMask:", xorMask)

	// Page data bytes

	fmt.Printf("%s", hex.Dump(data))
}
