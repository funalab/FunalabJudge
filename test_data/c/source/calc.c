#include<stdio.h>

void calc(int);
void print(int);

void calc(int n) {
	int i;
	int sum = 0;
	for (i = 1; i <= n; i++) {
		sum += i;
	}
	print(sum);
}
