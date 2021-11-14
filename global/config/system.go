package config

type System struct {
	Env       string `mapstructure:"env" json:"env" yaml:"env"`                       // 环境值
	Addr      int    `mapstructure:"addr" json:"addr" yaml:"addr"`                    // 端口值
	UrlPreFix string `mapstructure:"url-pre-fix" json:"urlPreFix" yaml:"url-pre-fix"` // api访问路径前缀
	SqlLog    bool   `mapstructure:"sql-log" json:"sqlLog" yaml:"sql-log"`            // sql语句显示
	Swagger   bool   `mapstaructure:"swagger" json:"swagger" yaml:"swagger"`          // 是否开启swagger
}
