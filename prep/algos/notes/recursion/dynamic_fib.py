def fib(n, memo={}):
    if n <= 1:
        return n
    if n in memo:
        return memo[n]

    memo[n] = fib(n - 1, memo) + fib(n - 2, memo)
    return memo[n]


def recursive_fib(n):
    if n <= 1:
        return n
    return recursive_fib(n - 1) + recursive_fib(n - 2)


print(fib(6))
print(fib(50))


# print(recursive_fib(50))
