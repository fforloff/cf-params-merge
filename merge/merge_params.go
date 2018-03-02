package merge

//MergedParams ...
func MergeParams(cft string, pfa []string) ([]Param, error) {

	var (
		res                []Param
		paramValue         string
		pv                 string
		err                error
		paramsFromTemplate map[string]interface{}
		paramsFromFiles    [][]Param
	)

	paramsFromTemplate, err = getParamsFromTemplate(string(cft))
	if err != nil {
		panic(err)
	}
	if len(pfa) > 0 {
		paramsFromFiles, err = getParamsFromFiles(pfa)
		if err != nil {
			panic(err)
		}
	}

	for name, valueInterface := range paramsFromTemplate {

		pv = getTemplateParamValue(valueInterface)
		if len(pv) > 0 {
			paramValue = pv
		}

		pv = getParamValueFromFiles(paramsFromFiles, name)
		if len(pv) > 0 {
			paramValue = pv
		}

		pv = getParamValueFromEnv(name)
		if len(pv) > 0 {
			paramValue = pv
		}

		if len(paramValue) > 0 {
			param := Param{
				ParameterKey:   name,
				ParameterValue: paramValue,
			}
			res = append(res, param)
		}
	}

	return res, nil
}
