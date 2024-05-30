package api

import (
	"questionplatform/api/handlers"
	"questionplatform/api/midware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.Default()
	u := r.Group("/user")
	{
		u.POST("/login", midware.Cors(), handlers.Login)

	}
	i := r.Group("/idiom")
	{
		i.GET("/getByid", midware.Cors(), midware.JWTAuth(), handlers.GetById)
		i.GET("/getByword", midware.Cors(), midware.JWTAuth(), handlers.GetByword)
		i.GET("/getRandom", midware.Cors(), midware.JWTAuth(), handlers.GetRandom)
	}
	c := r.Group("/collect")
	{
		c.POST("/get", midware.Cors(), midware.JWTAuth(), handlers.Collect)
		c.DELETE("/delete", midware.Cors(), midware.JWTAuth(), handlers.DelCollection)
		c.GET("/show", midware.Cors(), midware.JWTAuth(), handlers.GetCollections)
		c.GET("/check", midware.Cors(), midware.JWTAuth(), handlers.CheckIfCollected)
	}
	return r
}

// func InitRouter() *gin.Engine {

// 	r := gin.Default()
// 	u := r.Group("/user")
// 	{
// 		u.POST("/register", midware.Cors(), handlers.Rgesiter)
// 		u.POST("/login", midware.Cors(), handlers.Login)
// 		u.PUT("/changepass", midware.Cors(), midware.JWTAuth(), handlers.ChangePass)
// 		u.POST("/insertinfo", midware.Cors(), midware.JWTAuth(), handlers.InsertUserInfo)
// 		u.GET("/getinfo", midware.Cors(), midware.JWTAuth(), handlers.GetUserInfo)
// 	}
// 	i := r.Group("/idiom")
// 	{
// 		i.GET("/getByid", midware.Cors(), midware.JWTAuth(), handlers.GetById)
// 		i.GET("/getByword", midware.Cors(), midware.JWTAuth(), handlers.GetByword)
// 		i.GET("/getRandom", midware.Cors(), midware.JWTAuth(), handlers.GetRandom)
// 	}
// 	c := r.Group("/collection")
// 	{
// 		c.POST("/collect", midware.Cors(), midware.JWTAuth(), handlers.Collect)
// 		c.DELETE("/delete", midware.Cors(), midware.JWTAuth(), handlers.DelCollection)
// 		c.GET("/get", midware.Cors(), midware.JWTAuth(), handlers.GetCollections)
// 	}

// 	return r
// }
