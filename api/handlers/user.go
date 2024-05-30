package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"questionplatform/api/midware"
	"questionplatform/dao/mysql"
	"questionplatform/global"
	"questionplatform/model"
	"questionplatform/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Login(c *gin.Context) {
	var user1 model.User
	user1.Username = c.PostForm("username")
	user1.Password = c.PostForm("password")
	if user1.Password == "" || user1.Username == "" {
		utils.RespDiy(c, http.StatusBadRequest, "username or password cannot be empty")
		return
	}
	user, err := mysql.SearchUserByUsername(user1.Username)
	// 出现了不是 sql查询为空 的错误
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		global.Logger.Error("internal error in 'mysql.SearchUserByUsername' ", zap.Error(err))
		utils.RespDiy(c, http.StatusInternalServerError, "search error")
		return
	}

	// 用户不存在，直接注册
	if user.Username == "" {
		a, _ := utils.GetPwd(user1.Password)
		user1.Password = string(a)
		// 插入用户密码和username
		err = mysql.InsertUser(user1)
		if err != nil {
			global.Logger.Error("internal error in 'mysql.InsertUser' ", zap.Error(err))
			utils.RespDiy(c, http.StatusInternalServerError, "insert error")
		}
		// 查询获取 user 的 id
		user, err := mysql.SearchUserByUsername(user1.Username)
		// 出现了不是 sql查询为空 的错误
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			global.Logger.Error("internal error in 'mysql.SearchUserByUsername' ", zap.Error(err))
			utils.RespDiy(c, http.StatusInternalServerError, "search error")
			return
		}
		//此处jwt产生token
		token, err := midware.GenToken(user.Id)
		if err != nil {
			global.Logger.Error("internal error in 'midware.GenToken' ", zap.Error(err))
			utils.RespDiy(c, http.StatusInternalServerError, "generate token error")
			return
		}
		c.JSON(http.StatusOK, token)
		return
	}

	flag := utils.ComparePwd(user.Password, user1.Password)

	if !flag {
		utils.RespDiy(c, http.StatusBadRequest, "密码错误")
		return
	}
	//此处jwt产生token
	token, err := midware.GenToken(user.Id)
	if err != nil {
		global.Logger.Error("internal error in 'midware.GenToken' ", zap.Error(err))
		utils.RespDiy(c, http.StatusInternalServerError, "generate token error")
		return
	}
	c.JSON(http.StatusOK, token)
}

// func Rgesiter(c *gin.Context) {
// 	var user1 model.User
// 	user1.Username = c.PostForm("username")
// 	password := c.PostForm("password")

// 	if password == "" || user1.Username == "" {
// 		utils.RespDiy(c, http.StatusBadRequest, "username or password cannot be empty")
// 		return
// 	}
// 	a, _ := utils.GetPwd(password)
// 	user1.Password = string(a)

// 	//在数据库中找到该用户

// 	user, err := mysql.SearchUserByUsername(user1.Username)
// 	// 出现了不是 sql查询为空 的错误
// 	if err != nil && !errors.Is(err, sql.ErrNoRows) {
// 		global.Logger.Error("internal error in 'mysql.SearchIdIomByName' ", zap.Error(err))
// 		utils.RespDiy(c, http.StatusInternalServerError, "search error")
// 		return
// 	}
// 	if user.Username != "" {
// 		utils.RespDiy(c, http.StatusBadRequest, "用户已经存在")
// 		return
// 	}
// 	err = mysql.InsertUser(user1)
// 	if err != nil {
// 		global.Logger.Error("internal error in 'mysql.InsertUser' ", zap.Error(err))
// 		utils.RespDiy(c, http.StatusInternalServerError, "insert error")
// 	}
// 	utils.RespDiy(c, http.StatusOK, "register successfully")
// }

// func Login(c *gin.Context) {
// 	var user1 model.User
// 	user1.Username = c.PostForm("username")
// 	user1.Password = c.PostForm("password")
// 	if user1.Password == "" || user1.Username == "" {
// 		utils.RespDiy(c, http.StatusBadRequest, "username or password cannot be empty")
// 		return
// 	}
// 	user, err := mysql.SearchUserByUsername(user1.Username)
// 	// 出现了不是 sql查询为空 的错误
// 	if err != nil && !errors.Is(err, sql.ErrNoRows) {
// 		global.Logger.Error("internal error in 'mysql.SearchUserByUsername' ", zap.Error(err))
// 		utils.RespDiy(c, http.StatusInternalServerError, "search error")
// 		return
// 	}
// 	if user.Username == "" {
// 		utils.RespDiy(c, http.StatusBadRequest, "用户不存在")
// 		return
// 	}

// 	flag := utils.ComparePwd(user.Password, user1.Password)

// 	if !flag {
// 		utils.RespDiy(c, http.StatusBadRequest, "密码错误")
// 		return
// 	}
// 	//此处jwt产生token
// 	token, err := midware.GenToken(user.Id)
// 	if err != nil {
// 		global.Logger.Error("internal error in 'midware.GenToken' ", zap.Error(err))
// 		utils.RespDiy(c, http.StatusInternalServerError, "generate token error")
// 		return
// 	}
// 	c.JSON(http.StatusOK, token)
// }

func ChangePass(c *gin.Context) {

	//先要认证jwt，然后才进行后续的处理
	ID := c.MustGet("ID").(int)
	a := c.PostForm("newpass")
	//jwt读取id
	b, _ := utils.GetPwd(a)
	newpass := string(b)
	err := mysql.Changepass(ID, newpass)
	if err != nil {
		global.Logger.Error("internal error in 'mysql.Changepass' ", zap.Error(err))
		utils.RespDiy(c, http.StatusInternalServerError, "Change error")
		return
	}
	utils.RespDiy(c, http.StatusOK, "change password successfully")
}

func InsertUserInfo(c *gin.Context) {
	//jwt认证
	ID := c.MustGet("ID").(int)
	var user model.User
	user.Nickname = c.PostForm("nickname")
	user.Signature = c.PostForm("signature")
	user.Id = ID
	err := mysql.InsertUserInfo(user)
	if err != nil {
		global.Logger.Error("internal error in 'mysql.InsertUserInfo' ", zap.Error(err))
		utils.RespDiy(c, http.StatusInternalServerError, "insert userinfo error")
		return
	}
	utils.RespDiy(c, http.StatusOK, "insert successfully")
}

func GetUserInfo(c *gin.Context) {
	//jwt 认证
	ID := c.MustGet("ID").(int)
	info, err := mysql.GetUserInfo(ID)
	if err != nil {
		global.Logger.Error("internal error in 'mysql.GetUserInfo' ", zap.Error(err))
		utils.RespDiy(c, http.StatusInternalServerError, "get userinfo error")
		return
	}
	c.JSON(http.StatusOK, info)
}
