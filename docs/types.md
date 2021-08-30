# Types

## Strings

### Raw strings

Raw strings keep their content exactly as their written - it keeps whitespace and ignores escape codes.

```bake
print(`
hell\no
"world"
`)
```

```

hell\no
"world"

```

### Evaluated strings

Evaluated strings (regular strings) may not contain newlines, they do, however, allow for escape codes and inline evaluations.

```bake
let basic_string = "my string"
let formatted_string = "my\nmulti-line\nstring"
let evaluated_string = "1 + 2 = $(1 + 2)"
let escaped_evaluated_string = "1 + 1 + \$(1 + 2)"
```

```
my string
---------
my
multi-line
string
---------
1 + 2 = 3
---------
1 + 1 + $(1 + 2)
```

### Length

The length of strings may be retrieved by accessing the `length` property.

```bake
print("12345".length)
```

```
5
```

### Characters

A character may be retrieved by using the index syntax.

```bake
print("12345"[0])
```

```
1
```

## Arrays

Arrays may contain any elements - they're not bound to be of the same type.

```bake
let array = [1, 2, 3, 4, 5]
print(array)
```

```
[1, 2, 3, 4, 5]
```

### Length

The length of arrays may be retrieved by accessing the `length` property.

```bake
print([1, 2, 3, 4, 5].length)
```

```
5
```

### Elements

An element may be retrieved by using the index syntax.

```bake
print([1, 2, 3, 4, 5][0])
```

```
1
```

## Objects

Objects may be defined using the following syntax.

```bake
let a = {
  b: 1,
  c: "foo",
  d: {
    e: 2
  }
}

print(a)
```

```
{b: 1, c: "foo", d: {e: 2}}
```

### Values

An object's values may be accessed using their key, either via the index syntax or the selector syntax.

```bake
let object = {foo: "bar"}
print(object.foo)
print(object["foo"])
```

```
bar
bar
```
