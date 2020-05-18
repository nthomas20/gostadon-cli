package writer

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilya1st/rotatewriter"
)

func rotateLog(logWriter *rotatewriter.RotateWriter, rotateSignal chan (bool)) {
	for {
		doRotate := <-rotateSignal

		if doRotate == true && logWriter.RotationInProgress() == false {
			err := logWriter.Rotate(nil)

			fmt.Println(err)
		}
	}
}

func checkLogFileSize(logWriter *rotatewriter.RotateWriter, maxBytes int64, rotateSignal chan (bool)) {
	for {
		time.Sleep(10000 * time.Millisecond)

		if logWriter.RotationInProgress() == false {
			if fileStat, err := os.Stat(logWriter.Filename); err == nil {
				fileSize := fileStat.Size()

				if fileSize > maxBytes {
					rotateSignal <- true
				}
			}
		}
	}
}

// Configure : Configure a rotating log writer
func Configure(filename string, maxFiles int, maxBytes int64, prefix string, flagOptional ...int) (*log.Logger, error) {
	var (
		returnLogger *log.Logger
		err          error
		flags        = log.LstdFlags
	)

	rotateSignal := make(chan bool, 1)

	// We'll take an optional log flags parameter
	if len(flagOptional) > 0 {
		flags = flagOptional[0]
	}

	if logWriter, err := rotatewriter.NewRotateWriter(filename, maxFiles); err == nil {
		returnLogger = log.New(logWriter, prefix, flags|log.LUTC)
		go rotateLog(logWriter, rotateSignal)
		go checkLogFileSize(logWriter, maxBytes, rotateSignal)
	}

	return returnLogger, err
}
