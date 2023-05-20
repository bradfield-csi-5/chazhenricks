# recursive grid


def recursive_grid(m, n):
    if n == 1 and m == 1:
        return 1
    if n == 0 or m == 0:
        return 0

    return recursive_grid(m - 1, n) + recursive_grid(m, n - 1)


# print(recursive_grid(3, 2))
# print(recursive_grid(18, 18))


def dynamic_grid(m, n, memo={}):
    key = f"{str(m)},{str(n)}"

    if key in memo:
        return memo[key]
    if n == 1 and m == 1:
        return 1
    if n == 0 or m == 0:
        return 0

    memo[key] = dynamic_grid(m - 1, n, memo) + dynamic_grid(m, n - 1, memo)
    return memo[key]


print(dynamic_grid(1, 1))
print(dynamic_grid(2, 3))
print(dynamic_grid(3, 2))
print(dynamic_grid(3, 3))
print(dynamic_grid(18, 18))
