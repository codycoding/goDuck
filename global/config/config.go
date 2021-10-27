package config

//
// Server
//  @Description: 服务程序配置结构
//
type Server struct {
	System    System   `mapstructure:"system" json:"system" yaml:"system"`      // 程序配置结构
	Redis     Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`         // redis配置结构
	Zap       Zap      `mapstructure:"zap" json:"zap" yaml:"zap"`               // zap日志配置结构
	SlpDataDb Mysql    `mapstructure:"slpDataDb" json:"oneDb" yaml:"slpDataDb"` // mysql配置结构
	DgDb      Postgres `mapstructure:"DgDb" json:"dgDb" yaml:"dgDb"`            // postgres配置结构
}
