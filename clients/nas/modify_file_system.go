package nas

import (
	"github.com/morlay/goaliyun"
)

type ModifyFileSystemRequest struct {
	Description  string `position:"Query" name:"Description"`
	FileSystemId string `position:"Query" name:"FileSystemId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *ModifyFileSystemRequest) Invoke(client goaliyun.Client) (*ModifyFileSystemResponse, error) {
	resp := &ModifyFileSystemResponse{}
	err := client.Request("nas", "ModifyFileSystem", "2017-06-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type ModifyFileSystemResponse struct {
	RequestId goaliyun.String
}
