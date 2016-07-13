# Levels for logging

[![GoDoc](https://godoc.org/github.com/baijum/logger?status.svg)](https://godoc.org/github.com/baijum/logger)
[![Go Report Card](https://goreportcard.com/badge/github.com/baijum/logger)](https://goreportcard.com/report/github.com/baijum/logger)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://travis-ci.org/baijum/logger.svg?branch=master)](https://travis-ci.org/baijum/logger)

To get package:

	go get github.com/baijum/logger

Usage:

	if logger.Level <= logger.INFO {
		log.Printf("Some log message")
	}

	if logger.Level <= logger.WARNING {
		log.Printf("Some log message")
	}

