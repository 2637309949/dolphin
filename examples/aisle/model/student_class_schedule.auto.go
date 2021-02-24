// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// StudentClassSchedule defined
type StudentClassSchedule struct {
	//
	SCSId null.Int `xorm:"int(11) pk notnull autoincr 's_c_s_id'" json:"s_c_s_id" form:"s_c_s_id" xml:"s_c_s_id"`
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	//
	CsId null.Int `xorm:"int(11) 'cs_id'" json:"cs_id" form:"cs_id" xml:"cs_id"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	//
	Isdelete null.Int `xorm:"notnull 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	//
	ScmId null.Int `xorm:"int(11) 'scm_id'" json:"scm_id" form:"scm_id" xml:"scm_id"`
	//
	BaobanState null.Int `xorm:"int(11) 'baoban_state'" json:"baoban_state" form:"baoban_state" xml:"baoban_state"`
	//
	IvId null.Int `xorm:"int(11) 'iv_id'" json:"iv_id" form:"iv_id" xml:"iv_id"`
	//
	ScspId null.Int `xorm:"int(11) 'scsp_id'" json:"scsp_id" form:"scsp_id" xml:"scsp_id"`
	//
	BdKc null.Float `xorm:"float(10,2) 'bd_kc'" json:"bd_kc" form:"bd_kc" xml:"bd_kc"`
	//
	AttendDate null.Time `xorm:"datetime 'attend_date'" json:"attend_date" form:"attend_date" xml:"attend_date"`
	//
	AppStartTime null.Time `xorm:"datetime 'app_start_time'" json:"app_start_time" form:"app_start_time" xml:"app_start_time"`
	//
	AppEndTime null.Time `xorm:"datetime 'app_end_time'" json:"app_end_time" form:"app_end_time" xml:"app_end_time"`
	//
	StsId null.Int `xorm:"int(11) 'sts_id'" json:"sts_id" form:"sts_id" xml:"sts_id"`
	//
	KfId null.Int `xorm:"int(11) 'kf_id'" json:"kf_id" form:"kf_id" xml:"kf_id"`
	//
	KqState null.Int `xorm:"int(11) 'kq_state'" json:"kq_state" form:"kq_state" xml:"kq_state"`
	//
	BussType null.Int `xorm:"int(11) 'buss_type'" json:"buss_type" form:"buss_type" xml:"buss_type"`
	//
	IfKouHour null.Int `xorm:"int(11) 'if_kou_hour'" json:"if_kou_hour" form:"if_kou_hour" xml:"if_kou_hour"`
	//
	SctId null.Int `xorm:"int(11) 'sct_id'" json:"sct_id" form:"sct_id" xml:"sct_id"`
	//
	Sctmin null.Float `xorm:"float(50,2) default(40.00) 'sctmin'" json:"sctmin" form:"sctmin" xml:"sctmin"`
	//
	Csskqkc null.Float `xorm:"float(50,2) 'csskqkc'" json:"csskqkc" form:"csskqkc" xml:"csskqkc"`
	//
	CsskqKs null.Float `xorm:"float(50,2) 'csskq_ks'" json:"csskq_ks" form:"csskq_ks" xml:"csskq_ks"`
	//
	StuKtfkNum null.Float `xorm:"float(50,2) 'stu_ktfk_num'" json:"stu_ktfk_num" form:"stu_ktfk_num" xml:"stu_ktfk_num"`
	//
	StuKtzyNum null.Float `xorm:"float(50,2) 'stu_ktzy_num'" json:"stu_ktzy_num" form:"stu_ktzy_num" xml:"stu_ktzy_num"`
	//
	StuWczyNum null.Float `xorm:"float(50,2) 'stu_wczy_num'" json:"stu_wczy_num" form:"stu_wczy_num" xml:"stu_wczy_num"`
	//
	TeaPgNum null.Float `xorm:"float(50,2) 'tea_pg_num'" json:"tea_pg_num" form:"tea_pg_num" xml:"tea_pg_num"`
	//
	StuKtslNum null.Float `xorm:"float(50,2) 'stu_ktsl_num'" json:"stu_ktsl_num" form:"stu_ktsl_num" xml:"stu_ktsl_num"`
	//
	StuKtnrNum null.Float `xorm:"float(50,2) 'stu_ktnr_num'" json:"stu_ktnr_num" form:"stu_ktnr_num" xml:"stu_ktnr_num"`
	//
	StuPjNum null.Float `xorm:"float(50,2) 'stu_pj_num'" json:"stu_pj_num" form:"stu_pj_num" xml:"stu_pj_num"`
	//
	ClassId null.Int `xorm:"int(11) 'class_id'" json:"class_id" form:"class_id" xml:"class_id"`
	//
	SchId null.Int `xorm:"int(11) 'sch_id'" json:"sch_id" form:"sch_id" xml:"sch_id"`
	//
	OnePrice null.Float `xorm:"float(50,2) 'one_price'" json:"one_price" form:"one_price" xml:"one_price"`
	//
	AllPrice null.Float `xorm:"float(50,2) 'all_price'" json:"all_price" form:"all_price" xml:"all_price"`
	//
	RealKc null.Float `xorm:"float(50,2) 'real_kc'" json:"real_kc" form:"real_kc" xml:"real_kc"`
	//
	ChangeStuId null.Int `xorm:"int(11) 'change_stu_id'" json:"change_stu_id" form:"change_stu_id" xml:"change_stu_id"`
	//
	Remark null.String `xorm:"varchar(100) 'remark'" json:"remark" form:"remark" xml:"remark"`
}

// Parser defined
func (m *StudentClassSchedule) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *StudentClassSchedule) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StudentClassSchedule
func (m *StudentClassSchedule) TableName() string {
	return "student_class_schedule"
}
