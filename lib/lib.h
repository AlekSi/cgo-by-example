void f0(void);

void f1(int a);

char *f2(char *s);


struct s1 {
	int a;
};

struct s1 f31(struct s1 s);

struct s1 *f32(struct s1 *s);


struct s2 {
	int *p;
};

void f4(struct s2 s);


void f5(void (*f)(int));
