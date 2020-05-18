package logs

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/nthomas20/gostadon-cli/logs/writer"

	jsoniter "github.com/json-iterator/go"
)

// Write : Write a log entry
func Write(logger *log.Logger, o interface{}, debugOptional ...bool) {
	var (
		json = jsoniter.ConfigCompatibleWithStandardLibrary
	)

	debug := false

	if len(debugOptional) > 0 {
		debug = debugOptional[0]
	}

	// Marshal the structure to JSON
	if logRow, err := json.Marshal(o); err == nil {
		logger.Println(string(logRow[:]))

		if debug == true {
			fmt.Println(string(logRow[:]))
		}
	}
}

func logInt64Conv(s string, n string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		fmt.Println("Invalid integer conversion [", n, "]:", s)
		os.Exit(1)
	}

	return i
}

// StartLogging : Start the logging for a particular log
func StartLogging(name string, filename string, maxFilesString string, maxBytesString string) *log.Logger {
	maxFiles := logInt64Conv(maxFilesString, name+"_LOG_ROTATE_MAX_FILES")
	maxBytes := logInt64Conv(maxBytesString, name+"_LOG_ROTATE_BYTES")

	logger, err := writer.Configure(filename, int(maxFiles), maxBytes, "")

	if err != nil {
		fmt.Println("ERROR WITH ", name, " LOGGING:", err)
		os.Exit(1)
	}

	return logger
}
