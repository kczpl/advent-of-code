import time
from functools import wraps


def read_puzzle(day: int) -> str:
    with open(f"puzzles/input{day}.txt", "r") as file:
        return file.read().split("\n")


def read_test_puzzle(day: int) -> str:
    with open(f"puzzles/test{day}.txt", "r") as file:
        return file.read().split("\n")


def benchmark(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        start = time.perf_counter()
        result = func(*args, **kwargs)
        end = time.perf_counter()
        print(f"{func.__name__}: {result} ({(end - start) * 1000:.3f} ms)")
        return result

    return wrapper
