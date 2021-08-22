package main

import (
	"log"
	"os"

	"github.com/Thewbi/t3vm/loader/t3vmimage"
)

func main() {
	log.Println("Start ...")

	//file, err := os.Open("data/MrsPepper.t3")
	file, err := os.Open("data/helloworld/helloworld.t3")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	t3vmimage.Load(file)

	log.Println("End.")
}
