package oanda

type AuthToken string
type ClientId string

type Auth struct {
	Token    AuthToken
	ClientId ClientId
}
