package polardb

import (
	"github.com/morlay/goaliyun"
)

type CreateDBClusterRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBClusterDescription string `position:"Query" name:"DBClusterDescription"`
	Period               string `position:"Query" name:"Period"`
	DBInstanceStorage    string `position:"Query" name:"DBInstanceStorage"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	UsedTime             string `position:"Query" name:"UsedTime"`
	DBInstanceClass      string `position:"Query" name:"DBInstanceClass"`
	ClusterNetworkType   string `position:"Query" name:"ClusterNetworkType"`
	SecurityIPList       string `position:"Query" name:"SecurityIPList"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	PrivateIpAddress     string `position:"Query" name:"PrivateIpAddress"`
	Engine               string `position:"Query" name:"Engine"`
	VPCId                string `position:"Query" name:"VPCId"`
	DBType               string `position:"Query" name:"DBType"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	DBVersion            string `position:"Query" name:"DBVersion"`
	PayType              string `position:"Query" name:"PayType"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateDBClusterRequest) Invoke(client goaliyun.Client) (*CreateDBClusterResponse, error) {
	resp := &CreateDBClusterResponse{}
	err := client.Request("polardb", "CreateDBCluster", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateDBClusterResponse struct {
	RequestId        goaliyun.String
	DBClusterId      goaliyun.String
	OrderId          goaliyun.String
	ConnectionString goaliyun.String
	Port             goaliyun.String
}
