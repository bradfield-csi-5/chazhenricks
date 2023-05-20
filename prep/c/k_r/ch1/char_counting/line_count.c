#include <stdio.h>

int main() {
  int c, nl;
  nl = 0;

  while ((c = getchar()) != EOF)
    // siongle quotes will end up as ascii representation 
    if (c == '\n')
      ++nl;

  printf("%d\n", nl);
}
