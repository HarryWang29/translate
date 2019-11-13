package translate

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
	"translate/model"
	"translate/targets"
	"translate/types"
)

func Run(typ string, args model.CliArgs) error {
	//由于raw.githubusercontent.com域名可能存在问题，优先尝试请求规则
	ruleOnline, err := GetRules(args.RuleName, args.Target)
	if err != nil {
		return errors.Wrap(err, "GetRules")
	}
	log.Println("加载规则成功")
	//获取订阅配置
	subs, err := types.Run(typ, &args)
	if err != nil {
		return errors.Wrap(err, "types.Run")
	}
	log.Println("加载订阅成功")

	rule, err := targets.Run(args.Target, ruleOnline, subs)
	if err != nil {
		return errors.Wrap(err, "targets.Run")
	}
	log.Println("规则转化成功")
	fileName := typ + args.RuleName + strings.Title(args.Target) + time.Now().Format("150405") + "." + model.RuleFileType[args.Target]
	err = ioutil.WriteFile(fileName, rule, 0644)
	if err != nil {
		return errors.Wrap(err, "ioutil.WriteFile")
	}
	dir, _ := os.Getwd()
	log.Printf("配置文件生成成功，路径：%s", dir+"/"+fileName)
	return nil
}

func GetRules(name, target string) (string, error) {
	//优先查看本地是否有文件存在，若本地有规则文件，则优先使用本地
	fileName, ok := model.RuleFileName[name+target]
	if !ok {
		return "", fmt.Errorf("file name cache error,name:%s,target:%s", name, target)
	}
	_, err := os.Stat(fileName)
	if err == nil {
		//文件存在，则读取文件
		log.Println("本地存在规则文件，加载本地规则")
		b, e := ioutil.ReadFile(fileName)
		if e != nil {
			return "", errors.Wrap(e, "ioutil.ReadFile")
		}
		return string(b), nil
	} else if !os.IsNotExist(err) {
		//非不存在问题
		return "", errors.Wrap(err, "os.Stat")
	}
	log.Printf("本地未找到规则文件，从github加载规则, url:%s", model.RuleUrls[name+target])

	//若本地没有文件，则通过github获取规则
	b, err := httplib.Get(model.RuleUrls[name+target]).Bytes()
	if err != nil {
		return "", errors.Wrap(err, "httplib.Get")
	}
	return string(b), nil
}
