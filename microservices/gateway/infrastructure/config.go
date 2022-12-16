package infrastructure

type Service struct {
	Name string
	Host string
}

type Listen struct {
	Host string
	Port string
}

type Gateway struct {
	Environment string
	Listen
	Services []Service
}
