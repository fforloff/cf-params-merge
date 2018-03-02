package merge

func getParamValueFromFiles(pp [][]Param, n string) string {
	var (
		v  string
		ok bool
	)
	for _, p := range pp {
		if v, ok = thisParameterKeyHasValue(p, n); ok {
			return v
		}
	}
	return ""
}
