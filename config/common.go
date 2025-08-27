package config

type LogConfig struct {
	LogLevel     string `toml:"log_level" json:"log_level"`
	LogDir       string `toml:"log_dir" json:"log_dir"`
	LogFilename  string `toml:"log_filename" json:"log_filename"`
	MaxBackupNum int    `toml:"max_backup_num" json:"max_backup_num"`
	MaxFileSize  int    `toml:"max_file_size" json:"max_file_size"`
}

type DatabaseConfig struct {
	Host         string `toml:"host" json:"host"`
	Port         int    `toml:"port" json:"port"`
	Username     string `toml:"username" json:"username"`
	Password     string `toml:"password" json:"password"`
	DatabaseName string `toml:"database_name" json:"database_name"`
}
