#include <stdio.h>

static char daytab[2][13] = {
    {0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
    {0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
};

int day_of_year(int year, int month, int day) {
  int i, leap;

  // determine if its a leap year
  leap = year % 4 == 0 && year % 100 != 0 || year % 400 == 0;
  for (i = 0; i < month; i++)
    day += daytab[leap][i];  // add in each previous month, not including the current
  return day;
}

void month_day(int year, int yearday, int *pmonth, int *pday) {
  int i, leap;
  leap = year % 4 == 0 && year % 100 != 0 || year % 400 == 0;

  for (i = 1; yearday > daytab[leap][i]; i++) yearday -= daytab[leap][i];
  *pmonth = i;
  *pday = yearday;
}


int main() {
  int m, d;
  int yearday;

  yearday = day_of_year(2023, 6, 22);
  month_day(2023, yearday, &m, &d);

  printf("TODAY IS THE %d DAY OF THE YEAR\n", yearday);
  printf("TODAY IS THE %d OF %d\n", d, m);

  return 0;
}

char *month_name(int n) {
  static char *name[] = {"Illegal month", "January",   "February", "March",
                         "April",         "May",       "June",     "July",
                         "August",        "September", "October",  "November",
                         "December"};
  return (n < 1 || n > 12) ? name[0] : name[n];
}
