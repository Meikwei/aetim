/*
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-05-11 20:37:16
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-05-11 20:46:09
 * @FilePath: \user\pkg\common\db\cache\config.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package cache

import (
	"strings"
	"sync"

	"github.com/Meikwei/aetim/pkg/common/cachekey"
	"github.com/Meikwei/aetim/pkg/common/config"
)

var (
	once      sync.Once
	subscribe map[string][]string
)

func InitLocalCache(localCache *config.LocalCache) {
	once.Do(func() {
		list := []struct {
			Local config.CacheConfig
			Keys  []string
		}{
			{
				Local: localCache.User,
				Keys:  []string{cachekey.UserInfoKey, cachekey.UserGlobalRecvMsgOptKey},
			},
			// {
			// 	Local: localCache.Group,
			// 	Keys:  []string{cachekey.GroupMemberIDsKey, cachekey.GroupInfoKey, cachekey.GroupMemberInfoKey},
			// },
			// {
			// 	Local: localCache.Friend,
			// 	Keys:  []string{cachekey.FriendIDsKey, cachekey.BlackIDsKey},
			// },
			// {
			// 	Local: localCache.Conversation,
			// 	Keys:  []string{cachekey.ConversationKey, cachekey.ConversationIDsKey, cachekey.ConversationNotReceiveMessageUserIDsKey},
			// },
		}
		subscribe = make(map[string][]string)
		for _, v := range list {
			if v.Local.Enable() {
				subscribe[v.Local.Topic] = v.Keys
			}
		}
	})
}

func getPublishKey(topic string, key []string) []string {
	if topic == "" || len(key) == 0 {
		return nil
	}
	prefix, ok := subscribe[topic]
	if !ok {
		return nil
	}
	res := make([]string, 0, len(key))
	for _, k := range key {
		var exist bool
		for _, p := range prefix {
			if strings.HasPrefix(k, p) {
				exist = true
				break
			}
		}
		if exist {
			res = append(res, k)
		}
	}
	return res
}
