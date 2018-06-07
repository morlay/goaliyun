package cloudphoto

import (
	"github.com/morlay/goaliyun"
)

type CreateTransactionRequest struct {
	Ext       string `position:"Query" name:"Ext"`
	Size      int64  `position:"Query" name:"Size"`
	LibraryId string `position:"Query" name:"LibraryId"`
	StoreName string `position:"Query" name:"StoreName"`
	Force     string `position:"Query" name:"Force"`
	Md5       string `position:"Query" name:"Md.5"`
	RegionId  string `position:"Query" name:"RegionId"`
}

func (req *CreateTransactionRequest) Invoke(client goaliyun.Client) (*CreateTransactionResponse, error) {
	resp := &CreateTransactionResponse{}
	err := client.Request("cloudphoto", "CreateTransaction", "2017-07-11", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type CreateTransactionResponse struct {
	Code        goaliyun.String
	Message     goaliyun.String
	RequestId   goaliyun.String
	Action      goaliyun.String
	Transaction CreateTransactionTransaction
}

type CreateTransactionTransaction struct {
	Upload CreateTransactionUpload
}

type CreateTransactionUpload struct {
	Bucket          goaliyun.String
	FileId          goaliyun.String
	OssEndpoint     goaliyun.String
	ObjectKey       goaliyun.String
	SessionId       goaliyun.String
	AccessKeyId     goaliyun.String
	AccessKeySecret goaliyun.String
	StsToken        goaliyun.String
}
