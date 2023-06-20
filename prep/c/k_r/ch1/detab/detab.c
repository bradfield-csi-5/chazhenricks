#include <stdio.h>

#define SPACES 1 /* number of spaces to insert for tab */

int main() {
  int c;

  while ((c = getchar()) != EOF) {
    if (c == '\t') {
      for (int i = 0; i < SPACES; ++i) {
        printf(" ");
      }
    } else {
      putchar(c);
    }
  }
}
