package environtments

type Env struct {
	Host     string `env:"HOST" envDefault:"public:a1b2c3d4@tcp(127.0.0.1:3306)/ridnvil?charset=utf8mb4&parseTime=True&loc=Local"`
	HostOnly string `env:"HOST_ONLY" envDefault:"public:a1b2c3d4@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"`
}
