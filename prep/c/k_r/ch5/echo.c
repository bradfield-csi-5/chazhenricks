#include <stdio.h>

// first version, with array indexing

// int main(int argc, char *argv[]) {
//   int i;
//   //while i is less than the number of arguments passed in
//   //print the argument
//   //if it is not the last argument, print a space afterward, if it is the
//   last, print an empty string for (i = 1; i < argc; i++) printf("%s%s",
//   argv[i], (i < argc - 1) ? " " : ""); printf("\n");
//
//   return 0;
// }

int main(int argc, char *argv[]) {
  //count is going to be at least 1 if not more
  //while the count is greater than 0, we increase the argv pointer and print it 
  //we also do the same check if it is the last argument and printing the space or not  
  while (--argc > 0) printf("%s%s", *++argv, (argc > 1) ? " " : "");
  printf("\n");
  return 0; 
}
