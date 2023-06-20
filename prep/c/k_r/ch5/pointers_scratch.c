#include <stdio.h>

int main()
{
  //initialize an array of 10 ints
  int a[10];
  //set the first one to 25
  a[0] = 25;
  a[1] = 35;
  
  printf("A[0] IS: %d\n", a[0]);
  //initialize an int called pa - this will be a pointer to an element in a[].
  //initializing it with a * lets us store pointers in there.
  //but declaring it an int says "this pointer will point to an int, but its still a pointer"
  int *pa;

  // store the address of the item in a[0] to pa. 
  // weve already said that pa will hold pointers, so we dont need the * in this case:
 //welll only need the * again when we're trying to get the value out of the variable.  
  pa = &a[0];
  printf("*PA IS: %d\n", *pa);


  // this will just make a copy of the data stored in *pa and shove it in X
  // they wont equal the same thing if a[0] changes k
  int x = *pa;
  a[0] = 26;


  printf("X IS: %d\n", x);
  printf("MEMORY ADDRESS OF X IS: %d\n", &x);
  printf("MEMORY ADDRESS OF A[0] IS: %d\n", &a[0]);

  printf("A[0] IS: %d\n", a[0]);
  printf("*PA IS: %d\n", *pa);



  printf("*PA + 1 IS: %d\n", *(pa+1));
  int this_is_wrong;
  

}
