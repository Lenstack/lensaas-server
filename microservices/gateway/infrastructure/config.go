package infrastructure

type Route struct {
	Name     string
	Context  string
	Target   string
	Protocol string
}

type ListenAddress struct {
	Host string
	Port string
}

type GatewayConfig struct {
	Listen ListenAddress
	Routes []Route
}
