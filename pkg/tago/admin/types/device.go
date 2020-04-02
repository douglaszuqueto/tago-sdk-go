package types

// DeviceListResponse DeviceListResponse
type DeviceListResponse struct {
	Result []Device `json:"result"`
	Response
}

// DeviceGetResponse DeviceListResponse
type DeviceGetResponse struct {
	Result DeviceGet `json:"result"`
	Response
}

// DeviceTokenResponse DeviceTokenResponse
type DeviceTokenResponse struct {
	Result []DeviceToken `json:"result"`
	Response
}

// DeviceTag DeviceTag
type DeviceTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// DeviceBucket DeviceBucket
type DeviceBucket struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Device Device
type Device struct {
	ID   string      `json:"id"`
	Name string      `json:"name"`
	Tags []DeviceTag `json:"tags"`
}

// DeviceGet DeviceGet
type DeviceGet struct {
	ID             string       `json:"id"`
	Name           string       `json:"name"`
	Tags           []DeviceTag  `json:"tags"`
	Bucket         DeviceBucket `json:"bucket"`
	Connector      string       `json:"connector"`
	Active         bool         `json:"active"`
	ConnectorParse bool         `json:"connector_parse"`
	Visible        bool         `json:"visible"`
	Description    interface{}  `json:"description"`
	ParseFunction  interface{}  `json:"parse_function"`
	Profile        string       `json:"profile"`
	CreatedAt      string       `json:"created_at"`
	UpdatedAt      string       `json:"updated_at"`
	InspectedAt    string       `json:"inspected_at"`
	LastInput      string       `json:"last_input"`
	LastOutput     string       `json:"last_output"`
}

// DeviceToken DeviceToken
type DeviceToken struct {
	Name              string      `json:"name"`
	Token             string      `json:"token"`
	Type              string      `json:"type"`
	Permission        string      `json:"permission"`
	SerieNumber       string      `json:"serie_number"`
	VerificationCode  string      `json:"verification_code"`
	LastAuthorization interface{} `json:"last_authorization"`
	CreatedAt         string      `json:"created_at"`
	ExpireTime        string      `json:"expire_time"`
}
