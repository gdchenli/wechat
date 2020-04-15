package externalcontact

import (
	"github.com/gdchenli/wechat.v2/corp/core"
)

type TagListParameters struct {
	TagIds []string `json:"tag_id"` //选填，要查询的标签id，如果不填则获取该企业的所有客户标签，目前暂不支持标签组id
}

type TagGroup struct {
	GroupId    string `json:"group_id"`    //标签组id
	GroupName  string `json:"group_name"`  //标签组名称
	CreateTime int64  `json:"create_time"` //标签组创建时间
	Order      int    `json:"order"`       //标签组排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
	Deleted    bool   `json:"deleted"`     //标签组是否已经被删除，只在指定tag_id进行查询时返回
	Tags       []Tag  `json:"tag"`         //标签组内的标签列表
}

type Tag struct {
	Id          string `json:"id"`           //标签id
	Name        string `json:"name"`         //标签名称
	CreatedTime string `json:"created_time"` //标签创建时间
	Order       int    `json:"order"`        //标签排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
	Deleted     bool   `json:"deleted"`      //标签是否已经被删除，只在指定tag_id进行查询时返回
}

//获取企业标签库
func (clt *Client) TagList(para TagListParameters) (tagGroups []TagGroup, err error) {
	var result struct {
		core.Error
		TagGroups []TagGroup `json:"tag_group"`
	}

	incompleteURL := "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_corp_tag_list?access_token="
	if err = ((*core.Client)(clt)).PostJSON(incompleteURL, para, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	tagGroups = result.TagGroups
	return
}
