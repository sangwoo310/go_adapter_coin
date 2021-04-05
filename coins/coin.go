package coins

type ICoin interface {
	PublicKeyToAddress() error

	Sign() error
}

type CoinConfig struct {
	MainNetId string
}