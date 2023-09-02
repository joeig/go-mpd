package mpd_test

import (
	"fmt"
	"go.eigsys.de/go-mpd"
	"io"
	"log"
	"os"
	"path"
)

func mustOpenFixture(id string) io.ReadCloser {
	handle, err := os.Open(path.Join("testdata", id))
	if err != nil {
		log.Fatalf("%v", err)
	}

	return handle
}

var reader = mustOpenFixture("zencoder/segment_timeline_multi_period.mpd")

func ExampleMPD_Read() {
	var exampleMPD mpd.MPD
	if err := exampleMPD.Read(reader); err != nil {
		log.Fatalf("%v", err)
	}

	// Output:
}

func ExampleMPD_Bytes() {
	exampleMPD := mpd.MPD{
		XMLNS:         mpd.MPD2011Namespace,
		Profiles:      mpd.ISOFFOnDemand2011Profile,
		Type:          mpd.StaticPresentation,
		MinBufferTime: "PT2S",
		Period:        []mpd.Period{{ID: "0"}},
	}

	bytes, err := exampleMPD.Bytes()
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("%s", bytes)

	// Output: <?xml version="1.0" encoding="UTF-8"?>
	// <MPD xmlns="urn:mpeg:dash:schema:mpd:2011" profiles="urn:mpeg:dash:profile:isoff-on-demand:2011" type="static" minBufferTime="PT2S">
	//   <Period id="0"></Period>
	// </MPD>
}
