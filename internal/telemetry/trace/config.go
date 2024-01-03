package trace

type Config struct {
	Enable bool  `mapstructure:"enable"`
	Agent  Agent `mapstructure:"agent"`
}
