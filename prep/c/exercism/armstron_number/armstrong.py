def armstrong(number):
    running_total = 0
    digit_list = []
    orig = number

    while number > 0:
        digit_list.append(number % 10)
        number = number // 10

    power = len(digit_list)
    for digit in digit_list:
        exponent = digit**power
        running_total += exponent

    return running_total == orig


print(armstrong(10))
