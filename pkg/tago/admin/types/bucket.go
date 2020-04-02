package types

// BucketListResponse BucketListResponse
type BucketListResponse struct {
	Result []Bucket `json:"result"`
	Response
}

// BucketResponse BucketResponse
type BucketResponse struct {
	Result Bucket `json:"result"`
	Response
}

// BucketList BucketList
type BucketList struct {
	ID      string        `json:"id"`
	Name    string        `json:"name"`
	Profile string        `json:"profile"`
	Tags    []interface{} `json:"tags"`
}

// Bucket Bucket
type Bucket struct {
	ID                  string        `json:"id"`
	Name                string        `json:"name"`
	DataRetention       string        `json:"data_retention"`
	Tags                []interface{} `json:"tags"`
	DataRetentionIgnore []interface{} `json:"data_retention_ignore"`
	Description         interface{}   `json:"description"`
	LastBackup          interface{}   `json:"last_backup"`
	LastRetention       interface{}   `json:"last_retention"`
	Backup              bool          `json:"backup"`
	Visible             bool          `json:"visible"`
	CreatedAt           string        `json:"created_at"`
	UpdatedAt           string        `json:"updated_at"`
}
