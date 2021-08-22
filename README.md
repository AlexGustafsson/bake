<p align="center">
  <a href="https://github.com/AlexGustafsson/bake/blob/master/go.mod">
    <img src="https://shields.io/github/go-mod/go-version/AlexGustafsson/bake" alt="Go Version" />
  </a>
  <a href="https://github.com/AlexGustafsson/bake/releases">
    <img src="https://flat.badgen.net/github/release/AlexGustafsson/bake" alt="Latest Release" />
  </a>
  <br>
  <strong><a href="#quickstart">Quick Start</a> | <a href="https://alexgustafsson.github.io/bake">Documentation</a> | <a href="#contribute">Contribute</a></strong>
</p>

# Bake
### A cross-platform language and tool for building things - a better Make

⚠️ Bake is currently being actively developed. Until it reaches v1.0.0 breaking changes may occur in minor versions.

Bake is a new language and toolset for building things. Think of it like Make, with more tools to empower you to easily configure pragmatic builds of smaller projects. Make is a great tool which works wonders for smaller projects. It has a great, (mostly) readable syntax which is largely natural to work with. It does, however, suffer from some shortages. For example, it does not provide more advanced scripting features for more complex builds or configurations. Furthermore, it's not easily split into reusable models. Other common and arguably more complete tools such as Gradle, Mason etc. are not as simple and playful as a simple Makefile. The ambition of Bake is to become a better Make, building on the ideas, syntax and features provided by Make.

Bake has a couple of goals:

1. Provide a familiar syntax for Make users, but enable an even more general use with further scripting capabilities
2. Provide a uniform cross-platform experience without having to install Bake itself a la Gradle
3. Enable users to create libraries for common actions, easily imported a la Go
4. Be as fast as, or faster than Make

Bake also has a non-goal:

1. Be compatible with Make

## Quickstart
<a name="quickstart"></a>

⚠️ Bake is currently being actively developed. Not all features may be implemented.

_For a guide on how to get started, see the getting started in `docs/guide.md` or [here](https://alexgustafsson.github.io/bake/#/guide)._

First, install bake via a package manager like brew, or by downloading the latest release from GitHub.

```sh
brew install alexgustafsson/tap/bake
```

In this example we have a small C application we wish to compile. It consists of a library and a main executable.

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

To build the project, we create a `Bakefile` like the one below.

```
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
export func clean {
  shell rm lib.o lib.so main &>/dev/null || true
}
```

Lastly, running `bake` will execute the build. To execute a function such as `clean` or build a file, simply run `bake clean` or `bake lib.so`

## Table of contents

[Quickstart](#quickstart)<br/>
[Features](#features)<br />
[Contributing](#contributing)

<a id="features"></a>
## Features

Upcoming.

## Documentation

Please view [the docs on GitHub pages](https://alexgustafsson.github.io/bake) or see them in the docs folder.

To view them offline in your browser, first install docsify by executing `npm install --global docsify-cli` and then run `docsify serve docs` in the project's folder.

## Contributing
<a name="contributing"></a>

Any help with the project is more than welcome. The project is still in its infancy and not recommended for production.

For information on how to contribute, build and develop the project, please see the documentation in `docs/contribution.md` or [the docs on GitHub pages](https://alexgustafsson.github.io/bake/#/contribution).
