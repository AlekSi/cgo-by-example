all:
	$(CC) -c -o lib/lib.o lib/lib.c
	$(CC) -shared -o liblib.so lib/lib.o
	go run main.go
