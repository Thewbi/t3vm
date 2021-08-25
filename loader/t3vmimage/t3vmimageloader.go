package t3vmimage

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/Thewbi/t3vm/loader"
)

///
/// File format specification is: http://www.tads.org/t3spec/format.htm
/// Encoding specification is: http://www.tads.org/t3spec/bincode.htm
///
func Load(file *os.File) {
	log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	log.Println("Loading image!")

	//reader := bufio.NewReaderSize(file, 999999)
	//loader.ReadU1(reader)

	//
	// read the signature block which starts with a signature of 11 bytes
	//

	signature := loader.ReadNextBytes(file, 11)
	//fmt.Printf("%s", hex.Dump(signature))

	// T3-image\015\012\032
	expected_signature := []byte{'T', '3', '-', 'i', 'm', 'a', 'g', 'e', '\x0d', '\x0a', '\x1a'}
	//fmt.Printf("%s", hex.Dump(expected_signature))

	if !bytes.Equal(signature, expected_signature) {
		log.Println("Not a T3 image!")
		return
	}

	//
	// file format version number (The current file format version is 1 (0x0001))
	//

	fileFormatVersionNumber := loader.ReadUInt2FromFile_LE(file)
	//log.Println("T3 image version! ", fileFormatVersionNumber)

	if fileFormatVersionNumber != 0x0001 {
		log.Println("Not a T3 image for the expected version number 0x0001!")
		return
	}

	//
	// block of 32 reserved bytes which are ignored in version 0x0001 of the file format
	// 28 zero bytes
	// 4 bytes of build configuration hash code
	// (reserved for use by compilers, linkers, and other build tools, and up to
	// these tools to use as they wish; VM implementations should simply ignore this field)
	//

	reservedBytes := loader.ReadNextBytes(file, 32)
	reservedBytes = reservedBytes
	//fmt.Printf("%s", hex.Dump(reservedBytes))

	//
	// image file timestamp (ASCII string in C-Library asctime() formatting without zero termination)
	//

	imageFileTimestamp := loader.ReadNextBytes(file, 24)
	imageFileTimestamp = imageFileTimestamp
	//fmt.Println(string(imageFileTimestamp))

	//
	// data blocks
	//

	// Each data block starts with a 10-byte header describing the block
	// four byte - type of block
	// four byte - size of the block excluding the 10 byte header
	// two byte - flags

	done := false
	for !done {
		blockType := loader.ReadNextBytes(file, 4)
		blockTypeAsString := string(blockType)

		//fmt.Println("")
		//fmt.Println("")
		fmt.Println("Block found: ", blockTypeAsString)

		// do not process any more data after the EOF block was identified
		if blockTypeAsString == "EOF " {
			fmt.Println("End of File Block found!")
			break
		}

		blockSize := loader.ReadUInt4FromFile_LE(file)
		//fmt.Println("Block Size: ", blockSize)

		blockFlags := loader.ReadUInt2FromFile_LE(file)
		//fmt.Println("Block Flags: ", blockFlags)
		// prevent compiler error declared but not used
		blockFlags = blockFlags

		// consume block payload
		blockData := loader.ReadNextBytes(file, int(blockSize))
		// if blockTypeAsString == "ENTP" {
		// 	fmt.Printf("%s", hex.Dump(blockData))
		// }

		// DEBUG output first ten bytes of payload
		//firstBytes := make([]byte, 10)
		//copy(blockData[0:10], firstBytes[:])
		//fmt.Printf("%s", hex.Dump(firstBytes))

		if blockTypeAsString == "ENTP" {
			//fmt.Println("Entry Point Block found!")
			LoadEntryPointBlock(blockData, int(blockSize), int(blockFlags))
		} else if blockTypeAsString == "SYMD" {
			SymbolicNamesBlock(blockData, int(blockSize), int(blockFlags))
		} else if blockTypeAsString == "FNSD" {
			FunctionSetDependencyListBlock(blockData, int(blockSize), int(blockFlags))
		} else if blockTypeAsString == "CPDF" {
			ConstantPoolDefinitionBlock(blockData, int(blockSize), int(blockFlags))
		} else if blockTypeAsString == "CPPG" {
			ConstantPoolPageBlock(blockData, int(blockSize), int(blockFlags))
		} else if blockTypeAsString == "MCLD" {
			MetaClassDependencylistBlock(blockData, int(blockSize), int(blockFlags))
		} else if blockTypeAsString == "OBJS" {
			StaticObjectsBlock(blockData, int(blockSize), int(blockFlags))
		} else if blockTypeAsString == "EOF " {
			//fmt.Println("End of File Block found!")
			break
		} else {
			//fmt.Println("Unknown Block Found: ", blockTypeAsString)
		}
	}

}
