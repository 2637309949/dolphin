package svc

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	nh "net/http"

	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util/http"
	"github.com/2637309949/dolphin/platform/util/maps"
	"github.com/2637309949/dolphin/platform/util/slice"
	"github.com/gin-contrib/cache/persistence"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

// SvcHepler defined TODO
type SvcHepler struct {
	cacheStore persistence.CacheStore
	xlsx       *Xlsx
}

// GetCache defined TODO
func (svc *SvcHepler) GetCache(key string, value interface{}) error {
	return svc.cacheStore.Get(key, value)
}

// SetCache defined TODO
func (svc *SvcHepler) SetCache(key string, value interface{}, expire time.Duration) error {
	return svc.cacheStore.Set(key, value, expire)
}

// AddCache defined TODO
func (svc *SvcHepler) AddCache(key string, value interface{}, expire time.Duration) error {
	return svc.cacheStore.Add(key, value, expire)
}

// ReplaceCache defined TODO
func (svc *SvcHepler) ReplaceCache(key string, data interface{}, expire time.Duration) error {
	return svc.cacheStore.Replace(key, data, expire)
}

// DeleteCache defined TODO
func (svc *SvcHepler) DeleteCache(key string) error {
	return svc.cacheStore.Delete(key)
}

// IncrementCache defined TODO
func (svc *SvcHepler) IncrementCache(key string, data uint64) (uint64, error) {
	return svc.cacheStore.Increment(key, data)
}

// DecrementCache defined TODO
func (svc *SvcHepler) DecrementCache(key string, data uint64) (uint64, error) {
	return svc.cacheStore.Decrement(key, data)
}

// FlushCache defined TODO
func (svc *SvcHepler) FlushCache() error {
	return svc.cacheStore.Flush()
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
	tm := maps.MustMap(params)
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

	tm = maps.MustMap(cntSet[0])
	records := tm.MustInt64("records")
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

// Persist defined TODO
func (svc *SvcHepler) Persist(db *xorm.Session, ids ...string) (int64, error) {
	return new(types.SysAttachment).Persist(db, ids...)
}

// PersistFile defined TODO
func (svc *SvcHepler) PersistFile(db *xorm.Session, cb func([]types.SysAttachment) error, ids ...string) (int64, error) {
	return new(types.SysAttachment).PersistFile(db, cb, ids...)
}

// Remove defined TODO
func (svc *SvcHepler) Remove(db *xorm.Session, ids ...string) (int64, error) {
	return new(types.SysAttachment).Remove(db, ids...)
}

// RemoveFile defined TODO
func (svc *SvcHepler) RemoveFile(db *xorm.Session, cb func([]types.SysAttachment) error, ids ...string) (int64, error) {
	return new(types.SysAttachment).RemoveFile(db, cb, ids...)
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

// Get returns *HttpRequest with GET method.
func (svc *SvcHepler) Get(url string) *http.HttpRequest {
	return http.NewRequest(url, "GET")
}

// Post returns *HttpRequest with POST method.
func (svc *SvcHepler) Post(url string) *http.HttpRequest {
	return http.NewRequest(url, "POST")
}

// Put returns *HttpRequest with PUT method.
func (svc *SvcHepler) Put(url string) *http.HttpRequest {
	return http.NewRequest(url, "PUT")
}

// Delete returns *HttpRequest DELETE method.
func (svc *SvcHepler) Delete(url string) *http.HttpRequest {
	return http.NewRequest(url, "DELETE")
}

// Head returns *HttpRequest with HEAD method.
func (svc *SvcHepler) Head(url string) *http.HttpRequest {
	return http.NewRequest(url, "HEAD")
}

// Check defined TODO
func (svc *SvcHepler) Check(request *nh.Request) bool {
	xlsx := NewXlsx()
	if request.URL.Query().Get("__columns__") != "" {
		cstr := request.URL.Query().Get("__columns__")
		columns := []map[string]interface{}{}
		json.Unmarshal([]byte(cstr), &columns)
		xlsx.Header = columns
	}
	if request.URL.Query().Get("__name__") != "" {
		xlsx.FileName = request.URL.Query().Get("__name__")
	}
	svc.xlsx = xlsx
	return request.URL.Query().Get("__export__") == "true"
}

// SetOptionsetsFormat definedTODO
func (svc *SvcHepler) SetOptionsetsFormat(funk func(interface{}) func(interface{}) interface{}) {
	svc.xlsx.Format = funk
}

// PageExport defined TODO
func (svc *SvcHepler) PageExport(db *xorm.Engine, ctr, api, table string, params map[string]interface{}, formatters ...Formatter) (*types.ExportInfo, error) {
	var err error
	tm := maps.MustMap(params)
	page, size := tm.MustInt("page"), tm.MustInt("size")
	rowsSet := []map[string]interface{}{}
	for len(rowsSet) > 0 || page == 1 {
		params["page"] = page
		params["offset"] = (page - 1) * size
		rowsSet, err = db.SqlTemplateClient(fmt.Sprintf("%s_%s_select.tpl", ctr, api), &params).Query().List()
		if err != nil {
			return nil, err
		}
		for i := range formatters {
			rowsSet, err = formatters[i](db, rowsSet)
			if err != nil {
				return nil, err
			}
		}
		svc.xlsx.DumpRows(rowsSet...)
		page += 1
	}
	return svc.xlsx.ExportInfo(), nil
}

// ParseExcel defined TODO
func (svc *SvcHepler) ParseExcel(r io.Reader, sheet interface{}, header ...[]map[string]interface{}) ([]map[string]string, error) {
	if svc.xlsx == nil {
		svc.xlsx = NewXlsx()
	}
	svc.xlsx.SheetIndex = sheet
	return svc.xlsx.ParseExcel(r)
}

// NewSvcHepler defined TODO
func NewSvcHepler(cacheStore persistence.CacheStore) Svc {
	return &SvcHepler{cacheStore: cacheStore}
}
