package main

import (
	"flag"
	"fmt"
	"time"
)

var FUZZ = 2

var FUZZINESS = map[int]time.Duration{
	1: 5,
	2: 15,
	3: 30,
	4: 60, // every hour
	5: 60, // morning, day, evening, night
}

var WORDS = map[int]string{
	0:  "midnight",
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
	45: "quarter",
}

func hourWord(hours int) string {
	if hours > 12 {
		hours -= 12
	}

	return WORDS[hours]
}

func fuzzy(now time.Time, fuzz int) string {
	fuzziness, present := FUZZINESS[fuzz]
	if !present {
		return fmt.Sprintf("%s", "sometime somewhere")
	}

	t := now.Round(fuzziness * time.Minute)

	hours := t.Hour()
	minutes := t.Minute()

	glue := "past"

	if minutes > 35 {
		glue = "to"
		hours = t.Round(time.Hour).Hour()
	}

	if minutes == 0 {
		if (hours != 0) {
			return fmt.Sprintf("%s o'clock", hourWord(hours))
		} else {
			return fmt.Sprintf("%s", hourWord(hours))
		}
	} else {
		return fmt.Sprintf("%s %s %s", WORDS[minutes], glue, hourWord(hours))
	}
}

var fuzzy_func = []func(time.Time) string {
	abit_fuzzy,
	fuzzy,
	more_fuzzy,
	vague,
	abmiguous,
}

func main() {
	f := flag.Int(
		"f", 
		FUZZ,
		"fuzziness degree: 1 - a bit fuzzy, 2 - fuzzy, 3 - more fuzzy, 4 - vague, 5 - abmiguous")
	flag.Parse();

	//fmt.Println(fuzzy(time.Date(2024, 6, 6, 23, 40, 0, 0, time.UTC), *f))
	fmt.Println(fuzzy(time.Now(), *f))
}
