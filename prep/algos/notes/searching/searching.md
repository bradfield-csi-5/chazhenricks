# Searching 

- going through a collection of items and returning True/False if item exists
- Python has super handy `in` operator

```python
15 in [1,2,3,4,5]
 False 

 3 in [1,2,3,4,5]
 True 

```

_how_ search works is fun. There are two main ways to do a search - sequential and binary.


## Sequential Search 
- this is the typical, iterative, `O(n)` search 
- starts at the beginning, and touches each item until it finds the match. 

```python 
def sequential_search(items, target):

    for item in items: 
        if item == target:
            return True 
    
    return False

```
- best case (the item is the first one we see): `O(1)`
- worst case (item is last we see): `O(n)` 


## Binary Search 
- if the list is _ordered_ in some way - we can be a little more efficient with our searches. 
- we can start in the middle of the list, check if we're above or below the target, and discard half the remaining searches 

- an example of a _recursive_ binary search: 

```python 
def binary_search(items, target):
    if not items:  # list is empty, base case
        return False 
    
    midpoint = len(items) // 2 
    
    if items[midpoint] == target:  # found it! 
        return True 

    if target < items[midpoint]: 
        binary_search(items[0:midpoint], target) #search again with the midpoint as the upper bound 

    return binary_search(items[midpoint + 1:-1], target)  # otherwise search again, starting after the midpoint as lower
```
- `O(log n)` time 
> NOTE - in python using a slice `items[:index]` is `O(k)` time. If we pass a starting and ending index, however, we can make that constant. 
>Also note - depending on the size of the input, it _might_ be better to sort first, _then_ search. If we can sort once and will need to search many time, it could be worth it. But if we have a large input and are only searching once, it might be better to sequential. 

## Hashing 
- hash table == a collection of items stored in an ordered, indexed way 
- hash function == a function that takes an input and will, consistently, output an index to store that data in the hash table 
- one technique is to use the table size: `hash(item) = item % table_size`
- this leads to collisions: multiple items needing to share the same index value 

#### Hash Function Variations
- **Folding** 
- group numbers into even groups
- sum the groups 
- take remainder of summed_group % table_size  
So for phone number: 
```
770-715-3847
77 +  07 +  15 +  38 +  47 = 184 
184 % 11 = 8

```
- **mid-square**
- square number, take middle digits, use remainder as index 
```
44 

44^2 = 1936
93 % 11 = 5 
```

- **ordinal strings**
- ord value == the ascii representation of a character 
- sum ord values, use remainder 
```
cat 
ord(c) + ord(a) + ord(t) = 312 
312 % 11 = 4
```
- note anagrams will always have the same hash value - we can remedy this with a "position" offset

```python 
def ord_hash(string, tablesize):

    pos = 1
    for char in string:
        the_sum += ord(char + pos)
        pos += 1 
    return the_sum % tablesize
```

### Hash Collision
- if two items have same hash value, we need to put the item in a new slot
- **linear probing** == just keep moving to the next slot until we find an empty
  - problem with this is now we have to do a sequential search for every item 
  - there is also potential for hotspots 
- possible extension is the "slot" method == instead of the next value, we go by 2 or 3 slots 
  - `newhash = rehash(oldhash)`
  - `rehash(pos) = (pos + slot) % tablesize` where `slot` is spots to skip 
  - table should be a prime number size 
  - one more way is each item in the hash table is actually a collection.
    - when collisions happen, item added to the end of the current collection 





