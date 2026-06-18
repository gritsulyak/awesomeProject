//go:build test_integration

// go test --tags=test_integration ./test/...
package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/gritsulyak/awesomeProject/internal/application"
)

func TestMain(m *testing.M) {
	fmt.Println("TestMain")
	os.Exit(m.Run())
}

func createAndRunApp(ctx context.Context) *application.App {
	app := application.NewApp()
	if err := app.Start(ctx); err != nil {
		panic(err)
	}
	return app
}
