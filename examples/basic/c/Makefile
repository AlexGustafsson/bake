# Make sure Make doesn't interpret build as a path
.PHONY: build

# Default to building the main application
build: main

# Build the object files for the library
lib.o: lib.c lib.h
	gcc -c -o $@ $<

# Create the library archive on Darwin and Linux
lib.so: lib.o
ifeq ($(shell uname),Darwin)
	libtool -static -o $@ $^
else
	ar rcs $@ $^
endif

# Build the main application
main: main.c lib.so
	gcc -o $@ $^

# Delete any output files
clean:
	rm lib.o lib.so main &>/dev/null || true
