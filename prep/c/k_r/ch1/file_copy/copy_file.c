#include <stdio.h>
// the EOF char is just -1

int main() {
  int c;

  printf("%d\n", getchar() != EOF);
  // can do assignment inside a conditional statement
  while ((c = getchar()) != EOF) {
    // c will only equal getchar() in side this while looop
    putchar(c);
  }
  return 0;
}
