PAIRINGS = {"(": ")", "{": "}", "[": "]"}
OPENINGS = PAIRINGS.keys()


def is_balanced(parentheses):
    stack = []
    for paren in parentheses:
        if paren in OPENINGS:
            stack.append(paren)
        else:
            try:
                possible_opening = stack.pop()
                if paren == PAIRINGS[possible_opening]:
                    continue
                else:
                    return False
            except IndexError:
                return False
    return len(stack) == 0


print(is_balanced("(())"))
print(is_balanced("(()))"))
print(is_balanced("{[[((()))]]}"))
print(is_balanced("{[[((())]]}"))
print(is_balanced("(((((((((((((((())))))))))))))))"))
