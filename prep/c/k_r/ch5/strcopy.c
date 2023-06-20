// array version of copy from one string to another
void arr_strcopy(char *s, char *t) {
  int i;
  i = 0;
  while ((s[i] = t[i]) != '\0') i++;
}

// another version with pointers
//

void point_strcopy(char *s, char *t) {
  // both s and t are fetched first, s is set to t, that is compared to the \0
  // value, _then_ both are incremented doing this will also include the
  // terminating \0
  // writing this way is itomatic, and in reality we are just checking to see if both values = \0, which is actually just 0;
  while (*s++ = *t++)
    ;
}
