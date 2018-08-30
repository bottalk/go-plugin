# Bottalk Plugin package in Go

## Example usage
```
package main

import bottalk "github.com/bottalk/go-plugin"

func main() {

	plugin := bottalk.NewPlugin()
	plugin.Name = "Weather plugin"
	plugin.Description = "This plugin helps to fetch information about weather"

	plugin.Actions = map[string]bottalk.Action{"getWeather": bottalk.Action{
		Name:        "getWeather",
		Description: "This action fetches weather by cityId",
		Endpoint:    "/getWeather",
		Action: func() string {
			return "ok"
		},
		Params: map[string]string{"city": "Id of the city"},
	}}

	plugin.Run(":8080")
}}
```
