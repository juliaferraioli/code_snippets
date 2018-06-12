<!--- Copyright 2015 Google
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.--->

# Testing matrix operations with `go-cmp`

The source code in this directory is to accompany some blog posts about
[go-cmp](https://godoc.org/github.com/google/go-cmp/).

## To run:

```shell
$ go test matrix # -v if you want to see details
```

## To modify:

Interesting parts to modify might be the `matrix` package itself to add more
operations (which I will do periodically, but if you're looking for a real matrix
package, might I suggest 
[gonum/matrix](https://godoc.org/github.com/gonum/matrix)?), or the comparer
function in `matrix_test.go`.
