package merge

import "os"

func getParamValueFromEnv(n string) string {
	var (
		v  string
		ok bool
	)
	if v, ok = os.LookupEnv(n); ok {
		return v
	}
	return ""
}
