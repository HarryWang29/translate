package types

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/pkg/errors"
	"regexp"
	"strings"
	"translate/model"
	"translate/util"
)

var switchApi = make(map[string]func(args *model.CliArgs) api)

func init() {
	switchApi[model.Vmess] = NewVmess
	switchApi[model.SS] = NewSS
}

type api interface {
	Run() ([]model.Setting, error)
	getSub() error
}

func Run(typ string, args *model.CliArgs) ([]model.Setting, error) {
	if f, ok := switchApi[typ]; ok && args != nil {
		return f(args).Run()
	}
	return nil, fmt.Errorf("type(%s) is error or args is nil", typ)
}

//基类
type apiBase struct {
	Args       *model.CliArgs
	Configs    []string
	NpsboostV3 map[string]int
}

//获取订阅内容
func (a *apiBase) getSub() error {
	//为喵帕斯过滤v3节点
	if a.Args.Npsboost != "" {
		a.NpsboostV3 = make(map[string]int)
		//获取3j节点
		matched, err := a.dealSub(a.Args.Npsboost)
		if err != nil {
			return errors.Wrap(err, "获取喵帕斯3j地址失败")
		}
		//遍历放入map
		for _, value := range matched {
			decode := strings.ReplaceAll(value[1], "_", "+")
			tran := util.Base64Decode(decode)
			index := strings.Index(tran, ":")
			if index == -1 {
				continue
			}
			a.NpsboostV3[tran[:index]] = 0
		}
	}
	//处理正常订阅地址
	for _, value := range a.Args.SubLinks {
		matched, err := a.dealSub(value)
		if err != nil {
			return errors.Wrap(err, "dealSub")
		}
		for _, value := range matched {
			//直接解码
			v := ""
			if value[1] != model.SS {
				v = util.Base64Decode(value[2])
			} else {
				v = value[0]
			}
			a.Configs = append(a.Configs, v)
		}
	}
	return nil
}

func (a *apiBase) dealSub(url string) ([][]string, error) {
	resp, err := httplib.Get(url).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "httplib.Get")
	}
	//通过正则匹配规则
	re := regexp.MustCompile(`(.*)://(.*)`)
	matched := re.FindAllStringSubmatch(util.Base64Decode(string(resp)), -1)
	return matched, nil
}
