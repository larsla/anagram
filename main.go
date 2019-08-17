package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	all := make(map[string][]string)

	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		sorted := sortChars(word)
		if _, ok := all[sorted]; !ok {
			all[sorted] = make([]string, 0)
		}
		all[sorted] = append(all[sorted], word)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	highest := struct {
		word  string
		count int
	}{
		word:  "",
		count: 0,
	}
	for word, instances := range all {
		if len(instances) > highest.count {
			highest.count = len(instances)
			highest.word = word
		}
	}

	fmt.Println("Winner: ", highest.word, highest.count)
	for _, w := range all[highest.word] {
		fmt.Println("  ", w)
	}
}

func sortChars(word string) string {
	sorted := strings.Split(word, "")
	sort.Strings(sorted)

	return strings.Join(sorted, "")
}
