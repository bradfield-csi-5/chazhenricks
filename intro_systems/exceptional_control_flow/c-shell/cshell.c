#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/errno.h>
#include <unistd.h>

#define MAXARGS 128
#define MAXLINE 256

extern char **environ;

/* func prototypes */
void eval(char *cmdline);
int parseline(char *buf, char **argv);
int builtin_command(char **argv);
char *Fgets(char *ptr, int n, FILE *stream);
pid_t Fork(void);
void unix_error(char *msg);
void Fuck();
void Exit();

int main(int argc, char *argv[]) {
  char cmdline[MAXLINE];

  while (1) {
    printf("🐚 ");
    Fgets(cmdline, MAXLINE, stdin);
    if (feof(stdin)) {
      Exit();
    }

    eval(cmdline);
  }
  return 0;
}

/* -----------------
 * Shell Functions
 * -----------------
 */

/* eval is the function that will check each user input
 * will run through parselines to break up into an argv char array for
 * command/args will also check if command is a builtin before forking a child
 * process and exec the command passed in
 */
void eval(char *cmdline) {
  char *argv[MAXARGS];
  char buf[MAXLINE];
  int bg;
  pid_t pid;

  char **env = environ;
  // copy the user command passed in to a buffer we can mess with
  strcpy(buf, cmdline);
  // pass buf and empty argv to parse line to build up argv and check for
  // background tasks
  bg = parseline(buf, argv);

  // nothing passed in
  if (argv[0] == NULL) {
    return;
  }

  if (!builtin_command(argv)) {
    if ((pid = Fork()) == 0) {
      if (execvp(argv[0], argv) < 0) {
        // TODO
        // need to close the forked process if it fails before returning
        Fuck();
        printf("%s: command not found.\n", argv[0]);
        return;
      }
    }
  }

  if (!bg) {
    int status;
    if (waitpid(pid, &status, 0) < 0) {
      unix_error("waitfg: waitpid error");
    }
  }
  return;
}

// check for builtin commands
// will be a simple "if" check - possibly building out into a switch if that
// makes sense
//  this gets checked before fork/exec is run
//  returning 0 means not a builtin
//  returning 1 means it is a builtin
int builtin_command(char **argv) {
  if (!strcmp(argv[0], "farts")) {
    Exit();
  }
  if (!strcmp(argv[0], "exit")) {
    Exit();
  }
  if (!strcmp(argv[0], "quit")) {
    Exit();
  }

  if (!strcmp(argv[0], "&")) {
    // if we return 1, it is a builtin command
    return 1;
  }

  return 0;  // if we return 0 then its not a builtin command
}

/* parses the command line and builds argv array
 * the **argv passed in is empty and we are building it up to look like
 * what would be passed in to a command line program
 * this is one of those dual-return type functions:
 * we are building up argv as a pointer passed in, but returning a value for bg
 * to use in the main shell loop.
 * */

int parseline(char *buf, char **argv) {
  char *delim; /* points to first space delimiter */
  int argc;    /* number of args */
  int bg;      /* is this a background job? */

  buf[strlen(buf) - 1] = ' '; /* removes \n at end */

  // skip over any leading whitespace
  while (*buf && (*buf == ' ')) {
    buf++;
  }

  // build args list
  argc = 0;

  // strchr will search a string for a char and return the index.
  // we are searching the buf for the next space
  // that will indicate our delimeter
  while ((delim = strchr(buf, ' '))) {
    argv[argc++] = buf;  // we start argc at 0 so this will begin at argv[1]
    *delim = '\0';    // set the first space equal to \0, effectively ending the
                      // string there
    buf = delim + 1;  // set the buffer equal to one past the first space,
                      // meaning the next item in the arg array
    while (*buf && (*buf == ' ')) {
      buf++;
    }
  }
  // what we have effectively done is set argv[n] = to the buffer, but any time
  // there is a space, replace it with a null terminator and set the rest of the
  // buffer equal to the rest of the string we do this in a loop until the
  // string ls -a ends up like argv[1] = ls\0 argv[2] = -a\0

  argv[argc] = NULL;  // terminate the end of the string.

  if (argc == 0) {
    return 1;
  }

  // check for background jobs
  // sets bg = to the result of comparing the last item in argv list to &
  if ((bg = (*argv[argc - 1] == '&')) != 0) {
    // if it is, then the job should run in the background, so lets remove it
    // from our argv list
    argv[--argc] = NULL;
  }

  // be is set by the compare above, so it will either be:
  // 0: meaning no background jobs or
  // 1: there are background jobs
  return bg;
}

/* -----------------
 * Utility Functions
 * -----------------
 */

// error convient error message
void unix_error(char *msg) {
  fprintf(stderr, "%s: %s\n", msg, strerror(errno));
}

// creates a forked child process and retuns the PID
pid_t Fork(void) {
  pid_t pid;

  if ((pid = fork()) < 0) {
    unix_error("Fork error");
  }
  return pid;
}

char *Fgets(char *ptr, int n, FILE *stream) {
  char *rptr;

  if (((rptr = fgets(ptr, n, stream)) == NULL) && ferror(stream))
    unix_error("Fgets error");

  return rptr;
}

void Fuck() {
  printf("\n");
  printf("███████╗██╗   ██╗ ██████╗██╗  ██╗\n");
  printf("██╔════╝██║   ██║██╔════╝██║ ██╔╝\n");
  printf("█████╗  ██║   ██║██║     █████╔╝ \n");
  printf("██╔══╝  ██║   ██║██║     ██╔═██╗ \n");
  printf("██║     ╚██████╔╝╚██████╗██║  ██╗\n");
}

void Exit() {
  printf("\n");
  printf("💩Smell ya later💩\n");
  exit(0);
}
