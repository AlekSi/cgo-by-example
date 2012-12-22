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

struct s1 f31(struct s1 s) {
	s.a *= 2;
	return s;
}

struct s1 *f32(struct s1 *s) {
	s->a *= 2;
	return s;
}
