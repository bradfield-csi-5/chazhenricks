#include <stdio.h>

#include "my_getline.h"
#include "strindex.h"

#define MAXLINE 1000 /* max line input */

char pattern[] = "ould";

int main(void) {
  char line[MAXLINE];
  int found = 0;
  while (my_getline(line, MAXLINE) > 0) {
    if (strindex(line, pattern) >= 0) {
      printf("%s\n", line);
      found++;
    }
  }
  return found;
}
