"""Project Euler Problem 1: Multiples of 3 or 5.

Run with:
    uv run scripts/p001.py
"""

from pathlib import Path
import sys

if __package__ in (None, ""):
    sys.path.insert(0, str(Path(__file__).resolve().parents[1]))

from utils import timeit


@timeit
def solve(limit: int = 1_000) -> int:
    """Return the sum of multiples of 3 or 5 below ``limit``."""
    print(sum(number for number in range(limit) if number % 3 == 0 or number % 5 == 0))


if __name__ == "__main__":
    solve()
