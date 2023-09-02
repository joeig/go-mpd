package mpd_test

import (
	"encoding/xml"
	"errors"
	"fmt"
	"go.eigsys.de/go-mpd"
	"io"
	"log"
	"os"
	"path"
	"reflect"
	"testing"
)

type TestCase struct {
	fixture string
	wantMPD *mpd.MPD
	wantErr error
}

var testCases = []TestCase{
	{
		fixture: "zencoder/events.mpd",
		wantMPD: &mpd.MPD{
			XMLNS:                 "urn:mpeg:dash:schema:mpd:2011",
			Profiles:              "urn:mpeg:dash:profile:isoff-live:2011",
			Type:                  "dynamic",
			AvailabilityStartTime: "1970-01-01T00:00:00Z",
			MinBufferTime:         "PT1.97S",
			Period: []mpd.Period{{
				EventStream: []mpd.EventStream{{
					SchemeIdURI: "urn:example:eventstream",
					Value:       "eventstream",
					Timescale:   10,
					Event: []mpd.Event{
						{
							ID:               "event-0",
							PresentationTime: 100,
							Duration:         50,
						},
						{
							ID:               "event-1",
							PresentationTime: 200,
							Duration:         50,
						},
					},
				}},
			}},
			UTCTiming: []mpd.Descriptor{{}},
		},
	},
	{
		fixture: "zencoder/invalid.mpd",
		wantErr: errors.Join(mpd.ErrUnmarshalMPD, &xml.SyntaxError{
			Msg:  "unexpected EOF",
			Line: 3,
		}),
	},
	{
		fixture: "zencoder/location.mpd",
		wantMPD: &mpd.MPD{
			XMLNS:                 "urn:mpeg:dash:schema:mpd:2011",
			Profiles:              "urn:mpeg:dash:profile:isoff-live:2011",
			Type:                  "dynamic",
			AvailabilityStartTime: "1970-01-01T00:00:00Z",
			MinimumUpdatePeriod:   "PT5S",
			PublishTime:           "1970-01-01T00:00:00Z",
			Location:              []string{"https://example.com/location.mpd"},
		},
	},
}

func mustOpenFixture(fixture string) io.ReadCloser {
	handle, err := os.Open(path.Join("testdata", fixture))
	if err != nil {
		log.Fatalf("%v", err)
	}

	return handle
}

func TestMPD_Read(t *testing.T) {
	for _, testCase := range testCases {
		reader := mustOpenFixture(testCase.fixture)

		var testMPD mpd.MPD
		err := testMPD.Read(reader)

		if testCase.wantErr == nil && err != nil {
			t.Errorf("%s: unexpected error", testCase.fixture)
		}

		if testCase.wantErr != nil {
			if err == nil {
				t.Errorf("%s: error expected", testCase.fixture)
			} else {
				if !reflect.DeepEqual(err, testCase.wantErr) {
					t.Errorf("%s: wrong error", testCase.fixture)
				}
			}
		}

		if testCase.wantMPD != nil && !reflect.DeepEqual(testMPD, *testCase.wantMPD) {
			t.Errorf("%s: wrong MPD", testCase.fixture)
		}
	}
}

var exampleReader = mustOpenFixture("zencoder/segment_timeline_multi_period.mpd")

func ExampleMPD_Read() {
	var exampleMPD mpd.MPD
	if err := exampleMPD.Read(exampleReader); err != nil {
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
