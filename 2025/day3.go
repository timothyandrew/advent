package year2025

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/timothyandrew/advent/util"
)

func optimalJoltage(bank string, length int) (int, error) {
	digitStrings := strings.Split(bank, "")
	digits := make([]int, len(digitStrings))
	for i, digitString := range digitStrings {
		digit, err := strconv.Atoi(digitString)
		if err != nil {
			return 0, err
		}

		digits[i] = digit
	}

	n := 0
	max := 0
	maxIndex := -1

	for j := length; j > 0; j-- {
		// <j>th digit
		for i := maxIndex + 1; i < len(digits)-(j-1); i++ {
			if digits[i] > max {
				max = digits[i]
				maxIndex = i
			}
		}

		n += max * int(math.Pow(float64(10), float64(j-1)))
		max = 0
	}

	return n, nil
}

func (s Solutions) DayThreePartOne() error {
	lines, err := util.FilenameToLines("2025/input/3/puzzle.txt")
	if err != nil {
		return err
	}

	sum := 0
	for _, line := range lines {
		j, err := optimalJoltage(line, 2)
		if err != nil {
			return err
		}

		sum += j
	}

	fmt.Println(sum)

	return nil
}

func (s Solutions) DayThreePartTwo() error {
	lines, err := util.FilenameToLines("2025/input/3/puzzle.txt")
	if err != nil {
		return err
	}

	sum := 0
	for _, line := range lines {
		j, err := optimalJoltage(line, 12)
		if err != nil {
			return err
		}

		sum += j
	}

	fmt.Println(sum)

	return nil
}
