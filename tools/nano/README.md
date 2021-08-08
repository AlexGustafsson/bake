<p align="center">
  <img src="assets/logo-240x240.png" alt="Logo">
</p>
<p align="center">
  <a href="https://github.com/AlexGustafsson/bake/blob/master/go.mod">
    <img src="https://shields.io/github/go-mod/go-version/AlexGustafsson/bake" alt="Go Version" />
  </a>
  <a href="https://github.com/AlexGustafsson/bake/releases">
    <img src="https://flat.badgen.net/github/release/AlexGustafsson/bake" alt="Latest Release" />
  </a>
  <br>
  <strong><a href="#quickstart">Quick Start</a> | <a href="#contribute">Contribute</a> </strong>
</p>

# Bake configuration for nano
### A cross-platform language and tool for building things - a better Make

Note: Bake is currently being actively developed. Until it reaches v1.0.0 breaking changes may occur in minor versions.

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
[Installation](#installation)<br />
[Usage](#usage)<br />
[Contributing](#contributing)

<a id="features"></a>
## Features

* Syntax highlighting

<a id="installation"></a>
## Installation

### macOS

Make sure you've installed `nano` using `brew`. If you've just installed it, you may need to restart your terminal.

Add the following line to your `~/.nanorc`.

```js
include "/usr/local/share/nano/*.nanorc"
```

Now, download and copy `bake.nanorc` to the `/usr/local/share/nano` directory.

```sh
curl "https://raw.githubusercontent.com/AlexGustafsson/bake/main/tools/nano/bake.nanorc" > "/usr/local/share/nano/bake.nanorc"
```

## Contributing
<a name="contributing"></a>

Any help with the project is more than welcome. The project is still in its infancy and not recommended for production.