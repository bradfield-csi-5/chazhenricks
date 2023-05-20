# Lists

- two different types: Ordered and Unordered. 
- Un-ordered lists are like arrays, but the items aren't sorted
- Ordered Lists _are_sorted. Any time an item is added it is sorted to be inserted into the correct position. 

## Unordered List Abstract Data Type: 

- List() creates a new list that is empty. It needs no parameters and returns an empty list.
- add(item) adds a new item to the list. It needs the item and returns nothing. Assume the item is not already in the list.
- remove(item) removes the item from the list. It needs the item and modifies the list. Assume the item is present in the list.
- search(item) searches for the item in the list. It needs the item and returns a boolean value.
- is_empty() tests to see whether the list is empty. It needs no parameters and returns a boolean value.
- size() returns the number of items in the list. It needs no parameters and returns an integer.
- append(item) adds a new item to the end of the list making it the last item in the collection. It needs the item and returns nothing. Assume the item is not already in the list.
- index(item) returns the position of item in the list. It needs the item and returns the index. Assume the item is in the list.
- insert(pos, item) adds a new item to the list at position pos. It needs the item and returns nothing. Assume the item is not already in the list and there are enough existing items to have position pos.
- pop() removes and returns the last item in the list. It needs nothing and returns an item. Assume the list has at least one item.
- pop(pos) removes and returns the item at position pos. It needs the position and returns the item. Assume the item is in the list.


## Unordered List Implementation
- use a linked-list 
- node has value and reference to next node in the list, thats it. 

```python 
# The Node Class

class Node(object):
    def __init__(self, value):
        self.value = value
        self.next = None
                                
# The Unordered List
class UnorderedList(object):
    def __init__(self):
        self.head = None 

    def is_empty(self):
        return self.head is None

    def add(self, item):
        temp = Node(item)
        temp.next = self.head
        self.head = temp

    def size(self):
        current = self.head
        count = 0

        while current is not None:
            count = count + 1 
            current = current.next 

        return count

    def search(self, item):
        current = self.head

        while current is not None:
            if current.value == item:
                return True
            
            current = current.next

        return False

    def remove(self, item):
        current = self.head
        previous = None 

        while True:
            if current.value = item:
                break
            previous, current = current, current.next

        if previous is None:
            self.head = current.next 
        else:
            previous.next = current.next

    def append(self, item):
        current = self.head

        while current is not None:
            current = current.next

        current.next = Node(item)

    def insert(self, pos, item):
        if pos == 0:
            self.add(item)

        current = self.head
        previous = None
        index = 0

        while index < pos and current is not None: 
            previous, current = current, current.next
            index += 1 

        temp = Node(item)
        temp.next = current
        previous.next = temp 

    def pop(self): 
        current = self.head
        previous = None 

        while current.next is not None:
            previous, current = current, current.next

        previous.next = None 
        return current 

    def pop(self, pos):
        current = self.head
        previous = None 
        index = 0 

        while index < pos and current is not None: 
            previous, current = current, current.next 
            index += 1 

        previous.next = current.next 
        return current 

```


```python 
from unorderedlist import Node, UnorderedList

class OrderedList(UnorderedList):
    
    def search(self, item):
        current = self.head

        while current is not None:
            if current.value == item:
                return True
            if current.value > item:
                return False
            current = current.next

        return False 

    # add needs to be in order, so we need to traverse 
    def add(self, item):
        current = self.head
        previous = None 

        while current is not None:
            if current.value > item:
                break
            previous, current = current, current.next 

        temp = Node(item)

        if previous is None:
            temp.next, self.head = self.head, temp 
        elif:
            previous.next = temp 
            temp.next = current 

    def append(self, item):
        self.add(item)

    def insert(self, pos, item):
        self.add(item) 
``
