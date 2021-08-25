package t3vmimage

import (
	"fmt"
	"log"

	"github.com/Thewbi/t3vm/loader"
)

func LoadEntryPointBlock(data []byte, size int, flags int) {
	fmt.Println("LoadEntryPointBlock() size flags", size, flags)

	//fmt.Printf("%s", hex.Dump(data))

	// 4 byte - offset into the codepool (= a Constant Pool Definition Block (CPDF) that has a pool identifier set to the value 1 (= byte-code pool))
	// this is where code execution starts after parsing the t3 image file
	codePoolOffset := loader.ReadU4_Buffer_LE(data, 0)
	log.Println("codePoolOffset: ", codePoolOffset)

	//fmt.Printf("%s", hex.Dump(data))

	// 2 byte - method header size for all methods
	methodHeaderSize := loader.ReadU2_Buffer_LE(data, 4)
	log.Println("methodHeaderSize: ", methodHeaderSize)

	// 2 byte - exception table entry size for all exception tables
	exceptionTableEntrySize := loader.ReadU2_Buffer_LE(data, 6)
	log.Println("exceptionTableEntrySize:", exceptionTableEntrySize)

	// 2 byte - debugger line table entry size for all line tables
	debuggerLineTableEntrySize := loader.ReadU2_Buffer_LE(data, 8)
	log.Println("debuggerLineTableEntrySize:", debuggerLineTableEntrySize)

	// 2 byte - debug table header size for all debug tables
	debuggerTableHeaderSize := loader.ReadU2_Buffer_LE(data, 10)
	log.Println("debuggerTableHeaderSize:", debuggerTableHeaderSize)

	// 2 byte - debug table local symbol record header size for all debug tables
	debugTableLocalSymbolRecordHeaderSize := loader.ReadU2_Buffer_LE(data, 12)
	log.Println("debugTableLocalSymbolRecordHeaderSize:", debugTableLocalSymbolRecordHeaderSize)

	// 2 byte - debug records version number
	debugRecordsVersionNumber := loader.ReadU2_Buffer_LE(data, 14)
	log.Println("debugRecordsVersionNumber:", debugRecordsVersionNumber)

	//fmt.Printf("%s", hex.Dump(data))
}
