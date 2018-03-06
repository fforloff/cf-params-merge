package merge

func thisParameterIsInFiles(pp []Param, name string) (Param, bool) {
	var p Param
	for _, p := range pp {
		if p.ParameterKey == name {
			return p, true
		}
	}
	return p, false
}
