package environtments

type Env struct {
	Host     string `env:"HOST" envDefault:"root:v5jd7Sk61lL99Yumna@tcp(192.168.100.50:3002)/ridnvil?charset=utf8mb4&parseTime=True&loc=Local"`
	HostOnly string `env:"HOST_ONLY" envDefault:"root:v5jd7Sk61lL99Yumna@tcp(192.168.100.50:3002)/?charset=utf8mb4&parseTime=True&loc=Local"`
}
