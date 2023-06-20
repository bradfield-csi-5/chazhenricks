
#include <stdio.h>

void escape_char(char[], char[]);
int getlength(char[]);

int main() {
  char stuff[] = "this\nis\na cool thing";
  char other[1000] = {'\0'};

  escape_char(stuff, other);

  printf("STUFF: %s\n", stuff);
  printf("OTHER: %s\n", other);
  return 0;
}

int getlength(char str[]) {
  int i;
  i = 0;
  while (str[i] != '\0') {
    ++i;
  }
  return i;
}

void escape_char(char from[], char to[]) {

  int i, j;
  i = j = 0;
  while (from[i]) {
    switch (from[i]) {
    case '\n':
      to[j] = '\\';
      j++;
      to[j] = 'n';
      break;
    default:
      to[j] = from[i];
      break;
    }
    i++;
    j++;
  }
  to[j] = '\0';
}
