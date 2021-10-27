package config

type System struct {
	Env    string `mapstructure:"env" json:"env" yaml:"env"`            // 环境值
	Addr   int    `mapstructure:"addr" json:"addr" yaml:"addr"`         // 端口值
	SqlLog bool   `mapstructure:"sql-log" json:"sqlLog" yaml:"sql-log"` // sql语句显示
}
