# Levels for logging

To get package:

	go get github.com/baijum/logger

Usage:

	if logger.Level <= logger.INFO {
		log.Printf("Some log message")
	}

	if logger.Level <= logger.WARNING {
		log.Printf("Some log message")
	}

