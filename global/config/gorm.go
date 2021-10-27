package config

//
// Mysql
//  @Description: MySQL数据库连接配置
//
type Mysql struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                             // 服务器地址:端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                       // 高级配置
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`                     // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                 // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                 // 数据库密码
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`                  // 是否开启Gorm全局日志
	LogZap       string `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`                     // 是否通过zap写入日志文件
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
}

//
// Postgres
//  @Description: Postgres数据库连接配置
//
type Postgres struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`                             // 地址
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                             // 端口
	User         string `mapstructure:"user" json:"user" yaml:"user"`                             // 用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                 // 密码
	Dbname       string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`                       // 数据库
	Sslmode      string `mapstructure:"sslmode" json:"sslmode" yaml:"sslmode"`                    // ssl模式
	TimeZone     string `mapstructure:"timezone" json:"timezone" yaml:"timezone"`                 // 时区
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`                  // 是否开启Gorm全局日志
	LogZap       string `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`
}

func (p *Postgres) Dsn() string {
	return "host=" + p.Host + " user=" + p.User + " password=" + p.Password + " dbname=" + p.Dbname + " port=" + p.Port + " sslmode=" + p.Sslmode + " TimeZone=" + p.TimeZone
}
