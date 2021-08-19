In Make, there are phony rules. Such rules tell Make that a rule is not to be treated as a file. In Bake, phony does not exist. Instead, there are functions for tasks that do not have dependencies or aliases for other functions, rules or files.

## Alias rules

Here's a simple example of a phony build rule which ensures that some files are built.

```make
.PHONY build
build: build/my-app.exe build/my-app.dmg build/my-app
```

In Bake, one would instead use an alias. In order to make it accessible via the CLI, we also export it.

```bake
export alias build ["build/my-app.exe", "build/my-app.dmg", "build/my-app"]
```

## Functions

Here's an example of a phony rule which generates some API.

```make
.PHONY generate
generate:
  mkdir -p ./api
  api-generation-tool --input api.yml --output ./api
```

Such a rule in Bake would be called a function, as it does not care for dependencies.

```bake
export func generate {
  shell {
    mkdir -p ./api
    api-generation-tool --input api.yml --output ./api
  }
}
```
