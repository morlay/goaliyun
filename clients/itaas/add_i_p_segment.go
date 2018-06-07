package itaas

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type AddIPSegmentRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	Ipsegment   string `position:"Query" name:"Ipsegment"`
	Memo        string `position:"Query" name:"Memo"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *AddIPSegmentRequest) Invoke(client goaliyun.Client) (*AddIPSegmentResponse, error) {
	resp := &AddIPSegmentResponse{}
	err := client.Request("itaas", "AddIPSegment", "2017-05-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddIPSegmentResponse struct {
	RequestId goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
	Success   bool
	ErrorList AddIPSegmentErrorMessageList
}

type AddIPSegmentErrorMessage struct {
	ErrorMessage goaliyun.String
}

type AddIPSegmentErrorMessageList []AddIPSegmentErrorMessage

func (list *AddIPSegmentErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddIPSegmentErrorMessage)
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
