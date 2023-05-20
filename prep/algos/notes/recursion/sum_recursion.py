def iterative_sum(nums):
    total = 0
    for num in nums:
        total += num
    return total


print(iterative_sum([1, 3, 5, 7, 9]))


def recursion_sum(nums):
    if len(nums) == 0:
        return 0

    return nums[0] + recursion_sum(nums[1:])


print(recursion_sum([1, 3, 5, 7, 9]))
