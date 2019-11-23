package model

import (
	"fmt"
	"github.com/fatih/structs"
)

type Setting interface {
	ToSurge() (key, value string)
	ToQuantumultX() (name, value string)
	ToClash() map[string]interface{}
}

type VmessSetting struct {
	Add  string      `json:"add,omitempty" clash:"server" qx:"vmess"`
	Aid  interface{} `json:"aid,omitempty" clash:"alterId" qx:"-"`
	Host string      `json:"host,omitempty" clash:"-" qx:"obfs-host"`
	ID   string      `json:"id,omitempty" clash:"uuid" qx:"password"`
	Net  string      `json:"net,omitempty" clash:"-" qx:"obfs,omitempty"`      //手工处理
	Path string      `json:"path,omitempty" clash:"-" qx:"obfs-uri,omitempty"` //手工处理
	Port interface{} `json:"port,omitempty" clash:"port" qx:"-"`
	Ps   string      `json:"ps,omitempty" clash:"name" qx:"tag"`
	TLS  string      `json:"tls,omitempty" clash:"-" qx:"-"` //手工处理
	Type string      `json:"type,omitempty" clash:"cipher" qx:"method"`
	V    interface{} `json:"v,omitempty" clash:"-" qx:"-"`
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

func (t *VmessSetting) ToQuantumultX() (name, v string) {
	s := structs.New(t)
	s.TagName = "qx"
	m := s.Map()
	m[Vmess] = fmt.Sprintf("%v:%v", m[Vmess], t.Port)
	v = fmt.Sprintf("%v=%v,", Vmess, m[Vmess])
	delete(m, Vmess)
	for key, value := range m {
		v += fmt.Sprintf("%v=%v,", key, value)
	}
	return t.Ps, v[:len(v)-1]
}

func (t *VmessSetting) ToClash() map[string]interface{} {
	s := structs.New(t)
	s.TagName = Clash
	m := s.Map()
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
	Cipher   string `clash:"cipher" qx:"method"`
	Password string `clash:"password" qx:"password"`
	Domain   string `clash:"server" qx:"shadowsocks"`
	Port     string `clash:"port" qx:"-"`
	Name     string `clash:"name" qx:"tag"`
	Obfs     string `clash:"obfs,omitempty" qx:"obfs"`
	ObfsHost string `clash:"obfs-host,omitempty" qx:"obfs-host"`
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
	s := structs.New(t)
	s.TagName = Clash
	m := s.Map()
	m["type"] = SS
	return m
}

func (t *SSSetting) ToQuantumultX() (name, v string) {
	s := structs.New(t)
	s.TagName = "qx"
	m := s.Map()
	m["shadowsocks"] = fmt.Sprintf("%v:%v", m["shadowsocks"], t.Port)
	v = fmt.Sprintf("%v=%v,", "shadowsocks", m["shadowsocks"])
	delete(m, "shadowsocks")
	for key, value := range m {
		v += fmt.Sprintf("%v=%v,", key, value)
	}
	return t.Name, v[:len(v)-1]
}
