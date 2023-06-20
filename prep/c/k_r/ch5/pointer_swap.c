#include <stdio.h>

void swap(int *a, int *b);

int main() {
  int a, b;

  a = 10;
  b = 20;

  printf("I AM A: %d\n", a);
  printf("I AM B: %d\n", b);

  //pass is _directions_ to a and b 
  swap(&a, &b);

  printf("I AM A: %d\n", a);
  printf("I AM B: %d\n", b);
}

// clang wont even let this compile because the values never really get used;
//  void swap(int a, int b) {
//    int temp;
//
//    temp = a;
//    a = b;
//    b = temp;
//  }

void swap(int *a, int *b) {
  int temp = *a;
  *a = *b;
  *b = temp;
}
