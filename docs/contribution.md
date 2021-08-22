# Contributing

Any help with the project is more than welcome. The project is still in its infancy and not recommended for production.

## Project structure

### Preface

The project is structured as a multi-binary Go project.

The tool and grammar is inspired from several projects, mostly Go (grammar, lexer implementation), JavaScript (grammar), TypeScript (parser implementation), Makefile (grammar, features) and Docker (cli).

### Packages

| Name | Description |
| :--: | :---------: |
| `internal` | Internal code supporting the binaries, such as build information in `version` or Grapviz dot formatting of parse trees in `dot`. |
| `lexing` | The zero-dependency bake lexer. Takes care of handling an input source, providing a stream of tokens. |
| `parsing` | The bake parser. Uses `lexing` and forms a parse tree from `parsing/nodes`. |
| `semantics` | The bake semantic analysis. Traverses a parse tree from `parsing` and provides a semantically correct program. |
| `lsp` | The core of bagels - the bake language server. Uses `lexing`, `parsing` and `semantics` to provide all facets of a language server. |

Any code that is of shared interest for the community is kept public and supported as such. Private code kept in the `internal` package is not supported for third-party. The binaries that this prje

### Binaries

| Name | Description |
| :--: | :---------: |
| `bake` | The core bake parser, lexer, runtime and tool. Uses `lexing`, `parsing`, `semantics` and internal packages. |
| `bagels` | The language server for bake. Uses `lexing`, `parsing`, `semantics`, `lsp` and internal packages. |

### Bake resources

Bake documentation and resources are placed in the `docs` directory, the `examples` directory as well as the `stdlib` directory. As Bake supports imports, the top-level `stdlib` directory is intended for usage in Bake programs.

### Tools

The tools directory contains supported tooling, such as the VSCode plugin, nano syntax support and PrismJS grammar.

See the [editor](tools/editors.md) and [terminal](terminals/terminals.md) documentation for more information.

### Syntax

The syntax is documented [here](grammar.md). The grammar is kept up-to-date with the parser and lexer.

## Development

```shell
# Clone the repository
https://github.com/AlexGustafsson/upmon.git && cd upmon

# Show available commands
make help

## Building

# Build the project
make build

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
