#include <stdio.h>

int main()
{
        int total = 0;
                for(int i=0; i<999; i++){
                        for(int j=1; j<999; j++){
                                total += (i / j);
                        }
                }

	printf("%d \n", total);
        return 0;
}
