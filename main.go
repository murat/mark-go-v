package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

var (
	iteration int
	word      string
	length    int
)

func main() {
	flag.StringVar(&word, "word", "istanbul", "--word example")
	flag.IntVar(&length, "length", 100, "--length 100")
	flag.Parse()

	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	content := strings.ToLower(string(data))
	splittedContent := strings.Split(content, " ")

	wordList := map[string]map[string]int{}
	f := splittedContent[:len(splittedContent)-1]
	s := splittedContent[1:len(splittedContent)]

	for i, fw := range f {
		sw := s[i]

		if val, ok := wordList[fw]; ok {
			ct := 1
			for k, c := range val {
				if k == sw {
					ct = c + 1
				}
			}
			wordList[fw][sw] = ct
		} else {
			wordList[fw] = map[string]int{sw: 1}
		}
	}

	fmt.Printf("%v ", strings.ToLower(word))
	iteration = 0
	generateText(strings.ToLower(word), wordList)
}

func generateText(word string, m map[string]map[string]int) {
	iteration++
	found := m[word]

	sum := 0
	for _, v := range found {
		sum += v
	}

	rand.Seed(time.Now().UTC().UnixNano())
	randomInt := rand.Intn(sum)
	exportedK := ""

	for k, v := range found {
		randomInt -= v
		exportedK = k
		if randomInt < 0 {
			break
		}
	}

	fmt.Printf("%v ", exportedK)

	if iteration != length {
		generateText(exportedK, m)
	}
}
