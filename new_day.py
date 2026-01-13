#!/usr/bin/env python3
"""Generate a new day's solution file from template."""

import argparse
import os
from pathlib import Path


def main():
    parser = argparse.ArgumentParser(description="Create a new day's solution file")
    parser.add_argument("day", type=int, help="Day number (1-25)")
    args = parser.parse_args()

    if not 1 <= args.day <= 25:
        print("Day must be between 1 and 25")
        return

    root = Path(__file__).parent
    template = root / "solutions" / "template.py"
    target = root / "solutions" / f"day{args.day:02d}.py"

    input_dir = root / "inputs" / f"day{args.day:02d}"
    input_files = [input_dir / filename for filename in ["puzzle.txt", "sample.txt"]]

    if target.exists():
        print(f"Solution file for day {args.day} already exists")
        return

    content = template.read_text().replace("Day XX", f"Day {args.day:02d}").replace("XX", str(args.day))
    target.write_text(content)
    print(f"Created {target}")

    if not input_dir.exists():
        os.mkdir(input_dir)

    for input_file in input_files:
        if not input_file.exists():
            input_file.touch()
            print(f"Created empty input file: {input_file}")


if __name__ == "__main__":
    main()
