package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type RenewClusterForAdminRequest struct {
	ResourceOwnerId int64                               `position:"Query" name:"ResourceOwnerId"`
	RenewEcsDos     *RenewClusterForAdminRenewEcsDoList `position:"Query" type:"Repeated" name:"RenewEcsDo"`
	ClusterId       string                              `position:"Query" name:"ClusterId"`
	UserId          string                              `position:"Query" name:"UserId"`
	RegionId        string                              `position:"Query" name:"RegionId"`
}

func (req *RenewClusterForAdminRequest) Invoke(client goaliyun.Client) (*RenewClusterForAdminResponse, error) {
	resp := &RenewClusterForAdminResponse{}
	err := client.Request("emr", "RenewClusterForAdmin", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type RenewClusterForAdminRenewEcsDo struct {
	EcsId     string `name:"EcsId"`
	EcsPeriod string `name:"EcsPeriod"`
	EmrPeriod string `name:"EmrPeriod"`
}

type RenewClusterForAdminResponse struct {
	RequestId      goaliyun.String
	EcsOrderIdList goaliyun.String
	EmrOrderIdList goaliyun.String
}

type RenewClusterForAdminRenewEcsDoList []RenewClusterForAdminRenewEcsDo

func (list *RenewClusterForAdminRenewEcsDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RenewClusterForAdminRenewEcsDo)
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
