package config

//
// CorsConfig
//  @Description: 跨域配置
//
type CorsConfig struct {
	AllowOrigins     []string `mapstructure:"allow-origins" json:"allowOrigins" yaml:"allow-origins"`             // 准许跨域请求的网站
	AllowMethods     []string `mapstructure:"allow-methods" json:"allowMethods" yaml:"allow-methods"`             // 准许使用的请求方式
	AllowHeaders     []string `mapstructure:"allow-headers" json:"allowHeaders" yaml:"allow-headers"`             // 准许使用的请求表头
	ExposeHeaders    []string `mapstructure:"expose-headers" json:"exposeHeaders" yaml:"expose-headers"`          // 显示的请求表头
	AllowCredentials bool     `mapstructure:"allow-credentials" json:"allowCredentials" yaml:"allow-credentials"` // 凭证共享
	MaxAge           int      `mapstructure:"max-age" json:"maxAge" yaml:"max-age"`                               // 超时时间设定
}
