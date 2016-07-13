package logger

import (
	"os"
	"testing"
)

func TestDefalutLevel(t *testing.T) {
	if Level != INFO {
		t.Log("Wrong default log level", Level)
	}
}

func TestLevelValue(t *testing.T) {
	if DEBUG != 1 {
		t.Log("DEBUG value changed", DEBUG)
	}
	if INFO != 2 {
		t.Log("INFO value changed", INFO)
	}
	if WARNING != 4 {
		t.Log("WARNING value changed", WARNING)
	}
	if ERROR != 8 {
		t.Log("ERROR value changed", ERROR)
	}
}

func TestSetLevel(t *testing.T) {
	SetLevel(DEBUG)
	if Level != DEBUG {
		t.Log("Wrong log level", Level)
	}
}

func TestInitLogLevel(t *testing.T) {
	os.Setenv("LOG_LEVEL", "DEBUG")
	initLogLevel()
	if Level != DEBUG {
		t.Log("Log level not initialized", Level)
	}
	os.Setenv("LOG_LEVEL", "INFO")
	initLogLevel()
	if Level != INFO {
		t.Log("Log level not initialized", Level)
	}
	os.Setenv("LOG_LEVEL", "WARNING")
	initLogLevel()
	if Level != WARNING {
		t.Log("Log level not initialized", Level)
	}
	os.Setenv("LOG_LEVEL", "ERROR")
	initLogLevel()
	if Level != ERROR {
		t.Log("Log level not initialized", Level)
	}
}
