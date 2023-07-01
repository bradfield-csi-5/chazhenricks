#include <stdio.h>
#include <stdlib.h>
int main(){
  int *ip;
  ip = (int *) malloc(sizeof(int));

  *ip = 10; 

  printf("ip: %d\n", *ip);
  free(ip);
  printf("ip: %d\n", *ip);

  int i = 10;
  free(&i);

}
