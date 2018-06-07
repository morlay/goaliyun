package emr

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type CheckRenewClusterRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	RegionId        string `position:"Query" name:"RegionId"`
}

func (req *CheckRenewClusterRequest) Invoke(client goaliyun.Client) (*CheckRenewClusterResponse, error) {
	resp := &CheckRenewClusterResponse{}
	err := client.Request("emr", "CheckRenewCluster", "2016-04-08", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CheckRenewClusterResponse struct {
	RequestId      goaliyun.String
	ClusterId      goaliyun.String
	RenewEcsDOList CheckRenewClusterRenewEcsDOList
}

type CheckRenewClusterRenewEcsDO struct {
	EcsId                goaliyun.String
	EcsExpiredTime       goaliyun.String
	EmrExpiredTime       goaliyun.String
	EcsAutoRenew         goaliyun.String
	EmrAutoRenew         goaliyun.String
	EcsAutoRenewDuration goaliyun.Integer
	EmrAutoRenewDuration goaliyun.Integer
	HostGroupId          goaliyun.String
	HostGroupType        goaliyun.String
}

type CheckRenewClusterRenewEcsDOList []CheckRenewClusterRenewEcsDO

func (list *CheckRenewClusterRenewEcsDOList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CheckRenewClusterRenewEcsDO)
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
