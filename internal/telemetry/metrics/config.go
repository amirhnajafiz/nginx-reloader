package metrics

type Config struct {
	Enable bool `mapstructure:"enable"`
	Port   int  `mapstructure:"port"`
}
