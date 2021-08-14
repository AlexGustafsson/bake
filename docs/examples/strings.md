Strings come in two different forms - interpreted strings and raw strings.

Interpreted strings are written using double quotes and may contain escaped characters, variable substitutions and more.

```bake
let a = "standard string"
let b = "my non-$(a)"
let c = "multi-\nline-\nstring"
let d = "two plus three is $(2 + 3)"
```


Raw strings are written using backticks and will preserve all characters.

```bake
let d = `my $(verbatim) string\n
  with whitespace, without substitutions and escape characters
`
```

If you want to mix verbatim substitutions with evaluated substitutions, escape the dollar sign.

```bake
let e = "my \$(sometimes) $("eval" + "uated") string"
```
