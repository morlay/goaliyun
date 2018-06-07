package itaas

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type UpdateRoomNameRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
	Screencode  string `position:"Query" name:"Screencode"`
	Drname      string `position:"Query" name:"Drname"`
	RegionId    string `position:"Query" name:"RegionId"`
}

func (req *UpdateRoomNameRequest) Invoke(client goaliyun.Client) (*UpdateRoomNameResponse, error) {
	resp := &UpdateRoomNameResponse{}
	err := client.Request("itaas", "UpdateRoomName", "2017-05-05", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type UpdateRoomNameResponse struct {
	RequestId goaliyun.String
	ErrorCode goaliyun.Integer
	ErrorMsg  goaliyun.String
	Success   bool
	ErrorList UpdateRoomNameErrorMessageList
}

type UpdateRoomNameErrorMessage struct {
	ErrorMessage goaliyun.String
}

type UpdateRoomNameErrorMessageList []UpdateRoomNameErrorMessage

func (list *UpdateRoomNameErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]UpdateRoomNameErrorMessage)
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
