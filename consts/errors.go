package consts

import (
	"errors"
)

var (
	ErrServiceUnavailable             = errors.New("service unavailable")
	ErrMongoDBUnavailable             = errors.New("MongoDB unavailable")
	ErrNilMongoDBClient               = errors.New("nil MongoDB client")
	ErrNilRequest                     = errors.New("nil request")
	ErrNilRequestData                 = errors.New("nil request data")
	ErrMissingDUID                    = errors.New("missing DUID")
	ErrNilQueryArgs                   = errors.New("nil query arguments")
	ErrNilQueryResult                 = errors.New("nil query result")
	ErrInvalidDistinctResult          = errors.New("invalid distinct result")
	ErrInvalidDistinctFieldName       = errors.New("invalid distinct field name")
	ErrAtLeastOneImageAudioURL        = errors.New("requires at least 1 valid Document ImageURL or AudioURL")
	ErrInvalidDocumentDUID            = errors.New("invalid Document duid")
	ErrInvalidDocumentUUID            = errors.New("invalid Document uuid")
	ErrInvalidDocumentFUID            = errors.New("invalid Document fuid")
	ErrInvalidDocumentLastName        = errors.New("invalid Document LastName")
	ErrInvalidDocumentFirstName       = errors.New("invalid Document FirstName")
	ErrInvalidDocumentCallTypeName    = errors.New("invalid Document CallTypeName")
	ErrInvalidDocumentGroundType      = errors.New("invalid Document GroundType")
	ErrInvalidDocumentCity            = errors.New("invalid Document City")
	ErrInvalidDocumentState           = errors.New("invalid Document State")
	ErrInvalidDocumentProvince        = errors.New("invalid Document Province")
	ErrInvalidDocumentCountry         = errors.New("invalid Document Country")
	ErrInvalidDocumentOcean           = errors.New("invalid Document Ocean")
	ErrInvalidDocumentSensorType      = errors.New("invalid Document SensorType")
	ErrInvalidDocumentSensorName      = errors.New("invalid Document SensorName")
	ErrInvalidDocumentSamplingRate    = errors.New("invalid Document SamplingRate")
	ErrInvalidDocumentLatitude        = errors.New("invalid Document Latitude")
	ErrInvalidDocumentLongitude       = errors.New("invalid Document Longitude")
	ErrInvalidDocumentImageURLs       = errors.New("nil Document ImageURLs")
	ErrInvalidDocumentImageURL        = errors.New("invalid Document ImageURL")
	ErrInvalidDocumentAudioURLs       = errors.New("nil Document AudioURLs")
	ErrInvalidDocumentAudioURL        = errors.New("invalid Document AudioURL")
	ErrInvalidDocumentVideoURLs       = errors.New("nil Document VideoURLs")
	ErrInvalidDocumentVideoURL        = errors.New("invalid Document VideoURL")
	ErrInvalidDocumentFileURLs        = errors.New("nil Document FileURLs")
	ErrInvalidDocumentFileURL         = errors.New("invalid Document FileURL")
	ErrInvalidDocumentRecordTimestamp = errors.New("invalid Document RecordTimestamp")
	ErrInvalidDocumentCreateTimestamp = errors.New("invalid Document CreateTimestamp")
	ErrInvalidUpdateTimestamp         = errors.New("invalid Document UpdateTimestamp")
	ErrNilQueryTransaction            = errors.New("nil QueryTransaction")
	ErrInvalidFileMetadataParameters  = errors.New("invalid FileMetadataParameters")
	ErrUnreachableURI                 = errors.New("unreachable URI")
	ErrNoDocumentFound                = errors.New("no document found")
	ErrMediaType                      = errors.New("invalid media type")
)
