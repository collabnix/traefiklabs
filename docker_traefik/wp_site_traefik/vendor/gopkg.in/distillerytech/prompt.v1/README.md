# prompt
A simple go package for terminal prompts.

Check out the [docs](http://godoc.org/github.com/DistilleryTech/prompt).

## Example

```go
package main

import (
	"fmt"

	"gopkg.in/distillerytech/prompt.v1"
)

func main() {
	name := prompt.AskStringRequired("Enter your name:")
	fmt.Println(name)

	street := prompt.AskString("Enter your street:")
	fmt.Println(street)

	answer := prompt.AskStringLimit("Apples or oranges?", "apples", "oranges")
	fmt.Println(answer)

	age := prompt.AskInteger("Enter your age:")
	fmt.Println(age)

	nodes := prompt.AskIntegerDefault("Number of nodes [2]:", 2)
	fmt.Println(nodes)

	resp := prompt.Confirm("Are you sure?")
	fmt.Println(resp)
}


```
