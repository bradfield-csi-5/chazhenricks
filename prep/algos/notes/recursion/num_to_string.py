CHAR_FOR_INT = "0123456789abcdef"


def num_to_string(num, base):
    if num < base:
        return CHAR_FOR_INT[num]

    return num_to_string(num // base, base) + CHAR_FOR_INT[num % base]


print(num_to_string(7, 2))
