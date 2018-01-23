package fsmalexa

type Permissions struct {
	ConsentToken string `json:"consentToken"`
}

type Application struct {
	ApplicationID string `json:"applicationId"`
}

type User struct {
	UserID      string      `json:"userId"`
	AccessToken string      `json:"accessToken"`
	Permissions Permissions `json:"permissions"`
}

type Session struct {
	New         bool                   `json:"new"`
	SessionID   string                 `json:"sessionId"`
	Application Application            `json:"application"`
	Attributes  map[string]interface{} `json:"attributes"`
	User        User                   `json:"user"`
}

type SupportedInterfaces struct {
	AudioPlayer interface{} `json:"AudioPlayer"`
}

type Device struct {
	DeviceID            string              `json:"deviceId"`
	SupportedInterfaces SupportedInterfaces `json:"SupportedInterfaces"`
}

type System struct {
	Device         Device      `json:"device"`
	Application    Application `json:"application"`
	User           User        `json:"user"`
	APIEndpoint    string      `json:"apiEndpoint"`
	APIAccessToken string      `json:"apiAccessToken"`
}

type AudioPlayer struct {
	PlayerActivity       string `json:"playerActivity"`
	Token                string `json:"token"`
	OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
}

type Context struct {
	System      System      `json:"System"`
	AudioPlayer AudioPlayer `json:"AudioPlayer"`
}

type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type Request struct {
	Type        string `json:"type"`
	RequestID   string `json:"requestId"`
	Timestamp   string `json:"timestamp"`
	Reason      string `json:"reason"`
	Error       Error  `json:"error"`
	DialogState string `json:"dialogState"`
	Locale      string `json:"locale"`
	Intent      Intent `json:"intent"`
}

type MessageReceivedRequestBody struct {
	Version string  `json:"version"`
	Session Session `json:"session"`
	Context Context `json:"context"`
	Request Request `json:"request"`
}
