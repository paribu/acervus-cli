package credential

type Context int

const (
	LoginContext Context = iota
	RegisterContext
	RefreshContext
)

type Credential struct {
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
	Current      bool   `json:"current"`
}
