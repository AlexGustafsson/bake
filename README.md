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

Upcoming.

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

### Development

```sh
# Clone the repository
https://github.com/AlexGustafsson/bake.git && cd bake


# Build the project for the native target
# Yes, this will eventually be bootstrapped
make build
```

The project is written in Go. Its source is scattered in the `cmd` and `internal` directories.

```sh
## Building

# Build the project
make build/bake

# Build tools
make tools

## Code quality

# Format code
make format
# Lint code
make lint
# Vet the code
make vet

## Testing

# Run tests
make test
```

_Note: due to a bug (https://gcc.gnu.org/bugzilla/show_bug.cgi?id=93082, https://bugs.llvm.org/show_bug.cgi?id=44406, https://openradar.appspot.com/radar?id=4952611266494464), clang is required when building for macOS. GCC cannot be used._
