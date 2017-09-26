package Model

type TokenOauth struct {
	Entity
	AccessToken string
	RefeshToken string
	Scope string
	Client Client
	SecurityLevel string
}
