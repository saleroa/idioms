package boot

import (
	"io/ioutil"
	"log"
	"questionplatform/global"
	"strings"

	_ "github.com/go-sql-driver/mysql" // 导入您的数据库驱动
)

type IdiomJSON struct {
	Derivation   string `json:"derivation"`
	Example      string `json:"example"`
	Explanation  string `json:"explanation"`
	Pinyin       string `json:"pinyin"`
	Word         string `json:"word"`
	Abbreviation string `json:"abbreviation"`
}


func InitResource (path string) {

	// 读取SQL文件内容
	sqlContent, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// 分割SQL语句
	sqlStatements := strings.Split(string(sqlContent), ";")

	// 逐个执行SQL语句
	for _, statement := range sqlStatements {
		statement = strings.TrimSpace(statement)
		if statement == "" {
			continue
		}

		// 执行SQL语句
		_, err := global.DB.Exec(statement)
		if err != nil {
			global.Logger.Warn("failed to insert idiom to database")
		} 
	}
}
