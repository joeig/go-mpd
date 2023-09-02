package mpd

import (
	"bytes"
	"encoding/xml"
	"io"
)

type AdaptationSet struct {
	Items                     []string           `xml:",any"`
	FramePacking              []Descriptor       `xml:"FramePacking,omitempty"`
	AudioChannelConfiguration []Descriptor       `xml:"AudioChannelConfiguration,omitempty"`
	ContentProtection         []Descriptor       `xml:"ContentProtection,omitempty"`
	EssentialProperty         []Descriptor       `xml:"EssentialProperty,omitempty"`
	SupplementalProperty      []Descriptor       `xml:"SupplementalProperty,omitempty"`
	InbandEventStream         []EventStream      `xml:"InbandEventStream,omitempty"`
	Accessibility             []Descriptor       `xml:"Accessibility,omitempty"`
	Role                      []Descriptor       `xml:"Role,omitempty"`
	Rating                    []Descriptor       `xml:"Rating,omitempty"`
	Viewpoint                 []Descriptor       `xml:"Viewpoint,omitempty"`
	ContentComponent          []ContentComponent `xml:"ContentComponent,omitempty"`
	BaseURL                   []BaseURL          `xml:"BaseURL,omitempty"`
	Representation            []Representation   `xml:"Representation,omitempty"`
	Actuate                   string             `xml:"actuate,attr,omitempty"`
	Group                     uint               `xml:"group,attr,omitempty"`
	Lang                      string             `xml:"lang,attr,omitempty"`
	ContentType               ContentType        `xml:"contentType,attr,omitempty"`
	Par                       Ratio              `xml:"par,attr,omitempty"`
	MinBandwidth              uint               `xml:"minBandwidth,attr,omitempty"`
	MaxBandwidth              uint               `xml:"maxBandwidth,attr,omitempty"`
	MinWidth                  uint               `xml:"minWidth,attr,omitempty"`
	MaxWidth                  uint               `xml:"maxWidth,attr,omitempty"`
	MinHeight                 uint               `xml:"minHeight,attr,omitempty"`
	MaxHeight                 uint               `xml:"maxHeight,attr,omitempty"`
	MinFrameRate              FrameRate          `xml:"minFrameRate,attr,omitempty"`
	MaxFrameRate              FrameRate          `xml:"maxFrameRate,attr,omitempty"`
	SegmentAlignment          ConditionalUint    `xml:"segmentAlignment,attr,omitempty"`
	SubsegmentAlignment       ConditionalUint    `xml:"subsegmentAlignment,attr,omitempty"`
	SubsegmentStartsWithSAP   uint               `xml:"subsegmentStartsWithSAP,attr,omitempty"`
	Profiles                  string             `xml:"profiles,attr,omitempty"`
	Width                     uint               `xml:"width,attr,omitempty"`
	Height                    uint               `xml:"height,attr,omitempty"`
	Sar                       Ratio              `xml:"sar,attr,omitempty"`
	FrameRate                 FrameRate          `xml:"frameRate,attr,omitempty"`
	AudioSamplingRate         string             `xml:"audioSamplingRate,attr,omitempty"`
	MimeType                  string             `xml:"mimeType,attr,omitempty"`
	SegmentProfiles           string             `xml:"segmentProfiles,attr,omitempty"`
	Codecs                    string             `xml:"codecs,attr,omitempty"`
	MaximumSAPPeriod          float64            `xml:"maximumSAPPeriod,attr,omitempty"`
	StartWithSAP              uint               `xml:"startWithSAP,attr,omitempty"`
	MaxPlayoutRate            float64            `xml:"maxPlayoutRate,attr,omitempty"`
	CodingDependency          bool               `xml:"codingDependency,attr,omitempty"`
	ScanType                  VideoScan          `xml:"scanType,attr,omitempty"`
}

func (a *AdaptationSet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AdaptationSet

	var overlay struct {
		*T
		Actuate                 *string          `xml:"actuate,attr,omitempty"`
		SegmentAlignment        *ConditionalUint `xml:"segmentAlignment,attr,omitempty"`
		SubsegmentAlignment     *ConditionalUint `xml:"subsegmentAlignment,attr,omitempty"`
		SubsegmentStartsWithSAP *uint            `xml:"subsegmentStartsWithSAP,attr,omitempty"`
	}

	overlay.T = (*T)(a)
	overlay.Actuate = &overlay.T.Actuate
	overlay.SegmentAlignment = &overlay.T.SegmentAlignment
	overlay.SubsegmentAlignment = &overlay.T.SubsegmentAlignment
	overlay.SubsegmentStartsWithSAP = &overlay.T.SubsegmentStartsWithSAP

	return d.DecodeElement(&overlay, &start)
}

type BaseURL struct {
	Value                    string  `xml:",chardata"`
	ServiceLocation          string  `xml:"serviceLocation,attr,omitempty"`
	ByteRange                string  `xml:"byteRange,attr,omitempty"`
	AvailabilityTimeOffset   float64 `xml:"availabilityTimeOffset,attr,omitempty"`
	AvailabilityTimeComplete bool    `xml:"availabilityTimeComplete,attr,omitempty"`
}

type ConditionalUint string

type ContentComponent struct {
	Items         []string     `xml:",any"`
	Accessibility []Descriptor `xml:"Accessibility,omitempty"`
	Role          []Descriptor `xml:"Role,omitempty"`
	Rating        []Descriptor `xml:"Rating,omitempty"`
	Viewpoint     []Descriptor `xml:"Viewpoint,omitempty"`
	Lang          string       `xml:"lang,attr,omitempty"`
	ContentType   string       `xml:"contentType,attr,omitempty"`
	Par           Ratio        `xml:"par,attr,omitempty"`
}

type Descriptor struct {
	Items          []string `xml:",any"`
	CENCPSSH       []string `xml:"urn:mpeg:cenc:2013 pssh,omitempty"`
	CENCDefaultKID []string `xml:"urn:mpeg:cenc:2013 default_KID,attr,omitempty"`
	// FIXME: MSPRPro doesn't work properly
	MSPRPro     []string    `xml:"urn:microsoft:playready pro,attr,omitempty"`
	SchemeIdURI SchemeIdURI `xml:"schemeIdUri,attr"`
	Value       string      `xml:"value,attr,omitempty"`
}

type SchemeIdURI string

const AudioChannelConfigurationSchemeIdURI SchemeIdURI = "urn:mpeg:dash:23003:3:audio_channel_configuration:2011"

type EventStream struct {
	Items       []string    `xml:",any"`
	Event       []Event     `xml:"Event,omitempty"`
	Actuate     string      `xml:"actuate,attr,omitempty"`
	SchemeIdURI SchemeIdURI `xml:"schemeIdUri,attr"`
	Value       string      `xml:"value,attr,omitempty"`
	Timescale   uint        `xml:"timescale,attr,omitempty"`
}

func (e *EventStream) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T EventStream

	var overlay struct {
		*T
		Actuate *string `xml:"actuate,attr,omitempty"`
	}

	overlay.T = (*T)(e)
	overlay.Actuate = &overlay.T.Actuate

	return d.DecodeElement(&overlay, &start)
}

type Event struct {
	Items            []string       `xml:",any"`
	Value            string         `xml:",chardata"`
	SCTE35Signal     []SCTE35Signal `xml:"urn:scte:scte35:2014:xml+bin Signal,omitempty"`
	ID               string         `xml:"id,attr,omitempty"`
	PresentationTime uint64         `xml:"presentationTime,attr,omitempty"`
	Duration         uint64         `xml:"duration,attr,omitempty"`
	MessageData      string         `xml:"messageData,attr,omitempty"`
}

func (e *Event) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Event

	var overlay struct {
		*T
		PresentationTime *uint64 `xml:"presentationTime,attr,omitempty"`
	}

	overlay.T = (*T)(e)
	overlay.PresentationTime = &overlay.T.PresentationTime

	return d.DecodeElement(&overlay, &start)
}

type SCTE35Signal struct {
	Items        []string `xml:",any"`
	SCTE35Binary string   `xml:"urn:scte:scte35:2014:xml+bin Binary,omitempty"`
}

// FrameRate must match the pattern `[0-9]*[0-9](/[0-9]*[0-9])?`.
type FrameRate string

type MPD struct {
	Items                      []string     `xml:",any"`
	BaseURL                    []BaseURL    `xml:"BaseURL,omitempty"`
	Location                   []string     `xml:"Location,omitempty"`
	Period                     []Period     `xml:"Period"`
	Metrics                    []Metrics    `xml:"Metrics,omitempty"`
	EssentialProperty          []Descriptor `xml:"EssentialProperty,omitempty"`
	SupplementalProperty       []Descriptor `xml:"SupplementalProperty,omitempty"`
	UTCTiming                  []Descriptor `xml:"UTCTiming,omitempty"`
	XMLNS                      string       `xml:"xmlns,attr,omitempty"`
	Profiles                   string       `xml:"profiles,attr"`
	Type                       Presentation `xml:"type,attr,omitempty"`
	AvailabilityStartTime      string       `xml:"availabilityStartTime,attr,omitempty"`
	AvailabilityEndTime        string       `xml:"availabilityEndTime,attr,omitempty"`
	PublishTime                string       `xml:"publishTime,attr,omitempty"`
	MediaPresentationDuration  string       `xml:"mediaPresentationDuration,attr,omitempty"`
	MinimumUpdatePeriod        string       `xml:"minimumUpdatePeriod,attr,omitempty"`
	MinBufferTime              string       `xml:"minBufferTime,attr"`
	TimeShiftBufferDepth       string       `xml:"timeShiftBufferDepth,attr,omitempty"`
	SuggestedPresentationDelay string       `xml:"suggestedPresentationDelay,attr,omitempty"`
	MaxSegmentDuration         string       `xml:"maxSegmentDuration,attr,omitempty"`
	MaxSubsegmentDuration      string       `xml:"maxSubsegmentDuration,attr,omitempty"`
}

func (m *MPD) Read(reader io.ReadCloser, inputLimit int64) error {
	body, err := io.ReadAll(io.LimitReader(reader, inputLimit))
	if err != nil {
		return err
	}

	_ = reader.Close()

	if err := xml.Unmarshal(body, m); err != nil {
		return err
	}

	return nil
}

func (m *MPD) Bytes() ([]byte, error) {
	xmlData, err := xml.MarshalIndent(m, "", "  ")
	if err != nil {
		return nil, err
	}

	return append([]byte(xml.Header), xmlData...), nil
}

func (m *MPD) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T MPD

	var overlay struct {
		*T
		Type *Presentation `xml:"type,attr,omitempty"`
	}

	overlay.T = (*T)(m)
	overlay.Type = &overlay.T.Type

	return d.DecodeElement(&overlay, &start)
}

type Metrics struct {
	Items     []string     `xml:",any"`
	Reporting []Descriptor `xml:"Reporting"`
	Range     []Range      `xml:"Range,omitempty"`
	Metrics   string       `xml:"metrics,attr"`
}

type ContentType string

const (
	AudioContentType     ContentType = "audio"
	VideoContentType     ContentType = "video"
	SubtitlesContentType ContentType = "text"
)

type Period struct {
	Items              []string        `xml:",any"`
	BaseURL            []BaseURL       `xml:"BaseURL,omitempty"`
	EventStream        []EventStream   `xml:"EventStream,omitempty"`
	AdaptationSet      []AdaptationSet `xml:"AdaptationSet,omitempty"`
	ID                 string          `xml:"id,attr,omitempty"`
	Actuate            string          `xml:"actuate,attr,omitempty"`
	Start              string          `xml:"start,attr,omitempty"`
	Duration           string          `xml:"duration,attr,omitempty"`
	BitstreamSwitching bool            `xml:"bitstreamSwitching,attr,omitempty"`
}

func (p *Period) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Period

	var overlay struct {
		*T
		Actuate            *string `xml:"actuate,attr,omitempty"`
		BitstreamSwitching *bool   `xml:"bitstreamSwitching,attr,omitempty"`
	}

	overlay.T = (*T)(p)
	overlay.Actuate = &overlay.T.Actuate
	overlay.BitstreamSwitching = &overlay.T.BitstreamSwitching

	return d.DecodeElement(&overlay, &start)
}

type Presentation string

const (
	StaticPresentation  Presentation = "static"
	DynamicPresentation Presentation = "dynamic"
)

type Range struct {
	Starttime string `xml:"starttime,attr,omitempty"`
	Duration  string `xml:"duration,attr,omitempty"`
}

// Ratio must match the pattern `[0-9]*:[0-9]*`.
type Ratio string

type RepresentationBase struct {
	Items                     []string      `xml:",any"`
	FramePacking              []Descriptor  `xml:"FramePacking,omitempty"`
	AudioChannelConfiguration []Descriptor  `xml:"AudioChannelConfiguration,omitempty"`
	ContentProtection         []Descriptor  `xml:"ContentProtection,omitempty"`
	EssentialProperty         []Descriptor  `xml:"EssentialProperty,omitempty"`
	SupplementalProperty      []Descriptor  `xml:"SupplementalProperty,omitempty"`
	InbandEventStream         []EventStream `xml:"InbandEventStream,omitempty"`
	Profiles                  string        `xml:"profiles,attr,omitempty"`
	Width                     uint          `xml:"width,attr,omitempty"`
	Height                    uint          `xml:"height,attr,omitempty"`
	Sar                       Ratio         `xml:"sar,attr,omitempty"`
	FrameRate                 FrameRate     `xml:"frameRate,attr,omitempty"`
	AudioSamplingRate         string        `xml:"audioSamplingRate,attr,omitempty"`
	MimeType                  string        `xml:"mimeType,attr,omitempty"`
	SegmentProfiles           string        `xml:"segmentProfiles,attr,omitempty"`
	Codecs                    string        `xml:"codecs,attr,omitempty"`
	MaximumSAPPeriod          float64       `xml:"maximumSAPPeriod,attr,omitempty"`
	StartWithSAP              uint          `xml:"startWithSAP,attr,omitempty"`
	MaxPlayoutRate            float64       `xml:"maxPlayoutRate,attr,omitempty"`
	CodingDependency          bool          `xml:"codingDependency,attr,omitempty"`
	ScanType                  VideoScan     `xml:"scanType,attr,omitempty"`
}

type Representation struct {
	Items                     []string        `xml:",any"`
	FramePacking              []Descriptor    `xml:"FramePacking,omitempty"`
	AudioChannelConfiguration []Descriptor    `xml:"AudioChannelConfiguration,omitempty"`
	ContentProtection         []Descriptor    `xml:"ContentProtection,omitempty"`
	EssentialProperty         []Descriptor    `xml:"EssentialProperty,omitempty"`
	SupplementalProperty      []Descriptor    `xml:"SupplementalProperty,omitempty"`
	InbandEventStream         []EventStream   `xml:"InbandEventStream,omitempty"`
	BaseURL                   []BaseURL       `xml:"BaseURL,omitempty"`
	SegmentTemplate           SegmentTemplate `xml:"SegmentTemplate,omitempty"`
	ID                        string          `xml:"id,attr,omitempty"`
	Bandwidth                 uint            `xml:"bandwidth,attr"`
	QualityRanking            uint            `xml:"qualityRanking,attr,omitempty"`
	DependencyId              StringVector    `xml:"dependencyId,attr,omitempty"`
	MediaStreamStructureId    StringVector    `xml:"mediaStreamStructureId,attr,omitempty"`
	Profiles                  string          `xml:"profiles,attr,omitempty"`
	Width                     uint            `xml:"width,attr,omitempty"`
	Height                    uint            `xml:"height,attr,omitempty"`
	Sar                       Ratio           `xml:"sar,attr,omitempty"`
	FrameRate                 FrameRate       `xml:"frameRate,attr,omitempty"`
	AudioSamplingRate         string          `xml:"audioSamplingRate,attr,omitempty"`
	MimeType                  string          `xml:"mimeType,attr,omitempty"`
	SegmentProfiles           string          `xml:"segmentProfiles,attr,omitempty"`
	Codecs                    string          `xml:"codecs,attr,omitempty"`
	MaximumSAPPeriod          float64         `xml:"maximumSAPPeriod,attr,omitempty"`
	StartWithSAP              uint            `xml:"startWithSAP,attr,omitempty"`
	MaxPlayoutRate            float64         `xml:"maxPlayoutRate,attr,omitempty"`
	CodingDependency          bool            `xml:"codingDependency,attr,omitempty"`
	ScanType                  VideoScan       `xml:"scanType,attr,omitempty"`
}

type S struct {
	T *uint64 `xml:"t,attr,omitempty"`
	N uint64  `xml:"n,attr,omitempty"`
	D uint64  `xml:"d,attr"`
	R int     `xml:"r,attr,omitempty"`
}

func (s *S) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T S

	var overlay struct {
		*T
		R *int `xml:"r,attr,omitempty"`
	}

	overlay.T = (*T)(s)
	overlay.R = &overlay.T.R

	return d.DecodeElement(&overlay, &start)
}

type SegmentBase struct {
	Items                    []string `xml:",any"`
	Initialization           URL      `xml:"Initialization,omitempty"`
	RepresentationIndex      URL      `xml:"RepresentationIndex,omitempty"`
	Timescale                uint     `xml:"timescale,attr,omitempty"`
	PresentationTimeOffset   uint64   `xml:"presentationTimeOffset,attr,omitempty"`
	IndexRange               string   `xml:"indexRange,attr,omitempty"`
	IndexRangeExact          bool     `xml:"indexRangeExact,attr,omitempty"`
	AvailabilityTimeOffset   float64  `xml:"availabilityTimeOffset,attr,omitempty"`
	AvailabilityTimeComplete bool     `xml:"availabilityTimeComplete,attr,omitempty"`
}

func (s *SegmentBase) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T SegmentBase

	var overlay struct {
		*T
		IndexRangeExact *bool `xml:"indexRangeExact,attr,omitempty"`
	}

	overlay.T = (*T)(s)
	overlay.IndexRangeExact = &overlay.T.IndexRangeExact

	return d.DecodeElement(&overlay, &start)
}

type SegmentList struct {
	Items                    []string        `xml:",any"`
	Initialization           URL             `xml:"Initialization,omitempty"`
	RepresentationIndex      URL             `xml:"RepresentationIndex,omitempty"`
	SegmentTimeline          SegmentTimeline `xml:"SegmentTimeline,omitempty"`
	BitstreamSwitching       URL             `xml:"BitstreamSwitching,omitempty"`
	Actuate                  string          `xml:"actuate,attr,omitempty"`
	Duration                 uint            `xml:"duration,attr,omitempty"`
	StartNumber              uint            `xml:"startNumber,attr,omitempty"`
	Timescale                uint            `xml:"timescale,attr,omitempty"`
	PresentationTimeOffset   uint64          `xml:"presentationTimeOffset,attr,omitempty"`
	IndexRange               string          `xml:"indexRange,attr,omitempty"`
	IndexRangeExact          bool            `xml:"indexRangeExact,attr,omitempty"`
	AvailabilityTimeOffset   float64         `xml:"availabilityTimeOffset,attr,omitempty"`
	AvailabilityTimeComplete bool            `xml:"availabilityTimeComplete,attr,omitempty"`
}

func (s *SegmentList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T SegmentList

	var overlay struct {
		*T
		Actuate         *string `xml:"actuate,attr,omitempty"`
		IndexRangeExact *bool   `xml:"indexRangeExact,attr,omitempty"`
	}

	overlay.T = (*T)(s)
	overlay.Actuate = &overlay.T.Actuate
	overlay.IndexRangeExact = &overlay.T.IndexRangeExact

	return d.DecodeElement(&overlay, &start)
}

type SegmentTemplate struct {
	Items                    []string        `xml:",any"`
	SegmentTimeline          SegmentTimeline `xml:"SegmentTimeline,omitempty"`
	Media                    string          `xml:"media,attr,omitempty"`
	Index                    string          `xml:"index,attr,omitempty"`
	InitializationAttr       string          `xml:"initialization,attr,omitempty"`
	BitstreamSwitchingAttr   string          `xml:"bitstreamSwitching,attr,omitempty"`
	Duration                 uint            `xml:"duration,attr,omitempty"`
	StartNumber              *uint           `xml:"startNumber,attr,omitempty"`
	Timescale                uint            `xml:"timescale,attr,omitempty"`
	PresentationTimeOffset   uint64          `xml:"presentationTimeOffset,attr,omitempty"`
	IndexRange               string          `xml:"indexRange,attr,omitempty"`
	IndexRangeExact          bool            `xml:"indexRangeExact,attr,omitempty"`
	AvailabilityTimeOffset   float64         `xml:"availabilityTimeOffset,attr,omitempty"`
	AvailabilityTimeComplete bool            `xml:"availabilityTimeComplete,attr,omitempty"`
}

func (s *SegmentTemplate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T SegmentTemplate

	var overlay struct {
		*T
		IndexRangeExact *bool `xml:"indexRangeExact,attr,omitempty"`
	}

	overlay.T = (*T)(s)
	overlay.IndexRangeExact = &overlay.T.IndexRangeExact

	return d.DecodeElement(&overlay, &start)
}

type SegmentTimeline struct {
	Items []string `xml:",any"`
	S     []S      `xml:"S"`
}

// StringNoWhitespace must match the pattern `[^\r\n\t \p{Z}]*`.
type StringNoWhitespace string

type StringVector []string

func (s *StringVector) MarshalText() ([]byte, error) {
	result := make([][]byte, 0, len(*s))

	for _, v := range *s {
		result = append(result, []byte(v))
	}

	return bytes.Join(result, []byte(" ")), nil
}

func (s *StringVector) UnmarshalText(text []byte) error {
	for _, v := range bytes.Fields(text) {
		*s = append(*s, string(v))
	}

	return nil
}

type URL struct {
	Items     []string `xml:",any"`
	SourceURL string   `xml:"sourceURL,attr,omitempty"`
	Range     string   `xml:"range,attr,omitempty"`
}

type VideoScan string

const (
	ProgressiveVideoScan VideoScan = "progressive"
	InterlacedVideoScan  VideoScan = "interlaced"
	UnknownVideoScan     VideoScan = "unknown"
)
