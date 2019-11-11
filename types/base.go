package types

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/pkg/errors"
	"regexp"
	"translate/model"
	"translate/util"
)

var switchApi = make(map[string]func([]string) api)

func init() {
	switchApi[model.Vmess] = NewVmess
}

type api interface {
	Run() ([]*model.Setting, error)
	getSub() error
}

func Run(typ string, subLink []string) ([]*model.Setting, error) {
	if f, ok := switchApi[typ]; ok {
		return f(subLink).Run()
	}
	return nil, fmt.Errorf("type(%s) is error", typ)
}

//基类
type apiBase struct {
	SubLinks []string
	Configs  []string
}

//获取订阅内容
func (a *apiBase) getSub() error {
	for _, value := range a.SubLinks {
		resp, err := httplib.Get(value).Bytes()
		if err != nil {
			return errors.Wrap(err, "httplib.Get")
		}
		//通过正则匹配规则
		re := regexp.MustCompile(`.*://(.*)`)
		matched := re.FindAllStringSubmatch(util.Base64Decode(string(resp)), -1)
		for _, value := range matched {
			//直接解码
			a.Configs = append(a.Configs, util.Base64Decode(value[1]))
		}
	}
	return nil
}
