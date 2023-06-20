#include <limits.h>
#include <stdio.h>
#include <string.h>

int main() {
  printf("MIN SIGNED CHAR LIMIT: %d\n", SCHAR_MIN);
  printf("MAX SIGNED CHAR LIMIT: %d\n", SCHAR_MAX);
  printf("MIN UNSIGNED CHAR LIMIT: %d\n", CHAR_MIN);

  char s[10] = {'\0'};
  int len = strlen(s);
  printf("LENGTH: %d\n", len);

  printf("SINGLE QUOTES: %d\n", 'x');
  printf("DOUBLE QUOTES: %s\n", "x");

  enum escapes {
    BELL = '\a',
    BACKSPACE = '\b',
    TAB = '\t',
    NEWLINE = '\n',
    VTAB = '\v',
    RETURN = '\r'
  };

  enum escapes item;
  item = BELL;
  printf("ENUM 0: %d\n", item);
  printf("NEWLINE: %d\n", NEWLINE);

  enum Day {
    MONDAY = 123,
    TUESDAY,
    WEDNESDAY,
    THURSDAY,
    FRIDAY,
    SATURDAY,
    SUNDAY
  };
  enum Day today;

  today = MONDAY;

  printf("TODAY: %d\n", today);

  return 0;
}
