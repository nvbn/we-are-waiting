package main

import (
	"bufio"
	"flag"
	"github.com/nvbn/we-are-waiting/variants"
	"golang.org/x/crypto/ssh/terminal"
	"math/rand"
	"os"
	"time"
)

var MAX_POSITION = len(variants.All[0]) - 1

type human struct {
	variant  int
	position int
}

func watchApp(ch chan<- string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ch <- scanner.Text()
	}
	close(ch)
}

func getRandomHuman() *human {
	variant := rand.Intn(len(variants.All))
	return &human{
		variant:  variant,
		position: 0,
	}
}

func canMakeOlder(people []*human) (result []int) {
	for i, man := range people {
		if man.position < MAX_POSITION {
			result = append(result, i)
		}
	}

	return result
}

func getOldest(people []*human) int {
	oldest := 0

	for i, man := range people {
		if man.position > people[oldest].position {
			oldest = i
		}
	}

	return oldest
}

func updatePeople(people []*human, count int) []*human {
	addNew := rand.Intn(5) == 0
	toMakeOlder := canMakeOlder(people)

	if addNew || len(toMakeOlder) == 0 {
		people = append(people, getRandomHuman())
	} else {
		index := toMakeOlder[rand.Intn(len(toMakeOlder))]
		people[index].position += 1
	}

	if len(people) > count {
		oldest := getOldest(people)
		return append(people[:oldest], people[oldest+1:]...)
	} else {
		return people
	}
}

func printPeople(people []*human) {
	os.Stdout.WriteString("\r")
	for _, man := range people {
		os.Stdout.WriteString(variants.All[man.variant][man.position])
		os.Stdout.WriteString("  ")
	}
}

func main() {
	tick := flag.Int("tick", 3, "Tick in seconds")
	width, _, err := terminal.GetSize(0)
	if err != nil {
		width = 80
	}

	count := flag.Int("count", width/3, "Count of waiting people")
	flag.Parse()

	lines := make(chan string)
	people := make([]*human, 1)
	people[0] = getRandomHuman()

	go watchApp(lines)

	for {
		select {
		case line, isOpen := <-lines:
			os.Stdout.WriteString("\r\033[K")
			os.Stdout.WriteString(line)
			os.Stdout.WriteString("\n")

			if isOpen {
				printPeople(people)
			} else {
				return
			}
		case <-time.After(time.Duration(*tick) * time.Second):
			people = updatePeople(people, *count)
			printPeople(people)
		}
	}
}
