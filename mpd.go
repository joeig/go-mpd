// Package mpd parses, manipulates and builds MPEG-DASH (ISO/IEC 23009-1 5th edition) manifests.
package mpd

import (
	"errors"
	"go.eigsys.de/go-mpd/third_party/encoding/xml"
	"io"
)

type StringVector []string

type ListOf4CC []FourCC

// FourCC (4CC) as per latest ISO/IEC 14496-12.
type FourCC string

type UIntVector []uint

// SingleRFC7233Range must match the pattern `([0-9]*)(\-([0-9]*))?`.
type SingleRFC7233Range string

type VideoScan string

const (
	ProgressiveVideoScan VideoScan = "progressive"
	InterlacedVideoScan  VideoScan = "interlaced"
	UnknownVideoScan     VideoScan = "unknown"
)

// Codecs is a RFC6381 fancy-list without enclosing double quotes.
type Codecs string

type MIMEType string

const (
	VideoMP4MIMEType MIMEType = "video/mp4"
	AudioMP4MIMEType MIMEType = "audio/mp4"
	TextVTTMIMEType  MIMEType = "text/vtt"
)

type SchemeIDURI string

const (
	AudioChannelConfiguration2011SchemeIDURI SchemeIDURI = "urn:mpeg:dash:23003:3:audio_channel_configuration:2011"
	MP4Protection2011SchemeIDURI             SchemeIDURI = "urn:mpeg:dash:mp4protection:2011"
	FairPlaySchemeIDURI                      SchemeIDURI = "urn:uuid:94ce86fb-07ff-4f43-adb8-93d2fa968ca2"
	PlayReadySchemeIDURI                     SchemeIDURI = "urn:uuid:9a04f079-9840-4286-ab92-e65be0885f95"
	WidevineSchemeIDURI                      SchemeIDURI = "urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed"
)

type ContentEncoding string

const Base64ContentEncoding ContentEncoding = "base64"

// FrameRate must match the pattern `[0-9]+(/[1-9][0-9]*)?`.
type FrameRate string

type RFC6838ContentType string

const (
	TextRFC6838ContentType        RFC6838ContentType = "text"
	ImageRFC6838ContentType       RFC6838ContentType = "image"
	AudioRFC6838ContentType       RFC6838ContentType = "audio"
	VideoRFC6838ContentType       RFC6838ContentType = "video"
	ApplicationRFC6838ContentType RFC6838ContentType = "application"
	FontRFC6838ContentType        RFC6838ContentType = "font"
)

type Profile string

const (
	Live2011Profile      Profile = "urn:mpeg:dash:profile:isoff-live:2011"
	OnDemand2011Profile  Profile = "urn:mpeg:dash:profile:isoff-on-demand:2011"
	HbbTVLive2012Profile Profile = "urn:hbbtv:dash:profile:isoff-live:2012"
)

type Namespace string

const MPD2011Namespace Namespace = "urn:mpeg:dash:schema:mpd:2011"

type OperatingQualityMediaType string

const (
	VideoOperatingQualityMediaType OperatingQualityMediaType = "video"
	AudioOperatingQualityMediaType OperatingQualityMediaType = "audio"
	AnyOperatingQualityMediaType   OperatingQualityMediaType = "any"
)

// ListOfProfiles must be a comma-separated list.
type ListOfProfiles string

type OperatingBandwidthMediaType string

const (
	VideoOperatingBandwidthMediaType OperatingBandwidthMediaType = "video"
	AudioOperatingBandwidthMediaType OperatingBandwidthMediaType = "audio"
	AnyOperatingBandwidthMediaType   OperatingBandwidthMediaType = "any"
	AllOperatingBandwidthMediaType   OperatingBandwidthMediaType = "all"
)

type ContentType string

const (
	AudioContentType     ContentType = "audio"
	VideoContentType     ContentType = "video"
	SubtitlesContentType ContentType = "text"
)

type PresentationType string

const (
	StaticPresentationType  PresentationType = "static"
	DynamicPresentationType PresentationType = "dynamic"
)

// Ratio must match the pattern `[0-9]*:[0-9]*`.
type Ratio string

type AudioChannelConfiguration string

const DASH2011AudioChannelConfiguration AudioChannelConfiguration = "urn:mpeg:dash:23003:3:audio_channel_configuration:2011"

type Source string

const (
	ContentSource    Source = "content"
	StatisticsSource Source = "statistics"
	OtherSource      Source = "other"
)

type ProducerReferenceTimeType string

const (
	EncoderProducerReferenceTimeType     ProducerReferenceTimeType = "encoder"
	CapturedProducerReferenceTimeType    ProducerReferenceTimeType = "captured"
	ApplicationProducerReferenceTimeType ProducerReferenceTimeType = "application"
)

type AudioSamplingRate UIntVector

type SAPType uint

type Tag string

type SwitchingType string

const (
	MediaSwitchingType     SwitchingType = "media"
	BitstreamSwitchingType SwitchingType = "bitstream"
)

type RandomAccessType string

const (
	ClosedRandomAccessType  RandomAccessType = "closed"
	OpenRandomAccessType    RandomAccessType = "open"
	GradualRandomAccessType RandomAccessType = "gradual"
)

type PreselectionOrder string

const (
	UndefinedPreselectionOrder   PreselectionOrder = "undefined"
	TimeOrderedPreselectionOrder PreselectionOrder = "time-ordered"
	FullyOrdered                 PreselectionOrder = "fully-ordered"
)

var (
	ErrReadMPD      = errors.New("cannot read MPD")
	ErrUnmarshalMPD = errors.New("cannot unmarshal MPD")
	ErrMarshalMPD   = errors.New("cannot marshal MPD")
)

type AdaptationSet struct {
	RepresentationBase

	Accessibility    []Descriptor       `xml:"Accessibility,omitempty"`
	Role             []Descriptor       `xml:"Role,omitempty"`
	Rating           []Descriptor       `xml:"Rating,omitempty"`
	Viewpoint        []Descriptor       `xml:"Viewpoint,omitempty"`
	ContentComponent []ContentComponent `xml:"ContentComponent,omitempty"`
	BaseURL          []BaseURL          `xml:"BaseURL,omitempty"`
	SegmentBase      *SegmentBase       `xml:"SegmentBase,omitempty"`
	SegmentList      *SegmentList       `xml:"SegmentList,omitempty"`
	SegmentTemplate  *SegmentTemplate   `xml:"SegmentTemplate,omitempty"`
	Representation   []Representation   `xml:"Representation,omitempty"`
	ID               uint               `xml:"id,attr,omitempty"`
	Group            uint               `xml:"group,attr,omitempty"`
	Lang             string             `xml:"lang,attr,omitempty"`
	ContentType      ContentType        `xml:"contentType,attr,omitempty"`
	PAR              Ratio              `xml:"par,attr,omitempty"`
	MinBandwidth     uint               `xml:"minBandwidth,attr,omitempty"`
	MaxBandwidth     uint               `xml:"maxBandwidth,attr,omitempty"`
	MinWidth         uint               `xml:"minWidth,attr,omitempty"`
	MaxWidth         uint               `xml:"maxWidth,attr,omitempty"`
	MinHeight        uint               `xml:"minHeight,attr,omitempty"`
	MaxHeight        uint               `xml:"maxHeight,attr,omitempty"`
	MinFrameRate     FrameRate          `xml:"minFrameRate,attr,omitempty"`
	MaxFrameRate     FrameRate          `xml:"maxFrameRate,attr,omitempty"`

	// SegmentAlignment defaults to `false`.
	SegmentAlignment bool `xml:"segmentAlignment,attr,omitempty"`

	// SubsegmentAlignment defaults to `false`.
	SubsegmentAlignment bool `xml:"subsegmentAlignment,attr,omitempty"`

	// SubsegmentStartsWithSAP defaults to `0`.
	SubsegmentStartsWithSAP uint `xml:"subsegmentStartsWithSAP,attr,omitempty"`

	BitstreamSwitching      bool       `xml:"bitstreamSwitching,attr,omitempty"`
	InitializationSetRef    UIntVector `xml:"initializationSetRef,attr,omitempty"`
	InitializationPrincipal string     `xml:"initializationPrincipal,attr,omitempty"`
}

type BaseURL struct {
	Value                    string  `xml:",chardata"`
	ServiceLocation          string  `xml:"serviceLocation,attr,omitempty"`
	ByteRange                string  `xml:"byteRange,attr,omitempty"`
	AvailabilityTimeOffset   float64 `xml:"availabilityTimeOffset,attr,omitempty"`
	AvailabilityTimeComplete bool    `xml:"availabilityTimeComplete,attr,omitempty"`
	TimeShiftBufferDepth     string  `xml:"timeShiftBufferDepth,attr,omitempty"`

	// RangeAccess defaults to `false`.
	RangeAccess bool `xml:"rangeAccess,attr,omitempty"`
}

type ContentComponent struct {
	Items         []string           `xml:",any"`
	Accessibility []Descriptor       `xml:"Accessibility,omitempty"`
	Role          []Descriptor       `xml:"Role,omitempty"`
	Rating        []Descriptor       `xml:"Rating,omitempty"`
	Viewpoint     []Descriptor       `xml:"Viewpoint,omitempty"`
	Lang          string             `xml:"lang,attr,omitempty"`
	ContentType   RFC6838ContentType `xml:"contentType,attr,omitempty"`
	PAR           Ratio              `xml:"par,attr,omitempty"`
	Tag           Tag                `xml:"tag,attr,omitempty"`
}

type Descriptor struct {
	Items       []string    `xml:",any"`
	SchemeIDURI SchemeIDURI `xml:"schemeIdUri,attr,omitempty"`
	Value       string      `xml:"value,attr,omitempty"`
}

type EventStream struct {
	Items       []string    `xml:",any"`
	Event       []Event     `xml:"Event,omitempty"`
	SchemeIdURI SchemeIDURI `xml:"schemeIdUri,attr"`
	Value       string      `xml:"value,attr,omitempty"`
	Timescale   uint        `xml:"timescale,attr,omitempty"`

	// PresentationTimeOffset defaults to `0`.
	PresentationTimeOffset uint `xml:"presentationTimeOffset,attr,omitempty"`
}

type Event struct {
	Items           []string        `xml:",any"`
	Value           string          `xml:",chardata"`
	SelectionInfo   []SelectionInfo `xml:"SelectionInfo,omitempty"`
	SCTE35Signal    []SCTE35Signal  `xml:"urn:scte:scte35:2014:xml+bin Signal,omitempty"`
	ID              string          `xml:"id,attr,omitempty"`
	ContentEncoding ContentEncoding `xml:"contentEncoding,attr,omitempty"`

	// PresentationTime defaults to `0`.
	PresentationTime uint64 `xml:"presentationTime,attr,omitempty"`

	Duration    uint64 `xml:"duration,attr,omitempty"`
	MessageData string `xml:"messageData,attr,omitempty"`
}

type SelectionInfo struct {
	Selection     []Selection `xml:"Selection,omitempty"`
	SelectionInfo string      `xml:"selectionInfo,attr,omitempty"`
	ContactURL    string      `xml:"contactURL,attr"`
}

type Selection struct {
	DataEncoding ContentEncoding `xml:"dataEncoding,attr,omitempty"`
	Parameter    string          `xml:"parameter,attr"`
	Data         string          `xml:"data,attr"`
}

type SCTE35Signal struct {
	Items  []string `xml:",any"`
	Binary string   `xml:"urn:scte:scte35:2014:xml+bin Binary,omitempty"`
}

type MPD struct {
	Items                      []string               `xml:",any"`
	ProgramInformation         []ProgramInformation   `xml:"ProgramInformation,omitempty"`
	BaseURL                    []BaseURL              `xml:"BaseURL,omitempty"`
	Location                   []string               `xml:"Location,omitempty"`
	PatchLocation              []PatchLocation        `xml:"PatchLocation,omitempty"`
	ServiceDescription         []ServiceDescription   `xml:"ServiceDescription,omitempty"`
	InitializationSet          []InitializationSet    `xml:"InitializationSet,omitempty"`
	InitializationGroup        []UIntVWithID          `xml:"InitializationGroup,omitempty"`
	InitializationPresentation []UIntVWithID          `xml:"InitializationPresentation,omitempty"`
	ContentProtection          []ContentProtection    `xml:"ContentProtection,omitempty"`
	Period                     []Period               `xml:"Period"`
	Metrics                    []Metrics              `xml:"Metrics,omitempty"`
	EssentialProperty          []Descriptor           `xml:"EssentialProperty,omitempty"`
	SupplementalProperty       []Descriptor           `xml:"SupplementalProperty,omitempty"`
	UTCTiming                  []Descriptor           `xml:"UTCTiming,omitempty"`
	LeapSecondInformation      *LeapSecondInformation `xml:"LeapSecondInformation,omitempty"`
	XMLNS                      Namespace              `xml:"xmlns,attr,omitempty"`
	Profiles                   Profile                `xml:"profiles,attr"`

	// Type defaults to StaticPresentationType.
	Type PresentationType `xml:"type,attr,omitempty"`

	AvailabilityStartTime      string `xml:"availabilityStartTime,attr,omitempty"`
	AvailabilityEndTime        string `xml:"availabilityEndTime,attr,omitempty"`
	PublishTime                string `xml:"publishTime,attr,omitempty"`
	MediaPresentationDuration  string `xml:"mediaPresentationDuration,attr,omitempty"`
	MinimumUpdatePeriod        string `xml:"minimumUpdatePeriod,attr,omitempty"`
	MinBufferTime              string `xml:"minBufferTime,attr"`
	TimeShiftBufferDepth       string `xml:"timeShiftBufferDepth,attr,omitempty"`
	SuggestedPresentationDelay string `xml:"suggestedPresentationDelay,attr,omitempty"`
	MaxSegmentDuration         string `xml:"maxSegmentDuration,attr,omitempty"`
	MaxSubsegmentDuration      string `xml:"maxSubsegmentDuration,attr,omitempty"`
}

// New creates a new instance of MPD and sets the XML namespace to MPD2011Namespace.
func New() *MPD {
	return &MPD{XMLNS: MPD2011Namespace}
}

// Read creates a new instance of MPD and reads the content from an io.ReadCloser.
func Read(reader io.ReadCloser) (*MPD, error) {
	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.Join(ErrReadMPD, err)
	}

	_ = reader.Close()

	mpd := &MPD{}
	if err := xml.Unmarshal(body, mpd); err != nil {
		return nil, errors.Join(ErrUnmarshalMPD, err)
	}

	return mpd, nil
}

// Bytes marshals the MPD to an XML document with indentations.
func (m *MPD) Bytes() ([]byte, error) {
	xmlData, err := xml.MarshalIndent(m, "", "  ")
	if err != nil {
		return nil, errors.Join(ErrMarshalMPD, err)
	}

	return append([]byte(xml.Header), xmlData...), nil
}

type PatchLocation struct {
	Items []string `xml:",any"`
	TTL   float64  `xml:"ttl,omitempty"`
}

type InitializationSet struct {
	RepresentationBase

	Accessibility []Descriptor `xml:"Accessibility,omitempty"`
	Role          []Descriptor `xml:"Role,omitempty"`
	Rating        []Descriptor `xml:"Rating,omitempty"`
	Viewpoint     []Descriptor `xml:"Viewpoint,omitempty"`
	ID            uint         `xml:"id,attr"`

	// InAllPeriods defaults to `true`.
	InAllPeriods bool `xml:"inAllPeriods,attr,omitempty"`

	ContentType    RFC6838ContentType `xml:"contentType,attr,omitempty"`
	PAR            Ratio              `xml:"par,attr,omitempty"`
	MaxWidth       uint               `xml:"maxWidth,attr,omitempty"`
	MaxHeight      uint               `xml:"maxHeight,attr,omitempty"`
	MaxFrameRate   FrameRate          `xml:"maxFrameRate,attr,omitempty"`
	Initialization string             `xml:"initialization,attr,omitempty"`
}

type ServiceDescription struct {
	Items              []string             `xml:",any"`
	Scope              []Descriptor         `xml:"Scope,omitempty"`
	Latency            []Latency            `xml:"Latency,omitempty"`
	PlaybackRate       []PlaybackRate       `xml:"PlaybackRate,omitempty"`
	OperatingQuality   []OperatingQuality   `xml:"OperatingQuality,omitempty"`
	OperatingBandwidth []OperatingBandwidth `xml:"OperatingBandwidth,omitempty"`
	ID                 uint                 `xml:"id,attr,omitempty"`
}

type Latency struct {
	Items          []string          `xml:",any"`
	QualityLatency []UIntPairsWithID `xml:"QualityLatency,omitempty"`
	ReferenceID    uint              `xml:"referenceID,attr,omitempty"`
	Target         uint              `xml:"target,attr,omitempty"`
	Max            uint              `xml:"max,attr,omitempty"`
	Min            uint              `xml:"min,attr,omitempty"`
}

type PlaybackRate struct {
	Max float64 `xml:"max,attr,omitempty"`
	Min float64 `xml:"min,attr,omitempty"`
}

type OperatingQuality struct {
	// MediaType defaults to AnyMediaType.
	MediaType OperatingQualityMediaType `xml:"mediaType,attr,omitempty"`

	Min           uint   `xml:"min,attr,omitempty"`
	Max           uint   `xml:"max,attr,omitempty"`
	Target        uint   `xml:"target,attr,omitempty"`
	Type          string `xml:"type,attr,omitempty"`
	MaxDifference uint   `xml:"maxDifference,attr,omitempty"`
}

type OperatingBandwidth struct {
	// MediaType defaults to AllMediaType.
	MediaType OperatingBandwidthMediaType `xml:"mediaType,attr,omitempty"`

	Min    uint `xml:"min,attr,omitempty"`
	Max    uint `xml:"max,attr,omitempty"`
	Target uint `xml:"target,attr,omitempty"`
}

type UIntPairsWithID struct {
	UIntVector

	Type string `xml:"type,attr,omitempty"`
}

type UIntVWithID struct {
	UIntVector

	ID          uint               `xml:"id,attr"`
	Profiles    ListOfProfiles     `xml:"profiles,attr,omitempty"`
	ContentType RFC6838ContentType `xml:"contentType,attr,omitempty"`
}

type Metrics struct {
	Items     []string     `xml:",any"`
	Reporting []Descriptor `xml:"Reporting"`
	Range     []Range      `xml:"Range,omitempty"`
	Metrics   string       `xml:"metrics,attr"`
}

type Period struct {
	Items                []string             `xml:",any"`
	BaseURL              []BaseURL            `xml:"BaseURL,omitempty"`
	SegmentBase          *SegmentBase         `xml:"SegmentBase,omitempty"`
	SegmentList          *SegmentList         `xml:"SegmentList,omitempty"`
	SegmentTemplate      *SegmentTemplate     `xml:"SegmentTemplate,omitempty"`
	AssetIdentifier      *Descriptor          `xml:"AssetIdentifier,omitempty"`
	EventStream          []EventStream        `xml:"EventStream,omitempty"`
	ServiceDescription   []ServiceDescription `xml:"ServiceDescription,omitempty"`
	ContentProtection    []ContentProtection  `xml:"ContentProtection,omitempty"`
	AdaptationSet        []AdaptationSet      `xml:"AdaptationSet,omitempty"`
	Subset               []Subset             `xml:"Subset,omitempty"`
	SupplementalProperty []Descriptor         `xml:"SupplementalProperty,omitempty"`
	EmptyAdaptationSet   []AdaptationSet      `xml:"EmptyAdaptationSet,omitempty"`
	GroupLabel           []Label              `xml:"GroupLabel,omitempty"`
	Preselection         []Preselection       `xml:"Preselection,omitempty"`
	ID                   string               `xml:"id,attr,omitempty"`
	Actuate              string               `xml:"actuate,attr,omitempty"`
	Start                string               `xml:"start,attr,omitempty"`
	Duration             string               `xml:"duration,attr,omitempty"`

	// BitstreamSwitching defaults to `false`.
	BitstreamSwitching bool `xml:"bitstreamSwitching,attr,omitempty"`
}

type Range struct {
	StartTime string `xml:"starttime,attr,omitempty"`
	Duration  string `xml:"duration,attr,omitempty"`
}

type RepresentationBase struct {
	Items                     []string                `xml:",any"`
	FramePacking              []Descriptor            `xml:"FramePacking,omitempty"`
	AudioChannelConfiguration []*Descriptor           `xml:"AudioChannelConfiguration,omitempty"`
	ContentProtection         []ContentProtection     `xml:"ContentProtection,omitempty"`
	OutputProtection          *Descriptor             `xml:"OutputProtection,omitempty"`
	EssentialProperty         []Descriptor            `xml:"EssentialProperty,omitempty"`
	SupplementalProperty      []Descriptor            `xml:"SupplementalProperty,omitempty"`
	InbandEventStream         []EventStream           `xml:"InbandEventStream,omitempty"`
	Switching                 []Switching             `xml:"Switching,omitempty"`
	RandomAccess              []RandomAccess          `xml:"RandomAccess,omitempty"`
	GroupLabel                []Label                 `xml:"GroupLabel,omitempty"`
	Label                     []Label                 `xml:"Label,omitempty"`
	ProducerReferenceTime     []ProducerReferenceTime `xml:"ProducerReferenceTime,omitempty"`
	ContentPopularityRate     []ContentPopularityRate `xml:"ContentPopularityRate,omitempty"`
	Resync                    []Resync                `xml:"Resync,omitempty"`
	Profiles                  ListOfProfiles          `xml:"profiles,attr,omitempty"`
	Width                     uint                    `xml:"width,attr,omitempty"`
	Height                    uint                    `xml:"height,attr,omitempty"`
	SAR                       Ratio                   `xml:"sar,attr,omitempty"`
	FrameRate                 FrameRate               `xml:"frameRate,attr,omitempty"`
	AudioSamplingRate         *AudioSamplingRate      `xml:"audioSamplingRate,attr,omitempty"`
	MIMEType                  MIMEType                `xml:"mimeType,attr,omitempty"`
	SegmentProfiles           *ListOf4CC              `xml:"segmentProfiles,attr,omitempty"`
	Codecs                    Codecs                  `xml:"codecs,attr,omitempty"`
	ContainerProfiles         *ListOf4CC              `xml:"containerProfiles,attr,omitempty"`
	MaximumSAPPeriod          float64                 `xml:"maximumSAPPeriod,attr,omitempty"`
	StartWithSAP              uint                    `xml:"startWithSAP,attr,omitempty"`
	MaxPlayoutRate            float64                 `xml:"maxPlayoutRate,attr,omitempty"`
	CodingDependency          bool                    `xml:"codingDependency,attr,omitempty"`
	ScanType                  VideoScan               `xml:"scanType,attr,omitempty"`

	// SelectionPriority defaults to `1`.
	SelectionPriority uint `xml:"selectionPriority,attr,omitempty"`

	Tag Tag `xml:"tag,attr,omitempty"`
}

type ContentProtection struct {
	Descriptor

	MSPro          []string `xml:"urn:microsoft:playready mspr:pro,omitempty"`
	CENCPSSH       []string `xml:"urn:mpeg:cenc:2013 cenc:pssh,omitempty"`
	CENCDefaultKID string   `xml:"urn:mpeg:cenc:2013 cenc:default_KID,attr,omitempty"`
	Robustness     string   `xml:"robustness,attr,omitempty"`
	RefID          string   `xml:"refId,attr,omitempty"`
	Ref            string   `xml:"ref,attr,omitempty"`
}

type Resync struct {
	// Type defaults to `0`.
	Type SAPType `xml:"type,attr,omitempty"`

	DT    float32 `xml:"dT,attr,omitempty"`
	DIMax float32 `xml:"dImax,attr,omitempty"`

	// DIMin defaults to `0`.
	DIMin float32 `xml:"dImin,attr,omitempty"`

	// Marker defaults to `false`.
	Marker bool `xml:"marker,attr,omitempty"`
}

type ContentPopularityRate struct {
	PR                []PR   `xml:"PR"`
	Source            Source `xml:"source,attr"`
	SourceDescription string `xml:"source_description,attr,omitempty"`
}

type PR struct {
	PopularityRate uint   `xml:"popularityRate,attr,omitempty"`
	Start          uint64 `xml:"start,attr,omitempty"`

	// R defaults to `0`.
	R int `xml:"r,attr,omitempty"`
}

type Label struct {
	// ID defaults to `0`.
	ID uint `xml:"id,attr,omitempty"`

	Lang  string   `xml:"lang,attr,omitempty"`
	Items []string `xml:",chardata"`
}

type ProducerReferenceTime struct {
	Items     []string    `xml:",any"`
	UTCTiming *Descriptor `xml:"UTCTiming,omitempty"`
	ID        uint        `xml:"id,attr"`

	// Inband defaults to `false`.
	Inband bool `xml:"inband,attr,omitempty"`

	// Type defaults to EncoderProducerReferenceTimeType.
	Type ProducerReferenceTimeType `xml:"type,attr,omitempty"`

	ApplicationScheme string `xml:"applicationScheme,attr,omitempty"`
	WallClockTime     string `xml:"wallClockTime,attr"`
	PresentationTime  uint64 `xml:"presentationTime,attr"`
}

type Preselection struct {
	RepresentationBase

	Accessibility []Descriptor `xml:"Accessibility,omitempty"`
	Role          []Descriptor `xml:"Role,omitempty"`
	Rating        []Descriptor `xml:"Rating,omitempty"`
	Viewpoint     []Descriptor `xml:"Viewpoint,omitempty"`

	// ID defaults to `1`.
	ID string `xml:"id,attr,omitempty"`

	PreselectionComponents StringVector `xml:"preselectionComponents,attr"`
	Lang                   string       `xml:"lang,attr,omitempty"`

	// Order defaults to UndefinedPreselectionOrder.
	Order PreselectionOrder `xml:"order,attr,omitempty"`
}

type Representation struct {
	RepresentationBase

	BaseURL                []BaseURL           `xml:"BaseURL,omitempty"`
	ExtendedBandwidth      []ExtendedBandwidth `xml:"ExtendedBandwidth,omitempty"`
	SubRepresentation      []SubRepresentation `xml:"SubRepresentation,omitempty"`
	SegmentBase            *SegmentBase        `xml:"SegmentBase,omitempty"`
	SegmentList            *SegmentList        `xml:"SegmentList,omitempty"`
	SegmentTemplate        *SegmentTemplate    `xml:"SegmentTemplate,omitempty"`
	Bandwidth              uint                `xml:"bandwidth,attr"`
	ID                     string              `xml:"id,attr,omitempty"`
	QualityRanking         uint                `xml:"qualityRanking,attr,omitempty"`
	DependencyId           StringVector        `xml:"dependencyId,attr,omitempty"`
	AssociationId          StringVector        `xml:"associationId,attr,omitempty"`
	AssociationType        ListOf4CC           `xml:"associationType,attr,omitempty"`
	MediaStreamStructureId StringVector        `xml:"mediaStreamStructureId,attr,omitempty"`
}

type ExtendedBandwidth struct {
	Items     []string    `xml:",any"`
	ModelPair []ModelPair `xml:"ModelPair,omitempty"`

	// VBR defaults to `false`.
	VBR bool `xml:"vbr,attr,omitempty"`
}

type ModelPair struct {
	Items      []string `xml:",any"`
	BufferTime string   `xml:"bufferTime,attr"`
	Bandwidth  uint     `xml:"bandwidth,attr"`
}

type SubRepresentation struct {
	RepresentationBase

	Level            uint         `xml:"level,attr,omitempty"`
	DependencyLevel  UIntVector   `xml:"dependencyLevel,attr,omitempty"`
	Bandwidth        uint         `xml:"bandwidth,attr,omitempty"`
	ContentComponent StringVector `xml:"contentComponent,attr,omitempty"`
}

type Subset struct {
	Contains UIntVector `xml:"contains,attr"`
	ID       string     `xml:"id,attr,omitempty"`
}

type Switching struct {
	Interval uint `xml:"interval,attr"`

	// Type defaults to MediaSwitchingType.
	Type SwitchingType `xml:"type,attr,omitempty"`
}

type RandomAccess struct {
	Interval uint `xml:"interval,attr"`

	// Type defaults to ClosedRandomAccessType.
	Type RandomAccessType `xml:"type,attr,omitempty"`

	MinBufferTime string `xml:"minBufferTime,attr,omitempty"`
	Bandwidth     uint   `xml:"bandwidth,attr,omitempty"`
}

type S struct {
	T *uint64 `xml:"t,attr,omitempty"`
	N uint64  `xml:"n,attr,omitempty"`
	D uint64  `xml:"d,attr"`

	// R defaults to `0`.
	R int `xml:"r,attr,omitempty"`

	// K defaults to `1`.
	K uint64 `xml:"k,attr,omitempty"`
}

type SegmentBase struct {
	Items                  []string         `xml:",any"`
	Initialization         *URL             `xml:"Initialization,omitempty"`
	RepresentationIndex    *URL             `xml:"RepresentationIndex,omitempty"`
	FailoverContent        *FailoverContent `xml:"FailoverContent,omitempty"`
	Timescale              uint             `xml:"timescale,attr,omitempty"`
	EPTDelta               int              `xml:"eptDelta,attr,omitempty"`
	PDDelta                int              `xml:"pdDelta,attr,omitempty"`
	PresentationTimeOffset uint64           `xml:"presentationTimeOffset,attr,omitempty"`
	PresentationDuration   uint64           `xml:"presentationDuration,attr,omitempty"`

	// IndexRange defaults to `false`.
	IndexRange SingleRFC7233Range `xml:"indexRange,attr,omitempty"`

	IndexRangeExact          bool    `xml:"indexRangeExact,attr,omitempty"`
	AvailabilityTimeOffset   float64 `xml:"availabilityTimeOffset,attr,omitempty"`
	AvailabilityTimeComplete bool    `xml:"availabilityTimeComplete,attr,omitempty"`
}

type MultipleSegmentBase struct {
	SegmentBase

	SegmentTimeline    *SegmentTimeline `xml:"SegmentTimeline,omitempty"`
	BitstreamSwitching *URL             `xml:"BitstreamSwitching,omitempty"`
	Duration           uint             `xml:"duration,attr,omitempty"`
	StartNumber        uint             `xml:"startNumber,attr,omitempty"`
	EndNumber          uint             `xml:"endNumber,attr,omitempty"`
}

type FailoverContent struct {
	FCS []FCS `xml:"FCS"`

	// Valid defaults to `true`.
	Valid bool `xml:"valid,attr,omitempty"`
}

type FCS struct {
	T uint64 `xml:"t,attr"`
	D uint64 `xml:"d,attr,omitempty"`
}

type SegmentList struct {
	MultipleSegmentBase

	Items      []string     `xml:",any"`
	SegmentURL []SegmentURL `xml:"SegmentURL,omitempty"`
}

type SegmentURL struct {
	Items      []string           `xml:",any"`
	Media      string             `xml:"media,attr,omitempty"`
	MediaRange SingleRFC7233Range `xml:"mediaRange,attr,omitempty"`
	Index      string             `xml:"index,attr,omitempty"`
	IndexRange SingleRFC7233Range `xml:"indexRange,attr,omitempty"`
}

type SegmentTemplate struct {
	MultipleSegmentBase

	Media              string `xml:"media,attr,omitempty"`
	Index              string `xml:"index,attr,omitempty"`
	Initialization     string `xml:"initialization,attr,omitempty"`
	BitstreamSwitching string `xml:"bitstreamSwitching,attr,omitempty"`
}

type SegmentTimeline struct {
	Items []string `xml:",any"`
	S     []S      `xml:"S"`
}

type URL struct {
	Items     []string           `xml:",any"`
	SourceURL string             `xml:"sourceURL,attr,omitempty"`
	Range     SingleRFC7233Range `xml:"range,attr,omitempty"`
}

type ProgramInformation struct {
	Items              []string `xml:",any"`
	Title              string   `xml:"Title,omitempty"`
	Source             string   `xml:"Source,omitempty"`
	Copyright          string   `xml:"Copyright,omitempty"`
	Lang               string   `xml:"lang,attr,omitempty"`
	MoreInformationURL string   `xml:"moreInformationURL,attr,omitempty"`
}

type LeapSecondInformation struct {
	Items                           []string `xml:",any"`
	AvailabilityStartLeapOffset     int      `xml:"availabilityStartLeapOffset,attr,omitempty"`
	NextAvailabilityStartLeapOffset int      `xml:"nextAvailabilityStartLeapOffset,attr,omitempty"`
	NextLeapChangeTime              string   `xml:"nextLeapChangeTime,attr,omitempty"`
}
