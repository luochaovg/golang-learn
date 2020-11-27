package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	//log.Println("Aaa")
	//time.Sleep(time.Second)

	fileObj, err := os.OpenFile("./info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file err %v, \n", err)
		return
	}

	log.SetOutput(fileObj)

	for {
		log.Println("This is Test log")
		time.Sleep(time.Second)
	}
}
