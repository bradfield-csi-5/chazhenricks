#include <stdio.h>

// we need to check if the string t occurs at the end
// of string s
int strend(char *s, char *t) {
  // 1. start comparing chars
  // 2. if they match keep incrementing both
  // 3. if they dont match, just keep incrementing s
  // 4. if we hit the end of s and the end of t at the same time, return 1
  // 5. if any other situation occurs, return 0

  // start incrementing through s
  while (*s != '\0') {
    printf("BEGINING OF LOOP: *S IS: %c\n", *s);
    if (*s == *t) {
      s++;
      t++;
      if (*s == '\0' && *t == '\0') {
        printf("YAY WERE BOTH NULL");
        printf("*s IS: %c - *t IS: %c\n", *s, *t);
        return 1;
      }
    } else
      s++;
  }
  return 0;
}

//oz's solution from github
int oz_strend(char *s, char *t) {

  //create two new pointers from s and t. 
  //were gonne traverse these through to get to the end of each one 
  char *se = s, *te = t;
  //increment se until we get to a null character
  while (*se++)
    ;
  //same as above, increment until we get to null
  while (*te++)
    ;

  //were going to now decrement both se and te 
  //our check here is to make sure we havn't reached the begining of each word
  while (se > s && te > t)
    //if the two strings ever dont match while they both exist it means
    //out condition isn't true and we should return 0
    //(meaning as long as t has characters, they should match s)
    
    if (*--se != *--te) return 0;
  //if we make it through that loop without returning, that means that the entirety of
  //the string t exists at the end of the string s 
  return 1;
}

int main() {
  int ret;
  char s[] = "this thing";
  char t[] = "thing";

  ret = strend(s, t);
  printf("DID \"%s\" OCCUR AT THE END OF \"%s\"?: %d\n", t, s, ret);
}
