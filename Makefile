all:
	gofmt -e -s -w .
	go vet .
	$(CC) -g -fPIC -c -o lib/lib.o lib/lib.c
	$(CC) -g -fPIC -shared -o liblib.so lib/lib.o
	LD_LIBRARY_PATH=. go run main.go
	cp liblib.so wrapper/
	LD_LIBRARY_PATH=. go run main2.go
