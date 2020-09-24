package speedtest

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

// NewTest executes a test against speedtest.net, and blocks until it recieves
// results.
func NewTest() (Results, error) {
	var (
		r   = Results{}
		err error
	)

	// Execute the speedtest CLI. I originally wanted this to be pure Go, but
	// I ran into issues where existing libraries have trouble and I cannot be
	// bothered to make and maintain one myself. It's easier to just use this
	// tool installed on the operating system.
	cmd := exec.Command("speedtest", "-f", "jsonl")

	// lineParser will read the lines as they come in.
	lp := &lineParser{}
	lp.mapCallback = r.mapCallback
	cmd.Stdout = lp
	// TODO:
	// - Parse errors so we know what went wrong.
	cmd.Stderr = ioutil.Discard

	if err = cmd.Run(); err != nil {
		return r, fmt.Errorf("Failed to execute speedtest cli: %w", err)
	}

	if !r.found {
		return r, fmt.Errorf("Ran into issues")
	}

	return r, err
}
