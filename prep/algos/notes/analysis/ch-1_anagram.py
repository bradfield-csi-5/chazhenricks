from itertools import zip_longest
from collections import Counter

# Solution 1 - Checking offo


def anagram_checking_off(s1, s2):
    if len(s1) != len(s2):
        return False

    # make second string a list
    to_check_off = list(s2)

    # for each character in the first string
    for char in s1:
        for i, other_char in enumerate(
            to_check_off
        ):  # also loop through each item in the second string
            if (
                char == other_char
            ):  # if they are the same character, mark as none and hit the next iteration
                to_check_off[i] = None
                break
        else:
            return False
    return True


print(anagram_checking_off("abcd", "dcba"))
print(anagram_checking_off("abcd", "abcc"))

# This is O(n^2) - each caracter in s1 causes 2 loops to happen. Number of steps is a sum of integers from 1 to n


# Solution 2 - Sort and Compare
# approach here is to sort the strings and see if they have the same value


def anagram_sort_compare(s1, s2):
    for a, b in zip_longest(sorted(s1), sorted(s2)):
        if a != b:
            return False
    return True


print(anagram_sort_compare("abcd", "dcba"))
print(anagram_sort_compare("abcd", "abcc"))


# At first glance this looks like it would be O(n) - since we only loop through once
# However, sorted() _typically_ takes O(n log n) time, so our work looks like this:
# n log n + n log n + O(n)
# the n log n is what dominates, and theyre separate, so we can just drop them to a single O(n log n) time for this


# Solution 3 - Brute Force
# We can do this by creating a list of _all possible_ strings made of of letters from s1. And then check to see
# if s2 is in that list.
# The problem with this, is that there are n possibilites if first letters, second letters,
# third letters .... n letters.
# This translates to n * (n - 1) * (n - 2) ... * 3 * 2 * 1, which ends up being n!
# This is worse than 2^n - if a string is 20 characters long it will process 2,432,902,008,176,640,000 combinations.
# Processed at one per second will take 77,146,816,596 _years_ to process entire list.


# Solution 4 - Count and Compare
# Generate two identical counters - each a list of 0s for letter in the a column, b column etc...


def anagram_count_compare(s1, s2):
    c1 = [0] * 26
    c2 = [0] * 26

    for char in s1:
        pos = ord(char) - ord("a")
        c1[pos] += 1

    for char in s2:
        pos = ord(char) - ord("a")
        c2[pos] += 1

    for a, b in zip(c1, c2):
        if a != b:
            return False
    return True


anagram_count_compare("apple", "pleap")  # => True
anagram_count_compare("apple", "applf")  # => False


# can also be written more succenctly with the `collections.counter`
# which does the break-apart-and-store-in-dict thing


def anagram_with_counter(s1, s2):
    return Counter(s1) == Counter(s2)


print(anagram_with_counter("apple", "pleap"))  # => True
print(anagram_with_counter("apple", "applf"))  # => False
