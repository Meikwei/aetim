/*
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-05-10 21:36:37
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-05-12 22:31:40
 * @FilePath: \user\internal\user\user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
/*
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-05-10 21:36:37
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-05-11 20:28:34
 * @FilePath: \user\internal\user\user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package user

// import (
// 	"context"

// 	"github.com/Meikwei/aetim/pkg/common/config"
// 	"github.com/Meikwei/aetim/pkg/common/db/cache"
// 	"github.com/Meikwei/aetim/pkg/common/db/oceanbase"
// 	"github.com/Meikwei/aetim/pkg/common/db/oceanbase/controller"
// 	"github.com/Meikwei/aetim/pkg/common/db/oceanbase/models"
// 	"github.com/Meikwei/go-tools/db/oceanutil"
// 	"github.com/Meikwei/go-tools/db/redisutil"
// 	registry "github.com/Meikwei/go-tools/discovery"
// 	"google.golang.org/grpc"
// )
// type userServer struct{
//   db                       controller.UserDatabase
//   config *Config
// }

// type Config struct{
//   OceanbaseConfig config.OceanBase
//   RedisConfig        config.Redis
//   LocalCacheConfig   config.LocalCache
// }
// func Start(ctx context.Context, config *Config,client registry.SvcDiscoveryRegistry, server *grpc.Server) error{
//   oceanbaseCli,err:=oceanutil.NewOceanbase(ctx,config.OceanbaseConfig.Build())
//   if err != nil{
//     return nil
//   }
//   rdb, err := redisutil.NewRedisClient(ctx, config.RedisConfig.Build())
// 	if err != nil {
// 		return err
// 	}
//   users := make([]*models.User, 0)
//   userDB,err:=oceanbase.NewUserOceanbase(oceanbaseCli.GetDB())
//   userCache := cache.NewUserCacheRedis(rdb, &config.LocalCacheConfig, userDB, cache.GetDefaultOpt())
//   return nil;
// }
