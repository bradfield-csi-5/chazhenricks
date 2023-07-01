#include <stdio.h>
#include <string.h>

#define MAXLINE 1000 /* max line input */

int pointer_getline(char *s, int lim) {
  int c, i;

  for (i = 0; i < lim - 1 && (c = getchar()) != EOF && c != '\n'; i++) {
    s[i] = c;
  }

  if (c == '\n') {
    s[i] = c;
    ++i;
  }

  s[i] = '\0';
  return i;
}

int main(int argc, char *argv[]) {
  char line[MAXLINE];
  long lineno = 0;
  int c, except = 0, number = 0, found = 0;

  // we decrement the argc num by one as we verify that args are passed in
  //(*++argv)[0] increases the argv pointer to be argv[1], which is the first
  // argument argv[0] is the name of the program the first char in argv[1]
  // should
  // be a - if its an argument
  while (--argc > 0 && (*++argv)[0] == '-')
    // this will increment the actual argument char
    // so were looking at the first char not -
    while ((c = *++argv[0])) switch (c) {
        case 'x':
          except = 1;
          break;
        case 'n':
          number = 1;
          break;
        default:
          printf("illegal operation: get fucked \n");
          argc = 0;
          found = -1; 
          break;
      }
  
  //argc should have been incremented down to 1 at this point 
  //argv should also now be pointed to the pattern to match 
  if(argc != 1)
    printf("Usage: find -x -n pattern \n");
  else
    while (pointer_getline(line, MAXLINE) > 0) {
      lineno++;
      if ((strstr(line, *argv)!= NULL) != except) {
        if(number)
          printf("%ld ", lineno); /* print without newline to tak on to the begining */ 
        printf("%s\n", line);
        found++;
      }
    }
  return found;
}
