package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"questionplatform/global"
	"questionplatform/model"

	"github.com/go-redis/redis/v8"
)

// Collect 根据成语 id 收藏某个成语
func Collect(ctx context.Context, userId int, idiom model.Idiom) (err error) {
	uid := fmt.Sprintf("%s %d", global.CollectKeyPrefix, userId)
	cid := fmt.Sprintf("%s %d", global.CollectFieldPrefix, idiom.ID)

	collection := model.Collection{
		Id:   idiom.ID,
		Word: idiom.Name,
	}
	val, err := json.Marshal(collection)
	if err != nil {
		err = fmt.Errorf("failed to marshall struct to json: %w", err)
		return
	}

	err = global.Rdb.HSet(ctx, uid, cid, val).Err()
	if err != nil {
		err = fmt.Errorf("failed to collect idiom : %w", err)
	}

	return
}

// Collect 根据成语 id 取消收藏收藏某个成语
func Delete(ctx context.Context, userId, idiomId int) (err error) {
	uid := fmt.Sprintf("%s %d", global.CollectKeyPrefix, userId)
	cid := fmt.Sprintf("%s %d", global.CollectFieldPrefix, idiomId)

	err = global.Rdb.HDel(ctx, uid, cid).Err()
	if err != nil {
		err = fmt.Errorf("failed to delete collected idiom : %w", err)
	}
	return
}

// Show 查询所有收藏的成语
func Show(ctx context.Context, userId int) (collections []model.Collection, err error) {
	uid := fmt.Sprintf("%s %d", global.CollectKeyPrefix, userId)

	results, err := global.Rdb.HVals(ctx, uid).Result()
	if err != nil && err != redis.Nil {
		err = fmt.Errorf("failed to get your collections : %w", err)
		return
	}

	for _, result := range results {
		collection := model.Collection{}
		err = json.Unmarshal([]byte(result), &collection)
		if err != nil {
			err = fmt.Errorf("failed to unmarshall json to struct : %w", err)
			return
		}
		collections = append(collections, collection)
	}

	return
}

// CheckIfCollected 检查用户是否收藏了某个成语
func Check(ctx context.Context, userId, idiomId int) (bool, error) {
	uid := fmt.Sprintf("%s %d", global.CollectKeyPrefix, userId)
	cid := fmt.Sprintf("%s %d", global.CollectFieldPrefix, idiomId)

	// 使用HGET命令检查字段是否存在
	_, err := global.Rdb.HGet(ctx, uid, cid).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, fmt.Errorf("failed to check if idiom is collected: %w", err)
	}

	// 如果返回值非空，表示字段存在，用户已经收藏
	return true, nil
}
