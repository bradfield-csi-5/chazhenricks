#include <stdio.h>

// old version of strcat for reference

void arr_strcat(char s[], char t[]) {
  int i, j;
  // initialize index counters at zero
  i = j = 0;

  // first we need to get to the end of s
  // we traverse through s until we get to a null character
  while (s[i] != '\0') i++;

  // now that i is incremented until the end of s
  // we can start adding chars from the bginning of t onto the end of s
  while ((s[i++] = t[j++]) != '\0')
    ;
}

void pointer_strcat(char *s, char *t) {
  //increment the pointer until we reeach the end of the string
  while (*(++s))
    ;

  // copy begining of t onto the end of s
  //  will take the pre incremented pointer value of t
  //  and set it to the pre incremented value of s
  //  when *t = \0 the condition will evaluate to false and be done
  while ((*s++ = *t++))
    ;
}

int main() {
  char first[10] = "cool";
  char second[] = "man";

  pointer_strcat(first, second);
  printf("I AM A STRING: %s\n", first);
}
