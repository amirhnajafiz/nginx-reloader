package logger

type Config struct {
	Development bool   `mapstructure:"development"`
	Level       string `mapstructure:"level"`
}
