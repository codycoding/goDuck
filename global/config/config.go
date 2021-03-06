package config

//
// Server
//  @Description: 服务程序配置结构
//
type Server struct {
	System     System   `mapstructure:"system" json:"system" yaml:"system"`             // 程序配置结构
	Redis      Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`                // redis配置结构
	Zap        Zap      `mapstructure:"zap" json:"zap" yaml:"zap"`                      // zap日志配置结构
	MysqlDb    Mysql    `mapstructure:"mysqlDb" json:"mysqlDb" yaml:"mysqlDb"`          // mysql配置结构
	PostgresDb Postgres `mapstructure:"postgresDb" json:"postgresDb" yaml:"postgresDb"` // postgres配置结构
	JWT        JWT      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`                      // JWT配置结构
	Captcha    Captcha  `mapstructure:"captcha" json:"captcha" yaml:"captcha"`          // 验证码配置(权限应用专用)
	DingTalk   DingTalk `mapstructure:"dingtalk" json:"dingtalk" yaml:"dingtalk"`       // 钉钉应用配置
}
