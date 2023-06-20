#include <stdio.h>
#include <string.h>

void reverse(char s[]) {
  int c, i, j;
  for (i = 0, j = strlen(s) - 1; i < j; i++, j--) {
    c = s[i];
    s[i] = s[j];
    s[j] = c;
  }
}

void recursive_reverse(char s[], int start, int end) {
  if (start >= end) {
    return;
  }

  int c;
  c = s[start];
  s[start] = s[end];
  s[end] = c;

  start++;
  end--;
  recursive_reverse(s, start, end);
}

void itoa(int n, char s[]) {
  int i, sign;

  if ((sign = n) < 0) /* record sign */
    n = -n;

  i = 0;

  do {                     /* generate digits in reverse order */
    s[i++] = n % 10 + '0'; /* get next digit */
  } while ((n /= 10) > 0);

  if (sign < 0) s[i++] = '-';
  s[i] = '\0';
  reverse(s);
}

void recursive_itoa(int n, char s[]) {
  static int i = 0; /* making this static lets us not have to pass it through each time */
  if (n) {
    if (n < 0) {
      s[i++] = '-';
      n = -n; /* make sure rest of the recursive calls deal with a positive
               * number
               */
    }
    int d = n % 10;
    n /= 10;

    recursive_itoa(n, s);

    s[i++] = d + '0';
    s[i] = '\0';
  }
}

int main(void) {
  char name[] = "chaz";
  recursive_reverse(name, 0, strlen(name) - 1);

  int n = -15;
  char nums[10];
  int i = 0;
  recursive_itoa(n, nums);
  printf("ITOA: %s\n", nums);
  printf("%s\n", name);
}
