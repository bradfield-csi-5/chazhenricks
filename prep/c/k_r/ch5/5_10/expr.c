#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>

#define MAXOP 100
#define NUMBER '0'
#define MAXVAL 100

int sp = 0;
double stack[MAXVAL];

void push(double num) {
  printf("IM PUSHIN A THING DAD %f\n", num);
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

#define BUFSIZE 100

char buf[BUFSIZE];
int bufp = 0;

int getch(void) { return (bufp > 0) ? buf[--bufp] : getchar(); }

void ungetch(int c) {
  if (bufp >= BUFSIZE) {
    printf("TO MANY CHARACTERS\n");
  } else {
    buf[bufp++] = c;
  }
}

int getop(char *s) {
  printf("WHATS THE ARG? %s\n", s);
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

int main(int argc, char *argv[]) {
  int type;
  double op2; /* for division and subtraction */
  char s[MAXOP];
  double number;

  while (--argc > 0) {
    printf("ARGC %d\n", argc);
    type = getop(*++argv);
    printf("WHATS THE TYPE?: %d\n", type);
    switch (type) {
      case NUMBER:
        number = atof(*argv);
        printf("NUMBER: %f\n", number);
        push(number);
        break;
      case '+':
        push(pop() + pop());
        break;
      case 'x':
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
      default:
        break;
    }
  }

  printf("result: \t%.8g\n", pop());
  return 0;
}
