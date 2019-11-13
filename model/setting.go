package model

import (
	"fmt"
	"github.com/fatih/structs"
)

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

func (s *Setting) ToSurge() (key, value string) {
	value = fmt.Sprintf("%s, %s, %v, username=%s", s.Scheme, s.Add, s.Port, s.ID)
	if s.Net == "ws" {
		value += ", ws = true"
		value += ", ws-path = " + s.Path
		value += ", ws-headers = Host:" + s.Host
	}
	if s.TLS != "" {
		value += ", tls=true"
	} else {
		value += ", tls=false"
	}
	return s.Ps, value
}
