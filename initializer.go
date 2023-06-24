package common

// Original Idea:
// ==== Initializer Function Types ====
// type BackendInitializer[T any] func(modelPath string, opts ...InitializationOptionSetter) (*T, error)
// type BackendInitializerWithOptions[T any] func(modelPath string, initializationOptions InitializationOptions) (*T, error)

// Improved Form:
// ==== BackendInitializer Generic Struct ====
// Each go-* backend should export one of these where T is the wrapper for that backend.
// map[string]BackendInitializer can be created to easily initialize

// TODO: this really could just be any, but this looks margially cleaner?
type Backend interface{}

type BackendInitializer[T Backend] struct {
	DefaultInitializationOptions InitializationOptions
	Constructor                  func(modelPath string, initializationOptions InitializationOptions) (*T, error)
}

func (bi BackendInitializer[T]) Defaults(modelPath string) (*T, error) {
	return bi.Constructor(modelPath, bi.DefaultInitializationOptions)
}

func (bi BackendInitializer[T]) New(modelPath string, opts ...InitializationOptionSetter) (*T, error) {
	optionCopy := bi.DefaultInitializationOptions
	return bi.Constructor(modelPath, *MergeInitializationOptions(&optionCopy, opts...))
}

func (bi BackendInitializer[T]) NewWithInitializationOptions(modelPath string, initializationOptions InitializationOptions) (*T, error) {
	return bi.Constructor(modelPath, initializationOptions)
}
