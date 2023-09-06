package mpd_test

import (
	"aqwari.net/xml/xmltree"
	"encoding/xml"
	"errors"
	"github.com/google/go-cmp/cmp"
	"go.eigsys.de/go-mpd"
	"io"
	"log"
	"os"
	"path"
	"reflect"
	"testing"
)

func mustOpenFixture(fixture string) io.ReadCloser {
	handle, err := os.Open(path.Join("testdata", fixture))
	if err != nil {
		log.Fatalf("%v", err)
	}

	return handle
}

func TestNew(t *testing.T) {
	testMPD := mpd.New()
	wantMPD := &mpd.MPD{XMLNS: "urn:mpeg:dash:schema:mpd:2011"}

	if diff := cmp.Diff(testMPD, wantMPD); diff != "" {
		t.Errorf("wrong MPD: %s", diff)
	}
}

func TestRead(t *testing.T) {
	type TestCase struct {
		fixture string
		wantMPD *mpd.MPD
		wantErr error
	}

	testCases := []TestCase{
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
				MinBufferTime:         "PT2S",
				Location:              []string{"https://example.com/location.mpd"},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.fixture, func(t *testing.T) {
			reader := mustOpenFixture(testCase.fixture)
			testMPD, err := mpd.Read(reader)

			if testCase.wantErr == nil && err != nil {
				t.Error("unexpected error")
			}

			if testCase.wantErr != nil {
				if err == nil {
					t.Error("error expected")
				} else {
					if !reflect.DeepEqual(err, testCase.wantErr) {
						t.Error("wrong error")
					}
				}
			}

			if testCase.wantMPD != nil {
				if diff := cmp.Diff(testMPD, testCase.wantMPD); diff != "" {
					t.Errorf("wrong MPD: %s", diff)
				}
			}
		})
	}
}

type NopReadCloser struct{}

func (r NopReadCloser) Read(_ []byte) (n int, err error) {
	return 0, errors.New("error")
}

func (r NopReadCloser) Close() error {
	return nil
}

func TestRead_ErrReadMPD(t *testing.T) {
	reader := NopReadCloser{}
	if _, err := mpd.Read(reader); !errors.Is(err, mpd.ErrReadMPD) {
		t.Error("unexpected error")
	}
}

func TestRoundTrip(t *testing.T) {
	testCases := []string{
		"zencoder/adaptationset_switching.mpd",
		"zencoder/audio_channel_configuration.mpd",
		"zencoder/events.mpd",
		"zencoder/hbbtv_profile.mpd",
		"zencoder/inband_event_stream.mpd",
		"zencoder/live_profile.mpd",
		"zencoder/live_profile_dynamic.mpd",
		"zencoder/live_profile_multi_base_url.mpd",
		"zencoder/location.mpd",
		"zencoder/multiple_supplementals.mpd",
		"zencoder/newperiod.mpd",
		"zencoder/ondemand_profile.mpd",
		"zencoder/segment_list.mpd",
		"zencoder/segment_timeline.mpd",
		"zencoder/segment_timeline_multi_period.mpd",
		"zencoder/truncate.mpd",
		"zencoder/truncate_short.mpd",
	}

	for _, testCase := range testCases {
		t.Run(testCase, func(t *testing.T) {
			reader := mustOpenFixture(testCase)

			testMPD, err := mpd.Read(reader)
			if err != nil {
				t.Error("unexpected error")
			}

			output, err := testMPD.Bytes()
			if err != nil {
				t.Error("unexpected error")
			}

			outputDoc, err := xmltree.Parse(output)
			if err != nil {
				t.Error("unexpected error")
			}

			input, err := io.ReadAll(mustOpenFixture(testCase))
			if err != nil {
				t.Error("unexpected error")
			}

			inputDoc, err := xmltree.Parse(input)
			if err != nil {
				t.Error("unexpected error")
			}

			if !xmltree.Equal(inputDoc, outputDoc) {
				inputStr := xmltree.MarshalIndent(inputDoc, "", "  ")
				outputStr := xmltree.MarshalIndent(outputDoc, "", "  ")

				diff := cmp.Diff(inputStr, outputStr)
				t.Errorf("wrong MPD: %s", diff)
			}
		})
	}
}
