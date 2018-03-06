package merge

import "os"

func getParamValueFromEnv(n string) (Param, bool) {
	var (
		ok bool
		p  Param
		v  string
	)
	if v, ok = os.LookupEnv(n); ok {
		p.ParameterKey = n
		p.ParameterValue = v
		return p, true
	}
	return p, false
}
