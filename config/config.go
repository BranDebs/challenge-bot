package config

type Configer interface {
	UnmarshalKey(key string, val interface{}) error
	GetString(key string) string
}
