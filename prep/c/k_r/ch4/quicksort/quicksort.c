#include <stdio.h>

void qsort(int arr[], int left, int right);

int main() {
  int i;
  int nums[] = {1, 4, 3, 6, 5, 8};
  qsort(nums, 0, 5);

  for (i = 0; i < 6; i++) {
    printf("%d\n", nums[i]);
  }
}

void qsort(int arr[], int left, int right) {
  int i, last;
  void swap(int[], int, int);

  if (left >= right) {
    return;
  }

  swap(arr, left, (left + right) / 2);
  last = left;

  for (i = left + 1; i <= right; i++) {
    if (arr[i] < arr[left]) {
      swap(arr, ++last, i);
    }
  }

  swap(arr, left, last);
  qsort(arr, left, last - 1);
  qsort(arr, last + 1, right);
}

void swap(int arr[], int i, int j) {
  int temp;
  temp = arr[i];
  arr[i] = arr[j];
  arr[j] = temp;
}
