package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
)

func countWords(words chan string, wordsMap map[string]int) {
	for word := range words {
		_, ok := wordsMap[word]

		if ok {
			wordsMap[word]++
		} else {
			wordsMap[word] = 1
		}
	}
}

func convertTextFileToWordsMap(fileName string, wordsChl chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()

		// remove all non-character
		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {
			log.Fatal(err)
		}

		filteredWord := reg.ReplaceAllString(word, "")
		loweredWord := strings.ToLower(filteredWord)
		wordsChl <- loweredWord
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please input the folder path which contains text files you want to words count. Example: \"./week2/datasample\"")
	for scanner.Scan() {
		path := scanner.Text()
		wordsChl := make(chan string)
		wordsMap := make(map[string]int)
		var wg sync.WaitGroup

		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			wg.Add(1)
			go convertTextFileToWordsMap(path+"/"+file.Name(), wordsChl, &wg)
		}

		go countWords(wordsChl, wordsMap)
		wg.Wait()

		for k, v := range wordsMap {
			fmt.Printf("%s=>%d ", k, v)
		}
		fmt.Println("\n\nPlease input the folder path:")
	}
}
