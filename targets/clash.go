package targets

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"translate/model"
)

type Clash struct {
	apiBase
}

func NewClash(rule string, settings []model.Setting) api {
	return &Clash{
		apiBase: apiBase{
			Rule:     rule,
			Settings: settings,
		},
	}
}

func (t *Clash) Run() ([]byte, error) {
	//解析配置文件
	m := make(map[interface{}]interface{})

	err := yaml.Unmarshal([]byte(t.Rule), &m)
	if err != nil {
		return nil, errors.Wrap(err, "yaml.Unmarshal")
	}

	proxies := make([]map[string]interface{}, 0, len(t.Settings))
	proxyNames := make([]interface{}, 0, len(t.Settings))
	//转换为配置文件
	for _, value := range t.Settings {
		p := value.ToClash()
		proxies = append(proxies, p)
		proxyNames = append(proxyNames, p["name"])
	}
	m["Proxy"] = proxies

	proxyGroup := make([]interface{}, 0)
	for _, value := range m["Proxy Group"].([]interface{}) {
		proxy, ok := value.(map[interface{}]interface{})
		if !ok {
			continue
		}
		proxiesName, ok := proxy["proxies"].([]interface{})
		if !ok {
			continue
		}
		for i := 0; i < len(proxiesName); i++ {
			if proxiesName[i] == "1" {
				proxiesName = append(proxiesName[:i], proxyNames...)
				break
			}
		}
		proxy["proxies"] = proxiesName
		proxyGroup = append(proxyGroup, proxy)
	}
	m["Proxy Group"] = proxyGroup
	d, err := yaml.Marshal(&m)
	if err != nil {
		return nil, errors.Wrap(err, "yaml.Marshal")
	}
	return d, nil
}
