#include <stdio.h>

int main() {
  int c;

  while ((c = getchar()) != EOF) {
    if (c == '\t' || c == ' ' || c == '\n')
      putchar('\n');
    else
      putchar(c);
  }
  return 0;
}

