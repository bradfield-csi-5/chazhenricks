# write a funcion canSum(targetSum, numbers) that takes in a targetSum and an array of numbers as args
# function should return boolean indicating whether or not it is possible to generate the target sum using numbers from the array.
# you may use an element of the array as many times as needed
# you may asssume all numbers are non-megative.


def can_sum_recursive(target, nums, memo={}):
    if target in memo:
        return memo[target]

    if target == 0:
        return True

    if target < 0:
        return False

    for num in nums:
        remainder = target - num
        memo[target] = can_sum_recursive(remainder, nums, memo)
        if memo[target] is True:
            return True

    return False


print(can_sum_recursive(7, [2, 3]))
print(can_sum_recursive(7, [5, 3, 4, 7]))
print(can_sum_recursive(7, [2, 4]))
print(can_sum_recursive(300, [7, 14]))
