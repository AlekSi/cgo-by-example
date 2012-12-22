#include "lib.h"

#include <stdio.h>

void f1(int a) {
	printf("f1(%d)\n", a);
}

char *f2(char *s) {
	printf("f2(%s)\n", s);

	char *buf = malloc(100);
	sprintf(buf, "return f2(%s)", s);
	return buf;
}
