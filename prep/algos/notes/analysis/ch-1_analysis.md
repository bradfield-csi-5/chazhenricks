# Analysis 

## Section 1 - Big Picture

What makes a program better than others? It depends

- Typically we use speed and memory consumption
- Can be at odds 

Factors in time/space: 
- inputs
- data structures used
- iterations? 

ACTUAL: 
- how long computer to execute instruction
- computers architecture
- how many cores 
- language 
- how OS schedules processes
- other programs running at same time 

Because of this - Algo analysis aims to cut away all the random variables to condense into things that we can measure and control for. 

- Python function to calc sum of first n numbers

```python 
def sum(n):
    total = 0
    for i in range(n + 1):
        total += i
    return total 
```
- this will scale linerally - as N increases, so do the amount of loops

- Do the same, but with no looping: 

```python
def arithmetic_sum(n):
    start = time.time()
    total = n * (n + 1) // 2
    end = time.time()
    return total, end - start

```

The above example doesnt change depending on the input size. If n is 10000000 or 10, the same calculation happens. 
We call this Constant Time ( O(1) )


## Section 2 - Big O Notation 
- execution time == number of steps required to solve the problem 

- Big O === order of magnitude function 

#### Common Orders Of Magnitude Functions
- Constant === O(1)
- Logarithmic === O(log n) 
- Linear === O(n)
- Log Linear === O(n log n)
- Quadratic === O(n^2)
- Cubic === O(n^3) 
- Exponential === O(2^n) - this is _terrible_ 


What is the calculation of this function?: 
```python 
a = 5
b = 6
c = 10
for i in range(n):
   for j in range(n):
      x = i * i
      y = j * j
      z = i * j
for k in range(n):
   w = a * k + 45
   v = b * b
d = 33
```
- Start by counting assignments: 
Assignment operators = 3 
Assignments in first loops = 3n^2 
Assignments in second loop = 2n 
Assignment at end = 1 

Altogether: T(n) = 3 + 3n^2 + 2n + 1 = 3n^2 + 2n + 4 (combine constants together) 

Anything ^2 is going to overpower the rest as input grows, so we can ignore the constants and say this is `O(n^2)`


## Section 3 - Anagram Detection Example 
- see `ch1-analysis.py` for code examples. 

- main takeaway - the most common tradeoff in Big O performance is time vs space. 
Often in order to increase _time_ we make up for it in increased _space_ (or memory consumed) 

 
