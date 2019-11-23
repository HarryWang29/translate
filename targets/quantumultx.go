package targets

import (
	"bytes"
	"github.com/pkg/errors"
	"gopkg.in/ini.v1"
	"translate/model"
)

type QuantumultX struct {
	apiBase
	options ini.LoadOptions
}

func NewQuantumultX(rule string, settings []model.Setting) api {
	return &QuantumultX{
		apiBase: apiBase{
			Rule:     rule,
			Settings: settings,
		},
		options: ini.LoadOptions{UnparseableSections: []string{"policy", "filter_remote", "rewrite_remote", "filter_local", "server_local"}},
	}
}

func (t *QuantumultX) Run() ([]byte, error) {
	//解析配置文件
	f, err := ini.LoadSources(t.options, []byte(t.Rule))
	if err != nil {
		return nil, errors.Wrap(err, "ini.LoadSources")
	}

	//转换为配置文件
	names := ""
	local := ""
	for _, value := range t.Settings {
		name, v := value.ToQuantumultX()
		local += v + "\n"
		names += name + ","
	}
	names = names[:len(names)-1]
	//设置节点信息
	_, err = f.NewRawSection("server_local", local)
	if err != nil {
		return nil, errors.Wrap(err, "f.NewRawSection")
	}

	//设置健康策略
	body := f.Section("policy").Body()
	body += "\navailable=auto," + names
	_, err = f.NewRawSection("policy", body)
	if err != nil {
		return nil, errors.Wrap(err, "f.NewRawSection")
	}

	var d bytes.Buffer
	_, err = f.WriteTo(&d)
	if err != nil {
		return nil, errors.Wrap(err, "ini.WriteTo")
	}
	return d.Bytes(), nil
}
