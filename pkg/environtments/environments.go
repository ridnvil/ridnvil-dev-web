package environtments

type Env struct {
	Host     string `env:"HOST" envDefault:"root:a1b2c3d4@tcp(192.168.100.73:4006)/ridnvil?charset=utf8mb4&parseTime=True&loc=Local"`
	HostOnly string `env:"HOST_ONLY" envDefault:"root:a1b2c3d4@tcp(192.168.100.73:4006)/?charset=utf8mb4&parseTime=True&loc=Local"`
}
