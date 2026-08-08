# Project Euler

This repository contains standalone Python solutions for [Project Euler](https://projecteuler.net/) problems, managed with [uv](https://docs.astral.sh/uv/).

Each problem lives in exactly one file under `scripts/`. Start a new solution by adding a file named `pNNN.py`, where `NNN` is the zero-padded problem number. Keep all code needed for that solution in its file.

## Run a solution

```sh
uv run scripts/p001.py
```

`uv` creates the local virtual environment automatically. This project currently has no third-party dependencies; if a future solution needs one, add it with `uv add <package>`.

## Utilities

Use `@timeit` to print a solution's runtime after it executes:

```python
from utils import timeit


@timeit
def solve() -> int:
    return 42
```

## Legacy solutions

Earlier work has been preserved in `old/` and is not part of the active Python project.
