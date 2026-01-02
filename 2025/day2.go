package year2025

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/timothyandrew/advent/util"
)

var (
	productRangeRegex = regexp.MustCompile(`(\d+)-(\d+)`)
)

type productRange struct {
	min int
	max int
}

func isInvalid(n int) bool {
	s := strconv.Itoa(n)

	if len(s)%2 != 0 {
		return false
	}

	left := s[:len(s)/2]
	right := s[len(s)/2:]

	return left == right
}

func findInvalidIDs(r productRange) (invalidIDs []int) {
	for i := r.min; i <= r.max; i++ {
		if isInvalid(i) {
			invalidIDs = append(invalidIDs, i)
		}
	}

	return
}

func isInvalidNewRules(n int) bool {
	s := strconv.Itoa(n)

	for i := 1; i <= len(s)/2; i++ {
		candidate := strings.Builder{}
		candidate.WriteString(s[:i])

		for {
			if candidate.Len() > len(s) {
				break
			}

			if candidate.String() == s {
				return true
			}

			candidate.WriteString(s[:i])
		}
	}

	return false
}

func findInvalidIDsNewRules(r productRange) (invalidIDs []int) {
	for i := r.min; i <= r.max; i++ {
		if isInvalidNewRules(i) {
			invalidIDs = append(invalidIDs, i)
		}
	}

	return
}

func (s Solutions) DayTwoPartOne() error {
	lines, err := util.FilenameToLines("2025/input/2/puzzle.txt")
	if err != nil {
		return err
	}

	ranges := strings.Split(lines[0], ",")

	sum := 0
	for _, line := range ranges {
		matches := productRangeRegex.FindStringSubmatch(line)
		min, err := strconv.Atoi(matches[1])
		if err != nil {
			return err
		}
		max, err := strconv.Atoi(matches[2])
		if err != nil {
			return err
		}

		r := productRange{min: min, max: max}
		invalids := findInvalidIDs(r)
		for _, n := range invalids {
			sum += n
		}

	}
	fmt.Println(sum)

	return nil
}

func (s Solutions) DayTwoPartTwo() error {
	lines, err := util.FilenameToLines("2025/input/2/puzzle.txt")
	if err != nil {
		return err
	}

	ranges := strings.Split(lines[0], ",")

	sum := 0
	for _, line := range ranges {
		matches := productRangeRegex.FindStringSubmatch(line)
		min, err := strconv.Atoi(matches[1])
		if err != nil {
			return err
		}
		max, err := strconv.Atoi(matches[2])
		if err != nil {
			return err
		}

		r := productRange{min: min, max: max}
		invalids := findInvalidIDsNewRules(r)
		for _, n := range invalids {
			sum += n
		}

	}
	fmt.Println(sum)

	return nil
}
