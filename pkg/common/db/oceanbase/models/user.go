/*
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-05-09 20:24:26
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-05-11 22:22:58
 * @FilePath: \aet\server\user\pkg\common\db\oceanbase\relation\user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package models

// import (
// 	"gorm.io/gorm"
// )

// type User struct{
//   UserId uint64 `gorm:"primaryKey;autoIncrement;not null;comment:用户ID"`
//   UserName string `gorm:"size:50;not null;comment:用户名"`
//   UserNumber string `gorm:"size:50;not null;comment:用户编号"`
//   UserPassword string `gorm:"size:100;not null;comment:用户密码"`
//   UserPhone string `gorm:"size:20;unique;not null;comment:用户手机号"`
//   UserAvatar *string `gorm:"size:255;comment:用户头像"`
//   AppMangerLevel int32 `gorm:"comment:应用管理员等级"`
//   GlobalRecvMsgOpt int32 `gorm:"comment:全局接收消息选项"`
//   TimeUser TimeUser `gorm:"embedded"`
// }

// type UserModelInterface interface {
//   Create(users []*User) (*gorm.DB,error)
//   UpdateById(userId uint64 ,user *User)(*gorm.DB,error)
//   Take(userId uint64 ) (user *User,err error)
//   GetUserGlobalRecvMsgOpt(userId uint64) (opt int, err error)
// }
