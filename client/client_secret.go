package client

type ClientSecret struct {
	ID       int64  `json:"-" db:"pk"`
	ClientID string `json:"client_id" db:"client_id"`
	// Secret is client secret plaintex
	ClientSecretPlaintext string `json:"client_secret_plaintext" db:"client_secret_plaintext"`
}

func (ClientSecret) TableName() string {
	return "kyber_client_secret"
}
