package targets

import (
	"fmt"
	"translate/model"
)

var switchApi = make(map[string]func(string, []model.Setting) api)

func init() {
	switchApi[model.Clash] = NewClash
	switchApi[model.Surge3] = NewSurge3
	switchApi[model.QuantumultX] = NewQuantumultX
}

type api interface {
	Run() ([]byte, error)
}

func Run(target, rule string, settings []model.Setting) ([]byte, error) {
	if f, ok := switchApi[target]; ok {
		return f(rule, settings).Run()
	}
	return nil, fmt.Errorf("target(%s) is error", target)
}

type apiBase struct {
	Rule     string
	Settings []model.Setting
}
