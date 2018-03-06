package merge

//MergedParams...
func MergeParams(cft string, pfa []string) ([]Param, error) {

	var (
		res                []Param
		param              Param
		ok                 bool
		err                error
		paramsFromTemplate map[string]interface{}
		paramsFromFiles    [][]Param
	)

	paramsFromTemplate, err = getParamsFromTemplate(string(cft))
	if err != nil {
		return res, err
	}
	if len(pfa) > 0 {
		paramsFromFiles, err = loadParamsFromFiles(pfa)
		if err != nil {
			return res, err
		}
	}

	for name := range paramsFromTemplate {
		if param, ok = getParamValueFromEnv(name); ok {
			res = append(res, param)
			continue
		}
		if param, ok = getParamFromFiles(paramsFromFiles, name); ok {
			res = append(res, param)
		}
	}

	return res, nil
}
