def swap(arr: list, i: int, j: int):
    temp = arr[i]
    arr[i] = arr[j]
    arr[j] = temp


def partition(arr: list, start: int, end: int):

    pivot = arr[end]
    i = start - 1

    for j in range(start, end):
        if arr[j] <= pivot:
            i += 1
            arr[i], arr[j] = arr[j], arr[i]
    
    arr[i + 1], arr[end] = arr[end], arr[i + 1]
    return i + 1


def quicksort(arr: list, start: int, end: int):
    if start < end:
        pivot = partition(arr, start, end)
        print(pivot)
        quicksort(arr, start, pivot - 1)
        quicksort(arr, pivot + 1, end)


nums = [1, 7, 4, 1, 10, 9, -2]
print("UNSORTED")
print(nums)

size = len(nums)
quicksort(nums, 0, size - 1)

print("SORTED")
print(nums)
