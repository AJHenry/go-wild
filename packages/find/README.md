# find

> Find a file or directory by walking up parent directories

## Install

```
import "github.com/ajhenry/go-wild/find"
```

## Usage

```
/
└── Users
    └── ajhenry
        ├── unicorn.png
        └── foo
            └── bar
                ├── baz
                └── example.go
```

`main.go`

```go
package main

import (
	"fmt"
	"github.com/ajhenry/go-wild/find"
)

func main() {
	dir, err := find.Up("unicorn.png", Options{})
	//=> "/Users/ajhenry/", <nil>

    dir, err := find.Up("foo", Options{isDir: true})
	//=> "/Users/ajhenry/", <nil>

    dir, err := find.Up("ajhenry", Options{isDir: true, depth: 1})
	//=> "", <File not found in any parent directories>
}
```

## API

### Up(name string, options Options) (string, error)

Returns a string for the path or an `error` if it couldn't be found.

#### name

Type: `string`

Name of the file or directory to find.

#### options

Type: `interface`

##### cwd

Type: `string`\
Default: `os.Getwd()`

Directory to start from.

##### isDir

Type: `bool`\
Default: `false`

Determines if the file you are searching for is a directory.

### FileExists(name string, path string) bool

Returns a `bool` of whether the path exists for that file name.

### DirExists(name string, path string) bool

Returns a `bool` of whether the path exists for that directory.

#### name

Type: `string`

Name of the file/directory.

#### path

Type: `string`

Path to a file or directory.

## Based on findUp

This project is adapted from [findUp](https://github.com/sindresorhus/find-up).