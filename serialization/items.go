package serialization

type Items struct {
	ID          int    `json:"id"`
	Name        string `json:"name" gorm:"type:varchar(255);comment:'名称'"`
	Icon        string `json:"icon" gorm:"type:varchar(1024);comment:'图标'"`
	Price       int    `json:"price" gorm:"type:int(10);comment:'合成价'"`
	Description string `json:"description" gorm:"type:text;comment:'描述'"`
	PlainText   string `json:"plaintext" gorm:"type:varchar(255);comment:'简述'"`
	Sell        int    `json:"sell" gorm:"type:int(255);comment:'卖出价格'"`
	Total       int    `json:"total" gorm:"type:int(255);comment:'总价'"`
	Tag         string `json:"tag" gorm:"type:varchar(255);comment:'标签'"`
	Keywords    string `json:"keywords" gorm:"type:varchar(255);comment:'关键词'"`
}

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type ItemSons []struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	Children   []Item `json:"children"`
	ChildCount int    `json:"child_count"`
}
