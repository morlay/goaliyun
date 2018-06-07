package nas

import (
	"github.com/morlay/goaliyun"
)

type DeleteFileSystemRequest struct {
	FileSystemId string `position:"Query" name:"FileSystemId"`
	RegionId     string `position:"Query" name:"RegionId"`
}

func (req *DeleteFileSystemRequest) Invoke(client goaliyun.Client) (*DeleteFileSystemResponse, error) {
	resp := &DeleteFileSystemResponse{}
	err := client.Request("nas", "DeleteFileSystem", "2017-06-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteFileSystemResponse struct {
	RequestId goaliyun.String
}
