package externalcontact

import (
	"github.com/gdchenli/wechat.v2/corp/core"
)

//获取客户群列表
func (clt *Client) UserList(userid string) (externalUserids []string, err error) {
	var result struct {
		core.Error
		ExternalUserids []string `json:"external_userid"`
	}

	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list?userid=" +
		userid + "&access_token="
	if err = ((*core.Client)(clt)).GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	externalUserids = result.ExternalUserids
	return
}

type UserInfo struct {
	ExternalContact ExternalContact `json:"external_contact"`
	FollowUsers     []FollowUser    `json:"follow_user"` //添加了此外部联系人的企业成员
}

type ExternalContact struct {
	ExternalUserid  string          `json:"external_userid"`  //外部联系人的userid
	Name            string          `json:"name"`             //外部联系人的名称
	Position        string          `json:"position"`         //外部联系人头像，第三方不可获取
	Avatar          string          `json:"avatar"`           //外部联系人的类型，1表示该外部联系人是微信用户，2表示该外部联系人是企业微信用户
	Grander         int             `json:"grander"`          //外部联系人性别 0-未知 1-男性 2-女性
	Unionid         string          `json:"unionid"`          //外部联系人在微信开放平台的唯一身份标识（微信unionid），通过此字段企业可将外部联系人与公众号/小程序用户关联起来。仅当联系人类型是微信用户，且企业或第三方服务商绑定了微信开发者ID有此字段。
	CorpName        string          `json:"corp_name"`        //外部联系人的职位，如果外部企业或用户选择隐藏职位，则不返回，仅当联系人类型是企业微信用户时有此字段
	CorpFullName    string          `json:"corp_full_name"`   //外部联系人所在企业的简称，仅当联系人类型是企业微信用户时有此字段
	UserType        int             `json:"type"`             //外部联系人所在企业的主体名称，仅当联系人类型是企业微信用户时有此字段
	ExternalProfile ExternalProfile `json:"external_profile"` //外部联系人的自定义展示信息，可以有多个字段和多种类型，包括文本，网页和小程序，仅当联系人类型是企业微信用户时有此字段，字段详情见对外属性；
}

type ExternalProfile struct {
	ExternalAttrs []ExternalAttr `json:"external_attr"` //外部联系人的自定义展示信息
}

type ExternalAttr struct {
	AttrType    int                     `json:"type"`        //属性类型: 0-文本 1-网页 2-小程序
	Name        string                  `json:"name"`        //属性名称： 需要先确保在管理端有创建该属性，否则会忽略
	Text        ExternalAttrText        `json:"text"`        //文本类型的属性
	Web         ExternalAttrWeb         `json:"web"`         //网页类型的属性，url和title字段要么同时为空表示清除该属性，要么同时不为空
	Miniprogram ExternalAttrMiniprogram `json:"miniprogram"` //小程序类型的属性，appid和title字段要么同时为空表示清除改属性，要么同时不为空
}

type ExternalAttrText struct {
	Value string `json:"value"` //文本属性内容,长度限制12个UTF8字符
}

type ExternalAttrWeb struct {
	Url   string `json:"url"`   //网页的url,必须包含http或者https头
	Title string `json:"title"` //网页的展示标题,长度限制12个UTF8字符
}

type ExternalAttrMiniprogram struct {
	Appid    string `json:"appid"`    //小程序appid，必须是有在本企业安装授权的小程序，否则会被忽略
	Pagepath string `json:"pagepath"` //小程序的展示标题,长度限制12个UTF8字符
	Title    string `json:"title"`    //小程序的展示标题,长度限制12个UTF8字符
}

type FollowUser struct {
	Userid         string          `json:"userid"`           //添加了此外部联系人的企业成员userid
	Remark         string          `json:"remark"`           //该成员对此外部联系人的备注
	Description    string          `json:"description"`      //该成员对此外部联系人的描述
	Createtime     int64           `json:"createtime"`       //该成员添加此外部联系人的时间
	Tags           []FollowUserTag `json:"tags"`             //该成员添加此外部联系人所打标签的分组名称（标签功能需要企业微信升级到2.7.5及以上版本）
	RemarkCorpName string          `json:"remark_corp_name"` //该成员对此客户备注的企业名称
	RemarkMobiles  []string        `json:"remark_mobiles"`   //该成员对此客户备注的手机号码，第三方不可获取
	State          string          `json:"state"`            //该成员添加此客户的渠道，由用户通过创建「联系我」方式指定
}

type FollowUserTag struct {
	GroupName string `json:"group_name"` //该成员添加此外部联系人所打标签的分组名称（标签功能需要企业微信升级到2.7.5及以上版本）
	TagName   string `json:"tag_name"`   //该成员添加此外部联系人所打标签名称
	TagType   int    `json:"type"`       //该成员添加此外部联系人所打标签类型, 1-企业设置, 2-用户自定义
}

//获取客户群详情
func (clt *Client) UserDetail(externalUserid string) (userInfo UserInfo, err error) {
	var result struct {
		core.Error
		ExternalContact ExternalContact `json:"external_contact"`
		FollowUsers     []FollowUser    `json:"follow_user"`
	}

	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get?external_userid=" +
		externalUserid + "&access_token="
	if err = ((*core.Client)(clt)).GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	userInfo.ExternalContact = result.ExternalContact
	userInfo.FollowUsers = result.FollowUsers
	return
}
