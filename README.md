# Learning Go

### How to initialize a Go module?

- Use `go mod init learning-go` which will create go.mod file. 
- It helps to determine the module path and dependencies.

### How does `go run ./1-hello-world` work?

When you run `go run ./1-hello-world`, Go:
1. Looks at all `.go` files in the `./1-hello-world` directory
2. Compiles them together as a single program
3. Looks for a `main()` function in a `package main` declaration
4. Executes that `main()` function

**Note:** The filename `main.go` is just a convention - Go doesn't require it. What matters is:
- The file has `package main` at the top
- It contains a `func main()` function
- When you run `go run` on a directory, it compiles all `.go` files in that directory together
