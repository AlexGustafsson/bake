Bake files can be called Bakefile or use the '.bke' file extension, like this file does. Like with Make, Bake uses rules to match source files and destination files with a set of
commands centered on shell code to "glue" together other programs.

Each rule has a name. In most cases, this name is the file path of the rule's output.

```bake
"new-file.txt" {
  As previously mentioned, shell functions are at the core of Bake. These may be written
  using regular Bash or the like, with commands and programs available on the system.
  shell touch new-file.txt
}
```

Like with Make, there are some standard variables defined in each rule. These help you write more concise rules without duplicating names.

```bake
"new-file2.txt" {
  Here we use the '$(context.in)' instead of re-using the file's name,
  that way it's easily changable later on.
  shell touch $(context.in)
}
```

Sometimes a rule uses a lot of shell code. In cases like that you may use shell blocks instead. The entire shell block will be executed in the same session to allow for complex behavior.

```bake
"new-file3.txt" {
  shell {
    mkdir -p tmp
    cd tmp
    touch new-file3.txt
    cp new-file3.txt ../
    cd ../
    rm -r tmp
  }
}
```

Rules may have dependencies so that Bake knows when to re-run a rule if any of its dependencies have been changed. It also helps Bake know in which order to run rules. This rule has dependencies, which will make sure that new-file.txt is created before this rule is run (by running the rule for it).

```bake
// If the file already exists, it will be used as is and if it doesn't exist at all,
an error will be thrown.
"copied-file.txt" ["new-file.txt"] {
  // Just like with output files, input files can be referenced in Bake using the
  // context
  shell cp $(context.out) $(context.in)
}
```


If you have multiple output files, place them in a list.

```bake
["file1", "file2"] {
  shell touch file1 file2
}
```

The rule that is run first may be decided in several ways. If no rule is explictly chosen, the topmost rule will be run first. Another way to decide what rule to run is to create a function called main.

```bake
func main {
  print("Hello, world!")
}
```

If you want to alias a build rule, such as "build" for build/app you may use an alias like so, where each dependant rule is specified much like dependencies of a regular rule.

```bake
alias build ["build/app"]
```
