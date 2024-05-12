/*
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-05-11 20:42:09
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-05-11 20:45:25
 * @FilePath: \user\pkg\common\cachekey\user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package cachekey


const (
	UserInfoKey             = "USER_INFO:"
	UserGlobalRecvMsgOptKey = "USER_GLOBAL_RECV_MSG_OPT_KEY:"
)

func GetUserInfoKey(userID string) string {
	return UserInfoKey + userID
}

func GetUserGlobalRecvMsgOptKey(userID string) string {
	return UserGlobalRecvMsgOptKey + userID
}
