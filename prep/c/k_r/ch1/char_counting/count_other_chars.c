#include <stdio.h>

int main() {
  int c, line, tab, blanks;
  line = tab = blanks = 0;

  while ((c = getchar()) != EOF) {
    if (c == '\n')
      ++line;
    if (c == '\t')
      ++tab;
    if (c == ' ')
      ++blanks;
  }
  printf("lines tabs blanks\n");
  printf("%5d %4d %6d\n", line, tab, blanks);
}
