package nas

import (
	"github.com/morlay/goaliyun"
)

type CreateFileSystemRequest struct {
	Description  string `position:"Query" name:"Description"`
	ProtocolType string `position:"Query" name:"ProtocolType"`
	StorageType  string `position:"Query" name:"StorageType"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *CreateFileSystemRequest) Invoke(client goaliyun.Client) (*CreateFileSystemResponse, error) {
	resp := &CreateFileSystemResponse{}
	err := client.Request("nas", "CreateFileSystem", "2017-06-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateFileSystemResponse struct {
	RequestId    goaliyun.String
	FileSystemId goaliyun.String
}
