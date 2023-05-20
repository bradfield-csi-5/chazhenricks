from collections import deque


def is_palindrome(word):

    char_deque = deque(word)

    while len(char_deque) > 1:
        front = char_deque.popleft()
        rear = char_deque.pop()
        if front != rear:
            return False

    return True


word = "madam"
print(is_palindrome(word))
