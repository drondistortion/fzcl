package main

import (
	"fmt"
	"time"
)

var WORDS = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	15: "quarter",
	20: "twenty",
	30: "half",
}

func hourWord(hours int) string {
	if hours > 12 {
		hours -= 12
	}

	return WORDS[hours]
}

func fuzzy(now time.Time) string {
	hours := now.Hour()
	minutes := now.Minute()
	glue := "past"

	rounded := (minutes + 2) % 60 / 5 * 5
	if rounded > 35 {
		rounded = 60 - rounded
		hours++
		glue = "til"
	}

	if rounded == 25 || rounded == 35 {
		rounded = 30
	}

	if rounded == 0 {
		if minutes > 30 {
			hours++
		}

		return fmt.Sprintf("%s o'clock", hourWord(hours))
	} else {
		return fmt.Sprintf("%s %s %s", WORDS[rounded], glue, hourWord(hours))
	}
}

func main() {
	fmt.Println(fuzzy(time.Now()))
}
