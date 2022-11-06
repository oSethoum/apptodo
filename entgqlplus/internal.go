package entgqlplus

import "encoding/json"

func PrintJson(v interface{}) {
	buffer, _ := json.MarshalIndent(v, "", "   ")
	println(string(buffer))
}
