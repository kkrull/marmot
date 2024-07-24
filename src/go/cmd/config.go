package cmd

func AppConfig() Config {
	return &TheConfig{}
}

type TheConfig struct{}

func (config *TheConfig) MetaRepoHome() Setting[string] {
	return &StringSetting{
		Def: "/home/me/meta-default",
		Val: "/home/me/meta",
	}
}

type StringSetting struct {
	Def string
	Val string
}

func (setting *StringSetting) DefaultValue() string { return setting.Def }
func (setting *StringSetting) Value() string        { return setting.Val }

type Config interface {
	MetaRepoHome() Setting[string]
}

type Setting[V Value] interface {
	DefaultValue() V
	Value() V
}

type Value interface{ string }
