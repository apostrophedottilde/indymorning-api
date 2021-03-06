package user_test

import (
	"os"
	"testing"

	"github.com/apostrophedottilde/indymorning-api/project"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestProject_Cancel_From_Live(t *testing.T) {
	underTest := &project.GameProject{State: "LIVE"}

	if underTest.State != "LIVE" {
		t.Error("Expected initial state to be 'LIVE'")
	}

	underTest.Cancel()

	if underTest.State != "CANCELLED" {
		t.Error("Expected state to be 'CANCELLED' after invoking Cancel() function.")
	}
}

func TestProject_Cancel_From_Cancelled(t *testing.T) {
	underTest := &project.GameProject{State: "CANCELLED"}

	if underTest.State != "CANCELLED" {
		t.Error("Expected initial state to be 'CANCELLED'")
	}

	_, err := underTest.Cancel()

	if err.Error() != "Cannot transition a cancelled project into 'CANCELLED' state" {
		t.Error("Expected state to be 'CANCELLED' after invoking Cancel() function.")
	}
}
