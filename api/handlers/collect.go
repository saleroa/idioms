package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"questionplatform/dao/mysql"
	"questionplatform/dao/redis"
	"questionplatform/global"
	"questionplatform/utils"

	"strconv"

	RDB "github.com/go-redis/redis/v8"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Collect(c *gin.Context) {
	uid := c.MustGet("ID").(int)
	idiomId := c.Query("idiomId")
	if idiomId == "" {
		utils.RespDiy(c, http.StatusBadRequest, "idiomId cannot be empty")
		return
	}
	iid, _ := strconv.Atoi(idiomId)

	idiom, err := mysql.SearchIdiomById(iid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespDiy(c, http.StatusInternalServerError, "no such idiom")
			return
		}
		global.Logger.Error("internal error in 'mysql.SearchIdiomById' ", zap.Error(err))
		utils.RespDiy(c, http.StatusInternalServerError, "search idiom error")
		return
	}
	err = redis.Collect(c, uid, idiom)
	if err != nil {
		global.Logger.Error("internal error in 'redis.Collect' ", zap.Error(err))
		utils.RespDiy(c, http.StatusInternalServerError, "collect idiom error")
		return
	}
	utils.RespDiy(c, http.StatusOK, "collect successfully")
}

func DelCollection(c *gin.Context) {
	uid := c.MustGet("ID").(int)
	idiomId := c.Query("idiomId")

	if idiomId == "" {
		utils.RespDiy(c, http.StatusBadRequest, "idiomId cannot be empty")
		return
	}

	iid, _ := strconv.Atoi(idiomId)

	err := redis.Delete(c, uid, iid)
	if err != nil {
		global.Logger.Error("internal error in 'redis.Delete' ", zap.Error(err))
		utils.RespDiy(c, http.StatusInternalServerError, "del collection error")
		return
	}
	utils.RespDiy(c, http.StatusOK, "del collection successfully")
}

func CheckIfCollected(c *gin.Context) {
	uid := c.MustGet("ID").(int)
	idiomId := c.Query("idiomId")

	if idiomId == "" {
		utils.RespDiy(c, http.StatusBadRequest, "idiomId cannot be empty")
		return
	}
	iid, _ := strconv.Atoi(idiomId)

	flag, err := redis.Check(c, uid, iid)
	if err != nil {
		global.Logger.Error("internal error in 'redis.check' ", zap.Error(err))
		utils.RespDiy(c, http.StatusInternalServerError, "check if collected error")
		return
	}
	c.JSON(http.StatusOK, flag)
}

func GetCollections(c *gin.Context) {
	uid := c.MustGet("ID").(int)

	idioms, err := redis.Show(c, uid)
	if err != nil && !errors.Is(err, RDB.Nil) {
		global.Logger.Error("internal error in 'redis.show' ", zap.Error(err))
		utils.RespDiy(c, http.StatusInternalServerError, "show collections error")
		return
	}
	c.JSON(http.StatusOK, idioms)
}
