package cmd

/* Configuration */

func AppConfig() Config {
	return &theConfig{}
}

type Config interface {
	MetaRepoHome() Setting[string]
}

type theConfig struct{}

func (config *theConfig) MetaRepoHome() Setting[string] {
	return &StringSetting{
		Def: "/home/me/meta-default",
		Val: "/home/me/meta",
	}
}

/* Settings */

type StringSetting struct {
	Def string
	Val string
}

func (setting StringSetting) DefaultValue() string { return setting.Def }
func (setting StringSetting) Value() string        { return setting.Val }

type Setting[V Value] interface {
	DefaultValue() V
	Value() V
}

type Value interface{ string }
