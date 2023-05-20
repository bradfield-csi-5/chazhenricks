# Recursion

- method of solving problems where you break the problem down into repeatable steps and do those steps over and over again until you hit a base case
- in CS - that is a function calling itself (with a slightly different input each time) until a pre-determined base case is met. 

### Example - Sum Of Numbers
- see `sum_recursion.py` for code example 


## Three Laws Of Recursion

#### 1. Must Have A Base Case 
- this is the "exit hatch", without one we will get an infinite loop
- in the sum example, base case was an empty array. As long as the array wasn't empty, we kept going. 

#### 2. Must Change Its State And Move Toward Base Case
- each iteration must move toward the base case. 
- in the sum example - base case was empty array. On each call, we were able to remove an item from the array. 
- this moved us one step closer to the base case on each recursive call.
- calling the same function over and over again with no change in state is just as bad as no base case

#### 3. Must Call Itself, Recursively
- the point of this is to break down problems into chunks we can solve by calling the same thing over and over with one small change each time. 
- reduce is a great example - we call the same thing over and over again, with something small changing each time. 


#### Example - Integer To Any Base 

- say we want to convert a decimal number to its string representation. This can be in any base we want. 
- for now, lets stick with base 10. Say we want to represent 789 as a string. 
- we know our base 10 digits - `0123456789` - so any number less than 10 is easy to convert to a string this way
- so our recursive algorithm would be something like: 
1. reduce the number to a single digit 
2. convert digit to string with a lookup 
3. concatenate the single digit string to form final result

So how do we change state and make progress toward the base case? 

- since were dealing with numbers at the onset - we could resonably subtract or divide to diminish the current number toward a single digit
- subtraction seems frought - since the value to subtract by would change each time 
- we could divide by the base we're aiming for - and any remainder would be our single digit 

```python 
CHAR_FOR_INT = "0123456789abcdef"

def to_string(num, base):
    if num < base:
        return CHAR_FOR_INT[num]

    return to_string(num // base) + CHAR_FOR_INT[n % base] 

```
- note `//` will do division without the remainder (how we reduce the main number to smaller digits)
- `%` will do _just_ the remainder - the single digit we want. 
- so if we take a number 123 
- `123 // 10` will return just 12 (subtracting one 10s place digit)
- `123 % 10` will give us 3, the _remainder_ of dividing 123 by 10 

### Tower Of Hanoi 
- move a tower of five (or more!) disks from one peg to the third peg
- can only move one disk at a time 
- can only stack smaller disks on top of bigger disks

How do we go about this recursively? 
First think of a base case. 

How do we move 3 disks? Two disks? What about just one disk? 
Three steps to move from starting pole to goal pole using an intermediadery:

1. move tower of height - 1 to intermediate pole, using final pole
2. move base disk to final pole 
3. move tower of height - 1 to final pole using starting pole 

see`hanoi.py` for code 
