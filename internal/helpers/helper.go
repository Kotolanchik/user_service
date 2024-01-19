package helpers

import "github.com/fatih/structs"

func ToMap(data any) map[string]any {
	strct := structs.New(data)
	strct.TagName = "json"

	return strct.Map()
}
