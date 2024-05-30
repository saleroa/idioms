package mysql

import (
	"fmt"
	"questionplatform/global"
	"questionplatform/model"
)

// SearchIdiomById 根据成语的ID搜索成语
func SearchIdiomById(id int) (idiom model.Idiom, err error) {
	query := "SELECT * FROM idioms WHERE id = ?"
	row := global.DB.QueryRow(query, id)
	err = row.Scan(
		&idiom.ID,
		&idiom.Name,
		&idiom.Sound,
		&idiom.Explanation,
		&idiom.Provenance,
		&idiom.EmotionalColor,
		&idiom.Structure,
		&idiom.Synonyms,
		&idiom.Antonym,
		&idiom.Example,
	)
	if err != nil {
		err = fmt.Errorf("failed to SearchIdiomById: %w", err)
		return
	}
	return
}

// 根据 成语的文字 查询成语 , 使用全文索引
func SearchIdIomByName(word string) (idioms []model.Idiom, err error) {

	query := "SELECT * FROM idioms WHERE name LIKE ? LIMIT 8"

	rows, _ := global.DB.Query(query, "%"+word+"%")

	// 遍历查询结果
	for rows.Next() {
		var idiom model.Idiom
		err = rows.Scan(
			&idiom.ID,
			&idiom.Name,
			&idiom.Sound,
			&idiom.Explanation,
			&idiom.Provenance,
			&idiom.EmotionalColor,
			&idiom.Structure,
			&idiom.Synonyms,
			&idiom.Antonym,
			&idiom.Example,
		)
		if err != nil {
			err = fmt.Errorf("failed to SearchIdomByName: %w", err)
			return
		}
		idioms = append(idioms, idiom)

	}
	return
}

// GetRandomOne 随机获取一个成语
func GetRandomOne() (idiom model.Idiom, err error) {
	query := "SELECT * FROM idioms ORDER BY RAND() LIMIT 1"
	row := global.DB.QueryRow(query)
	err = row.Scan(
		&idiom.ID,
		&idiom.Name,
		&idiom.Sound,
		&idiom.Explanation,
		&idiom.Provenance,
		&idiom.EmotionalColor,
		&idiom.Structure,
		&idiom.Synonyms,
		&idiom.Antonym,
		&idiom.Example,
	)
	if err != nil {
		err = fmt.Errorf("failed to SearchIdiomById: %w", err)
		return
	}
	return
}
