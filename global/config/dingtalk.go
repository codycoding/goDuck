package config

type DingTalk struct {
	AppKey    string `mapstructure:"app-key" json:"appKey" yaml:"app-key"`
	AppSecret string `mapstructure:"app-secret" json:"appSecret" yaml:"app-secret"`
	CacheKey  string `mapstructure:"cache-key" json:"cacheKey" yaml:"cache-key"`
}
