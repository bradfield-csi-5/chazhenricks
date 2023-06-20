#include <stdio.h>
#include <string.h>

int main() {

  char a = 0101;

  printf("a = %d\n", a << 1);

  unsigned char j = 255;
  j = j + 10;
  printf("%u\n", j);

  char k = 127;
  k = k + 10;
  printf("%u\n", k);

  int prices[] = {1, 2, 3, 4, 5};
  for (int i = 0; i < 5; i++) {
    printf("%d\n", prices[i]);
  }

  int stuff[5];
  for (int i = 0; i < 5; i++) {
    printf("STUFF: %d\n", stuff[i]);
  }

  char stringtest[] = "fuck";
  printf("STRINGTEST: %s\n", stringtest);
  printf("LAST CHAR: %c\n", stringtest[3]);
  printf("LENGTH: %lu\n", strlen(stringtest));

  /* pointer shit */
  int age = 33;
  printf("AGE POINTER: %p\n", &age);

  int *address = &age;
  printf("VALUE OF VARIABLE ADDRESS: %p\n", address);
  printf("VALUE OF POINTER OF ADDRESS: %u\n", *address);

  int butts;
  int *buttaddress = &butts;
  *buttaddress = 37;
  printf("BUTTS: %d\n", butts);

  char neat[] = "neat stuff";
  char *neat_address = neat;
  *neat_address = 'b';
  printf("NEAT POINTER: %p\n", neat_address);
  printf("STRING: %s\n", neat);
  printf("FIRST CHASR: %c\n", *neat);
  printf("SECOND CHASR: %c\n", *(neat + 1));






  char name[10];
  printf("WHAT IS IT THO: %d\n", name[0]);

  return 0;
}
