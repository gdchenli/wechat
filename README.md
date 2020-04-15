### 示例
```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gdchenli/wechat.v2/corp/addresslist"
	"github.com/gdchenli/wechat.v2/corp/core"
	"github.com/gdchenli/wechat.v2/corp/externalcontact"
)


var customerSecret = "客户联系secret"
var corpSecret = "应用secrest"
var corpid = "corpid"
var AccessTokenServer = core.NewDefaultAccessTokenServer(corpid, customerSecret, nil)

func main() {
	externalClient := externalcontact.NewClient(AccessTokenServer, http.DefaultClient)
	//外部联系人标签列表
	params := externalcontact.TagListParameters{TagIds: []string{"tagid"}}
	tagGroups, err := externalClient.TagList(params)
	if err != nil {
		fmt.Println(err)
	}
	marshal, _ := json.Marshal(tagGroups)
	fmt.Printf("外部联系人标签列表 %v\n", string(marshal))

	//客户群列表
	groupChatListparams := externalcontact.GroupChatListParameters{Limit: 100}
	groupChats, err := externalClient.GroupChatList(groupChatListparams)
	if err != nil {
		fmt.Println(err)
	}
	groupChatsBytes, _ := json.Marshal(groupChats)
	fmt.Printf("客户群列表 %v\n", string(groupChatsBytes))

	//客户群详细
	groupChatDetailparams := externalcontact.GroupChatDetailParameters{ChatId: "ChatId"}
	groupChatDetail, err := externalClient.GroupChatDetail(groupChatDetailparams)
	if err != nil {
		fmt.Println(err)
	}
	groupChatDetailBytes, _ := json.Marshal(groupChatDetail)
	fmt.Printf("客户群详细 %v\n", string(groupChatDetailBytes))

	//获取客户列表
	externalUserids, err := externalClient.UserList("ChenLi")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("获取客户列表%v\n", externalUserids)

	//获取客户详情
	userInfo, err := externalClient.UserDetail("ChatId")
	if err != nil {
		fmt.Println(err)
	}
	userInfoBytes, _ := json.Marshal(userInfo)
	fmt.Printf("获取客户详情 %v\n", string(userInfoBytes))

	//获取离职成员的客户列表
	unassignedListParameters := externalcontact.UnassignedListParameters{PageSize: 100}
	unassignedList, err := externalClient.UnassignedList(unassignedListParameters)
	if err != nil {
		fmt.Println(err)
	}
	unassignedListBytes, _ := json.Marshal(unassignedList)
	fmt.Printf("获取离职成员的客户列表 %v\n", string(unassignedListBytes))

	addressListclient := addresslist.NewClient(AccessTokenServer, http.DefaultClient)

	//部门列表
	departmentList, err := addressListclient.DepartmentList(0)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v\n", departmentList)

	//获取部门的员工
	userList, err := addressListclient.UserSimpleList(1, false, 0)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%+v\n", userList)

	enterpriseUserInfo, err := addressListclient.UserInfo("ChenLi")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v\n", enterpriseUserInfo)
}

```