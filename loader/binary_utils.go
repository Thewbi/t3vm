package loader

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"strings"
)

// read one byte unsigned
func ReadUInt1FromFile(file *os.File) uint32 {
	bytes := make([]byte, 1)
	binary.Read(file, binary.BigEndian, bytes)
	return uint32(bytes[0])
}

// read two byte unsigned (big endian)
func ReadUInt2FromFile(file *os.File) uint16 {
	bytes := make([]byte, 2)
	binary.Read(file, binary.BigEndian, bytes)
	return uint16(bytes[1]) | uint16(bytes[0])<<8
}

// read two byte unsigned (little endian)
func ReadUInt2FromFile_LE(file *os.File) uint16 {
	bytes := make([]byte, 2)
	binary.Read(file, binary.LittleEndian, bytes)
	//return uint16(bytes[0]) | uint16(bytes[1])<<8
	return binary.LittleEndian.Uint16(bytes)
}

// read four byte unsigned
func ReadUInt4FromFile(file *os.File) uint32 {
	bytes := make([]byte, 4)
	binary.Read(file, binary.BigEndian, bytes)
	return binary.BigEndian.Uint32(bytes)
}

// read four byte unsigned
func ReadUInt4FromFile_LE(file *os.File) uint32 {
	bytes := make([]byte, 4)
	binary.Read(file, binary.LittleEndian, bytes)
	return binary.LittleEndian.Uint32(bytes)
}

func ReadUInt8FromFile(file *os.File) uint64 {
	bytes := make([]byte, 8)
	binary.Read(file, binary.BigEndian, bytes)
	return binary.BigEndian.Uint64(bytes)
}

func ReadInt4FromFile(file *os.File) int32 {
	bytes := make([]byte, 4)
	binary.Read(file, binary.BigEndian, bytes)
	uivalue := binary.BigEndian.Uint32(bytes)
	return int32(uivalue)
}

func ReadBytesFromFile(file *os.File, size int) []byte {
	bytes := make([]byte, size)
	binary.Read(file, binary.BigEndian, bytes)
	return bytes
}

func ReadU1(reader *bufio.Reader) (uint8, error) {
	bytes := make([]byte, 1)
	n, err := reader.Read(bytes)
	if err != nil {
		//panic(err)
		return 0, err
	}
	if n != 1 {
		//panic("Cannot read bytes!")
		return 0, err
	}
	return bytes[0], err
}

// read four byte unsigned
func ReadU4(reader *bufio.Reader) uint32 {
	bytes := make([]byte, 4)
	n, err := reader.Read(bytes)
	if err != nil {
		panic(err)
	}
	if n != 4 {
		panic("readU4 - Cannot read bytes!")
	}
	return binary.BigEndian.Uint32(bytes)
}

func ReadU2_Buffer(buffer []byte, offset int) uint16 {
	slice := make([]byte, 2)

	// copy is not working for some reason
	//copy(buffer[offset:(offset+2)], slice)

	for i, j := offset, 0; i < offset+2; i, j = i+1, j+1 {
		slice[j] = buffer[i]
	}

	//fmt.Printf("%s", hex.Dump(slice))

	return binary.LittleEndian.Uint16(slice)
}

func ReadU4_Buffer(buffer []byte, offset int) uint32 {
	slice := make([]byte, 4)

	// copy is not working for some reason
	//copy(buffer[offset:offset+4], slice[:])

	for i, j := offset, 0; i < offset+4; i, j = i+1, j+1 {
		slice[j] = buffer[i]
	}

	//fmt.Printf("%s", hex.Dump(slice))

	return binary.LittleEndian.Uint32(slice)
}

func ReadZeroTerminatedString(reader *bufio.Reader) string {

	// string builder
	var b strings.Builder
	b.Grow(32)

	// read until the zero delimiter is reached
	bytes, err := reader.ReadBytes(0)
	if err != nil {
		panic(err)
	}
	b.Write(bytes)

	return b.String()

}

func ReadZeroTerminatedString3(reader *bufio.Reader) string {

	// string builder
	var b strings.Builder
	b.Grow(32)

	// read byte-wise until the zero-termination is reached
	for {
		c, err := reader.ReadByte()
		if err != nil {
			panic(err)
		}
		if c == 0 {
			return b.String()
		}
		b.WriteByte(c)
	}
}

func ReadZeroTerminatedString2(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	var b strings.Builder
	b.Grow(32)

	i := 0
	for i < number && bytes[i] != 0 {

		//fmt.Print(bytes[i])
		b.WriteByte(bytes[i])

		i++
	}

	fmt.Print(b.String())

	return bytes
}

func ReadNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}
