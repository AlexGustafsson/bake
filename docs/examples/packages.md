A package is a reusable bundle of functions, hooks and rules you may share with others. A package always starts with a package declaration.

```bake
package basic
```

This package may then be published in a repository such as GitHub and imported in other Bake files. The system borrows a lot from go. The package name is expected to be the same as the directory it is in.

```bake
import (
  "github.com/AlexGustafsson/bake/examples/basic"
)
```

Packages work like any other Bake file, but may also export functionality.

```bake
func print_something_private {
  print("This is private")
}
```

Exported rules and functions are prefixed with the export keyword.

```bake
export func print_something {
  print_something_private()
}
```

To use imported functionality, you must refer to the package's name like so.

```bake
basic::print_something()
```
