package lcdbackpack

import logger "github.com/sconklin/go-logger"

// You can manage verbosity of log output
// in the package by changing last parameter value.
var loglcd = logger.NewPackageLogger("lcd-backpack",
	// logger.DebugLevel,
	logger.InfoLevel,
)
