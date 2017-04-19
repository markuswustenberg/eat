package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"time"
)

const (
	confused    = "🤔"
	openMouth   = "😬"
	closedMouth = "😐"
	dead        = "💀"
)

func main() {
	for {
		files, err := ioutil.ReadDir(".")
		if err != nil {
			face(confused)
			break
		}

		eaten := false
		for _, file := range files {
			if file.IsDir() || file.Size() == 0 || file.Name() == "eat" {
				continue
			}
			eat(file)
			eaten = true
		}

		if !eaten {
			me, err := os.Stat("eat")
			if err != nil {
				face(confused)
				break
			}
			eat(me)
			face(dead)
			break
		}
	}
	fmt.Println()
}

func eat(file os.FileInfo) {
	currentSize := file.Size()
	tick := 0
	for currentSize > 0 {
		if tick%2 == 0 {
			face(openMouth)
		} else {
			face(closedMouth)
		}

		biteSize := int64(math.Min(float64(currentSize), 1024*1024))
		currentSize -= biteSize
		os.Truncate(file.Name(), currentSize)
		time.Sleep(time.Second)
		tick++
	}
}

func face(face string) {
	fmt.Printf("\r%v ", face)
}
