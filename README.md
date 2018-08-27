# Bottalk Plugin package in Go

## Example usage
```
package main

import "github.com/bottalk/go-plugin"

func main() {

	plugin := bottalk.NewPlugin()

	plugin.Actions = []bottalk.Action{{
		Name:        "Test action",
		Description: "My description",
		Endpoint:    "/hello",
		Action: func() string {
			return "ok"
		},
	}}

	plugin.Run(":8080")
}
```
