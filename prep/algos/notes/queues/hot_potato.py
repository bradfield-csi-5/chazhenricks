from collections import deque


def hot_potato(names, num):
    # create a queue
    queue = deque()

    # add names to queue
    for name in names:
        queue.appendleft(name)

    # until there is one child left
    while len(queue) > 1:
        # for num number of times, go from end of list to begining
        for _ in range(num):
            queue.appendleft(queue.pop())
        # after num times, that child is left off
        queue.pop()

    return queue.pop()


names = ("Chaz", "Alex", "Steve", "Scott", "Michael", "David")
print(hot_potato(names, 9))
