/*  industryminutes.c
    
    Calculate time span from given start and end point 
    in industry minute format.
    
    marthjod@gmail.com 2013-02-10

*/

#include <stdio.h>

int main(int argc, char* argv[]) {

	int start = 0;
	int end = 0;
	int hours = 0;
	float minutes_frac = 0.0;

    // no input checks
	if (argc < 3) {	
		printf("Start (HHmm): ");
		scanf("%i", &start);
		printf("End (HHmm): ");	
		scanf("%i", &end);
	} else {
		start = atoi(argv[1]);
		end = atoi(argv[2]);
	}
	
	hours = (end - start)/100;
	minutes_frac = (float) ((end - start) % 100) / 60;

	printf("%.2f\n", hours + minutes_frac);
	
	return 0;
}
