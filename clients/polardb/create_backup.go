package polardb

import (
	"github.com/morlay/goaliyun"
)

type CreateBackupRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *CreateBackupRequest) Invoke(client goaliyun.Client) (*CreateBackupResponse, error) {
	resp := &CreateBackupResponse{}
	err := client.Request("polardb", "CreateBackup", "2017-08-01", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateBackupResponse struct {
	RequestId goaliyun.String
}
