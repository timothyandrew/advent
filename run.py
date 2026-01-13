#!/usr/bin/env python3
"""Runner script for Advent of Code 2025 solutions."""

import argparse
import importlib
import time


def main():
    parser = argparse.ArgumentParser(description="Run Advent of Code 2025 solutions")
    parser.add_argument("day", type=int, help="Day number (1-25)")
    parser.add_argument("--part", "-p", type=int, choices=[1, 2], help="Run only part 1 or 2")
    args = parser.parse_args()

    module_name = f"solutions.day{args.day:02d}"
    try:
        module = importlib.import_module(module_name)
    except ModuleNotFoundError:
        print(f"Solution for day {args.day} not found")
        return

    print(f"=== Day {args.day} ===")

    if args.part in (None, 1) and hasattr(module, "part1"):
        start = time.perf_counter()
        result = module.part1()
        elapsed = time.perf_counter() - start
        print(f"Part 1: {result} ({elapsed:.3f}s)")

    if args.part in (None, 2) and hasattr(module, "part2"):
        start = time.perf_counter()
        result = module.part2()
        elapsed = time.perf_counter() - start
        print(f"Part 2: {result} ({elapsed:.3f}s)")


if __name__ == "__main__":
    main()
