"""Utility functions for Advent of Code 2025."""

from pathlib import Path


def read_input(day: int, type: str) -> str:
    """Read the input file for a given day as a string."""
    path = Path(__file__).parent / "inputs" / f"day{day:02d}/{type}.txt"
    return path.read_text().strip()


def read_lines(day: int, type: str) -> list[str]:
    """Read the input file for a given day as a list of lines."""
    return read_input(day, type).splitlines()


def read_ints(day: int) -> list[int]:
    """Read the input file for a given day as a list of integers."""
    return [int(line) for line in read_lines(day)]
