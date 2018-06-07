package itaas

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type RemoveIPSegmentRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	Uuid        string `position:"Query" name:"Uuid"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *RemoveIPSegmentRequest) Invoke(client goaliyun.Client) (*RemoveIPSegmentResponse, error) {
	resp := &RemoveIPSegmentResponse{}
	err := client.Request("itaas", "RemoveIPSegment", "2017-05-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RemoveIPSegmentResponse struct {
	RequestId goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
	Success   bool
	ErrorList RemoveIPSegmentErrorMessageList
}

type RemoveIPSegmentErrorMessage struct {
	ErrorMessage goaliyun.String
}

type RemoveIPSegmentErrorMessageList []RemoveIPSegmentErrorMessage

func (list *RemoveIPSegmentErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RemoveIPSegmentErrorMessage)
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
