// For C.free in Go and malloc in C.
#include <stdlib.h>

void f1(int a);

char *f2(char *s);

struct s1 {
	int a;
};

struct s1 f3(struct s1 s);

struct s1 *f4(struct s1 *s);
