#include <stdio.h>

int main()
{
	int c, nl, blank, tab;
	nl = 0;
	tab = 0;
	blank = 0;

	while((c = getchar()) != EOF) {
		if (c == '\n')
			++nl;
		if (c == '\t')
			++tab;
		if(c == ' ')
			++blank;
	}
	printf("%s", "new lines - ");
	printf("%d\n", nl);

	printf("%s", "tabs - ");
	printf("%d\n", tab);

	printf("%s", "blanks - ");
}
