package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"time"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("🤔")
	}
	for _, file := range files {
		if file.IsDir() || file.Name() == "eat" {
			continue
		}
		currentSize := file.Size()
		tick := 0
		for currentSize > 0 {
			var face string
			if tick%2 == 0 {
				face = "😐"
			} else {
				face = "😬"
			}
			fmt.Printf("\r%v ", face)

			biteSize := int64(math.Min(float64(currentSize), 1024*1024))
			currentSize -= biteSize
			os.Truncate(file.Name(), currentSize)
			time.Sleep(time.Second)
			tick++
		}
	}
	fmt.Println()
}
