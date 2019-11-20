package targets

import (
	"bytes"
	"github.com/pkg/errors"
	"gopkg.in/ini.v1"
	"strings"
	"translate/model"
)

type Surge struct {
	apiBase
	options ini.LoadOptions
}

func NewSurge3(rule string, settings []model.Setting) api {
	return &Surge{
		apiBase: apiBase{
			Rule:     rule,
			Settings: settings,
		},
		options: ini.LoadOptions{UnparseableSections: []string{"Rule", "Header Rewrite", "URL Rewrite", "Script"}},
	}
}

func (t *Surge) Run() ([]byte, error) {
	//解析配置文件
	f, err := ini.LoadSources(t.options, []byte(t.Rule))
	if err != nil {
		return nil, errors.Wrap(err, "ini.LoadSources")
	}

	proxies := make([]string, 0, len(t.Settings))
	proxyNames := make([]string, 0, len(t.Settings))
	//删除样例
	f.Section("Proxy").DeleteKey("1")
	f.Section("Proxy").DeleteKey("2")
	f.Section("Proxy").DeleteKey("3")
	f.Section("Proxy").DeleteKey("4")
	//转换为配置文件
	for _, value := range t.Settings {
		name, p := value.ToSurge()
		proxies = append(proxies, p)
		f.Section("Proxy").Key(name).SetValue(p)
		proxyNames = append(proxyNames, name)
	}

	proxy := ""
	for _, value := range f.Section("Proxy Group").Keys() {
		index := strings.Index(value.String(), "1,")
		if index != -1 {
			proxy = value.String()[:index] + strings.Join(proxyNames, ",")
			//自动测试最后需要增加测速网址
			if value.Name() == "Auto" {
				proxy += ", url = http://bing.com/"
			}
			value.SetValue(proxy)
		}
	}
	var d bytes.Buffer
	_, err = f.WriteTo(&d)
	if err != nil {
		return nil, errors.Wrap(err, "ini.WriteTo")
	}
	return d.Bytes(), nil
}
