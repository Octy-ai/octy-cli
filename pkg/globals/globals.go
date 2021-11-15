package globals

// versioning
const (
	ApiVersion = "v1-beta"
	CliVersion = "v0.1.0-pre-alpha"
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
	DeleteTemplatesRoute  = APIRootURL + "v1/retention/messaging/templates/delete"
	GetChurnReportRoute   = APIRootURL + "v1/retention/churn_prediction/report"
)

// thrid party
const (
	SentryDSN = "https://e23880e34ee840209803f0635c93ddbb@sentry.io/2121131"
)
