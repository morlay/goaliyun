package cms

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type GetContactsRequest struct {
	GroupName string `position:"Query" name:"GroupName"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *GetContactsRequest) Invoke(client goaliyun.Client) (*GetContactsResponse, error) {
	resp := &GetContactsResponse{}
	err := client.Request("cms", "GetContacts", "2018-03-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type GetContactsResponse struct {
	RequestId      goaliyun.String
	Success        bool
	Code           goaliyun.Integer
	Message        goaliyun.String
	HttpStatusCode goaliyun.Integer
	Datapoints     GetContactsDatapoints
}

type GetContactsDatapoints struct {
	Name     goaliyun.String
	Contacts GetContactsContactList
}

type GetContactsContactList []goaliyun.String

func (list *GetContactsContactList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
