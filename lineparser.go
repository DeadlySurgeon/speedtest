package speedtest

import (
	"encoding/json"
	"fmt"
)

// lineParser implements io.Writer so that it may have STDOUT piped to it with
// io.Copy. It then watches each line written to it, as the speedtest cli
// outputs the results as jsonl.
type lineParser struct {
	buffer      []byte
	mapCallback func(messageType string, bytes []byte) error
}

// Write fulfilles the io.Writer implementation.
func (lp *lineParser) Write(p []byte) (n int, err error) {
	// Read each of the bytes.
	for i := 0; i < len(p); i++ {
		// If it's not a new line.
		// Note:
		// - If it's windows, we're fucked as windows lines to use `\r`.
		if p[i] != '\n' {
			// Store it in the buffer and move along.
			lp.buffer = append(lp.buffer, p[i])
			continue
		}

		// Call our parser.
		if err = lp.parseBuffer(); err != nil {
			return -1, err
		}

		// Reset buffer.
		lp.buffer = []byte{}
	}

	// Return that we read all of the provided bytes.
	return len(p), nil
}

// parseBuffer is a breakdown of reading the buffer we've built up, broken out
// of the Write function to enable better testing.
func (lp *lineParser) parseBuffer() error {
	if len(lp.buffer) == 0 {
		return nil
	}
	// Create a store for our data. Since we don't know the types ahead of time
	// it makes more since to just store it all as a map and then just convert
	// the values over after so we don't have to call json.Unmarshal twice.
	// While I think that our jsonl will be consistent with the order of fields
	// it isn't something that we should rely on for parsing.
	var data map[string]interface{}

	// Unmarshal our data into our storage.
	if err := json.Unmarshal(lp.buffer, &data); err != nil {
		return fmt.Errorf("Failed to parse output: %w", err)
	}

	messagetype, ok := data["type"].(string)
	if !ok {
		return fmt.Errorf("Message didn't contain type")
	}

	// Let our caller know that we've got a line to read.
	if err := lp.mapCallback(messagetype, lp.buffer); err != nil {
		return fmt.Errorf("Failed to parse output: %w", err)
	}

	return nil
}

func (r *Results) mapCallback(messageType string, bytes []byte) error {
	if messageType == "result" {
		if err := json.Unmarshal(bytes, r); err != nil {
			return err
		}
		r.found = true
	}
	return nil
}
