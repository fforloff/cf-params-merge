package merge

import (
	"os"
	"strconv"
)

func getTemplateParamValue(i interface{}) string {
	dv := ""
	if vv, ok := i.(map[string]interface{}); ok {
		if v, ok := vv["Default"]; ok {
			switch vTyped := v.(type) {
			case string:
				dv = vTyped
			case int:
				dv = strconv.Itoa(vTyped)
			case int64:
				dv = strconv.FormatInt(vTyped, 10)
			case float64:
				dv = strconv.FormatFloat(vTyped, 'f', -1, 64)
			case bool:
				if vTyped {
					dv = "true"
				} else {
					dv = "false"
				}
			default:
				// need better error handling here
				os.Exit(1)
			}
		}
	}
	return dv
}
