package main

import (
	"fmt"
	"github.com/agentabstract/go-fitz"
	"io"
	"log"
	"os"
	"runtime"
	"time"
)

func main(){

	file, err:=os.Open("bigtest.pdf")
	if err != nil{
		panic(err)
	}
	defer file.Close()

	bb, err:=io.ReadAll(file)
	if err != nil{
		panic(err)
	}

	for i:=0;i<2000;i++{
		doc, err := fitz.NewFromMemory(bb)
		if err != nil{
			panic(err)
		}
		var textversion string
		for i := 0; i < doc.NumPage(); i++ {
			txt, _ := doc.Text(i)
			textversion += txt
		}
		log.Println(len(textversion))
		PrintMemUsage()
		err = doc.Close()
		if err != nil{
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
