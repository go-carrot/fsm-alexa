package fsmalexa

type Permissions struct {
	ConsentToken string `json:"consentToken,omitempty"`
}

type Application struct {
	ApplicationID string `json:"applicationId,omitempty"`
}

type User struct {
	UserID      string      `json:"userId,omitempty"`
	AccessToken string      `json:"accessToken,omitempty"`
	Permissions Permissions `json:"permissions,omitempty"`
}

type Session struct {
	New         bool                   `json:"new,omitempty"`
	SessionID   string                 `json:"sessionId,omitempty"`
	Application Application            `json:"application,omitempty"`
	Attributes  map[string]interface{} `json:"attributes,omitempty"`
	User        User                   `json:"user,omitempty"`
}

type SupportedInterfaces struct {
	AudioPlayer interface{} `json:"AudioPlayer,omitempty"`
}

type Device struct {
	DeviceID            string              `json:"deviceId,omitempty"`
	SupportedInterfaces SupportedInterfaces `json:"SupportedInterfaces,omitempty"`
}

type System struct {
	Device         Device      `json:"device,omitempty"`
	Application    Application `json:"application,omitempty"`
	User           User        `json:"user,omitempty"`
	APIEndpoint    string      `json:"apiEndpoint,omitempty"`
	APIAccessToken string      `json:"apiAccessToken,omitempty"`
}

type AudioPlayer struct {
	PlayerActivity       string `json:"playerActivity,omitempty"`
	Token                string `json:"token,omitempty"`
	OffsetInMilliseconds int    `json:"offsetInMilliseconds,omitempty"`
}

type Context struct {
	System      System      `json:"System,omitempty"`
	AudioPlayer AudioPlayer `json:"AudioPlayer,omitempty"`
}

type Error struct {
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}

type Request struct {
	Type        string `json:"type,omitempty"`
	RequestID   string `json:"requestId,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	Reason      string `json:"reason,omitempty"`
	Error       Error  `json:"error,omitempty"`
	DialogState string `json:"dialogState,omitempty"`
	Locale      string `json:"locale,omitempty"`
	Intent      Intent `json:"intent,omitempty"`
}

type MessageReceivedRequestBody struct {
	Version string  `json:"version,omitempty"`
	Session Session `json:"session,omitempty"`
	Context Context `json:"context,omitempty"`
	Request Request `json:"request,omitempty"`
}
