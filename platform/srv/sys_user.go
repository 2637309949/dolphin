// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_user.go

package srv

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// SysUserTODO defined srv
func SysUserTODO(ginCtx *gin.Context, db *xorm.Engine, actCtx context.Context, params struct{}) (interface{}, error) {
	actCtx, cancel := context.WithTimeout(actCtx, 5*time.Second)
	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			select {
			case <-ctx.Done():
				logrus.Infoln("child process interrupt...")
				return
			default:
				logrus.Infoln("child job...")
			}
		}
	}(actCtx)
	defer cancel()
	<-actCtx.Done()
	logrus.Infoln("main process exit!")
	return nil, errors.New("no implementation found")
}

// SysUserGetOrgsFromInheritance defined srv
func SysUserGetOrgsFromInheritance(db *xorm.Engine, cn string) ([]string, error) {
	idst := struct {
		IDS string `xorm:"ids"`
	}{}
	_, err := db.SQL(fmt.Sprintf(`select IFNULL(GROUP_CONCAT(id), '') ids, is_delete from sys_org where is_delete=0 and inheritance like "%v" group by is_delete`, "%"+cn+"%")).Get(&idst)
	if err != nil {
		return nil, err
	}
	// if id type...
	ids := []string{}
	for _, v := range strings.Split(idst.IDS, ",") {
		ids = append(ids, fmt.Sprintf("'%v'", v))
	}
	return ids, nil
}

// SysUserGetUserRolesByUID defined
func SysUserGetUserRolesByUID(db *xorm.Engine, ids string) ([]map[string]interface{}, error) {
	roles, err := db.SqlTemplateClient("sys_role_user.tpl", &map[string]interface{}{"uids": template.HTML(ids)}).Query().List()
	if err != nil {
		return nil, err
	}
	return roles, err
}

// SysUserGetUserOrgsByUID defined
func SysUserGetUserOrgsByUID(db *xorm.Engine, ids string) ([]map[string]interface{}, error) {
	orgs, err := db.SqlTemplateClient("sys_user_org.tpl", &map[string]interface{}{"ne": template.HTML("<>"), "oids": template.HTML(ids)}).Query().List()
	if err != nil {
		return nil, err
	}
	return orgs, err
}
