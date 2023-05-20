# Trees
- hierarchical representation of data 
- files systems / html elements examples 

- **node** - element in the tree 
- **edge** - connection between two nodes 
- **root** - top level node 
- **path** - ordered list of nodes connected by edges (how it got from A to D)
- ** children** - a node that has incoming edges from a node is a _child_ of that node 
- **parent** - a node that has outgoing edges to another node
- **sibling** - two nodes that have incoming edges from the same parent node 
- **subtree** - set of nodes comprised of a parent and its edges/children / typically used to break down a large tree into smaller tree problems 
- **leaf node** - a node that has no outgoing edges. The last node in a path. 
- **level** - number of edges from parent to node. 
- **height** - maximum level of any node in the tree. (the biggest level in the whole tree is the height of the tree)

### Formal Definitions 
1. A tree consists of a set of nodes and a set of edges that connect pairs of nodes. A tree has the following properties: 
  - one node is the root 
  - every node N except the root node, is connectd by an edge from exactly one other node. 
  - a unique path traverses from root to each node (no double dippin)
  - if each node has a max of two children - then we say its a binary tree 
2. A tree is either empty or consists a root and zero or more subtrees, each of which is also a tree. The root of each subtree is connected to a parent through an edge. 


## Representing A Tree 
- how to represent a node in a binary search tree: 
```python 
class Node(object):
    def __init__(self, val):
        self.val = val
        self.left = left 
        self.right = right 

    def insert_left(self, child):
        if self.left is None: 
            self.left = child 
        else:
            child.left = self.left
            self.left = child 

    def insert_right(self, child):
        if self.right is None: 
            self.right = child 
        else:
            child.right = self.right
            self.right = child 
```

#### List of Lists 
- another common way to represent trees in python is with a "list of lists" 
- first element == value, 
- second element == left child
- third element == right child

```python 
tree = [
    'a', # root
    [
        'b',  # left child
        ['d', [],[]],
        ['e',[],[]]
    ],
    [
        'c',  # right child
        ['f',[],[]],
        []

    ]
]


def insert_left(root, child_val):
    subtree = root.pop(1)
    if len(subtree) > 1:
        root.insert(1, [child_val, subtree,[]])
    else: 
        root.insert(1, [child_val, [], []])
    return root

def insert_right(root, child_val):
    subtree = root.pop(2)
    if len(subtree) > 1:
        root.insert(2, [child_val,[], subtree])
    else: 
        root.insert(2, [child_val, [], []])
    return root

def get_root_val(root):
    return root[0]

def set_root_val(root, new_val):
    root[0] = new_val

def get_left_child(root):
    return root[1]

def get_right_child(root):
    return root[2]

```
- this works, and is fairly simple, but tough to see tree structure when represented as list of lists


### Map-based representation (preferred)
- use maps/dicts/objects - whatever your language calls it. 
- example in python: 
```python
{
    'val': 'A',
    'left': {
        'val': 'B',
        'left': {'val': 'D'},
        'right': {'val': 'E'}
    },
    'right': {
        'val': 'C',
        'right': {'val': 'F'}
    }
}

### or with multiple children

{
    'val': 'A',
    'children': [
        {
            'val': 'B',
            'children': [
                {'val': 'D'},
                {'val': 'E'},
            ]
        },
        {
            'val': 'C',
            'children': [
                {'val': 'F'},
                {'val': 'G'},
                {'val': 'H'}
            ]
        }
    ]
}

```

## Parse Trees 
- used to construct real-world representations. Like sentences or math expressions. 
```python
{
    'val': 'Sentence',
    'left': {
        'val': 'Noun Phrase',
        'left': {
            'val': 'Proper Noun'
            'left': {
                    'val': "Homer"
                }
        },
    },
    'right': {
        'val': 'Verb Phrase',
        'left': {
            'val': 'Verb'
            'left':{
                'val': 'Hit'
            },
        },
        'right': {
            'val': 'Noun Phrase',
            'left': {
                'val': 'Bart'
            }
        }
    }
}


```

- see `expression_parse.py` for a code example of an expression parser in python


## Tree Traversal 
- three main methods for traversing a tree: 
1. preorder
2. in order 
3. post order 

- #### Pre Order 
- visit root node first 
- recursive preorder of the left tree, then recursive preorder of right tree 
- think of as reading a book front to back. 
- top root is book 
- first level subtrees are chapters
- chapter subtreees are sections etc... 

```python 
def preorder(node):
    if node: 
        print(node['val'])
        preorder(node.get('left'))
```
    preorder(node.get('right'))

#### In Order 
- recursive inorder of left subtrees
- then root 
- then recursive inorder or right subtree. 
- this will read the tree from left to right 

```python
def inorder(node):
    if node: 
        inorder(node.get('left'))
        print(node['val'])
        inorder(node.get('right'))
```

### Post Order 
- recursive post order of left 
- recursive post order of right 
- root 
- think of like the parse tree 
- evaluate left side, then the right side, then the root 
```python
def postorder(node):
    if node: 
        postorder(node.get('left'))
        postorder(node.get('right'))
        print(node['val'])
```
 
## Priority Queues with Binary Heap 
 - a self ordering queue 
 - min heap == smallest always at front 
 - max heap == largest always at front 
 - enqueues/deques in `O(log n)`

#### Abstract Data Type

- BinaryHeap() = creates new empty heap
- insert(k) = adds new item
- find_min() = returns item with min value 
- del_min() = removes min value and returns it 
- is_empty() = boolean is empty
- size()
- built_heap(list) - builds a new heap from list of keys 

### Structure 
- in order to keep log add/remove we need to keep tree balanced (almost same on left and right)
- we fill in leaf nodes from left to right 
- can represent with a single list:
  - left child of parent `p` is at `2p` 
  - right child at `2p + 1` 
  - if a node is at position `n` its parent is at `n//2`

### Binary Heap Implementation 

```python 
class BinaryHeap(object):
    def __init__(self): 
        self.items = [0] # added to make the integer division easier

    def __len__(self):
        return len(self.items) - 1 # accounts for the leading 0

    # we want bigger items on bottom. 
    # when new item added to list we will "bubble" it up to the spot in the list where it fits
    def percolate_up(self):
        i = len(self)
        while i // 2 > 0:
            if self.items[i] < self.items[i //2]:
                self.items[i], self.items[i // 2] = self.items[i // 2], self.items[i]
        i = i // 2 

    # add items to end of list then call percolate_up to self balance 
    def insert(self, item):
        self.items.append(item)
        self.percolate_up()

    def delete_min(self):
        return_value = self.items[1] # grab first item in heap
        self.items[1] = self.items(len(self))  # replace top of heap with last item on heap
        self.items.pop() #remove redundant item from end of heap since its now at begining
        self.percolate_down(1)
        return return_value

    def percolate_down(self, index):
        # while there are still children (parent is i // 2, so children check would be i * 2 ) 
        while index * 2 <= len(self):
            min_child = self.min_child(index)
            if self.items[i] > self.items[min_child]:
                self.items[i], self.items[min_child] = self.items[min_child], self.items[i]
            i = min_child

    def min_child(seld, index):
        # if a right child exists (which will be the smaller item in a min heap) we want to return that 
        # these are checks to see if a right child exists and to make sure the right child is less than left child
        # if right child index is out of bounds
        if i * 2 + 1 > len(self):
            return i * 2 
        #if for whatever reason the left and right are switched
        if self.items[i * 2] < self.items[i * 2 + 1 ]: 
            return i * 2 
        return i * 2 + 1 

    def build_heap(self, a_list): 
        i = len(a_list) // 2 
        self.items = [0] + a_list  # add entire list to the heap 
        # sort the list by percolating_down the entire thing
        while i > 0: 
            self.percolate_down(i)
            i = i - 1 


```


