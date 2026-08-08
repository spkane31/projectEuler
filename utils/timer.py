"""Timing helpers."""

from collections.abc import Callable
from functools import wraps
from time import perf_counter
from typing import ParamSpec, TypeVar


Parameters = ParamSpec("Parameters")
Result = TypeVar("Result")


def timeit(function: Callable[Parameters, Result]) -> Callable[Parameters, Result]:
    """Print a function's elapsed runtime in seconds after it runs."""

    @wraps(function)
    def wrapper(*args: Parameters.args, **kwargs: Parameters.kwargs) -> Result:
        started_at = perf_counter()
        try:
            return function(*args, **kwargs)
        finally:
            elapsed = perf_counter() - started_at
            print(f"{function.__name__} took {elapsed:.6f} seconds")

    return wrapper
