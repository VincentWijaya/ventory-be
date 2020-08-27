package user

type (
	UserRepository interface{}
)

type Module struct {
	userRepo  UserRepository
	jwtSecret string
}

func New(jwtSecret string) *Module {
	return &Module{
		// userRepo:  user,
		jwtSecret: jwtSecret,
	}
}
