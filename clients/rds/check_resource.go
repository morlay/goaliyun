package rds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CheckResourceRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SpecifyCount         string `position:"Query" name:"SpecifyCount"`
	EngineVersion        string `position:"Query" name:"EngineVersion"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	DBInstanceClass      string `position:"Query" name:"DBInstanceClass"`
	Engine               string `position:"Query" name:"Engine"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	DBInstanceUseType    string `position:"Query" name:"DBInstanceUseType"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CheckResourceRequest) Invoke(client goaliyun.Client) (*CheckResourceResponse, error) {
	resp := &CheckResourceResponse{}
	err := client.Request("rds", "CheckResource", "2014-08-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CheckResourceResponse struct {
	RequestId    goaliyun.String
	SpecifyCount goaliyun.String
	Resources    CheckResourceResourceList
}

type CheckResourceResource struct {
	DBInstanceAvailable goaliyun.String
	Engine              goaliyun.String
	EngineVersion       goaliyun.String
}

type CheckResourceResourceList []CheckResourceResource

func (list *CheckResourceResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CheckResourceResource)
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
