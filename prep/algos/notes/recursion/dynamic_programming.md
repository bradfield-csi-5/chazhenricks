#Dynamic Programming 

Dynamic Programming is a technique similiar to recursion, but will handle sub-problems one at a time instead of overlapping. 
This is tough to think of abstractly - so lets take an example. 

### Fibbonaci Numbers 

The next number in the sequence is the sum of the previous two numbers. 
Recursivly, it could look like this: 

```python 

def recursive_fib(n):
    if n <= 1:
        return n 
    return recursive_fib(n - 1) + recursive_fib(n - 2)

```
While this works, it has a run-time of `O(2^n)` - we end up with a number of redundat calls. 

If we think of a recursion as a "top down" approach - we start with the end and break it up into smaller and smaller sub problems, then Dynamic Programming is the "bottom up" approach. 

From a DP perspective, we want to start at the base case and keep iterating until we hit the nth item. 
That could look like this: 

```python
def dynamic_fib(n):
    a,b = 0,1  # base case 

    for _ in range(n):  # iterate until we hit the nth time 
        a, b = a + b, a   # the next item is always the addition of the previous two  

    return a 
```

While the above is mostly an iterative approach - one technique of dtynamic programming is memoization. 
Memoization is a way of storing calculations so we dont need to calculate them again. In the case of our recursive Fibbonaci - we end up calulcating the same sub trees a few times.

One way we could get around this is to create a fast lookup object (a dictionary in python) and store our result in the dictionary as we calculate it. We can use the fact that python can do default arguments 

From a runtime perspective, we're now at `O(n)` for time and `O(1)` for space 


## Grid Traveler 

- We travel on a 2d grid. Try to get from top left to bottom right. Can only move down or right. 
- How many ways can we get from top to bottom? 

- we can only go to the right and down - so we are essentially making our grid one unit smaller in either direction on each iteration. 
- a base case would be: 
1. The grid is 1,1 - meaning we cant go anywhere else 
2. The grid is in an un-movable state (essentially one of the sides is 0)

a basic recursive version could be something like: 


```python 
def recursive_grid(m,n):
    if m == 1 and n == 1: 
        return 1 
    if m == 0 or n == 0:
        return 0

    return recursive_grid(m - 1, n) + recursive_grid(m, n - 1)
```
- this _works_ but puts us in the same exponential time space as fibonacci
- a good rule of thumb for recursive Big O is - however many times we are returning the recursive function, that is the base. 
- In this example - we are returning 2 calls to the recursive function, and no significant changes to n or m to make it better than linear, so our runtime is `O(2^n+m)` - big yikes for anything above 6 or 7 

### Memoize the calculations 
- Same somewhat formulaic steps to memoize stuff in this instance as well:
1. create a default empty object argument for the memo 
2. check if your key (i.e. the new state of your arguments) exists in the memo 
3. if yes, return `memo[key]`
4. if no set `memo[key] = the previously returned recursive call` 
5. return `memo[key]`


```python 
def dynamic_grid(m,n, memo={}):
    key = f"{str(m)},{str(n)}"

    if key in memo: 
        return memo[key]

    if n == 1 and m == 1:
        return 1 
    if n == 0 or m == 0:
        return 0 
    memo[key] = dynamic_grid(m - 1, n, memo) + dynamic_grid(m, n - 1, memo)
    return memo[key]
```


### Memoization Recipe
1. Make it work 
- visualize the problem as a tree (break down calls into tree structure)
- implement the tree using recursion 
- leaves of the tree are the base cases 
- test it - once we implement it, make sure its correct

2. Make it efficint 
- add a memo object (default empty object)
- come up with a key that is some string representation of the new state of the arguments 
- add base case checking the memo object (if key in memo)
- store the recursive return value into the memo under the key 
- return memo[key]




