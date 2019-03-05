package user_test

import (
	"os"
	"testing"

	"github.com/apostrophedottilde/go-forum-api/forum"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestProject_Cancel_From_Live(t *testing.T) {
	underTest := &forum.Forum{State: "LIVE"}

	if underTest.State != "LIVE" {
		t.Error("Expected initial state to be 'LIVE'")
	}

	underTest.Cancel()

	if underTest.State != "CANCELLED" {
		t.Error("Expected state to be 'CANCELLED' after invoking Close() function.")
	}
}

func TestProject_Cancel_From_Cancelled(t *testing.T) {
	underTest := &forum.Forum{State: "CANCELLED"}

	if underTest.State != "CANCELLED" {
		t.Error("Expected initial state to be 'CANCELLED'")
	}

	_, err := underTest.Cancel()

	if err.Error() != "Cannot transition a cancelled project into 'CANCELLED' state" {
		t.Error("Expected state to be 'CANCELLED' after invoking Close() function.")
	}
}
