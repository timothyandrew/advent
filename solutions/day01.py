"""Advent of Code 2025 - Day 01."""

from utils import read_input, read_lines

def parse_line(line):
    if line[0] == 'R':
        return int(line[1:])
    else:
        return -1 * int(line[1:])

def part1():
    """Solve part 1."""
    lines = [parse_line(line) for line in read_lines(1, "puzzle")]

    position = 50
    result = 0

    for line in lines:
        position = (position + line) % 100

        if position == 0:
            result += 1

    return result

def part2():
    """Solve part 2."""
    lines = read_lines(1, "sample")
    pass


if __name__ == "__main__":
    print("Part 1:", part1())
    print("Part 2:", part2())
