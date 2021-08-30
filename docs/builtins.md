# Builtins

Builtins are globally available functions that are included in bake. These functions provide you with the functionality to print things, easily create ranges etc. These functions are runtime specific, and not implemented in bake itself.

## print

The `print` builtin function prints the given values to the standard output, separated by a space. It also prints a trailing newline.

```bake
print("hello", "world")
```

```
hello world
```

## range

The `range` builtin function creates an array representing the specified range. It takes two arguments, the inclusive `start` and the exclusive `stop` of the range.

```bake
print(range(1, 5))
```

```
[1, 2, 3, 4]
```

## env

The `env` builtin object holds all of the environment variables of the current session.

```bake
print(env.CC)

if env["DEBUG"] == "1" {
  print("debugging")
}
```
