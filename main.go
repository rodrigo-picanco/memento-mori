package main

import (
	"strconv"
	"time"
)

func main() {
	density := 3
	mem := New("1995-12-29", "80")
	elapsed := mem.CalcElapsedWeeks(time.Now())
        lines := mem.Expectation / density
        items := 52 *  density

        println("MEMEMENTO MORI 💀")

	for i := 0; i < lines ; i++ {
		for j := 0; j < items; j++ {
                        item := i*items+j
			if item <= elapsed {
				print("✅")
			} else {

				print("☑️")
			}
		}
		println()
	}
}

type MementoMori struct {
	Birth       time.Time
	Expectation int
}

func New(birth, expectation string) *MementoMori {
	return &MementoMori{
		Birth:       parseBirth(birth),
		Expectation: parseExpecation(expectation),
	}
}

func parseExpecation(e string) int {
	res, err := strconv.Atoi(e)
	if err != nil {
		panic(err)
	}
	return res

}

func (m *MementoMori) CalcWeek() int {
	return m.Expectation * 52
}

func (m *MementoMori) CalcElapsedWeeks(now time.Time) int {
	return int(now.Sub(m.Birth).Hours() / 24 / 7)
}

func parseBirth(birth string) time.Time {
	res, err := time.Parse(time.DateOnly, birth)
	if err != nil {
		panic(err)
	}
	return res
}
