# Deques 

## Intro
- **Deque** -  double ended queue
- items can be added/removed from the front or back of the line 
- combination stack and queue 

### Abstract Data Type 
- `deque()` - creates a new empty instance 
- `add_front(item)` - adds to the front of the line 
- `add_rear(item)` - adds to back of the line 
- `remove_front()` - removes from front and returns item 
- `remove_rear()` - removes from back and returns item 
- `is_empty()`
- `size`


### Implementation

```python 
class Deque(object):
    def __init__(self):
        self._items = []

    def is_empty(self):
        return self._items == []

    def add_front(self,item):
        self._items.append(item)

    def add_rear(iself,tem):
        self._items.insert(0, item)

    def remove_rear(self):
        self._items.pop(0)

    def remove_front(self):
        self._items.pop()

```
