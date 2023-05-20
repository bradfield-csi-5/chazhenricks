# Stacks

## Section 1 - Intro 

- Stacks | Queues | Dequeues | Lists == Data structures with items ordered according to how they are added/removed
- Once added, stays in same position === Linear Data Structures 
- Differences occur in _where_ items are added / removed 


### Stacks 
- adding and removal of items occur at the same end. 
- think stack of books or stack of plates. 
- "last in, first out" 
- **top** == part of stack items are added/removed
- **base** == other end of stack 

- order of removal is exact opposite of oder of addition 
- Examples: 
- web browser history 
- call stack of application code 


### Abstract Data Type 
- **abstract data type** == (ADT) logical description of how we view the data and allowed operations without regard as to how theyre implemented. 
- **data structure** == the implementation of the abstract data type. Thing of the ADT as an interface and the Data Structure as the class that implements it. 

This is the Abstract Data Type for a Stack: 

- `Stack()` == creates a new, empty stack 
- `push(item)` == adds a given item to the top and returns nothing 
- `pop()` == removes and returns the top item from the stack 
- `peek()` == returns the top item without removing it 
- `is_empty()` == returns boolean representing if stack has any items 
- `size()` == returns number of items in the stack as an integer. 



## Section 2 - Implementation 
 
- Is a python List a stack? 
- We can use a list as a stack, but a list allows more operations that we have defined in our Abstract Data type
- By using a defined data structure, we can more effectively communicate what we are trying to do as well as give us some rails to not fuck up  

Here is an implementation of a Stack in python: 
```python 
class Stack: 
    def __init__(self):
        self._items = []
    
    def is_empty(self):
        return not bool(self._items)

    def push(self, item):
        self._items.append(item)

    def pop(self):
        return self._items.pop()

    def peek():
        return self._items[-1]

    def size(self):
        return len(self._items)
```
- Note - we could implement the stack by inserting/removing from the front of the stack as well using `insert(0, item)` and `pop(0)`   
- We would end up hurting the performance, as `append()` is `O(1)` and `insert(0)` is  `O(n)` 


## Section 3 - Balanced Parentheses
- In many programming languages, the use of parenthesis needs to be balanced. Meaning that for every `(` we need a cooresponding `)` 
- See `balanced_parentheses.py` for a code example. 


## Section 4 - Converting Number Bases
- the conversion of a decimal number to a binary number is a matter of constantly taking the remainder of the division of a decimal number by two and adding it to a stack. 



