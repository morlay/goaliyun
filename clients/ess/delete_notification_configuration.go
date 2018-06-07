package ess

import (
	"github.com/morlay/goaliyun"
)

type DeleteNotificationConfigurationRequest struct {
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	NotificationArn      string `position:"Query" name:"NotificationArn"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DeleteNotificationConfigurationRequest) Invoke(client goaliyun.Client) (*DeleteNotificationConfigurationResponse, error) {
	resp := &DeleteNotificationConfigurationResponse{}
	err := client.Request("ess", "DeleteNotificationConfiguration", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DeleteNotificationConfigurationResponse struct {
	RequestId goaliyun.String
}
