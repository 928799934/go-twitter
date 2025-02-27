package types

import "github.com/928799934/go-twitter/resources"

// {"data":{"id":"1895027516701843457","expires_after_secs":86399}}
type InitOutput struct {
	Data   resources.UploadInit     `json:"data"`
	Errors []resources.PartialError `json:"errors,omitempty"`
}

func (r *InitOutput) HasPartialError() bool {
	return len(r.Errors) != 0
}

/*
	{
	    "data": {
	        "id": "1880028106020515840",
	        "media_key": "13_1880028106020515840",
	        "size": 1024,
	        "expires_after_secs": 86400,
	        "processing_info": {
	            "state": "pending",
	            "check_after_secs": 1
	        }
	    }
	}
*/
type FinalizeOutput struct {
	Data   resources.UploadFinalize `json:"data"`
	Errors []resources.PartialError `json:"errors,omitempty"`
}

func (r *FinalizeOutput) HasPartialError() bool {
	return len(r.Errors) != 0
}

/*
	{
	    "data":{
	        "id":"1880028106020515840",
	        "media_key":"13_1880028106020515840",
	        "processing_info":{
	            "state":"uploading" // state transition flow is pending -> in_progress -> [failed|succeeded]
	        }
	    }
	}
*/
type StatusOutput struct {
	Data   resources.UploadFinalize `json:"data"`
	Errors []resources.PartialError `json:"errors,omitempty"`
}

func (r *StatusOutput) HasPartialError() bool {
	return len(r.Errors) != 0
}
