// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://gopkg.in/chanxuehong/wechat.v1 for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/v1/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package externalcontact

import (
	"net/http"

	"github.com/gdchenli/wechat.v2/corp/core"
)

type Client core.Client

func NewClient(srv core.AccessTokenServer, clt *http.Client) *Client {
	return (*Client)(core.NewClient(srv, clt))
}
