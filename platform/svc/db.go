package svc

import (
	"container/list"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util/slice"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

type (
	// Db defined TODO
	Db interface {
		PageSearch(*xorm.Engine, string, string, string, map[string]interface{}, ...Formatter) (*types.PageList, error)
		TreeSearch(*xorm.Engine, string, string, string, map[string]interface{}, ...Formatter) (interface{}, error)
		GetOptions(*xorm.Engine, ...string) (map[string]map[string]interface{}, error)
		InRole(*xorm.Engine, string, ...string) bool
		InAdmin(*xorm.Engine, string, ...string) bool
	}
	// Formatter defined TODO
	Formatter func(*xorm.Engine, []map[string]interface{}) ([]map[string]interface{}, error)
	// TypeMap Defined TODO
	TypeMap map[string]interface{}
)

// MustInt64 defined TODO
func (q *TypeMap) MustInt64(key string) int64 {
	i, _ := (*q)[key].(int64)
	return i
}

// MustInt defined TODO
func (q *TypeMap) MustInt(key string) int {
	i, _ := (*q)[key].(int)
	return i
}

// MustFloat32 defined TODO
func (q *TypeMap) MustFloat32(key string) float32 {
	i, _ := (*q)[key].(float32)
	return i
}

// MustFloat32 defined TODO
func (q *TypeMap) MustFloat64(key string) float64 {
	i, _ := (*q)[key].(float64)
	return i
}

// MustBool defined TODO
func (q *TypeMap) MustBool(key string) bool {
	i, _ := (*q)[key].(bool)
	return i
}

// MustString defined TODO
func (q *TypeMap) MustString(key string) string {
	i, _ := (*q)[key].(string)
	return i
}

// InRole defined TODO
func (svc *SvcHepler) InRole(db *xorm.Engine, userId string, role ...string) bool {
	var cnt int
	role, _ = slice.RemoveStringDuplicates(role)
	if _, err := db.SQL(
		fmt.Sprintf(`select 
		count(distinct(role_id)) cnt 
		from sys_role_user 
		left join sys_role on sys_role.id=sys_role_user.role_id 
		where sys_role.code in (%v) and sys_role_user.user_id = ? and sys_role_user.is_delete != 1`,
			strings.Join(funk.Map(role, func(id string) string { return fmt.Sprintf("'%v'", id) }).([]string), ",")),
		userId).Get(&cnt); err != nil {
		logrus.Error(err)
		return false
	}
	return cnt == len(role)
}

// InAdmin defined TODO
func (svc *SvcHepler) InAdmin(db *xorm.Engine, userId string, role ...string) bool {
	var cnt int
	role, _ = slice.RemoveStringDuplicates(append(role, types.AdminRole.Code.String))
	if _, err := db.SQL(
		fmt.Sprintf(`select 
		count(sys_role_user.id) cnt 
		from sys_role_user 
		left join sys_role on sys_role.id=sys_role_user.role_id 
		where sys_role_user.user_id = ? and sys_role.code in (%v) and sys_role_user.is_delete != 1`,
			strings.Join(funk.Map(role, func(id string) string { return fmt.Sprintf("'%v'", id) }).([]string), ",")),
		userId).Get(&cnt); err != nil {
		logrus.Error(err)
		return false
	}
	return cnt > 0
}

// PageSearch defined TODO
func (svc *SvcHepler) PageSearch(db *xorm.Engine, ctr, api, table string, params map[string]interface{}, formatters ...Formatter) (*types.PageList, error) {
	tm := TypeMap(params)
	page := tm.MustInt("page")
	size := tm.MustInt("size")
	params["offset"] = (page - 1) * size
	rowsSet, err := db.SqlTemplateClient(fmt.Sprintf("%s_%s_select.tpl", ctr, api), &params).Query().List()
	if err != nil {
		return nil, err
	}
	cntSet, err := db.SqlTemplateClient(fmt.Sprintf("%s_%s_count.tpl", ctr, api), &params).Query().List()
	if err != nil {
		return nil, err
	}
	var plt types.PageList
	if len(rowsSet) == 0 {
		plt.Data = []map[string]interface{}{}
		return &plt, nil
	}
	records, _ := cntSet[0]["records"].(int64)
	var totalpages int64 = 0
	if records < int64(size) {
		totalpages = 1
	} else if records%int64(size) == 0 {
		totalpages = records / int64(size)
	} else {
		totalpages = records/int64(size) + 1
	}
	plt.Page = page
	plt.Size = size
	plt.Data = rowsSet
	plt.TotalRecords = int(records)
	plt.TotalPages = int(totalpages)
	for i := range formatters {
		data, err := formatters[i](db, plt.Data)
		if err != nil {
			return nil, err
		}
		plt.Data = data
	}
	return &plt, nil
}

// TreeSearch defined TODO
func (svc *SvcHepler) TreeSearch(db *xorm.Engine, controller, api, table string, q map[string]interface{}, formatters ...Formatter) (interface{}, error) {
	rowsSet, err := db.SqlTemplateClient(fmt.Sprintf("%s_%s.tpl", controller, api), &q).Query().List()
	if err != nil {
		return nil, err
	}

	valueFiled, parentField, nameFiled := "id", "parent", "name"
	treeNodeList, originNodeList := list.New(), list.New()
	root, rootArr := "", []*types.TreeNode{}
	paramsArr, parentValue := rowsSet, root

	for _, params := range paramsArr {
		value, parent, name := "", "", ""
		if params[valueFiled] != nil {
			value = fmt.Sprintf("%v", params[valueFiled])
		}
		if params[parentField] != nil {
			parent = fmt.Sprintf("%v", params[parentField])
		}
		if params[nameFiled] != nil {
			name = fmt.Sprintf("%v", params[nameFiled])
		}
		if (parentValue == "" && parent == "") || value == parentValue || parent == parentValue {
			node := &types.TreeNode{
				ID:     value,
				Name:   name,
				Parent: parent,
				Tag:    params,
				Nodes:  make([]*types.TreeNode, 0),
			}
			treeNodeList.PushBack(node)
			rootArr = append(rootArr, node)
		} else {
			originNodeList.PushBack(&types.TreeNode{
				ID:     value,
				Name:   name,
				Parent: parent,
				Tag:    params,
				Nodes:  make([]*types.TreeNode, 0),
			})
		}
	}

	for ele := treeNodeList.Front(); ele != nil; ele = ele.Next() {
		treeNode := ele.Value.(*types.TreeNode)
		originEle := originNodeList.Front()
		if originEle == nil {
			break
		}
		for originEle != nil {
			originNextEle := originEle.Next()
			originNode := originEle.Value.(*types.TreeNode)
			if originNode.Parent == treeNode.ID {
				treeNodeList.InsertAfter(originNode, ele)
				treeNode.Nodes = append(treeNode.Nodes, originNode)
				originNodeList.Remove(originEle)
			}
			originEle = originNextEle
		}
	}
	return rootArr, nil
}

// GetOptions defined TODO
func (svc *SvcHepler) GetOptions(db *xorm.Engine, keys ...string) (map[string]map[string]interface{}, error) {
	var optSets []types.SysOptionset
	if err := db.Where("is_delete = 0").In("code", keys).Find(&optSets); err != nil {
		return nil, err
	}
	optMap := make(map[string]map[string]interface{})
	for i := range optSets {
		var options []struct {
			Value interface{} `json:"value"`
			Text  string      `json:"text"`
		}
		if err := json.Unmarshal([]byte(optSets[i].Value.String), &options); err != nil {
			return nil, err
		}
		if optMap[optSets[i].Code.String] == nil {
			optMap[optSets[i].Code.String] = map[string]interface{}{}
		}
		for j := range options {
			optMap[optSets[i].Code.String][options[j].Text] = options[j].Value
		}
	}
	return optMap, nil
}
