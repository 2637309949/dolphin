package svc

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util/http/client"
	"github.com/2637309949/dolphin/platform/util/maps"
	"github.com/2637309949/dolphin/platform/util/slice"
	"github.com/gin-contrib/cache/persistence"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

type (
	// Store defined TODO
	Store interface {
		Get(key string, value interface{}) error
		Set(key string, value interface{}, expire time.Duration) error
		Add(key string, value interface{}, expire time.Duration) error
		Replace(key string, data interface{}, expire time.Duration) error
		Delete(key string) error
		Increment(key string, data uint64) (uint64, error)
		Decrement(key string, data uint64) (uint64, error)
		Flush() error
	}

	// Report defined TODO
	Report interface {
		Check(*http.Request) bool
		PageExport(*xorm.Engine, string, string, string, map[string]interface{}, ...Formatter) (*types.ExportInfo, error)
		ParseExcel(r io.Reader, sheet interface{}, header ...[]map[string]interface{}) ([]map[string]string, error)
		SetOptionsetsFormat(func(interface{}) func(interface{}) interface{})
	}

	// XReport defined TODO
	XReport struct {
		xlsx *Xlsx
	}

	// DB defined TODO
	DB interface {
		PageSearch(*xorm.Engine, string, string, string, map[string]interface{}, ...Formatter) (*types.PageList, error)
		TreeSearch(*xorm.Engine, string, string, string, map[string]interface{}, ...Formatter) (interface{}, error)
		GetOptions(*xorm.Engine, ...string) (map[string]map[string]interface{}, error)
		InRole(*xorm.Engine, string, ...string) bool
		InAdmin(*xorm.Engine, string, ...string) bool
		Persist(db *xorm.Session, ids ...string) (int64, error)
		PersistFile(db *xorm.Session, cb func([]types.SysAttachment) error, ids ...string) (int64, error)
		Remove(db *xorm.Session, ids ...string) (int64, error)
		RemoveFile(db *xorm.Session, cb func([]types.SysAttachment) error, ids ...string) (int64, error)
	}

	// XDB defined TODO
	XDB struct{}

	// Formatter defined TODO
	Formatter func(*xorm.Engine, []map[string]interface{}) ([]map[string]interface{}, error)

	// XClient defined TODO
	Client interface {
		// Get returns *HttpRequest with GET method.
		Get(url string) *client.HttpRequest
		// Post returns *HttpRequest with POST method.
		Post(url string) *client.HttpRequest
		// Put returns *HttpRequest with PUT method.
		Put(url string) *client.HttpRequest
		// Delete returns *HttpRequest DELETE method.
		Delete(url string) *client.HttpRequest
		// Head returns *HttpRequest with HEAD method.
		Head(url string) *client.HttpRequest
	}

	// XClient defined TODO
	XClient struct{}

	// ServiceContext defined TODO
	ServiceContext struct {
		Report Report
		Store  Store
		Client Client
		DB     DB
	}
)

// XReport defined TODO
// ======================================================================
// ======================================================================
// ======================================================================
// Check defined TODO
func (x *XReport) Check(request *http.Request) bool {
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
	x.xlsx = xlsx
	return request.URL.Query().Get("__export__") == "true"
}

// SetOptionsetsFormat definedTODO
func (svc *XReport) SetOptionsetsFormat(funk func(interface{}) func(interface{}) interface{}) {
	svc.xlsx.Format = funk
}

// PageExport defined TODO
func (x *XReport) PageExport(db *xorm.Engine, ctr, api, table string, params map[string]interface{}, formatters ...Formatter) (*types.ExportInfo, error) {
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
		x.xlsx.DumpRows(rowsSet...)
		page += 1
	}
	return x.xlsx.ExportInfo(), nil
}

// ParseExcel defined TODO
func (x *XReport) ParseExcel(r io.Reader, sheet interface{}, header ...[]map[string]interface{}) ([]map[string]string, error) {
	if x.xlsx == nil {
		x.xlsx = NewXlsx()
	}
	x.xlsx.SheetIndex = sheet
	return x.xlsx.ParseExcel(r)
}

// NewXReport defined TODO
func NewXReport() *XReport {
	return &XReport{}
}

// XClient defined TODO
// ======================================================================
// ======================================================================
// ======================================================================
// Get returns *HttpRequest with GET method.
func (x *XClient) Get(url string) *client.HttpRequest {
	return client.NewRequest(url, "GET")
}

// Post returns *HttpRequest with POST method.
func (x *XClient) Post(url string) *client.HttpRequest {
	return client.NewRequest(url, "POST")
}

// Put returns *HttpRequest with PUT method.
func (x *XClient) Put(url string) *client.HttpRequest {
	return client.NewRequest(url, "PUT")
}

// Delete returns *HttpRequest DELETE method.
func (x *XClient) Delete(url string) *client.HttpRequest {
	return client.NewRequest(url, "DELETE")
}

// Head returns *HttpRequest with HEAD method.
func (x *XClient) Head(url string) *client.HttpRequest {
	return client.NewRequest(url, "HEAD")
}

// NewXClient defined TODO
func NewXClient() *XClient {
	return &XClient{}
}

// XDB defined TODO
// ======================================================================
// ======================================================================
// ======================================================================
// InRole defined TODO
func (x *XDB) InRole(db *xorm.Engine, userId string, role ...string) bool {
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
func (x *XDB) InAdmin(db *xorm.Engine, userId string, role ...string) bool {
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
func (x *XDB) PageSearch(db *xorm.Engine, ctr, api, table string, params map[string]interface{}, formatters ...Formatter) (*types.PageList, error) {
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
func (x *XDB) TreeSearch(db *xorm.Engine, controller, api, table string, q map[string]interface{}, formatters ...Formatter) (interface{}, error) {
	rowsSet, err := db.SqlTemplateClient(fmt.Sprintf("%s_%s.tpl", controller, api), &q).Query().List()
	if err != nil {
		return nil, err
	}

	valueFiled, parentField, nameFiled := "id", "parent", "name"
	treeNodeList, originNodeList := list.New(), list.New()
	root, rootArr := "", []*types.TreeNode{}
	paramsArr, parentValue := rowsSet, root

	for i := range paramsArr {
		value, parent, name := "", "", ""
		if paramsArr[i][valueFiled] != nil {
			value = fmt.Sprintf("%v", paramsArr[i][valueFiled])
		}
		if paramsArr[i][parentField] != nil {
			parent = fmt.Sprintf("%v", paramsArr[i][parentField])
		}
		if paramsArr[i][nameFiled] != nil {
			name = fmt.Sprintf("%v", paramsArr[i][nameFiled])
		}
		if (parentValue == "" && parent == "") || value == parentValue || parent == parentValue {
			node := &types.TreeNode{
				ID:     value,
				Name:   name,
				Parent: parent,
				Tag:    paramsArr[i],
				Nodes:  make([]*types.TreeNode, 0),
			}
			treeNodeList.PushBack(node)
			rootArr = append(rootArr, node)
		} else {
			originNodeList.PushBack(&types.TreeNode{
				ID:     value,
				Name:   name,
				Parent: parent,
				Tag:    paramsArr[i],
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
func (x *XDB) Persist(db *xorm.Session, ids ...string) (int64, error) {
	return new(types.SysAttachment).Persist(db, ids...)
}

// PersistFile defined TODO
func (x *XDB) PersistFile(db *xorm.Session, cb func([]types.SysAttachment) error, ids ...string) (int64, error) {
	return new(types.SysAttachment).PersistFile(db, cb, ids...)
}

// Remove defined TODO
func (x *XDB) Remove(db *xorm.Session, ids ...string) (int64, error) {
	return new(types.SysAttachment).Remove(db, ids...)
}

// RemoveFile defined TODO
func (x *XDB) RemoveFile(db *xorm.Session, cb func([]types.SysAttachment) error, ids ...string) (int64, error) {
	return new(types.SysAttachment).RemoveFile(db, cb, ids...)
}

// GetOptions defined TODO
func (x *XDB) GetOptions(db *xorm.Engine, keys ...string) (map[string]map[string]interface{}, error) {
	var optSets []types.SysOptionset
	if err := db.Where("is_delete != 1").In("code", keys).Find(&optSets); err != nil {
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

// NewXDB defined TODO
func NewXDB() *XDB {
	return &XDB{}
}

// NewServiceContext defined TODO
// ======================================================================
// ======================================================================
// ======================================================================
func NewServiceContext(cache persistence.CacheStore) *ServiceContext {
	return &ServiceContext{
		Store:  cache,
		Report: NewXReport(),
		Client: NewXClient(),
		DB:     NewXDB(),
	}
}
