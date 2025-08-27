package config

type Config struct {
	Log      LogConfig      `toml:"log" json:"log" mapstructure:"log"`
	Database DatabaseConfig `toml:"database" json:"database" mapstructure:"database"`

	EmailSender EmailSenderConfig `toml:"email_sender" json:"email_sender" mapstructure:"email_sender"`
}
