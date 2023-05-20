#include <stdio.h>
#include <math.h>

int main()
{
  int orig, n, digit;
  int count, running_total; 
  running_total = count = 0;
  printf("Enter a number: ");
  scanf("%d", &n);
  orig = n;

  /* get length of input */ 
  while(n != 0)
  {
    n = n/10;
    count++;
  }

  n = orig 

  /* calculate running_total */ 
  while(n != 0)
  {
    digit = n % 10; 
    digit = pow(digit, count);
    running_total += digit;
    n = n / 10;
  }

  return running_total == orig;
  

}

