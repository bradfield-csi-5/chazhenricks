#include <stdio.h>

int binsearch(int, int[], int);
int binsearch_singlecheck(int, int[], int);

int main(void) {
  int stuff[] = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9};
  int result = binsearch(5, stuff, 10);
  printf("RESULT %d\n", result);

  int result_single = binsearch_singlecheck(5, stuff, 9);
  printf("RESULT %d\n", result_single);
  return 0;
}

int binsearch(int target, int searchable[], int length) {
  int low, high, mid;

  low = 0;
  high = length - 1;

  while (low <= high) {
    mid = (low + high) / 2;
    /* if our target is less than the midpoint */
    if (target < searchable[mid]) {
      high = mid; /* make the high end the midpoint */
                      /* if the target is bigger than the mid point */
    } else if (target > searchable[mid]) {
      low = mid + 1; /* make the low end the mid point */
    } else {
      return mid; /* we found it */
    }
  }
  return -1; /* no match */
}
int binsearch_singlecheck(int target, int searchable[], int length) {
  int low, high, mid;

  low = 0;
  high = length - 1;

  while (low < high) {
    printf("LOW: %d\n", low);
    printf("HIGH: %d\n", high);
    mid = (low + high) / 2;
    printf("MID: %d\n", mid);
    /* if our target is less than the midpoint */
    if (target < searchable[mid]) {
      high = mid; /* make the high end the midpoint */
                  /* if the target is bigger than the mid point */
    } else {
      low = mid + 1;
    }
  }

  if (mid == target) {
    return mid;
  }
  return -1; /* no match */
}
