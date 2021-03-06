# Editor plugins

## VSCode, VSCodium and Friends

### Features

* Syntax highlighting
* Syntax highlighting injection in Markdown
* Bagels language server support
* Snippets

### Installation

```shell
make build vscode
code --install-extension "build/bake-lsp-x.x.x.vsix"
```

## Nano

### Features

* Syntax highlighting

### Installation

#### macOS

Make sure you've installed `nano` using `brew`. If you've just installed it, you may need to restart your terminal.

Add the following line to your `~/.nanorc`.

```
include "/usr/local/share/nano/*.nanorc"
```

Now, download and copy `bake.nanorc` to the `/usr/local/share/nano` directory.

```shell
curl "https://raw.githubusercontent.com/AlexGustafsson/bake/main/tools/nano/bake.nanorc" > "/usr/local/share/nano/bake.nanorc"
```
