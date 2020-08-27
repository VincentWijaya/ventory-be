package middleware

type Module struct {
	jwtSecret string
}

func New(jwtSecret string) *Module {
	return &Module{
		jwtSecret: jwtSecret,
	}
}
