#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// int main(int argc, char **argv) {
//   int c;
//   int isLower = 1;
//
//
//   if (--argc > 0 && (*++argv)[0] == '-') {
//     int flag = *++argv[0];
//     switch (flag) {
//       case 'u':
//         isLower = 0;
//         break;
//       case 'l':
//         isLower = 1;
//         break;
//       default:
//         printf("Usage: ./case -u for uppercase");
//         argc = -1;
//         break;
//     }
//   }
//
//   printf("WHAT IS ARGC? %d\n", argc);
//   if (argc == -1) {
//     return -1;
//   } else {
//     while ((c = getchar()) != EOF) {
//       if (isLower) {
//         putchar(tolower(c));
//       } else {
//         putchar(toupper(c));
//       }
//     }
//   }
//   return 0;
// }
//

int main(int argc, char **argv) {
  int c;
  int Status = EXIT_SUCCESS;
  int func;
  int (*convert[2])(int) = {toupper, tolower};
  int(*compare)(const char *, const char *) = strcmp;


  if (argc > 0) {
    if (compare((*argv + 2), "CASE") == 0) {
      func = 0;
    } else {
      func = 1;
    }
  }

  while ((c = getchar()) != EOF) {
    putchar(convert[func](c));
  }

  return Status;
}
