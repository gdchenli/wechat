package externalcontact

import "github.com/gdchenli/wechat.v2/corp/core"

type UnassignedListParameters struct {
	PageId   int `json:"page_id"`   //选填，分页查询，要查询页号，从0开始
	PageSize int `json:"page_size"` //必填，每次返回的最大记录数，默认为1000，最大值为1000
}

type UnassignedList struct {
	UnassignedInfos []UnassignedInfo `json:"info"`
	IsLast          bool             `json:"is_last"`
}

type UnassignedInfo struct {
	HandoverUserId string //离职成员的userid
	ExternalUserId string //外部联系人userid
	DimissionTime  int64  //成员离职时间
}

//获取离职成员的客户列表
func (clt *Client) UnassignedList(para UnassignedListParameters) (unassignedList UnassignedList, err error) {
	var result struct {
		core.Error
		Infos  []UnassignedInfo `json:"info"`
		IsLast bool             `json:"is_last"`
	}

	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_unassigned_list?access_token="
	if err = ((*core.Client)(clt)).PostJSON(incompleteURL, para, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	unassignedList.UnassignedInfos = result.Infos
	unassignedList.IsLast = result.IsLast
	return
}

type UserTransferParameters struct {
	ExternalUserid string `json:"external_userid"` //外部联系人的userid，注意不是企业成员的帐号
	HandoverUserid string `json:"handover_userid"` //离职成员的userid
	TakoverUserid  string `json:"takover_userid"`  //接替成员的userid
}

//离职成员的外部联系人再分配
func (clt *Client) UnassignedUserTransfer(para UserTransferParameters) (flag bool, err error) {
	var result struct {
		core.Error
	}

	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/transfer?access_token="
	if err = ((*core.Client)(clt)).PostJSON(incompleteURL, para, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}

	return true, nil
}

type GroupChatTransferParameters struct {
	ChatIdList []string `json:"chat_id_list"` //必填，需要转群主的客户群ID列表。取值范围： 1 ~ 100
	NewOwner   string   `json:"new_owner"`    //必填，新群主ID
}

type FailedChatList struct {
	ChatId string `json:"chat_id"`
	core.Error
}

//离职成员的群再分配
func (clt *Client) UnassignedGroupChatTransfer(para GroupChatTransferParameters) (failedChatList FailedChatList, err error) {
	var result struct {
		core.Error
		FailedChatList FailedChatList `json:"failed_chat_list"`
	}

	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/transfer?access_token="
	if err = ((*core.Client)(clt)).PostJSON(incompleteURL, para, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	failedChatList = result.FailedChatList
	return failedChatList, nil
}
