#include <stdio.h>

int main(void) {
  int int_arr[] = {1, 2, 3, 4, 5, 6};
  int *item = int_arr;
  int *item2 = int_arr + 1;
  printf("item: %d - address:%p\n", *item, item);  // item: 1 - address:0x16b062d50
  printf("size of *item is %lu\n", sizeof(*item));  // 4 bytes
  printf("item2: %d - address:%p\n", *item, item);  // item: 2 - address:0x16b062d54
  printf("difference between item and item2: %ld\n", item2 - item);  // 1
}
