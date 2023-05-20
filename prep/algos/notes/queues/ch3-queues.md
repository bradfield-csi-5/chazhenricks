# Queues

- **queue** - ordered collection where items are added to one end and removed from the other end.
- think a line at a counter.
- first in, first out

### Abstract Data Type

- `queue()` - creates a new empty queue
- `enqueue(item)` - adds item to rear of the queue
- `dequeue()` - removes item from front of the queue and returns the item
- `is_empty` - returns a boolean of if queue is empty
- `size()` - returns an iteger of number of items in queue

### Implementation

```python
class Queue(object):
    def __init__(self):
        self._items = []

    def is_empty():
        return self._items == []

    def enqueue(item):
        self._items.insert(0, item);

    def dequeue():
        return self._items.pop()

    def size():
        return len(self._items)
```
