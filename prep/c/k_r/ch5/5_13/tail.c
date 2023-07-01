#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define NUMBER 0
#define MAXLINES 5000   /* max lines to be sorted */
#define MAXLEN 1000     /* max length of an array input line */
#define ALLOCSIZE 10000 /* size of available space */

char *lineptr[MAXLINES]; /* array of pointers to the lines to be sorted */

static char allocbuf[ALLOCSIZE];
static char *allocp = allocbuf;

// detemrine if input is a number or not
// probably overkill but whatever
int isNumber(char *s) {
  int i, c;
  i = 0;

  c = s[i];
  if (!isdigit(c) && c != '.' && c != '-') return *s;

  if (c == '-') {
    c = s[++i];
    if (!isdigit(c)) {
      return *s;
    }
  }

  return NUMBER;
}

char *alloc(int n) {
  if (allocbuf + ALLOCSIZE - allocp >= n) {
    allocp += n;
    return allocp - n;
  } else
    return NULL;
}

void afree(char *p) {
  if (p >= allocbuf && p < allocbuf + ALLOCSIZE) {
    allocp = p;
  }
}

int my_getline(char *s, int lim) {
  int c, i;

  /* get the next character until we hit a newline or eof */
  for (i = 0; i < lim - 1 && (c = getchar()) != EOF && c != '\n'; ++i) {
    s[i] = c;
  }
  if (c == '\n') {
    s[i] = c;
    i++;
  }

  s[i] = '\0';

  return i;
}

int readlines(char *lineptr[], int maxlines) {
  int len, nlines;
  char *p, line[MAXLINES]; /* initialize an empty pointer and an array to hold
                              the read in line */
  nlines = 0;

  // read in lines from stdin
  while ((len = my_getline(line, MAXLEN)) > 0)
    // read in lines until we run out of room or there are no new lines
    if (nlines >= maxlines || (p = alloc(len)) == NULL)
      return -1;
    else {
      line[len - 1] = '\0';
      strcpy(p, line);
      *++lineptr = p;
      nlines++;
    }
  return nlines;
}

void writelines(char *lineptr[], int nlines) {
  // count backwards from the newlienes number while counting up from the
  // lineptr
  while (nlines-- > 0) {
    printf("%s\n", *(lineptr--));
  }
}

int main(int argc, char *argv[]) {
  int n = 10;
  int c, type, nlines;
  int status = 0;
  // increment the first argv pointer, _then_ access its first element
  while (--argc > 0 && (*++argv)[0] == '-') {
    type = isNumber(*argv);
    switch (type) {
      case NUMBER:
        n = atoi(++argv[0]);
        break;
      default:
        argc = 0;
        printf("Usage: ./tail -n\n");
    }

  }

  nlines = readlines(lineptr, MAXLINES);
  writelines((lineptr + nlines), n);
  return status;
}
