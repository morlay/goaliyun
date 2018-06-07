package cloudphoto

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type TagPhotoRequest struct {
	LibraryId   string                  `position:"Query" name:"LibraryId"`
	Confidences *TagPhotoConfidenceList `position:"Query" type:"Repeated" name:"Confidence"`
	StoreName   string                  `position:"Query" name:"StoreName"`
	PhotoId     int64                   `position:"Query" name:"PhotoId"`
	TagKeys     *TagPhotoTagKeyList     `position:"Query" type:"Repeated" name:"TagKey"`
	RegionId    string                  `position:"Query" name:"RegionId"`
}

func (req *TagPhotoRequest) Invoke(client goaliyun.Client) (*TagPhotoResponse, error) {
	resp := &TagPhotoResponse{}
	err := client.Request("cloudphoto", "TagPhoto", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type TagPhotoResponse struct {
	Code      goaliyun.String
	Message   goaliyun.String
	RequestId goaliyun.String
	Action    goaliyun.String
}

type TagPhotoConfidenceList []float64

func (list *TagPhotoConfidenceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]float64)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type TagPhotoTagKeyList []string

func (list *TagPhotoTagKeyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
