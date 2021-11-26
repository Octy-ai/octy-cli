package globals

import (
	"errors"
)

// versioning
const (
	ApiVersion = "v1-beta"
	CliVersion = "v0.7.0-pre-alpha"
)

// Octy Docs links

const (
	RootURL          = "https://octy.ai/"
	Docs             = RootURL + "docs"
	SupportTicketURL = RootURL + "support"
	LimitsURL        = Docs + "/api#limits"
)

// Octy api routes
const (
	APIRootURL            = "https://api.octy.ai/"
	AuthRoute             = APIRootURL + "v1/account/authenticate"
	SetAccConfigRoute     = APIRootURL + "v1/configurations/account/set"
	GetAccConfigRoute     = APIRootURL + "v1/configurations/account"
	SetAlgoConfigRoute    = APIRootURL + "v1/configurations/retention/algorithms/set"
	GetAlgoConfigRoute    = APIRootURL + "v1/configurations/retention/algorithms"
	GetEventTypesRoute    = APIRootURL + "v1/retention/events/types"
	CreateEventTypesRoute = APIRootURL + "v1/retention/events/types/create"
	DeleteEventTypesRoute = APIRootURL + "v1/retention/events/types/delete"
	GetSegmentsRoute      = APIRootURL + "v1/retention/segments"
	CreateSegmentRoute    = APIRootURL + "v1/retention/segments/create"
	DeleteSegmentsRoute   = APIRootURL + "v1/retention/segments/delete"
	GetTemplatesRoute     = APIRootURL + "v1/retention/messaging/templates"
	CreateTemplatesRoute  = APIRootURL + "v1/retention/messaging/templates/create"
	UpdateTemplatesRoute  = APIRootURL + "v1/retention/messaging/templates/update"
	DeleteTemplatesRoute  = APIRootURL + "v1/retention/messaging/templates/delete"
	GetChurnReportRoute   = APIRootURL + "v1/retention/churn_prediction/report"
	GetResourceFormats    = APIRootURL + "v1/admin/application/resources/format"
	CreateProfiles        = APIRootURL + "v1/retention/profiles/create"
)

// Errors
var (
	ErrUnknownError = errors.New("unknown error occurred")
)

// thrid party
const (
	SentryDSN = "https://e23880e34ee840209803f0635c93ddbb@sentry.io/2121131"
)

// Accepted boolean representations when parsing csv files
var BoolRepresentations = map[string]bool{
	"True":  true,
	"TRUE":  true,
	"true":  true,
	"t":     true,
	"T":     true,
	"False": true,
	"FALSE": true,
	"false": true,
	"f":     true,
	"F":     true,
}

// Accepted null representations when parsing csv files
var NullRepresentations = map[string]bool{
	"":     true,
	" ":    true,
	"n/a":  true,
	"--":   true,
	"none": true,
	"None": true,
	"NaN":  true,
}

// limits
const (
	UploadLimit = 100000
)
