package config

type LogConfig struct {
	LogLevel     string `toml:"log_level" json:"log_level" mapstructure:"log_level" default:"debug"`
	LogDir       string `toml:"log_dir" json:"log_dir" mapstructure:"log_dir"`
	LogFilename  string `toml:"log_filename" json:"log_filename" mapstructure:"log_filename"`
	MaxBackupNum int    `toml:"max_backup_num" json:"max_backup_num" mapstructure:"max_backup_num" default:"10"` // 保留日志文件最大数量
	MaxFileSize  int    `toml:"max_file_size" json:"max_file_size" mapstructure:"max_file_size" default:"100"`   // 单个文件最大大小，单位MB
}

type DatabaseConfig struct {
	Host         string `toml:"host" json:"host" mapstructure:"host" default:"127.0.0.1"`
	Port         int    `toml:"port" json:"port" mapstructure:"port" default:"3306"`
	Username     string `toml:"username" json:"username" mapstructure:"username" default:"root"`
	Password     string `toml:"password" json:"password" mapstructure:"password"`
	DatabaseName string `toml:"database_name" json:"database_name" mapstructure:"database_name"`
}

type EmailSenderConfig struct {
	From      string        `toml:"from" json:"from" mapstructure:"from"`
	SmtpAuth  EmailSmtpAuth `toml:"smtp_auth" json:"smtp_auth" mapstructure:"smtp_auth"`
	DefaultTo []string      `toml:"default_to" json:"default_to" mapstructure:"default_to"`
}
type EmailSmtpAuth struct {
	Identity string `toml:"identity" json:"identity" mapstructure:"identity"`
	Username string `toml:"username" json:"username" mapstructure:"username"`
	Password string `toml:"password" json:"password" mapstructure:"password"`
	Host     string `toml:"host" json:"host" mapstructure:"host" default:"smtp.126.com"`
}
