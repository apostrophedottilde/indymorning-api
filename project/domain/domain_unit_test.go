package domain_test

import (
	"os"
	"testing"

	"github.com/apostrohedottilde/indymorning/api/project/domain"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestProject_Cancel_From_Live(t *testing.T) {
	underTest := &domain.GameProject{State: "LIVE"}

	if underTest.State != "LIVE" {
		t.Error("Expected initial state to be 'LIVE'")
	}

	underTest.Cancel()

	if underTest.State != "CANCELLED" {
		t.Error("Expected state to be 'CANCELLED' after invoking Cancel() function.")
	}
}

func TestProject_Cancel_From_Cancelled(t *testing.T) {
	underTest := &domain.GameProject{State: "CANCELLED"}

	if underTest.State != "CANCELLED" {
		t.Error("Expected initial state to be 'CANCELLED'")
	}

	_, err := underTest.Cancel()

	if err.Error() != "Cannot transition a cancelled project into 'CANCELLED' state" {
		t.Error("Expected state to be 'CANCELLED' after invoking Cancel() function.")
	}
}
