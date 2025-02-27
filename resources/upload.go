package resources

type UploadInit struct {
	ID               *string `json:"id"`
	MediaKey         *string `json:"media_key"`
	ExpiresAfterSecs *int    `json:"expires_after_secs"`
}

type UploadFinalize struct {
	ID               *string `json:"id"`
	MediaKey         *string `json:"media_key"`
	ExpiresAfterSecs *int    `json:"expires_after_secs,omitempty"`
	Size             *int    `json:"size,omitempty"`
	ProcessingInfo   *struct {
		State          *string `json:"state"` // state transition flow is pending -> in_progress -> [failed|succeeded]
		CheckAfterSecs *int    `json:"check_after_secs"`
	} `json:"processing_info,omitempty"`
}
