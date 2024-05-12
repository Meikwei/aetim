package oceanbase

// import (
// 	"github.com/Meikwei/aetim/pkg/common/db/oceanbase/models"
// 	"github.com/Meikwei/go-tools/db/oceanutil"
// 	"gorm.io/gorm"
// )
// func NewUserOceanbase(db *gorm.DB)(models.UserModelInterface,error){
//   return &UserOcean{coll:db},nil
// }
// type UserOcean struct{
//   coll *gorm.DB
// }
// func (u *UserOcean)Create(users []*models.User) (*gorm.DB,error) {
//   return oceanutil.InsertMany(u.coll, users)
// }

// func (u *UserOcean)UpdateById(userId uint64 ,user *models.User) (*gorm.DB,error) {
//   return oceanutil.UpdateMany(u.coll, user, "user_id = ?", userId)
// }

// func (u *UserOcean)Take(userId uint64 ) (user *models.User,err error){
//   foundUser := new(models.User)
//   _,error:=oceanutil.FindOne[models.User](u.coll,foundUser,"user_id = ?",userId);
//   return foundUser,error
// }

// func (u *UserOcean) GetUserGlobalRecvMsgOpt(userId uint64) (opt int, err error) {
//   foundUser := new(models.User)
//   _,err=oceanutil.FindOne[models.User](u.coll,foundUser,"user_id = ?",userId);
//   return int(foundUser.GlobalRecvMsgOpt),err
// }
