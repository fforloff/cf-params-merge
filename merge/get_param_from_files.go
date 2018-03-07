package merge

func getParamFromFiles(ppp [][]Param, n string) (Param, bool) {
	var (
		ok bool
		p  Param
	)
	for _, pp := range ppp {
		if p, ok = thisParameterIsInFiles(pp, n); ok {
			return p, true
		}
	}
	return p, false
}
