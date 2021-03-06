Variables are useful for calculated rules etc. There are generally two ways to use variables, one for values never leaving Bake such as numeric expressions and parameters.

```bake
let one = 1
let two = 2
let three = one + three
print(three)
```

The other way to use variables is for values taken from the output of a shell function.

```bake
let one = 1
let two = 2
// Here we use the arbitrary precision calculator bc and evaluated strings to calculate the sum
shell echo "$(one) + $(two)" | bc
// We may then access the shell's output via the standard output and cast it as a number
let three = context.shell.stdout.number
print(three)
```

Using the same mechanics, we may create a utility function that calculates the sum of two arbitrary large values.

```bake
func shell_return(a, b) {
  shell echo "$(a) + $(b)" | bc
  return context.shell.stdout.number
}
```

Environment variables are accessible via the global `env` object.

```bake
func shell_example {
  shell echo "I'm using $(env.CC) as the compiler"

  // To default to another value if the environment variable does not exist, use
  // boolean expressions
  shell echo "I'm using $(env.CC || "gcc") as the compiler"

  // One may also use optional assignment, which will set the value only if none already exist.
  // In this case, it's set only for the context's shell - but may be set globally by using the
  // global `env` object
  context.shell.env.CC ?= "gcc"
  shell echo "$CC" is in use now, which is not necessarily the same as $(env.CC)
}
```
