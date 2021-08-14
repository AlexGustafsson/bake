To enable you to act on rules being run, you may use hooks. The before hook will run before matching rules. In this case, a wildcard is used to match any rule.

```bake
@before.rule "*" {
  print("Running $(context.rule)")
  shell echo -e "\e[34m"
}
```

Hooks may also be used to trigger existing rules. Here, we're creating such a rule.

```bake
rule log_after {
  shell echo -e "\e[0m" && clear
  print("Ran $(context.rule)")
}
```


The after hook will run after rules. Here we specify the previously created rule instead of writing a new body.

```bake
@after.rule "*" log_after
```

One may also use the same type of hooks to act on commands being run.

```bake
@before.command "*" {
  print(context.command)
}

@before.command "gcc *" {
  print("Compiled code")
}
```
