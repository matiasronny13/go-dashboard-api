package model

type EnvironmentSettings struct {
	EnvironmentName  string `env:"ENVIRONMENT_NAME"`
	ConnectionString string `env:"PG_CONNECTION" envDefault:"none"`
}
