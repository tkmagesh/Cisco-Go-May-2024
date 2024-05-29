package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Incididunt laborum velit minim ex excepteur nostrud reprehenderit occaecat occaecat ea occaecat dolore in irure Non nulla incididunt occaecat enim eiusmod irure dolor Velit ut voluptate adipisicing elit officia aute aute et magna magna deserunt dolor Exercitation ipsum consequat tempor reprehenderit Laborum Lorem fugiat sit voluptate eiusmod Aliqua pariatur do ad commodo deserunt cupidatat incididunt sint ad in do incididuntAliquip laboris mollit qui ex sint esse Eu adipisicing laborum in mollit non commodo deserunt consectetur Amet excepteur commodo incididunt nostrud commodo ipsum Velit et enim sunt anim ea excepteur cupidatat Velit Lorem incididunt aute mollit id esse cillum reprehenderit occaecat officia adipisicing Occaecat irure duis sint laboris velit Irure eu nostrud esse sit excepteur occaecat ea culpa esse tempor quis et mollit deseruntNostrud minim incididunt sunt dolore sit irure fugiat irure Nisi commodo anim non culpa Ut consectetur nisi esse est Lorem sunt est laboris sunt laborum"

	/* find the "word size that occurs the most" and print the word size and the count
	ex:
		4 letter words occurs the most with 30 occurrances
	*/

	wordsCountBySize := make(map[int]int)
	words := strings.Split(str, " ")
	for _, word := range words {
		key := len(word)
		wordsCountBySize[key]++
	}
	var maxCount, maxSize int
	for size, count := range wordsCountBySize {
		if count > maxCount {
			maxCount, maxSize = count, size
		}
	}
	fmt.Println(wordsCountBySize)
	fmt.Printf("%d letter words occur the most with %d occurrances\n", maxSize, maxCount)
}
