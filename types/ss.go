package types

import (
	"github.com/pkg/errors"
	"net/url"
	"strings"
	"translate/model"
	"translate/util"
)

type SS struct {
	apiBase
}

func NewSS(args *model.CliArgs) api {
	return &SS{
		apiBase: apiBase{
			Args: args,
		},
	}
}

func (t *SS) Run() ([]model.Setting, error) {
	err := t.getSub()
	if err != nil {
		return nil, errors.Wrap(err, "t.getSub")
	}
	ret := make([]model.Setting, 0, len(t.Configs))
	for _, value := range t.Configs {
		u, err := url.Parse(value)
		if err != nil {
			return nil, errors.Wrap(err, "url.Parse")
		}
		ss := &model.SSSetting{}
		//解析加密方式+密码
		cipherAndPwd := strings.Split(util.Base64Decode(u.User.Username()), ":")
		ss.Cipher = cipherAndPwd[0]
		ss.Password = cipherAndPwd[1]

		//获取服务器信息
		hostAndPort := strings.Split(u.Host, ":")
		ss.Domain = hostAndPort[0]
		ss.Port = hostAndPort[1]

		//获取参数部分
		values := u.Query()
		ss.Name = u.Fragment

		plugin := values.Get("plugin")
		if plugin != "" {
			m := t.dealPlugin(plugin)
			ss.Obfs = m["obfs"]
			ss.ObfsHost = m["obfs-host"]
		}
		ret = append(ret, ss)
	}
	return ret, nil
}

func (t *SS) dealPlugin(plugin string) map[string]string {
	obfs := strings.Split(plugin, ";")
	m := make(map[string]string)
	for _, value := range obfs {
		kv := strings.Split(value, "=")
		if len(kv) > 1 {
			m[kv[0]] = kv[1]
		}
	}
	return m
}
