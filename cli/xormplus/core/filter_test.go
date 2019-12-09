package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type quoterOnly struct {
	Dialect
}

func (q *quoterOnly) Quote(item string) string {
	return "[" + item + "]"
}

func TestQuoteFilter_Do(t *testing.T) {
	f := QuoteFilter{}
	sql := "SELECT `COLUMN_NAME` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = ? AND `TABLE_NAME` = ? AND `COLUMN_NAME` = ?"
	res := f.Do(sql, new(quoterOnly), nil)
	assert.EqualValues(t,
		"SELECT [COLUMN_NAME] FROM [INFORMATION_SCHEMA].[COLUMNS] WHERE [TABLE_SCHEMA] = ? AND [TABLE_NAME] = ? AND [COLUMN_NAME] = ?",
		res,
	)
}
