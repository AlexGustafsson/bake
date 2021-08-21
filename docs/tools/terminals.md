# Terminal tools

## Fish Autocompletion

Simply add the following to your config.

```fish
source (bake autocomplete --shell fish | psub)
```

You may also simply store the output of `bake autocomplete --shell fish` and source that file if you wish.
