package speedtest

import (
	"errors"
	"io"
	"strings"
	"testing"
)

func TestLineParserWrite(t *testing.T) {
	var (
		validJSON = `{"type":"result", "field": "thing"}` + "\n"
		tests     = []struct {
			name          string
			err           bool
			lp            lineParser
			input         string
			bufferCleared bool
		}{
			{
				name:          "Valid Test",
				err:           false,
				bufferCleared: true,
				lp: lineParser{
					mapCallback: func(messageType string, bytes []byte) error {
						if messageType != "result" {
							t.Fatalf("Expected message type \"result\". Got \"%v\"\n", messageType)
						}
						if len(bytes) != len(validJSON)-1 {
							t.Fatalf("Expected byte length of %v. Got %v\n", len(validJSON)-1, len(bytes))
						}
						return nil
					},
				},
				input: validJSON,
			},
			{
				name: "Parse Buffer Error",
				err:  true,
				lp: lineParser{
					mapCallback: func(string, []byte) error {
						return errors.New("Intential Error")
					},
				},
				input: validJSON,
			},
			{
				name: "Empty Buffer",
				err:  false,
				lp: lineParser{
					mapCallback: func(string, []byte) error {
						return nil
					},
				},
				bufferCleared: true,
				input:         validJSON + "\n",
			},
			{
				name: "No Line Termination",
				err:  false,
				lp: lineParser{
					mapCallback: func(string, []byte) error {
						t.Fatalf("Callback got called despite no termination.")
						return nil
					},
				},
				input: validJSON[:len(validJSON)-1],
			},
		}
	)

	for _, test := range tests {
		// Extra careful about scope. I blame my coworker at C1.
		test := test
		t.Run(test.name, func(t *testing.T) {
			read, err := io.Copy(&test.lp, strings.NewReader(test.input))
			if (err != nil) != test.err {
				t.Fatalf("Expected err to be %v\nGot: %v\n", test.err, err)
			}
			if err == nil && (int(read) != len(test.input)) {
				t.Fatalf("Expected %v bytes to be read, only read %v\n", len(test.input), read)
			}

			if test.bufferCleared != (len(test.lp.buffer) == 0) {
				t.Fatalf("Expected the buffered to be cleared, it wasn't.\nLength: %v\n", len(test.lp.buffer))
			}
		})
	}
}

func TestParseBuffer(t *testing.T) {
	var (
		validJSON            = `{"type":"result", "field": "thing"}`
		invalidJSON          = `{"name": "bob"`
		invalidJSONStructure = `{"name": "bob"}`
		tests                = []struct {
			name string
			err  bool
			lp   lineParser
		}{
			{
				name: "Valid Test",
				err:  false,
				lp: lineParser{
					buffer: []byte(validJSON),
					mapCallback: func(messageType string, bytes []byte) error {
						if messageType != "result" {
							t.Fatalf("Expected message type \"result\". Got \"%v\"\n", messageType)
						}

						if validJSON != string(bytes) {
							t.Fatalf("Invalid input of bytes")
						}

						return nil
					},
				},
			},
			{
				name: "Invalid JSON",
				err:  true,
				lp: lineParser{
					buffer: []byte(invalidJSON),
				},
			},
			{
				name: "Invalid JSON Structure",
				err:  true,
				lp: lineParser{
					buffer: []byte(invalidJSONStructure),
				},
			},
			{
				name: "Error in MapCallback",
				err:  true,
				lp: lineParser{
					buffer: []byte(validJSON),
					mapCallback: func(string, []byte) error {
						return errors.New("Intential Error")
					},
				},
			},
		}
	)

	for _, test := range tests {
		// Extra careful about scope. I blame my coworker at C1.
		test := test
		t.Run(test.name, func(t *testing.T) {
			err := test.lp.parseBuffer()
			if (err != nil) != test.err {
				t.Fatalf("Expected err to be %v\nGot: %v\n", test.err, err)
			}
		})
	}
}

func TestMapCallback(t *testing.T) {
	t.Run("InvalidJSON", func(t *testing.T) {
		r := Results{}
		if err := r.mapCallback("result", []byte("Invalid JSON :D")); err == nil {
			t.Fatalf("Expected there to be an error unmarshalling an invalid JSON Object.")
		}
	})
}
