package types

import (
	"encoding/json"
	"github.com/pkg/errors"
	"translate/model"
)

type Vmess struct {
	apiBase
}

func NewVmess(args *model.CliArgs) api {
	return &Vmess{
		apiBase: apiBase{
			Args: args,
		},
	}
}

func (t *Vmess) Run() ([]model.Setting, error) {
	err := t.getSub()
	if err != nil {
		return nil, errors.Wrap(err, "t.getSub")
	}
	ret := make([]model.Setting, 0, len(t.Configs))
	for _, value := range t.Configs {
		v := &model.VmessSetting{}
		err := json.Unmarshal([]byte(value), v)
		if err != nil {
			return nil, errors.Wrap(err, "json.Unmarshal")
		}
		if len(t.NpsboostV3) == 0 {
			ret = append(ret, v)
		} else {
			if _, ok := t.NpsboostV3[v.Add]; ok {
				ret = append(ret, v)
			}
		}
	}
	return ret, nil
}
