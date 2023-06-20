#include <stdio.h>

// char *s is the same thing as char s[]
// an array name is the same as a pointer to the first item in the array

int my_strlen(char *s) {
  int n;

  // while the pointer does not equal a null character, increment it
  // this is the same as i = 0; s[i] != \0; i++
  // we can increment a pointer and it will refer to the next item in the array
  //
  for (n = 0; *s != '\0'; s++) {
    n++;
    // a pointer is a memory address, which is a numerical value.
    // incrementing the memory address by one goes to the next value.
    // in an array, memory is allocated sequentially, were just going to the
    // next memory address
    printf("WHAT IS *s?: %p\n", s);
  }

  return n;
}

int another_strlen(char *s) {
  char *p = s;
  while (*p != '\0') p++;

  //this works because s is the _beginning_ of the array, and p is the end of the array. 
  //since \0 is the last character of the array, if we increment p until the end, the difference between p and s will be the size
  //of the array. 
  return p - s;
}

int main() {
  int len;
  char thing[] = "I am a cool boy";
  len = my_strlen(thing);
  printf("LENGTH: %d\n", len);
}
