# Guide

## Installation

### Using Docker

```shell
git clone https://github.com/AlexGustafsson/bake.git && cd bake
docker build -t bake .
docker run -it bake help
```

### Using Homebrew

Upcoming.

```shell
brew install alexgustafsson/tap/bake
```

### Downloading a pre-built release

Download the latest release from [here](https://github.com/AlexGustafsson/bake/releases).

### Build from source

Clone the repository.

```shell
git clone https://github.com/AlexGustafsson/bake.git && cd bake
```

Optionally check out a specific version.

```shell
git checkout v0.1.0
```

Build the application.

```shell
make build
```

## Creating your first build

In this example we have a small C application we wish to compile. It consists of a library and a main executable. Below you may find the source code for the library.
```c
// lib.h

#ifndef LIB_H
#define LIB_H

void hello_world();

#endif
```

```c
// lib.c

#include <stdio.h>
#include "lib.h"

void hello_world() {
  printf("Hello, world!\n");
}
```

The source for the main entrypoint of the application is shown below.

```c
// main.c

#include "lib.h"

void main() {
  hello_world();
}
```

In order to build this, we first need to compile the library so that we may link it alongside the object for the entrypoint. This functionality is part of the standard library, so we may include it.

```bake
// Bakefile

// Import the standard libray
import (
  "github.com/AlexGustafsson/bake/stdlib/c"
)
```

We may then use that functionality to build `lib.so`.

```bake
// Bakefile

// Build a static library for lib.c using standard library functions
"lib.so" ["lib.c", "lib.h"] : c::static_library
```

We can then write a rule for compiling and linking the main executable.

```bake
// Bakefile

// Build the main application
"main" ["main.c", "lib.so"] {
  shell gcc -o $@ $<
}
````

Lastly, for convenience we can add a `clean` function to remove all of the compiled files.

```bake
// Bakefile

// Delete any output files
func clean {
  shell rm lib.o lib.so main &>/dev/null || true
}
````

Our final build script looks like this.

```bake
// Bakefile

// Import the standard libray
import (
  "github.com/AlexGustafsson/bake/stdlib/c"
)

// Build a static library for lib.c using standard library functions
"lib.so" ["lib.c", "lib.h"] : c::static_library

// Build the main application
"main" ["main.c", "lib.so"] {
  shell gcc -o $@ $<
}

// Delete any output files
func clean {
  shell rm lib.o lib.so main &>/dev/null || true
}
```
