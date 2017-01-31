build:
	rm -f helloworld.dylib
	gcc -c -fPIC helloworld.c
	gcc -dynamiclib -o libhelloworld.dylib helloworld.o
	rm -f helloworld.o