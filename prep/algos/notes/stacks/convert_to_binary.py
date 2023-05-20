DIGITS = "0123456789abcdef"


def convert_to_base(decimal_number, base):
    remainder_stack = []
    while decimal_number > 0:
        new_remainder = decimal_number % base
        remainder_stack.append(new_remainder)
        decimal_number = decimal_number // base

    binary_digits = []
    while remainder_stack:
        binary_digits.append(DIGITS[remainder_stack.pop()])

    return "".join(binary_digits)


print(convert_to_base(25, 2))
print(convert_to_base(25, 16))
