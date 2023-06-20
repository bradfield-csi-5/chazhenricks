
#include <stdio.h>

#define MAXLINE 1000 /*  maximum input line length */

int get_line(char line[], int maxline);
void reverse(char line[], int length);

int main() {
  int len;            /* current line length */
  char line[MAXLINE]; /* current input line */

  while ((len = get_line(line, MAXLINE)) > 0) {
    reverse(line, len);
    printf("%s\n", line);
  }
  return 0;
}

/* get_line function: read line into s[] and return length of s[] */
int get_line(char s[], int lim) {
  int c, i;

  for (i = 0; i < lim - 1 && (c = getchar()) != EOF && c != '\n'; ++i) {
    s[i] = c;
  }
  if (c == '\n') {
    s[i] = c;
    ++i;
    /* '\0' is the null char */
  }
  s[i] = '\0';
  return i;
}

/* reverse line */
void reverse(char line[], int length) {
  int b, e;
  b = 0;
  e = length - 1;
  while (b < e) {
    char temp = line[b];
    line[b] = line[e];
    line[e] = temp;
    ++b;
    --e;
  }
}
