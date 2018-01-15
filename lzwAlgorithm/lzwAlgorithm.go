package lzwAlgorithm

import (
	"fmt"
	"sync"
)

func initialCompressDictionary() map[string]int {
	dictionary := make(map[string]int)

	for i := 0; i < 256; i++ {
		dictionary[string(i)] = i
	}
	return dictionary
}

func CompressLZW(wg *sync.WaitGroup, testStr string, result chan<- []int) {
	defer wg.Done()
	code := 256
	dictionary := initialCompressDictionary()
	fmt.Println("Cheburek")
	currChar := ""
	res := make([]int, 0)

	for _, c := range []byte(testStr) {
		phrase := currChar + string(c)
		if _, isTrue := dictionary[phrase]; isTrue {
			currChar = phrase
		} else {
			res = append(res, dictionary[currChar])
			dictionary[phrase] = code
			code++
			currChar = string(c)
		}
	}
	if currChar != "" {
		res = append(res, dictionary[currChar])
	}
	result <- res
}

// Inits dictionary for decode a given sequence of numbers
func initialDecompressDictionary() map[int]string {
	dictionary := make(map[int]string)
	for i := 0; i < 256; i++ {
		dictionary[i] = string(i)
	}
	return dictionary
}

func decodeLZW(indx int, currSymbol string, currCode int, dict map[int]string) string {
	decodedStr := ""
	if num, availible := dict[indx]; availible {
		decodedStr = num
	} else if indx == currCode {
		decodedStr = currSymbol + currSymbol[:1]
	} else {
		panic(fmt.Sprintf("Bad compressed element: %d", indx))
	}
	return decodedStr
}

func DecompressLZW(wg *sync.WaitGroup, compressed []int, result chan<- string) {
	defer wg.Done()
	fmt.Println("kek")
	dictionary := initialDecompressDictionary()
	currentChar := string(compressed[0])
	res := currentChar
	currentCode := 256
	for _, element := range compressed[1:] {
		decodedStr := decodeLZW(element, currentChar, currentCode, dictionary)
		res += decodedStr
		dictionary[currentCode] = currentChar + decodedStr[:1]
		currentCode++
		currentChar = decodedStr
	}
	result <- res
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
}
=======
}
>>>>>>> ccbac0d... Rock'n'Roll babe.
=======
}
>>>>>>> 57d1e8c... Add little changes and refactor code.
=======
}
>>>>>>> ccbac0d... Rock'n'Roll babe.
