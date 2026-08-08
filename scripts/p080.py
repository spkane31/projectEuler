"""
It is well known that if the square root of a natural number is not an integer, then it is irrational.
The decimal expansion of such square roots is infinite without any repeating pattern at all.

The square root of two is 1.4142135623730951, and the digital sum of the first one hundred decimal digits is 475.

For the first one hundred natural numbers, find the total of the digital sums of the first one hundred
decimal digits for all the irrational square roots.

Run with:
    uv run scripts/p080.py
"""

from pathlib import Path
import sys

if __package__ in (None, ""):
    sys.path.insert(0, str(Path(__file__).resolve().parents[1]))

from utils import timeit
from decimal import Decimal, getcontext

getcontext().prec = 105


@timeit
def solve(
    limit: int = 100,  # 100 is rational so no need to go over
) -> int:
    result = 0
    for n in range(1, limit):
        root = Decimal(n).sqrt()

        if int(root) ** 2 == n:
            continue

        root_sum = sum_decimal_digits(root)
        result += root_sum
    print(f"Result: {result}")


def sum_decimal_digits(n: Decimal, limit: int = 100) -> int:
    return sum([int(char) for char in f"{n:.105f}".replace(".", "")[:limit]])


if __name__ == "__main__":
    solve(limit=100)
