# Advent of Code 2025

My solutions to [Advent of Code 2025](https://adventofcode.com/2025) in Python.

## Project Structure

```
advent/
├── inputs/              # Puzzle inputs (day01.txt, day02.txt, etc.)
├── solutions/           # Solution files (day01.py, day02.py, etc.)
│   └── template.py      # Template for new solutions
├── utils.py             # Helper functions
├── new_day.py           # Generate new day's files
└── run.py               # Run solutions with timing
```

## Usage

### Starting a New Day

```bash
python new_day.py <day>
```

This creates:
- `solutions/dayXX.py` - Solution file from template
- `inputs/dayXX.txt` - Empty input file

### Running Solutions

```bash
python run.py <day>         # Run both parts
python run.py <day> -p 1    # Run only part 1
python run.py <day> -p 2    # Run only part 2
```

### Utility Functions

Import helpers in your solution files:

```python
from utils import read_input, read_lines, read_ints

# read_input(day)  - Returns raw input as a string
# read_lines(day)  - Returns list of lines
# read_ints(day)   - Returns list of integers (one per line)
```

## Example Workflow

```bash
# Start day 1
python new_day.py 1

# Paste puzzle input into inputs/day01.txt

# Edit solutions/day01.py with your solution

# Run it
python run.py 1
```
