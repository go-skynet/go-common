package common

import (
	"testing"
)

type MockBackend struct {
}

func (MockBackend) Name() string {
	return "MOCK"
}

func (MockBackend) Close() error {
	return nil
}

var MockBackendInitializer BackendInitializer[MockBackend] = BackendInitializer[MockBackend]{
	DefaultInitializationOptions: InitializationOptions{},
	Constructor: func(modelPath string, initializationOptions InitializationOptions) (*MockBackend, error) {
		return &MockBackend{}, nil
	},
}

func TestBI(t *testing.T) {
	mb, err := MockBackendInitializer.New("fake", EnableLowVRAM, SetModelSeed(4))
	if err != nil {
		t.Fatalf("MockBackendInitializer.New error %s", err.Error())
	}

	if mb.Name() != "MOCK" {
		t.Fatalf("MockBackendInitializer.New name mismatch %s", mb.Name())
	}
}
