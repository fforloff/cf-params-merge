package merge

import (
	"github.com/fforloff/cfmingle/utils"
)

func loadParamsFromFiles(fs []string) ([][]Param, error) {
	var res [][]Param
	utils.ReverseSlice(fs)
	for _, f := range fs {
		pp, err := loadParamsFromAFile(f)
		if err != nil {
			return res, err
		}
		res = append(res, pp)
	}
	return res, nil
}
