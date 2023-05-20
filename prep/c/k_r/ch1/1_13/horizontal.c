/* WTF I DO NOT UNDERSTAND THIS EXERCISE
 * THIS SOLUTION IS RIPPED FROM THE OFFICAL SOLUTION BOOK
 * I STILL DONT FUKIN GET IT
 */

#include <stdio.h>

#define MAXHIST 15 /* max length of a histogram */
#define MAXWORD 11 /* max length of a word */
#define IN 1       /* inside a word */
#define OUT 0      /*outside a word */

int main() {
  int c, i, nc, state;
  int len;         /* length of each bar */
  int max_value;   /* maximum value for wl[] */
  int overflow;    /* number of overflow words */
  int wl[MAXWORD]; /* word length counters */

  state = OUT;
  nc = 0;       /* num of chars in a word */
  overflow = 0; /* number of words overflow */

  // initialize word length counters
  for (i = 0; i < MAXWORD; ++i)
    wl[i] = 0;

  while ((c = getchar()) != EOF) {
    if (c == ' ' || c == '\n' || c == '\t') {
      state = OUT;
      if (nc > 0) {
        if (nc < MAXWORD) {
          ++wl[nc];
        } else {
          ++overflow;
        }
      }
      nc = 0;
    } else if (state == OUT) {
      state = IN;
      nc = 1; /* begining of a new word */
    } else {
      ++nc; /* inside a word */
    }
  }

  max_value = 0;
  for (i = 1; i < MAXWORD; ++i) {
    if (wl[i] > max_value) {
      max_value = wl[i];
    }
  }

  for (i = 1; i < MAXWORD; ++i) {
    printf("%5d - %5d : ", i, wl[i]);
    if (wl[i] > 0) {
      if ((len = wl[i] * MAXHIST / max_value) <= 0) {
        len = 1;
      } else {
        len = 0;
      }
      while (len > 0) {
        putchar('*');
        --len;
      }
      putchar('\n');
    }
  }
  if (overflow > 0) {
    printf("There are %d word >= %d\n", overflow, MAXWORD);
  }
}
