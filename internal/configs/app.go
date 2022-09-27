package configs

type App struct {
	Environment string `envconfig:"ENV" default:"develop"`
	Port        int    `envconfig:"PORT" required:"true"`
	JWTSecret   string `envconfig:"JWT_SECRET" required:"true"`
}
