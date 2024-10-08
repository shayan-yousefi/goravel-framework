package console

import (
	"log"
	"os"
	"testing"

	supportdocker "github.com/goravel/framework/support/docker"
	"github.com/goravel/framework/support/env"
)

var testDatabaseDocker *supportdocker.Database

func TestMain(m *testing.M) {
	if !env.IsWindows() {
		var err error
		testDatabaseDocker, err = supportdocker.InitDatabase()
		if err != nil {
			log.Fatalf("Init docker error: %s", err)
		}
	}

	exit := m.Run()

	if !env.IsWindows() {
		defer func() {
			if err := testDatabaseDocker.Stop(); err != nil {
				log.Fatalf("Stop docker error: %s", err)
			}
		}()
	}

	os.Exit(exit)
}
