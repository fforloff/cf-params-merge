package merge

import (
	"github.com/fforloff/cfmingle/utils"
)

func getParamsFromFiles(fs []string) ([][]Param, error) {
	var res [][]Param
	utils.ReverseSlice(fs)
	for _, f := range fs {
		pp, err := getParamsFromFile(f)
		if err != nil {
			return res, err
		}
		res = append(res, pp)
	}
	return res, nil
}
