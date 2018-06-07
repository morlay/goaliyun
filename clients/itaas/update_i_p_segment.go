package itaas

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type UpdateIPSegmentRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	Uuid        string `position:"Query" name:"Uuid"`
	Ipsegment   string `position:"Query" name:"Ipsegment"`
	Memo        string `position:"Query" name:"Memo"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *UpdateIPSegmentRequest) Invoke(client goaliyun.Client) (*UpdateIPSegmentResponse, error) {
	resp := &UpdateIPSegmentResponse{}
	err := client.Request("itaas", "UpdateIPSegment", "2017-05-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateIPSegmentResponse struct {
	RequestId goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
	Success   bool
	ErrorList UpdateIPSegmentErrorMessageList
}

type UpdateIPSegmentErrorMessage struct {
	ErrorMessage goaliyun.String
}

type UpdateIPSegmentErrorMessageList []UpdateIPSegmentErrorMessage

func (list *UpdateIPSegmentErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]UpdateIPSegmentErrorMessage)
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
