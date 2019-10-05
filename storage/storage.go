package storage

// Explains how interfaces and pointer to structs work
// https://stackoverflow.com/questions/44370277/type-is-pointer-to-interface-not-interface-confusion

type Client interface {
	Set(string, interface{}) error
	Get(string) (interface{}, error)
	Exists(string) (bool, error)
	Cleanup() error
}
