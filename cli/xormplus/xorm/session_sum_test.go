// Copyright 2017 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xorm

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/2637309949/dolphin/cli/xormplus/builder"
)

func isFloatEq(i, j float64, precision int) bool {
	return fmt.Sprintf("%."+strconv.Itoa(precision)+"f", i) == fmt.Sprintf("%."+strconv.Itoa(precision)+"f", j)
}

func TestSum(t *testing.T) {
	assert.NoError(t, prepareEngine())

	type SumStruct struct {
		Int   int
		Float float32
	}

	var (
		cases = []SumStruct{
			{1, 6.2},
			{2, 5.3},
			{92, -0.2},
		}
	)

	var i int
	var f float32
	for _, v := range cases {
		i += v.Int
		f += v.Float
	}

	assert.NoError(t, testEngine.Sync2(new(SumStruct)))

	cnt, err := testEngine.Insert(cases)
	assert.NoError(t, err)
	assert.EqualValues(t, 3, cnt)

	colInt := testEngine.ColumnMapper.Obj2Table("Int")
	colFloat := testEngine.ColumnMapper.Obj2Table("Float")

	sumInt, err := testEngine.Sum(new(SumStruct), colInt)
	assert.NoError(t, err)
	assert.EqualValues(t, int(sumInt), i)

	sumFloat, err := testEngine.Sum(new(SumStruct), colFloat)
	assert.NoError(t, err)
	assert.Condition(t, func() bool {
		return isFloatEq(sumFloat, float64(f), 2)
	})

	sums, err := testEngine.Sums(new(SumStruct), colInt, colFloat)
	assert.NoError(t, err)
	assert.EqualValues(t, 2, len(sums))
	assert.EqualValues(t, i, int(sums[0]))
	assert.Condition(t, func() bool {
		return isFloatEq(sums[1], float64(f), 2)
	})

	sumsInt, err := testEngine.SumsInt(new(SumStruct), colInt)
	assert.NoError(t, err)
	assert.EqualValues(t, 1, len(sumsInt))
	assert.EqualValues(t, i, int(sumsInt[0]))
}

func TestSumCustomColumn(t *testing.T) {
	assert.NoError(t, prepareEngine())

	type SumStruct struct {
		Int   int
		Float float32
	}

	var (
		cases = []SumStruct{
			{1, 6.2},
			{2, 5.3},
			{92, -0.2},
		}
	)

	assert.NoError(t, testEngine.Sync2(new(SumStruct)))

	cnt, err := testEngine.Insert(cases)
	assert.NoError(t, err)
	assert.EqualValues(t, 3, cnt)

	sumInt, err := testEngine.Sum(new(SumStruct),
		"CASE WHEN `int` <= 2 THEN `int` ELSE 0 END")
	assert.NoError(t, err)
	assert.EqualValues(t, 3, int(sumInt))
}

func TestCount(t *testing.T) {
	assert.NoError(t, prepareEngine())

	type UserinfoCount struct {
		Departname string
	}
	assert.NoError(t, testEngine.Sync2(new(UserinfoCount)))

	colName := testEngine.ColumnMapper.Obj2Table("Departname")
	var cond builder.Cond = builder.Eq{
		"`" + colName + "`": "dev",
	}

	total, err := testEngine.Where(cond).Count(new(UserinfoCount))
	assert.NoError(t, err)
	assert.EqualValues(t, 0, total)

	cnt, err := testEngine.Insert(&UserinfoCount{
		Departname: "dev",
	})
	assert.NoError(t, err)
	assert.EqualValues(t, 1, cnt)

	total, err = testEngine.Where(cond).Count(new(UserinfoCount))
	assert.NoError(t, err)
	assert.EqualValues(t, 1, total)
}

func TestSQLCount(t *testing.T) {
	assert.NoError(t, prepareEngine())

	type UserinfoCount2 struct {
		Id         int64
		Departname string
	}

	type UserinfoBooks struct {
		Id     int64
		Pid    int64
		IsOpen bool
	}

	assertSync(t, new(UserinfoCount2), new(UserinfoBooks))

	total, err := testEngine.SQL("SELECT count(id) FROM userinfo_count2").
		Count()
	assert.NoError(t, err)
	assert.EqualValues(t, 0, total)
}
