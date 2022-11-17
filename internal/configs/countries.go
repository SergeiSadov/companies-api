package configs

type Countries struct {
	AlloweCountryCode string `envconfig:"ALLOWED_COUNTRY_CODE" default:"CY"`
}
