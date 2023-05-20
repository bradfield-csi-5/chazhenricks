def howSum(target, nums):
    if target == 0:
        return []
    if target < 0:
        return None

    for num in nums:
        remainder = target - num
        result = howSum(remainder, nums)
        if result is not None:
            return [*result, remainder]

    return None


print(howSum(7, [2, 3]))
