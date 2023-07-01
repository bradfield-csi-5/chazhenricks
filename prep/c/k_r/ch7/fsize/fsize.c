#include <dirent.h>
#include <fcntl.h>
#include <math.h>
#include <stdio.h>
#include <string.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>
#include <errno.h>

void fsize(char *);

int main(int argc, char **argv) {
  // if no arguments passed in, assume we want to run on current directory
  if (argc == 1)
    fsize("./");
  else
    while (--argc > 0) fsize(*++argv);

  return 0;
}

void dirwalk(char *, void (*fcn)(char *));

void fsize(char *name) {
  struct stat stbuf;  // stat is the info retreived by passing the inode number
                      // into stat function


  //stat will return -1 if there is an error, and will dump star info into stbuf on success
  int err; 
  err = stat(name, &stbuf);

  if(err < 0){
    fprintf(stderr, "fsize: cant access %s with error %s\n", name, strerror(errno));
    return;
  }
  if((stbuf.st_mode & S_IFMT) == S_IFDIR)
    dirwalk(name, fsize);

  printf("%lld %s\n", stbuf.st_size, name);
}

#define MAX_PATH 1024


void dirwalk(char *dir, void (*fcn)(char *)){
  char name[MAX_PATH];
  struct dirent *dp; //directory entity 
  DIR *dfd; 

  if((dfd = opendir(dir)) == NULL){
    fprintf(stderr, "dirwalk cant open %s\n", dir);
    return;
  }

  while((dp = readdir(dfd)) != NULL){
    if(strcmp(dp->d_name, ".") == 0 || strcmp(dp->d_name, "..") == 0)
      continue;
    if(strlen(dir) + strlen(dp->d_name) + 2 > sizeof(name))
      fprintf(stderr, "dirwalk: name %s/%s is too long\n", dir, dp->d_name);
    else{
      sprintf(name, "%s%s", dir, dp->d_name);
      (*fcn)(name);
    }
  }
  closedir(dfd);
}
