package mpd_test

import (
	"fmt"
	"go.eigsys.de/go-mpd"
	"log"
)

func ExampleNew() {
	example := mpd.New()
	example.MinBufferTime = "PT2S"
	example.Period = []mpd.Period{{ID: "period-0"}}
	// Output:
}

var exampleReader = mustOpenFixture("zencoder/segment_timeline_multi_period.mpd")

func ExampleRead() {
	example, err := mpd.Read(exampleReader)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("%s", example.Type)
	// Output: static
}

func ExampleMPD_Bytes() {
	example := mpd.New()
	example.Profiles = mpd.OnDemand2011Profile
	example.Type = mpd.StaticPresentationType
	example.MinBufferTime = "PT2S"
	example.Period = []mpd.Period{{ID: "period-0"}}

	exampleBytes, err := example.Bytes()
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("%s", exampleBytes)
	// Output: <?xml version="1.0" encoding="UTF-8"?>
	// <MPD xmlns="urn:mpeg:dash:schema:mpd:2011" profiles="urn:mpeg:dash:profile:isoff-on-demand:2011" type="static" minBufferTime="PT2S">
	//   <Period id="period-0"></Period>
	// </MPD>
}
