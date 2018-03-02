package merge

import "github.com/awslabs/goformation"

func getParamsFromTemplate(f string) (map[string]interface{}, error) {
	t, err := goformation.Open(string(f))
	if err != nil {
		return nil, err
	}
	return t.Parameters, nil
}
