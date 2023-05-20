#include <stdio.h>

#define IN  1
#define OUT 0

int main() {
  int c, line, word, chars, state;
  line = word = chars = 0;
  state = OUT;
  while((c = getchar()) != EOF){
    ++chars;
    if(c == '\n')
      ++line;
    if(c == ' ' || c == '\n' || c == '\t')
      state = OUT;
    else if(state == OUT){
      state = IN;
      ++word;
    }
  }
  printf("LINE WORD CHARS\n");
  printf("%4d %4d %4d\n", line, word, chars);
}
