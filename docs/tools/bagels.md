# Bagels

Bagels is the official bake language server. The name comes from "bake language server", or "BakeLS" which conveniently sounds a bit like "bagels".

## Features

* Fast reporting of parse errors on every keystroke

## Installation

After installing bagels using one of the alternatives below, follow your editor's implementation details on using `build/bagels` with communication over JSON RPC v2 over the standard streams.

The [VSCode plugin](tools/editors.md) supports the language server natively.

### Using Homebrew

Upcoming.

```shell
brew install alexgustafsson/tap/bagels
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
make build/bagels
```
