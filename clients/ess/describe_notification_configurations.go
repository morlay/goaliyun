package ess

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeNotificationConfigurationsRequest struct {
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeNotificationConfigurationsRequest) Invoke(client goaliyun.Client) (*DescribeNotificationConfigurationsResponse, error) {
	resp := &DescribeNotificationConfigurationsResponse{}
	err := client.Request("ess", "DescribeNotificationConfigurations", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeNotificationConfigurationsResponse struct {
	RequestId                       goaliyun.String
	NotificationConfigurationModels DescribeNotificationConfigurationsNotificationConfigurationModelList
}

type DescribeNotificationConfigurationsNotificationConfigurationModel struct {
	ScalingGroupId    goaliyun.String
	NotificationArn   goaliyun.String
	NotificationTypes DescribeNotificationConfigurationsNotificationTypeList
}

type DescribeNotificationConfigurationsNotificationConfigurationModelList []DescribeNotificationConfigurationsNotificationConfigurationModel

func (list *DescribeNotificationConfigurationsNotificationConfigurationModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeNotificationConfigurationsNotificationConfigurationModel)
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

type DescribeNotificationConfigurationsNotificationTypeList []goaliyun.String

func (list *DescribeNotificationConfigurationsNotificationTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
