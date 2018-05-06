package oanda

type Options struct {
	auth *Auth
	mode *ApiMode
}

func newOptions(token AuthToken, id ClientId, mode *ApiMode) *Options {
	return &Options{
		mode: mode,
		auth: &Auth{
			Token:    token,
			ClientId: id,
		},
	}
}
