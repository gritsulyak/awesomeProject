//go:build test_integration

// go test --tags=test_integration ./test/...
package test

import (
	"context"
	"fmt"
	"github.com/BigDwarf/testci/internal/application"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("TestMain")
	os.Exit(m.Run())
}

func createAndRunApp(_ context.Context) *application.App {
	app := application.NewApp()
	app.Start()
	return app
}
