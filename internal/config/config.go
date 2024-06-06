package config

type Config struct {
	Type    string `env:"NR_TYPE" envDefault:"download"`
	Address string `env:"NR_ADDRESS,required"`
}
