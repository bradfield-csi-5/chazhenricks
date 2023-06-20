#include <stdio.h>

// strncpy
// copy n chars from t to s
void my_strncopy(char *s, char *t, int n) {
  int i;
  for (i = 0; i < n; (*s++ = *t++)) i++;
  *s = '\0';
}

// strncat
// concatenate n chars from t onto s

void my_strncat(char *s, char *t, int n) {
  int i;
  // get to the end of s
  while (*s != '\0') s++;

  for (i = 0; i < n; (*s++ = *t++)) i++;
  *s = '\0';
}

// strncm
// compare n characters between t and s
// return <0 if s< t; 0 if s==t; an >0 if s > t
int my_strncmp(char *s, char *t, int n) {
  // initialize a counter at 0 so we can track our comparisons to n
  int i = 0;

  // make copies of the array pointers passed in
  // we want to compare our subset, not the actual array pointes
  char *se = s, *te = t;

  // while i is less than n,
  // and the value of *se is equal to the value of *te, increment them by 1.
  // the loop will terminate when *se and *te are different
  // we determine the sorting rank by what the fitst differing character is.
  for (i = 0; i < n && *se == *te; se++, te++, i++)
    ;
  printf("WHAT IS *se?: %c\n", *se);
  printf("WHAT IS *se?: %c\n", *se);
  return *se - *te;
}

int main() {
  char s[10];
  char t[] = "fantastic";

  my_strncopy(s, t, 3);
  printf("THIS IS THE NEW STRING %s\n", s);

  char neat[] = "ok";
  char cool[] = "eeolo";
  my_strncat(neat, cool, 2);
  printf("THIS IS THE NEW STRING %s\n", neat);

  char first[] = "aaaaaaaa";
  char second[] = "zzzzzzzzzzz";
  int same = my_strncmp(first, second, 3);
  printf("Are the first 3 chars of \"%s\" and \"%s\" the same?: %d\n", first,
         second, same);
}
