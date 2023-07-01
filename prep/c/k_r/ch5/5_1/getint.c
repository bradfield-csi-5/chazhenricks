#include <ctype.h>
#include <math.h>
#include <stdio.h>

int getch(void);
void ungetch(int);
int getint(int *pn);

#define BUFSIZE 100
#define SIZE 5
char buf[BUFSIZE];
int bufp = 0;

int getch(void) { return (bufp > 0) ? buf[--bufp] : getchar(); }

void ungetch(int c) {
  if (bufp >= BUFSIZE)
    printf("too many chars in buf \n");
  else
    buf[bufp++] = c;
}

int getint(int *pn) {
  int c, sign;

  // gets the next item from input until the item is no longer a space
  while (isspace(c = getch())) {
    ;
  }

  // if c is not a digit AND
  // c is EOF AND
  // c is not a + sign AND
  // c is not a - sign
  if (!isdigit(c) && c != EOF && c != '+' && c != '-') {
    ungetch(c); /* push the item back to the buffer and return 0 instead*/
    return 0;
  }

  // if c is a minus sign, set sign to -1. Otherwise set it to 1.
  //  we will multiply the number by the sign to make sure we et the correct
  //  positive or negative
  sign = (c == '-') ? -1 : 1;

  // if our current char is a + or a -, then fetch the _next_ character
  if (c == '+' || c == '-')
    c = getch();

  if(!isdigit(c)){
    ungetch(c);
    return 0;
  }

  for (*pn = 0; isdigit(c); c = getch()) {
    *pn = 10 * *pn + (c - '0');
  }

  *pn *= sign;
  if (c != EOF) ungetch(c);
  return c;
}

int main(void) {
  int n, array[SIZE], getint(int *);
 // we pass in the pointer to array[n] to getint
 // the getint function will do all of the reading of stdin to see what the number is 
 // and it will set *pn (the pointer to array[n]) to something
 //
  for (n = 0; n < SIZE && getint(&array[n]) != EOF; n++) {
    printf("the %d index is %d\n", n, array[n]);
  }
}
