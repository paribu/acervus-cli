package credential

type Context int

const (
	LoginContext Context = iota
	RegisterContext
	RefreshContext
)

type Credential struct {
	Email        string `json:"email"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	Current      bool   `json:"current"`
}
