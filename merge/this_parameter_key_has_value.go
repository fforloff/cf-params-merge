package merge

func thisParameterKeyHasValue(p []Param, name string) (string, bool) {
	for _, v := range p {
		if v.ParameterKey == name {
			return v.ParameterValue, true
		}
	}
	return "", false
}
