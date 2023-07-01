#include <ctype.h>
#include <stdio.h>

#define SIZE 5

int getfloat(float *pn, FILE *stream);

int main() {
  float n; /* the thing were gonna dump the value into */ 
  int ret; /* return of the function */ 
  while (1) {
    ret = getfloat(&n, stdin); /* pass pointer of n and stdin to the functoin */ 
    printf("ret: %d, n: %f\n", ret, n);
  }
}

int getfloat(float *pn, FILE *stream) {
  int c, sign;
  float power;

  // skip whitespace
  while (isspace(c = getc(stream)))
    ;

  //if we get any weird values, return 0; 
  if (!isdigit(c) && c != EOF && c != '-' && c != '+' && c != '.') {
    ungetc(c, stream);
    return 0;
  }
 
  sign = (c == '-') ? -1 : 1;

  if (c == '+' || c == '-') c = getc(stream);

  if (!isdigit(c)) {
    ungetc(c, stream);
    return 0;
  }
  
  //we build up the left side of the . place like we did with the int version 
  for (*pn = 0; isdigit(c); c = getc(stream)) {
    //for each iteration, we're adding to the 10s spot. 
    // 123 
    // 1st 10 * 0 + 1  (1) 
    // 2nd 10 * 1 + 2 (12)
    // 3rd 10 * 12 + 3 (123) 
    *pn = 10.0 * *pn + (c - '0');
  }

  //if we hit a decimal, jump over it
  if (c == '.')
    c = getc(stream);

  //now were doing the right side of the decimal
  for (power = 1.0; isdigit(c); c = getc(stream)) {
    //we keep building up the number, but now keep track of the power. 
    *pn = 10 * *pn + (c - '0');
    power *= 10.0;
  }
 
  // divide the now large number by the number of decimal places 
  // ex 12.34 
  // *pn = 1234 
  // pow = 100 
  // 1234 / 100 = 12.34
  *pn /= power;
  *pn *= sign;

  if (c != EOF) {
    ungetc(c, stream);
  }
  return c;
}
