#include <stdio.h>
// will return negative is s < t
// reuturn = if s == t
// return positive if s > t

// array version
int arr_strcmp(char *s, char *t) {
  int i; /* initialize an index counter */

  // test here will start at begining of array, and copy chars over one at a
  // time
  for (i = 0; s[i] == t[i]; i++)
    // if the source reaches the null character, return 0;
    if (s[i] == '\0') return 0;

  return s[i] - t[i];
}

int point_strcmp(char *s, char *t) {
  for (; *s == *t; s++, t++)
    if (*s == '\0') return 0;
  return *s - *t;
}


int main()
{
  char first[] = "aaaa";
  char next[] = "zzz";
  int ret = point_strcmp(first, next);
  printf("WHATS THE NUMBER?: %d\n", ret);
}
