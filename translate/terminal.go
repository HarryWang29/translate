package translate

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
	"translate/model"
)

type terminal struct {
	apiBase
}

func NewTerminal(typ string, args model.CliArgs) api {
	return &terminal{
		apiBase: apiBase{
			typ:  typ,
			args: args,
		},
	}
}

func (t *terminal) Run() error {
	rule, err := t.translate(t.typ, t.args)
	if err != nil {
		return errors.Wrap(err, "translate")
	}
	//终端转换直接写文件
	fileName := t.typ + t.args.RuleName + strings.Title(t.args.Target) + time.Now().Format("150405") + "." + model.RuleFileType[t.args.Target]
	err = ioutil.WriteFile(fileName, rule, 0644)
	if err != nil {
		return errors.Wrap(err, "ioutil.WriteFile")
	}
	dir, _ := os.Getwd()
	log.Printf("配置文件生成成功，路径：%s", dir+"/"+fileName)
	return nil
}
