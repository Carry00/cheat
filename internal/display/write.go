package display

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cheat/cheat/internal/config"
)

// Write writes output either directly to stdout, or through a pager,
// depending upon configuration. If AI processing is enabled, the output
// will be processed by AI before being displayed.
func Write(out string, conf config.Config) {
	// Process with AI if enabled
	var processedOut string
	var err error
	
	if conf.AIEnabled {
		processedOut, err = processWithAI(out, conf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "AI processing failed: %v\n", err)
			processedOut = out // Fallback to original output
		}
	} else {
		processedOut = out
	}

	// if no pager was configured, print the output to stdout and exit
	if conf.Pager == "" {
		fmt.Print(processedOut)
		os.Exit(0)
	}

	// otherwise, pipe output through the pager
	parts := strings.Split(conf.Pager, " ")
	pager := parts[0]
	args := parts[1:]

	// configure the pager
	cmd := exec.Command(pager, args...)
	cmd.Stdin = strings.NewReader(processedOut)
	cmd.Stdout = os.Stdout

	// run the pager and handle errors
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write to pager: %v\n", err)
		os.Exit(1)
	}
}
