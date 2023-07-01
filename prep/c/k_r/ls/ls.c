#include <dirent.h>
#include <errno.h>
#include <fcntl.h>
#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>
#include <wchar.h>

void standard_print(char *path, int *count, int type);
void list_contents(char *name, int hidden);

int main(int argc, char **argv) {
  int opt, hidden = 0;
  char *path = "./";

  while ((opt = getopt(argc, argv, "a")) != -1) {
    switch (opt) {
      case 'a':
        hidden = 1;
        break;
      default:
        printf("only -a now. Come back later\n");
        argc = -1;
        break;
    }
  }
  while (--argc > 0 && (*++argv)[0] == '-')
    ;

  while (--argc > 0) {
    path = *argv++;
  }
  list_contents(path, hidden);
  return 0;
}

void list_contents(char *name, int hidden) {
  int count = 0;
  DIR *dirp;
  struct dirent *dentp;
  int longest_filename = 10;

  if ((dirp = opendir(name)) == NULL) printf("oops. err opening %s\n", name);

  while ((dentp = readdir(dirp)) != NULL) {
    if (!hidden) {
      if (strcmp(dentp->d_name, ".") == 0 || strcmp(dentp->d_name, "..") == 0)
        continue;
    }
    standard_print(dentp->d_name, &count, dentp->d_type);
  }

  closedir(dirp);
}
#define GREEN "\x1B[32m"
#define NORMAL "\x1B[0m"

void standard_print(char *path, int *count, int type) {
  switch (type) {
    case DT_DIR:
      printf("%s %-10s", GREEN, path);
      printf("%s", NORMAL);
      break;
    case DT_REG:
    default:
      printf("%-10s", path);
      break;
  }

  // increment the number of things weve printed
  (*count)++;

  // max 4 per row
  if ((*count % 4) == 0) {
    printf("\n");
  }
}
