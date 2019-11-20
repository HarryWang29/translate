package model

const (
	Web = "web"
)

//协议枚举
const (
	Vmess = "vmess"
	SS    = "ss"
)

const (
	Clash        = "clash"
	QuantumultX  = "quantumultx"
	Shadowrocket = "shadowrocket"
	Surge3       = "surge3"
)

const (
	ConnersHua = "ConnersHua"
)

const (
	Yaml = "yaml"
	Conf = "conf"
)

//记录规则路径
var RuleUrls = make(map[string]string)

//规则文件名
var RuleFileName = make(map[string]string)

//各软件配置文件后缀
var RuleFileType = make(map[string]string)

func init() {
	//初始化规则github地址
	RuleUrls[ConnersHua+Clash] = "https://raw.githubusercontent.com/ConnersHua/Profiles/master/Clash/Pro.yaml"
	RuleUrls[ConnersHua+QuantumultX] = "https://raw.githubusercontent.com/ConnersHua/Profiles/master/Quantumult/X/Pro.conf"
	RuleUrls[ConnersHua+Shadowrocket] = "https://raw.githubusercontent.com/ConnersHua/Profiles/master/Shadow/Pro.conf"
	RuleUrls[ConnersHua+Surge3] = "https://raw.githubusercontent.com/ConnersHua/Profiles/master/Surge/Surge3.conf"

	//初始化配置文件后缀
	RuleFileType[Clash] = Yaml
	RuleFileType[QuantumultX] = Conf
	RuleFileType[Shadowrocket] = Conf
	RuleFileType[Surge3] = Conf

	////初始化规则本地文件名
	setRuleFileName(ConnersHua, Clash)
	setRuleFileName(ConnersHua, QuantumultX)
	setRuleFileName(ConnersHua, Shadowrocket)
	setRuleFileName(ConnersHua, Surge3)

}

func setRuleFileName(ruleName, target string) {
	RuleFileName[ruleName+target] = ruleName + target + "." + RuleFileType[target]
}
