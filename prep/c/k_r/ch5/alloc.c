#include <stdio.h>
#define ALLOCSIZE 10000  /* size of available space */ 

static char allocbuf[ALLOCSIZE]; /* private buffer to manage storage in this program */ 
static char *allocp = allocbuf; /* keep track of pointer of currently availble space */ 

// if we get a pointer back, were good to go 
// if we get 0 back, there is not enough space 
// NOTE: C guarentees that 0 is not a valid memory address, so this is safe
char *alloc(int n)   /* we are going to return a pointer to a chunk of n chars */ 
{
  //this chek works because allocbuf == a pointer to the begining of our alloc buffer array, or allocbuf[0]
  //allocp is the pointer to the next free space, think alloc[i] 
  //we're essentiallt checking if MAX - current position is greater than the number of things were asking for. 
  //so if capacity is 10, current storage is 7, and we ask for 2, its gonna fit and we're good
  //but if we ask for 4, it wont fit and we're no longer good 
  if(allocbuf + ALLOCSIZE - allocp >= n) /* check to see if it fits */ 
  {
    allocp += n; /* update our pointer to the new spot. think of like incrementing i for arr[i] */ 
    return allocp - n; /* we want to return the pointer address of the first free position, so that client can use */
  }else 
    return NULL; 
}

//this functoin will free up space, or in practice, reset the pointer to a spot that is ok to overwrite
void afree(char *p)
{
  //saying, if the memory address passed in is greater than the start of our array 
  //as well as if p is less than the _end_ of the array 
  //essentially saying "make sure this pointer fits in the range of what we're keeping track of"
  if(p >= allocbuf && p < allocbuf + ALLOCSIZE)
  {
  allocp = p; 
  }
}
