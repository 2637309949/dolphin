package svc

import (
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/util/slice"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

// Db defined TODO
type Db interface {
	PageSearch(db *xorm.Engine, ctr, api, table string, params map[string]interface{}) (*model.PageList, error)
	TreeSearch(db *xorm.Engine, controller, api, table string, q map[string]interface{}) (interface{}, error)
	GetOptions(db *xorm.Engine, keys ...string) (map[string]map[string]interface{}, error)
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
	role, _ = slice.RemoveStringDuplicates(append(role, model.AdminRole.Code.String))
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
func (svc *SvcHepler) PageSearch(db *xorm.Engine, ctr, api, table string, params map[string]interface{}) (*model.PageList, error) {
	page, ok := params["page"].(int)
	if !ok {
		return nil, errors.New("not found page")
	}
	size, ok := params["size"].(int)
	if !ok {
		return nil, errors.New("not found size")
	}
	params["offset"] = (page - 1) * size
	rowsSet, err := db.SqlTemplateClient(fmt.Sprintf("%s_%s_select.tpl", ctr, api), &params).Query().List()
	if err != nil {
		return nil, err
	}
	cntSet, err := db.SqlTemplateClient(fmt.Sprintf("%s_%s_count.tpl", ctr, api), &params).Query().List()
	if err != nil {
		return nil, err
	}
	var plt model.PageList
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
	return &plt, nil
}

// TreeSearch defined TODO
func (svc *SvcHepler) TreeSearch(db *xorm.Engine, controller, api, table string, q map[string]interface{}) (interface{}, error) {
	rowsSet, err := db.SqlTemplateClient(fmt.Sprintf("%s_%s.tpl", controller, api), &q).Query().List()
	if err != nil {
		return nil, err
	}

	valueFiled, parentField, nameFiled := "id", "parent", "name"
	treeNodeList, originNodeList := list.New(), list.New()
	root, rootArr := "", []*model.TreeNode{}
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
			node := &model.TreeNode{
				ID:     value,
				Name:   name,
				Parent: parent,
				Tag:    params,
				Nodes:  make([]*model.TreeNode, 0),
			}
			treeNodeList.PushBack(node)
			rootArr = append(rootArr, node)
		} else {
			originNodeList.PushBack(&model.TreeNode{
				ID:     value,
				Name:   name,
				Parent: parent,
				Tag:    params,
				Nodes:  make([]*model.TreeNode, 0),
			})
		}
	}

	for ele := treeNodeList.Front(); ele != nil; ele = ele.Next() {
		treeNode := ele.Value.(*model.TreeNode)
		originEle := originNodeList.Front()
		if originEle == nil {
			break
		}
		for originEle != nil {
			originNextEle := originEle.Next()
			originNode := originEle.Value.(*model.TreeNode)
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
	var optSets []model.SysOptionset
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
