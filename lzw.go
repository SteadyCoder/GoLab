// LZW Data Compression in Golang
package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"./lzwAlgorithm"
	// "github.com/pkg/profile"
)

func readLine(path string) string {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	var result string = ""
	for scanner.Scan() {
		result = result + "\n" + scanner.Text()
	}

	return result
}

func main() {



	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup

	wg.Add(2)

	fmt.Print("Input file length:\"")

	testStr := readLine("testReadHard")
	fmt.Print(len(testStr))
	fmt.Println("\"")
	testStrLen := len(testStr)
	part1 := testStr[:testStrLen/2]
	part2 := testStr[testStrLen/2:]

	start := time.Now()

	compressed1 := make(chan []int, 1)
	compressed2 := make(chan []int, 1)

	go lzwAlgorithm.CompressLZW(&wg, part1, compressed1)
	go lzwAlgorithm.CompressLZW(&wg, part2, compressed2)
	wg.Wait()

	comp1 := <-compressed1
	comp2 := <-compressed2
	wg.Add(2)

	uncompressed1 := make(chan string, 1)
	uncompressed2 := make(chan string, 1)

	go lzwAlgorithm.DecompressLZW(&wg, comp1, uncompressed1)
	go lzwAlgorithm.DecompressLZW(&wg, comp2, uncompressed2)

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Binomial took ", elapsed)

}