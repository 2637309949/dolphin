// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// ArticleInfo defined 文章信息
type ArticleInfo struct {
	*Article
	// 地址
	URL null.String `json:"url" xml:"url"`
}

func (r *ArticleInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalArticleInfo(data []byte) (ArticleInfo, error) {
	var r ArticleInfo
	err := json.Unmarshal(data, &r)
	return r, err
}
