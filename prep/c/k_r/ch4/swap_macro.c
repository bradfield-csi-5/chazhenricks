#include <stdio.h>

#define swap(t, first, second) \
                               \ 
  {                            \
    t temp;                    \
    temp = first;              \
    first = second;            \
    second = temp;             \
  }

#define paste(front, back) front ## back

int main(void) {
  int a, b;
  a = 10;
  b = 20;

  char myname[] = "chaz";
  swap(int, a, b);
  printf("THIS IS A: %d\n", a);
  printf("THIS IS B: %d\n", b);
  printf("MY NAME IS %s\n", paste(my, name));
}
