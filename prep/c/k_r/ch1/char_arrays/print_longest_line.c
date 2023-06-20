#include <stdio.h>

#define MAXLINE 1000 /*  maximum input line length */

int get_line(char line[], int maxline);
void copy(char to[], char from[]);

/* print longest input line*/
int main() {
  int len;               /* current line length */
  int max;               /* max seen so far */
  char line[MAXLINE];    /* current input line */
  char longest[MAXLINE]; /* longest saved here */

  max = 0;
  printf("I AM THE VALUE OF LINE OUT OF WHILE\n");
  printf("%s\n", line);
  while ((len = get_line(line, MAXLINE)) > 0) {
    printf("I AM THE VALUE OF LINE\n");
  printf("%s\n", line);
    if (len > max) {
      max = len;
      copy(longest, line);
    }
  }
  /* there was a line */
  if (max > 0) {
    printf("LONGEST LINE IS\n");
    printf("%s", longest);
  }
  return 0;
}

/* get_line function: read line into s[] and return length of s[] */
int get_line(char s[], int lim) {
  int c, i;
  printf("THIS IS LINE IN GET_LINE: %s\n", s);

  for (i = 0; i < lim - 1 && (c = getchar()) != EOF && c != '\n'; ++i) {
    s[i] = c;
  }
  if (c == '\n') {
    s[i] = c;
    ++i;
  }
  /* '\0' is the end of line char */
  s[i] = '\0';
  return i;
}

/* copy 'from' into 'to' */
void copy(char to[], char from[]) {
  int i;
  i = 0;
  while ((to[i] = from[i]) != '\0') {
    ++i;
  }
}
