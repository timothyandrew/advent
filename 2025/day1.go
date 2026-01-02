package year2025

import (
	"fmt"
	"strconv"

	"github.com/timothyandrew/advent/util"
)

type direction int

const (
	left direction = iota
	right
)

type dial struct {
	position int
	zeroes   int
}

type instruction struct {
	direction direction
	steps     int
}

func (d *dial) ApplyInstructionCountZeroes(i *instruction) {
	newPosition := d.position
	if i.direction == left {
		newPosition -= i.steps
	} else {
		newPosition += i.steps
	}

	newPosition = util.Mod(newPosition, 100)

	if newPosition == 0 {
		d.zeroes++
	}

	d.position = newPosition
}

func (d *dial) ApplyInstructionCountZeroPasses(i *instruction) {
	newPosition := d.position
	distanceToZero := 0

	if i.direction == left {
		newPosition -= i.steps
		distanceToZero = d.position
	} else {
		newPosition += i.steps
		distanceToZero = 100 - d.position
	}

	if i.steps >= distanceToZero {
		// Count an extra zero when normalizing to the nearest zero, _except_ if
		// the dial is currently at zero, in which case no normalization is needed.
		if distanceToZero > 0 {
			d.zeroes++
		}

		normalized := i.steps - distanceToZero
		d.zeroes += normalized / 100
	}

	d.position = util.Mod(newPosition, 100)
}

func (s Solutions) DayOnePartOne() error {
	lines, err := util.FilenameToLines("2025/input/1/puzzle.txt")
	if err != nil {
		return err
	}

	dial := dial{position: 50}
	for _, line := range lines {
		i := instruction{direction: left}

		if line[0] == 'R' {
			i.direction = right
		}

		i.steps, err = strconv.Atoi(line[1:])
		if err != nil {
			return err
		}

		dial.ApplyInstructionCountZeroes(&i)
	}

	fmt.Println(dial.zeroes)

	return nil
}

func (s Solutions) DayOnePartTwo() error {
	lines, err := util.FilenameToLines("2025/input/1/puzzle.txt")
	if err != nil {
		return err
	}

	dial := dial{position: 50}
	for _, line := range lines {
		i := instruction{direction: left}

		if line[0] == 'R' {
			i.direction = right
		}

		i.steps, err = strconv.Atoi(line[1:])
		if err != nil {
			return err
		}

		dial.ApplyInstructionCountZeroPasses(&i)
	}

	fmt.Println(dial.zeroes)

	return nil
}
