package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"questionplatform/dao/mysql"
	"questionplatform/global"
	"questionplatform/model"
	"questionplatform/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetByword(c *gin.Context) {
	word := c.Query("word")
	if word == "" {
		utils.RespDiy(c, http.StatusBadRequest, "word can not be empty")
		return
	}
	// 调用函数查询

	idioms, err := mysql.SearchIdIomByName(word)
	// 出现了不是 sql查询为空 的错误
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		global.Logger.Error("internal error in 'mysql.SearchIdIomByName' ", zap.Error(err))
		utils.RespDiy(c, http.StatusInternalServerError, "search idiom error")
		return
	}

	if len(idioms) == 0 {
		idioms = append(idioms, model.NilIdiom)
	}

	c.JSON(http.StatusOK, idioms)
}

func GetById(c *gin.Context) {
	idiomId := c.Query("idiomId")
	if idiomId == "" {
		utils.RespDiy(c, http.StatusBadRequest, "idiomId can not be empty")
		return
	}
	id, _ := strconv.Atoi(idiomId)

	idiom, err := mysql.SearchIdiomById(id)
	// 出现了不是 sql查询为空 的错误
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		global.Logger.Error("internal error in 'mysql.SearchIdIomById' ", zap.Error(err))
		utils.RespDiy(c, http.StatusInternalServerError, "search idiom error")
		return
	}

	if idiom == (model.Idiom{}) {
		idiom = model.NilIdiom
	}

	c.JSON(http.StatusOK, idiom)

}

func GetRandom(c *gin.Context) {

	idiom, err := mysql.GetRandomOne()
	if err != nil {
		global.Logger.Error("internal error in 'mysql.GetRandomOne' ", zap.Error(err))
		utils.RespDiy(c, http.StatusInternalServerError, "search idiom error")
		return
	}
	c.JSON(http.StatusOK, idiom)
}
