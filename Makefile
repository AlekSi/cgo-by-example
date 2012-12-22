all:
	$(CC) -g -c -o lib/lib.o lib/lib.c
	$(CC) -g -shared -o liblib.so lib/lib.o
	go run main.go
