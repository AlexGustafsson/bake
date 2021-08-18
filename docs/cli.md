# CLI

```
Usage: bake [global options] command [command options] [arguments]

A cross-platform language and tool for building things

Version: v0.1.0, build 7a306aa. Built Wed Aug 18 10:52:20 CEST 2021 using go version go1.16.5 darwin/amd64

Options:
  --verbose   Enable verbose logging (default: false)
  --help, -h  show help (default: false)

Commands:
  autocomplete  Print autocompletion for a shell
  format        Format a Bakefile
  init          Initialize a bake
  lex           Lex a Bakefile
  parse         Parse a Bakefile
  version       Show the application's version
  help          Shows a list of commands or help for one command

Run 'bake help command' for more information on a command.
```

## Autocomplete

```
Usage: bake autocomplete [options]

Print autocompletion for a shell

Options:
   --shell value  Name of the shell to print an autocompletion script for
```

## Format

```
Usage: bake format [options]

Format a Bakefile

Options:
   --input value, -i value  Path to input file
```

## Init

```
Usage: bake init [options]

Initialize a bake
```

## Lex

```
Usage: bake lex [options]

Lex a Bakefile

Options:
   --input value, -i value  Path to input file
```

## Parse

```
Usage: bake parse [options]

Parse a Bakefile

Options:
   --input value, -i value  Path to input file
   --type value, -t value   Output type (default: dot)
```

## Version

```
Usage: bake version [options]

Show the application's version
```

## Help

```
Usage: bake help [options] [command]

Shows a list of commands or help for one command

Options:
   --help, -h  show help (default: false)
```
