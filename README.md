# spdx-go-model-3

Generated Golang code for the SPDX Spec version 3

All bindings in this repository are generated using
[shacl2code](https://github.com/JPEWdev/shacl2code) at the time the package is
built with the `generate-bindings` script. The binding should not be manually
edited.

**NOTE:** The bindings are pretty low level, intended for more directly
manipulating SPDX files. While they are fully functions, they lack higher level
helper functions that may be useful for creating SPDX documents. If you want a
higher level approach, please see the
[SPDX  Golang Tools](https://github.com/spdx/tools-golang) (however this repo
doesn't yet support SPDX 3)

## Usage

Each version of the SPDX spec has a module named `v{MAJOR}_{MINOR}_{MICRO}`
that contains the bindings for that version under the `github.com/spdx/spdx-go-model`
top level modules. It is recommend you import this with an alias like:

```golang
import (
    "github.com/spdx/spdx-go-model/spdx_v3_0_1"
)

...
p := spdx_v3_0_1.MakePerson()
...
```

## Examples

For examples of how to use the bindings, see the `examples_test.go` file

## Testing

This repository has support for running tests against the generated bindings using
`go test`.
