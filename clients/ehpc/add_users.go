package ehpc

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type AddUsersRequest struct {
	ReleaseInstance string            `position:"Query" name:"ReleaseInstance"`
	ClusterId       string            `position:"Query" name:"ClusterId"`
	Users           *AddUsersUserList `position:"Query" type:"Repeated" name:"User"`
	RegionId        string            `position:"Query" name:"RegionId"`
}

func (req *AddUsersRequest) Invoke(client goaliyun.Client) (*AddUsersResponse, error) {
	resp := &AddUsersResponse{}
	err := client.Request("ehpc", "AddUsers", "2018-04-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type AddUsersUser struct {
	Name     string `name:"Name"`
	Group    string `name:"Group"`
	Password string `name:"Password"`
}

type AddUsersResponse struct {
	RequestId goaliyun.String
}

type AddUsersUserList []AddUsersUser

func (list *AddUsersUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddUsersUser)
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
