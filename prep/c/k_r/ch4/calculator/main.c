#include <stdio.h>
#include <stdlib.h> /* for atof() */

#include "calc.h"
#define MAXOP 100 /* max size of operand */

int main(void) {
  int type;
  double op2; /* for division and subtraction */
  char s[MAXOP];
  double number;

  while ((type = getop(s)) != EOF) {
    printf("WHATS THE TYPE?: %d\n", type);
    switch (type) {
      case NUMBER:
        number = atof(s);
        printf("NUMBER: %f\n", number);
        push(number);
        break;
      case '+':
        push(pop() + pop());
        break;
      case '*':
        push(pop() * pop());
        break;
      case '-':
        printf("IM THE MINUS SIGN\n");
        op2 = pop();
        push(pop() - op2);
        break;
      case '/':
        op2 = pop();
        if (op2 != 0.0) {
          push(pop() / op2);
        } else {
          printf("you cant divide by 0 ya dingus\n");
        }
        break;
      case '%':
        op2 = pop();
        if (op2 != 0.0) {
          push((int)pop() % (int)op2);
        } else {
          printf("you cant divide by 0 ya dingus\n");
        }
        break;
      case '\n':
        printf("result: \t%.8g\n", pop());
        break;
        printf("idk man you fucked something up\n");
      default:
        break;
    }
  }
  return 0;
}
