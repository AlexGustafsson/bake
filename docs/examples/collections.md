Collections in bake come in the form of objects and arrays. Objects are key-value stores and arrays a series of elements.

A common object in bake is the builtin `env`, in which all the environment variables are defined.

```bake
// Print the configured C compiler
print(env.CC)
```

```
gcc
```

One may also access objects using the index syntax.

```bake
print(env["MY-ENV-VAR"])
```

Objects may also be defined using the following syntax.

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

Arrays are defined in two main ways.

```bake
// Either, via directly declaring its contents.
let a = [1, 2, 3, 4]
// Or by using the range builtin
let b = range(1, 5)
```

The elements in an array may be accessed using the same index syntax, with integer indices.

```bake
let a = [1, 2, 3, 4]
print(a[0])
```

```
1
```
