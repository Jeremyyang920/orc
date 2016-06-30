package orc

import (
	"testing"

	"code.simon-critchley.co.uk/orc/proto"
)

func TestCursor(t *testing.T) {

	r, err := Open("./examples/orc-file-11-format.orc")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	// Select a single column from the file.
	c := r.Select("boolean1")

	// Call Stripes to trigger reading the first stripe.
	s := c.Stripes()
	if !s {
		t.Errorf("Test failed, expected true, got false")
	}

	// Call Next to initialise the readers.
	n := c.Next()
	if !n {
		t.Errorf("Test failed, expected true, got false")
	}

	// There should be a data stream available for reading.
	stream := c.streams.get(streamName{1, proto.Stream_DATA})
	if stream == nil {
		t.Errorf("Test failed, got nil stream")
	}

	// There should also be a row index.
	stream = c.streams.get(streamName{1, proto.Stream_ROW_INDEX})
	if stream == nil {
		t.Errorf("Test failed, got nil stream")
	}

	if err := c.Err(); err != nil {
		t.Fatal(err)
	}

}