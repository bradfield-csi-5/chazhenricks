#include <stdio.h>
#include <ctype.h>

int main()
{
	int c, blanks; 
	blanks = 0;

	while((c = getchar()) != EOF){
	
		if(isspace(c)){
			blanks = blanks + 1;
			if(blanks <= 1){
				putchar(c);
			}
		}
		else{
			putchar(c);
			blanks = 0;
		}
	}
	return 0;
}












