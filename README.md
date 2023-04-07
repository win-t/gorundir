# What is this?

simple utility to `go run` some your package when you are outside the module directory

# Example

```bash
mkdir -p /tmp/test

cat <<EOF > /tmp/test/main.go
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}
EOF

cat <<EOF > /tmp/test/go.mod
module hello-world

go 1.20
EOF

cd ~

go run winto.dev/gorundir@latest /tmp/test
```
