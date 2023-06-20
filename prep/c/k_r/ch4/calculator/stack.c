#include <stdio.h>

#include "calc.h"

#define MAXVAL 100
int sp = 0;
double stack[MAXVAL];

void push(double num) {
  if (sp < MAXVAL) {
    stack[sp] = num;
    sp++;
  } else {
    printf("Wtf are you doing? Thats too many numbers\n");
  }
}

double pop(void) {
  if (sp > 0) {
    return stack[--sp];
  } else {
    printf("stack empty ding dong\n");
    return 0.0;
  }
}
