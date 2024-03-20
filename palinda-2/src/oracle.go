// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

var nonsense = []string{
	"The moon is dark.",
	"The sun is bright.",
}

var prophecies = []string{
	"Tomorrow will be a good day.",
	"The sun will rise in the east.",
	"The moon will be full.",
	"You will fail your MDI tenta.",
}

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := make(chan string)
	answers := make(chan string)

	go questionReceiver(questions, answers)
	go predictionGenerator(answers)
	go answerPrinter(answers)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}

func questionReceiver(questions <-chan string, answers chan<- string) {
	for {
		question := <-questions
		answer := make(chan string)
		go prophecy(question, answer)
		answers <- <-answer
	}
}

func predictionGenerator(answers chan<- string) {
	for {
		time.Sleep(time.Duration(10+rand.Intn(10)) * time.Second)
		answers <- prophecies[rand.Intn(len(prophecies))]
	}
}

func answerPrinter(answers <-chan string) {
	for {
		fmt.Println("\r" + <-answers)
		fmt.Print(prompt)
	}
}

func prophecy(question string, answer chan<- string) {
	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)

	longestWord := ""
	words := strings.Fields(question)
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
}

func init() {
	rand.Seed(time.Now().Unix())
}
