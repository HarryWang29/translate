package model

import "github.com/fatih/structs"

type Setting struct {
	Scheme string `json:"-" clash:"type"`
	//v2ray相关
	VmessSetting
	//
}

func (s *Setting) ToClash() map[string]interface{} {
	structs.DefaultTagName = Clash
	m := structs.Map(s.VmessSetting)
	if s.TLS != "" {
		m["tls"] = true
	} else {
		m["tls"] = false
	}
	if s.Net == "ws" {
		m["network"] = s.Net
		m["ws-path"] = s.Path
		m["ws-headers"] = map[string]interface{}{"Host": s.Host}
	}
	m["type"] = s.Scheme
	//todo 确认cipher字段
	return m
}
