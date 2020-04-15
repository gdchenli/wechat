package externalcontact

import "github.com/gdchenli/wechat.v2/corp/core"

type GroupChatListParameters struct {
	StatusFilter int         `json:"status_filter"` //选填
	OwnerFilter  OwnerFilter `json:"owner_filter"`  //选填
	Offset       int         `json:"offset"`        //分页，偏移量
	Limit        int         `json:"limit"`         //必填，分页，预期请求的数据量，取值范围 1 ~ 1000
}

type OwnerFilter struct {
	UseridList  []string `json:"userid_list"`  //用户ID列表。最多100个
	PartyidList []int    `json:"partyid_list"` //部门ID列表。最多100个
}

type GroupChatList struct {
	ChatId string `json:"chat_id"` //客户群ID
	Status int    `json:"status"`  //客户群状态。 0 - 正常 1 - 跟进人离职 2 - 离职继承中 3 - 离职继承完成
}

//获取客户群列表
func (clt *Client) GroupChatList(para GroupChatListParameters) (groupChats []GroupChatList, err error) {
	var result struct {
		core.Error
		GroupChats []GroupChatList `json:"group_chat_list"`
	}

	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/list?access_token="
	if err = ((*core.Client)(clt)).PostJSON(incompleteURL, para, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	groupChats = result.GroupChats
	return
}

type GroupChatDetailParameters struct {
	ChatId string `json:"chat_id"` //必填，客户群ID
}

type GroupChat struct {
	ChatId     string       `json:"chat_id"`     //客户群ID
	Name       string       `json:"name"`        //群名
	Owner      string       `json:"owner"`       //群主ID
	CreateTime int64        `json:"create_time"` //群的创建时间
	Notice     string       `json:"notice"`      //群公告
	Members    []MemberList `json:"member_list"` //群成员列表
}

type MemberList struct {
	Userid    string `json:"userid"`     //群成员id
	UserType  int    `json:"type"`       //成员类型。 1 - 企业成员 2 - 外部联系人
	JoinTime  int64  `json:"join_time"`  //入群时间
	JoinScene int    `json:"join_scene"` //入群方式。 1 - 由成员邀请入群（直接邀请入群） 2 - 由成员邀请入群（通过邀请链接入群） 3 - 通过扫描群二维码入群
}

//获取客户群详情
func (clt *Client) GroupChatDetail(para GroupChatDetailParameters) (groupChat GroupChat, err error) {
	var result struct {
		core.Error
		GroupChat GroupChat `json:"group_chat"`
	}

	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/get?access_token="
	if err = ((*core.Client)(clt)).PostJSON(incompleteURL, para, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	groupChat = result.GroupChat
	return
}
