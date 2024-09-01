package timer

import (
	"time"

	"github.com/nestorlai1994/nesoli/logger"
)

func Timer(done chan bool, msg string, timeInterval time.Duration) {
	startTime := time.Now()
	go func() {
		for {
			select {
			case <-done:
				logger.Info("Elapsed time: " + time.Since(startTime).String())
				return
			default:
				// Update and print the progress or time counter
				// You can use a library like "github.com/cheggaaa/pb/v3" for progress bar
				// Or simply print the elapsed time using time.Since() function
				// Example: fmt.Println("Elapsed time:", time.Since(startTime))
				logger.Info(msg + time.Since(startTime).String())
			}
			time.Sleep(timeInterval) // Adjust the sleep duration as needed
		}
	}()
}
