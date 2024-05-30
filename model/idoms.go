package model

type Idiom struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Sound          string `json:"sound"`
	Explanation    string `json:"explanation"`
	Provenance     string `json:"provenance"`
	EmotionalColor string `json:"emotional_color"`
	Structure      string `json:"structure"`
	Synonyms       string `json:"synonyms"`
	Antonym        string `json:"antonym"`
	Example        string `json:"example"`
}

var NilIdiom = Idiom{
	ID:             31717,
	Name:           "空空如也",
	Sound:          "成语发音：kōng kōng rú yě",
	Explanation:    "成语解释：形容一无所有，空无所有。",
	Provenance:     "成语出处：清·曹雪芹《红楼梦》第二十五回：‘那僧答道：‘大善哉！大地之上，人物还须济度，岂因妖物败坏空空如也。’",
	EmotionalColor: "感情色彩：中性成语",
	Structure:      "成语结构：复合式成语",
	Synonyms:       "近义词：空虚无物",
	Antonym:        "", // 反义词字段保持为空字符串
	Example:        "成语例句：这篇文章虽然篇幅很长，但实际上内容空空如也，没有一点有用的信息。",
}
