package model

import (
	"fmt"
	"github.com/fatih/structs"
)

type Setting interface {
	ToSurge() (key, value string)
	ToClash() map[string]interface{}
}

type VmessSetting struct {
	Add  string      `json:"add,omitempty" clash:"server"`
	Aid  interface{} `json:"aid,omitempty" clash:"alterId"`
	Host string      `json:"host,omitempty" clash:"-"`
	ID   string      `json:"id,omitempty" clash:"uuid"`
	Net  string      `json:"net,omitempty" clash:"-"`  //手工处理
	Path string      `json:"path,omitempty" clash:"-"` //手工处理
	Port interface{} `json:"port,omitempty" clash:"port"`
	Ps   string      `json:"ps,omitempty" clash:"name"`
	TLS  string      `json:"tls,omitempty" clash:"-"` //手工处理
	Type string      `json:"type,omitempty" clash:"cipher"`
	V    interface{} `json:"v,omitempty" clash:"-"`
}

func (t *VmessSetting) ToSurge() (key, value string) {
	value = fmt.Sprintf("%s, %s, %v, username=%s", Vmess, t.Add, t.Port, t.ID)
	if t.Net == "ws" {
		value += ", ws = true"
		value += ", ws-path = "
		if t.Path == "" {
			value += "/"
		} else {
			value += t.Path
		}
		value += ", ws-headers = Host:" + t.Host
	}
	if t.TLS != "" {
		value += ", tls=true"
	} else {
		value += ", tls=false"
	}
	return t.Ps, value
}

func (t *VmessSetting) ToClash() map[string]interface{} {
	structs.DefaultTagName = Clash
	m := structs.Map(t)
	if t.TLS != "" {
		m["tls"] = true
	} else {
		m["tls"] = false
	}
	if t.Net == "ws" {
		m["network"] = t.Net
		m["ws-path"] = t.Path
		m["ws-headers"] = map[string]interface{}{"Host": t.Host}
	}
	m["type"] = Vmess
	//todo 确认cipher字段
	return m
}

type SSSetting struct {
	Cipher   string `clash:"cipher"`
	Password string `clash:"password"`
	Domain   string `clash:"server"`
	Port     string `clash:"port"`
	Name     string `clash:"name"`
	Obfs     string `clash:"obfs,omitempty"`
	ObfsHost string `clash:"obfs-host,omitempty"`
}

func (t *SSSetting) ToSurge() (key, value string) {
	value = fmt.Sprintf("%s, %s, %v, encrypt-method=%s, password=%s", SS, t.Domain, t.Port, t.Cipher, t.Password)
	if t.Obfs != "" {
		value += ", obfs = " + t.Obfs
	}
	if t.ObfsHost != "" {
		value += ", obfs-host=" + t.ObfsHost
	}
	return t.Name, value
}

func (t *SSSetting) ToClash() map[string]interface{} {
	structs.DefaultTagName = Clash
	m := structs.Map(t)
	m["type"] = SS
	return m
}
