# Levels for logging

[![GoDoc](https://godoc.org/github.com/baijum/logger?status.svg)](https://godoc.org/github.com/baijum/logger)
[![Go Report Card](https://goreportcard.com/badge/github.com/baijum/logger)](https://goreportcard.com/report/github.com/baijum/logger)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://travis-ci.org/baijum/logger.svg?branch=master)](https://travis-ci.org/baijum/logger)

To get package:

	go get github.com/baijum/logger

Environment variable "LOG_LEVEL" set the log level.  These are the
possible values: DEBUG, INFO, WARNING, and ERROR

Log level can be set from the code like this:

	logger.SetLevel(logger.WARNING)

To log a message, compare the current level with available levels:

	if logger.Level <= logger.INFO {
		log.Printf("Some log message")
	}

	if logger.Level <= logger.WARNING {
		log.Printf("Some log message")
	}
