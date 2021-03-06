// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package api

import (
	"aisle/srv"
	"aisle/svc"
	"aisle/types"

	"github.com/spf13/viper"
)

// Name project
var Name = "aisle"
var NopHandlerFunc = func(ctx *Context) { ctx.Next() }

// Controller defined
type Controller struct {
	Method       string
	RelativePath string
	Auth,
	Roles,
	Cache,
	Interceptor,
	Handler HandlerFunc
}

// Organ defined
type Organ struct {
	Name string
	Srv  *srv.Organ
	Add,
	BatchAdd,
	Del,
	BatchDel,
	Update,
	BatchUpdate,
	Page,
	Get Controller
}

// NewOrgan defined
func NewOrgan() *Organ {
	ctr := &Organ{Name: "organ", Srv: srv.NewOrgan()}
	ctr.Add.Method = "POST"
	ctr.Add.RelativePath = "/organ/add"
	ctr.Add.Auth = Auth("token")
	ctr.Add.Roles = NopHandlerFunc
	ctr.Add.Cache = NopHandlerFunc
	ctr.Add.Interceptor = NopHandlerFunc
	ctr.Add.Handler = ctr.OrganAdd
	ctr.BatchAdd.Method = "POST"
	ctr.BatchAdd.RelativePath = "/organ/batch_add"
	ctr.BatchAdd.Auth = Auth("token")
	ctr.BatchAdd.Roles = NopHandlerFunc
	ctr.BatchAdd.Cache = NopHandlerFunc
	ctr.BatchAdd.Interceptor = NopHandlerFunc
	ctr.BatchAdd.Handler = ctr.OrganBatchAdd
	ctr.Del.Method = "DELETE"
	ctr.Del.RelativePath = "/organ/del"
	ctr.Del.Auth = Auth("token")
	ctr.Del.Roles = NopHandlerFunc
	ctr.Del.Cache = NopHandlerFunc
	ctr.Del.Interceptor = NopHandlerFunc
	ctr.Del.Handler = ctr.OrganDel
	ctr.BatchDel.Method = "PUT"
	ctr.BatchDel.RelativePath = "/organ/batch_del"
	ctr.BatchDel.Auth = Auth("token")
	ctr.BatchDel.Roles = NopHandlerFunc
	ctr.BatchDel.Cache = NopHandlerFunc
	ctr.BatchDel.Interceptor = NopHandlerFunc
	ctr.BatchDel.Handler = ctr.OrganBatchDel
	ctr.Update.Method = "PUT"
	ctr.Update.RelativePath = "/organ/update"
	ctr.Update.Auth = Auth("token")
	ctr.Update.Roles = NopHandlerFunc
	ctr.Update.Cache = NopHandlerFunc
	ctr.Update.Interceptor = NopHandlerFunc
	ctr.Update.Handler = ctr.OrganUpdate
	ctr.BatchUpdate.Method = "PUT"
	ctr.BatchUpdate.RelativePath = "/organ/batch_update"
	ctr.BatchUpdate.Auth = Auth("token")
	ctr.BatchUpdate.Roles = NopHandlerFunc
	ctr.BatchUpdate.Cache = NopHandlerFunc
	ctr.BatchUpdate.Interceptor = NopHandlerFunc
	ctr.BatchUpdate.Handler = ctr.OrganBatchUpdate
	ctr.Page.Method = "GET"
	ctr.Page.RelativePath = "/organ/page"
	ctr.Page.Auth = Auth("token")
	ctr.Page.Roles = NopHandlerFunc
	ctr.Page.Cache = NopHandlerFunc
	ctr.Page.Interceptor = NopHandlerFunc
	ctr.Page.Handler = ctr.OrganPage
	ctr.Get.Method = "GET"
	ctr.Get.RelativePath = "/organ/get"
	ctr.Get.Auth = Auth("token")
	ctr.Get.Roles = NopHandlerFunc
	ctr.Get.Cache = NopHandlerFunc
	ctr.Get.Interceptor = NopHandlerFunc
	ctr.Get.Handler = ctr.OrganGet
	return ctr
}

// OrganRoutes defined
func OrganRoutes(rg *RouterGroup) {
	group, instance := rg.Group(viper.GetString("http.prefix")), OrganInstance
	group.Handle(instance.Add.Method, instance.Add.RelativePath, instance.Add.Auth, instance.Add.Roles, instance.Add.Cache, instance.Add.Interceptor, instance.Add.Handler)
	group.Handle(instance.BatchAdd.Method, instance.BatchAdd.RelativePath, instance.BatchAdd.Auth, instance.BatchAdd.Roles, instance.BatchAdd.Cache, instance.BatchAdd.Interceptor, instance.BatchAdd.Handler)
	group.Handle(instance.Del.Method, instance.Del.RelativePath, instance.Del.Auth, instance.Del.Roles, instance.Del.Cache, instance.Del.Interceptor, instance.Del.Handler)
	group.Handle(instance.BatchDel.Method, instance.BatchDel.RelativePath, instance.BatchDel.Auth, instance.BatchDel.Roles, instance.BatchDel.Cache, instance.BatchDel.Interceptor, instance.BatchDel.Handler)
	group.Handle(instance.Update.Method, instance.Update.RelativePath, instance.Update.Auth, instance.Update.Roles, instance.Update.Cache, instance.Update.Interceptor, instance.Update.Handler)
	group.Handle(instance.BatchUpdate.Method, instance.BatchUpdate.RelativePath, instance.BatchUpdate.Auth, instance.BatchUpdate.Roles, instance.BatchUpdate.Cache, instance.BatchUpdate.Interceptor, instance.BatchUpdate.Handler)
	group.Handle(instance.Page.Method, instance.Page.RelativePath, instance.Page.Auth, instance.Page.Roles, instance.Page.Cache, instance.Page.Interceptor, instance.Page.Handler)
	group.Handle(instance.Get.Method, instance.Get.RelativePath, instance.Get.Auth, instance.Get.Roles, instance.Get.Cache, instance.Get.Interceptor, instance.Get.Handler)
}

// OrganInstance defined
var OrganInstance = NewOrgan()

// SyncModel defined
func (dol *Dolphin) SyncModel() error {
	mseti := dol.Manager.ModelSet()
	mseti.Add(new(types.AboutUs))
	mseti.Add(new(types.AboutusContentPic))
	mseti.Add(new(types.AchievementSendEmailMsg))
	mseti.Add(new(types.ActiveLesson))
	mseti.Add(new(types.ActiveLessonProType))
	mseti.Add(new(types.ActiveLessonStudent))
	mseti.Add(new(types.ActiveTeacher))
	mseti.Add(new(types.ActiveUser))
	mseti.Add(new(types.AddAgreementMb))
	mseti.Add(new(types.AndroidHotfix))
	mseti.Add(new(types.AppError))
	mseti.Add(new(types.AppMessageFile))
	mseti.Add(new(types.AppMessageNotification))
	mseti.Add(new(types.AppMessageSch))
	mseti.Add(new(types.AppMessageStu))
	mseti.Add(new(types.Area))
	mseti.Add(new(types.AutoCsTimeSet))
	mseti.Add(new(types.BranchCompany))
	mseti.Add(new(types.BusinessDepartment))
	mseti.Add(new(types.ButtonPicture))
	mseti.Add(new(types.BuysctGlGivesct))
	mseti.Add(new(types.Casting))
	mseti.Add(new(types.CcRewardUser))
	mseti.Add(new(types.CcsHomeworkCheckFile))
	mseti.Add(new(types.ChangeTeaHisTea))
	mseti.Add(new(types.ChangeTeaNewTea))
	mseti.Add(new(types.ChannelPlan))
	mseti.Add(new(types.ChannelSchoolPlan))
	mseti.Add(new(types.CheckFlowPool))
	mseti.Add(new(types.CheckFlowSet))
	mseti.Add(new(types.CheckFlowSetFloor))
	mseti.Add(new(types.CheckFlowSetFloorCasting))
	mseti.Add(new(types.CheckFlowSetFloorUser))
	mseti.Add(new(types.ClassAllotCustomer))
	mseti.Add(new(types.ClassChangeTea))
	mseti.Add(new(types.ClassChangeTeaHis))
	mseti.Add(new(types.ClassChangeTeaNew))
	mseti.Add(new(types.ClassContent))
	mseti.Add(new(types.ClassContentFile))
	mseti.Add(new(types.ClassFeedback))
	mseti.Add(new(types.ClassFeedbackFile))
	mseti.Add(new(types.ClassManage))
	mseti.Add(new(types.ClassManageStageCourse))
	mseti.Add(new(types.ClassManageTea))
	mseti.Add(new(types.ClassManagerPt))
	mseti.Add(new(types.ClassProductCategories))
	mseti.Add(new(types.ClassSchedule))
	mseti.Add(new(types.ClassScheduleStudent))
	mseti.Add(new(types.ClassScheduleTask))
	mseti.Add(new(types.ClassTimeRange))
	mseti.Add(new(types.ClassTimeRangeTea))
	mseti.Add(new(types.ClassTimeResource))
	mseti.Add(new(types.ClassType))
	mseti.Add(new(types.ClassTypeOrgan))
	mseti.Add(new(types.ClassTypePt))
	mseti.Add(new(types.ClassTypeStageCourse))
	mseti.Add(new(types.ClassYpSc))
	mseti.Add(new(types.ClassroomTimeResource))
	mseti.Add(new(types.CmStandardOperation))
	mseti.Add(new(types.Complaint))
	mseti.Add(new(types.ComplaintDeal))
	mseti.Add(new(types.Course))
	mseti.Add(new(types.CourseProductType))
	mseti.Add(new(types.CourseVideo))
	mseti.Add(new(types.CourseVideoFile))
	mseti.Add(new(types.CoverComplaintPeople))
	mseti.Add(new(types.CpcCity))
	mseti.Add(new(types.CpcCt))
	mseti.Add(new(types.CsChangeTeacher))
	mseti.Add(new(types.CsExecuteContent))
	mseti.Add(new(types.CsKfHourDetail))
	mseti.Add(new(types.CsStuTeacher))
	mseti.Add(new(types.CsTaskFiles))
	mseti.Add(new(types.CsVisit))
	mseti.Add(new(types.CspvTea))
	mseti.Add(new(types.CssChageStudent))
	mseti.Add(new(types.CssChangeKf))
	mseti.Add(new(types.CssCsTask))
	mseti.Add(new(types.CssCsTaskFile))
	mseti.Add(new(types.CssDeleteRecord))
	mseti.Add(new(types.CssHFFile))
	mseti.Add(new(types.CustomerServiceProcess))
	mseti.Add(new(types.CustomerServiceProcessVisit))
	mseti.Add(new(types.DataDeletionBackups))
	mseti.Add(new(types.DataDictionary))
	mseti.Add(new(types.DataDictionaryCopy))
	mseti.Add(new(types.DateTable))
	mseti.Add(new(types.DegreeDeposit))
	mseti.Add(new(types.Deletestudentbackups))
	mseti.Add(new(types.Department))
	mseti.Add(new(types.DepositIn))
	mseti.Add(new(types.DepositOut))
	mseti.Add(new(types.DepositOutFile))
	mseti.Add(new(types.DepositOutFlow))
	mseti.Add(new(types.DoRefapply))
	mseti.Add(new(types.EcOfYj))
	mseti.Add(new(types.EcUserTargetplan))
	mseti.Add(new(types.EducationalResearch))
	mseti.Add(new(types.EducationalResearchFile))
	mseti.Add(new(types.EducationalResearchSchool))
	mseti.Add(new(types.EmailMessage))
	mseti.Add(new(types.EmailSendSet))
	mseti.Add(new(types.EmailSendUser))
	mseti.Add(new(types.Enterprises))
	mseti.Add(new(types.EnterprisesOrgan))
	mseti.Add(new(types.ExamPaper))
	mseti.Add(new(types.ExamPaperBlock))
	mseti.Add(new(types.ExamPaperCourse))
	mseti.Add(new(types.ExamPaperIntroduction))
	mseti.Add(new(types.ExamPaperProType))
	mseti.Add(new(types.ExportLog))
	mseti.Add(new(types.ExternalPlace))
	mseti.Add(new(types.FailApply))
	mseti.Add(new(types.FailApplyFile))
	mseti.Add(new(types.FailStudentVisit))
	mseti.Add(new(types.FamilyShareNumbe))
	mseti.Add(new(types.FavourableProductType))
	mseti.Add(new(types.Fee))
	mseti.Add(new(types.FeeStandard))
	mseti.Add(new(types.FeeStandardCity))
	mseti.Add(new(types.FeeStandardCourse))
	mseti.Add(new(types.Files))
	mseti.Add(new(types.FqorderZx))
	mseti.Add(new(types.FreezeOffApply))
	mseti.Add(new(types.FriendlyLinkSet))
	mseti.Add(new(types.GaiFile))
	mseti.Add(new(types.Gift))
	mseti.Add(new(types.GiftCity))
	mseti.Add(new(types.GiftExchange))
	mseti.Add(new(types.GiftExchangeJfsc))
	mseti.Add(new(types.GjssAppIntroduction))
	mseti.Add(new(types.HolidaySet))
	mseti.Add(new(types.HomeEmailMessage))
	mseti.Add(new(types.HomePhoneMessage))
	mseti.Add(new(types.HomeWeixinMessage))
	mseti.Add(new(types.HotfixFiles))
	mseti.Add(new(types.HourAllotMake))
	mseti.Add(new(types.HourAllotRatio))
	mseti.Add(new(types.IcbcPayList))
	mseti.Add(new(types.InviteVisit))
	mseti.Add(new(types.Invoice))
	mseti.Add(new(types.JfscFlSyTuFile))
	mseti.Add(new(types.JfscFlSytu))
	mseti.Add(new(types.JfscOrder))
	mseti.Add(new(types.JfscOrderGe))
	mseti.Add(new(types.JfscSyTu))
	mseti.Add(new(types.JfscSyTuFile))
	mseti.Add(new(types.KfChangeHistory))
	mseti.Add(new(types.LabelInfo))
	mseti.Add(new(types.LanguagePack))
	mseti.Add(new(types.LanguagePackCity))
	mseti.Add(new(types.LanguagePackOrder))
	mseti.Add(new(types.LbfJl))
	mseti.Add(new(types.LearningparkBackgroundMap))
	mseti.Add(new(types.LearningparkBook))
	mseti.Add(new(types.LearningparkBookFileid))
	mseti.Add(new(types.Level))
	mseti.Add(new(types.LevelUnit))
	mseti.Add(new(types.LgClasstype))
	mseti.Add(new(types.Loginlanguageset))
	mseti.Add(new(types.MaBusiness))
	mseti.Add(new(types.MaTeacher))
	mseti.Add(new(types.MarketActivity))
	mseti.Add(new(types.MarketActivityHead))
	mseti.Add(new(types.MarketActivityMateriel))
	mseti.Add(new(types.MarketActivityProcess))
	mseti.Add(new(types.MarketExpend))
	mseti.Add(new(types.MarketFeeBudget))
	mseti.Add(new(types.MarketFeeBudgetFiles))
	mseti.Add(new(types.MarketFeeCost))
	mseti.Add(new(types.MarketFeeName))
	mseti.Add(new(types.MarketInviteStu))
	mseti.Add(new(types.MarketModel))
	mseti.Add(new(types.MarketModelMateriel))
	mseti.Add(new(types.MarketModelProcess))
	mseti.Add(new(types.MarketerPlanTarget))
	mseti.Add(new(types.MarketingActivitiesCity))
	mseti.Add(new(types.MarketingActivitiesSchool))
	mseti.Add(new(types.MaterialApply))
	mseti.Add(new(types.MaterialBudget))
	mseti.Add(new(types.MaterialTable))
	mseti.Add(new(types.Materiel))
	mseti.Add(new(types.MaterielAdd))
	mseti.Add(new(types.MaterielMinus))
	mseti.Add(new(types.MenuPicture))
	mseti.Add(new(types.MessageNotice))
	mseti.Add(new(types.NetSchoolPlan))
	mseti.Add(new(types.NetUserPlan))
	mseti.Add(new(types.NetworkActivity))
	mseti.Add(new(types.NetworkActivityExpend))
	mseti.Add(new(types.NetworkActivityHead))
	mseti.Add(new(types.NetworkActivityOrgan))
	mseti.Add(new(types.NetworkDetail))
	mseti.Add(new(types.NewKxRecord))
	mseti.Add(new(types.NewStuTransition))
	mseti.Add(new(types.OfDepositDetail))
	mseti.Add(new(types.OfOverdue))
	mseti.Add(new(types.OfOverdueSct))
	mseti.Add(new(types.OldGiveGift))
	mseti.Add(new(types.OldPushStudent))
	mseti.Add(new(types.OnSale))
	mseti.Add(new(types.OnSaleCity))
	mseti.Add(new(types.OnSaleFiles))
	mseti.Add(new(types.OrderDeleteLog))
	mseti.Add(new(types.OrderFiles))
	mseti.Add(new(types.OrderForm))
	mseti.Add(new(types.OrderFromStuIntegral))
	mseti.Add(new(types.OrderFromStudyCard))
	mseti.Add(new(types.OrderProduct))
	mseti.Add(new(types.Organ))
	mseti.Add(new(types.OrganAddAgreement))
	mseti.Add(new(types.OrganSchool))
	mseti.Add(new(types.OrganSchoolBussType))
	mseti.Add(new(types.OtmClassroom))
	mseti.Add(new(types.OtomSysSet))
	mseti.Add(new(types.OverseasMaPt))
	mseti.Add(new(types.PaActiveReservation))
	mseti.Add(new(types.PaActivityClassFile))
	mseti.Add(new(types.PaActivityClassStu))
	mseti.Add(new(types.ParFile))
	mseti.Add(new(types.ParStudnetClassType))
	mseti.Add(new(types.ParentalEvaluation))
	mseti.Add(new(types.PartTimeUser))
	mseti.Add(new(types.PartTimeUserCity))
	mseti.Add(new(types.ParticipantEvaluateInfo))
	mseti.Add(new(types.ParticipantEvaluateTrainer))
	mseti.Add(new(types.ParticipantTable))
	mseti.Add(new(types.Partner))
	mseti.Add(new(types.PartnerOrgan))
	mseti.Add(new(types.PcSctGiveMx))
	mseti.Add(new(types.PeriodViolateAbnormalApply))
	mseti.Add(new(types.PhoneMessage))
	mseti.Add(new(types.Phonehistory))
	mseti.Add(new(types.PrearrangedCourses))
	mseti.Add(new(types.ProductType))
	mseti.Add(new(types.ProtocolType))
	mseti.Add(new(types.ReTeachingTime))
	mseti.Add(new(types.RecipientInfo))
	mseti.Add(new(types.RefYishiFile))
	mseti.Add(new(types.Refund))
	mseti.Add(new(types.RefundCheckFiles))
	mseti.Add(new(types.RefundDetailReason))
	mseti.Add(new(types.RefundFiles))
	mseti.Add(new(types.RefundFlow))
	mseti.Add(new(types.RefundProcessSet))
	mseti.Add(new(types.RefundStuClassType))
	mseti.Add(new(types.RefundStuTextBook))
	mseti.Add(new(types.RegularDictionary))
	mseti.Add(new(types.ReplaceGift))
	mseti.Add(new(types.RpsCity))
	mseti.Add(new(types.RpsSch))
	mseti.Add(new(types.SaleLanguageModel))
	mseti.Add(new(types.SaleSchool))
	mseti.Add(new(types.SatisfyVisitMode))
	mseti.Add(new(types.SatisfyVisitModel))
	mseti.Add(new(types.SatisfyVisitModelContent))
	mseti.Add(new(types.ScClassTea))
	mseti.Add(new(types.SchAchievementDivide))
	mseti.Add(new(types.SchIcbcPosNum))
	mseti.Add(new(types.SchMarketPlan))
	mseti.Add(new(types.SchTaTargetplan))
	mseti.Add(new(types.SchTargetsignPlan))
	mseti.Add(new(types.School))
	mseti.Add(new(types.SchoolImgFile))
	mseti.Add(new(types.SchoolOfLimit))
	mseti.Add(new(types.ScoreItem))
	mseti.Add(new(types.SctHourDetail))
	mseti.Add(new(types.ServiceProVisitFile))
	mseti.Add(new(types.Sheng))
	mseti.Add(new(types.StuAllotTmk))
	mseti.Add(new(types.StuBfJl))
	mseti.Add(new(types.StuBussType))
	mseti.Add(new(types.StuClassContent))
	mseti.Add(new(types.StuClassFeedbacFile))
	mseti.Add(new(types.StuClassFeedback))
	mseti.Add(new(types.StuClassTypeAssBook))
	mseti.Add(new(types.StuClassTypeSusHour))
	mseti.Add(new(types.StuClassTypeTextbook))
	mseti.Add(new(types.StuConfirmHour))
	mseti.Add(new(types.StuCourseVideo))
	mseti.Add(new(types.StuCtTimeRange))
	mseti.Add(new(types.StuExchangeGift))
	mseti.Add(new(types.StuFailToLive))
	mseti.Add(new(types.StuFinishclass))
	mseti.Add(new(types.StuGwc))
	mseti.Add(new(types.StuGwcGef))
	mseti.Add(new(types.StuKwJfx))
	mseti.Add(new(types.StuLaostuGift))
	mseti.Add(new(types.StuLoss))
	mseti.Add(new(types.StuMonthYpHour))
	mseti.Add(new(types.StuParent))
	mseti.Add(new(types.StuPcHour))
	mseti.Add(new(types.StuProductType))
	mseti.Add(new(types.StuSatisfyVisit))
	mseti.Add(new(types.StuSatisfyVisitContent))
	mseti.Add(new(types.StuScore))
	mseti.Add(new(types.StuScoreItem))
	mseti.Add(new(types.StuToReopened))
	mseti.Add(new(types.StuTypeStudyCard))
	mseti.Add(new(types.StuUseTextbook))
	mseti.Add(new(types.StuUserDate))
	mseti.Add(new(types.StuWxdqDate))
	mseti.Add(new(types.Stucasting))
	mseti.Add(new(types.StuclasstypeExtrafee))
	mseti.Add(new(types.Student))
	mseti.Add(new(types.StudentAddAgreement))
	mseti.Add(new(types.StudentCastingTable))
	mseti.Add(new(types.StudentCc))
	mseti.Add(new(types.StudentClass))
	mseti.Add(new(types.StudentClassSchedule))
	mseti.Add(new(types.StudentClassType))
	mseti.Add(new(types.StudentClassTypeUseIntegral))
	mseti.Add(new(types.StudentClassYpk))
	mseti.Add(new(types.StudentClassYpk1))
	mseti.Add(new(types.StudentClassYpk2))
	mseti.Add(new(types.StudentClassYpk3))
	mseti.Add(new(types.StudentClassYpk31))
	mseti.Add(new(types.StudentClassYpk32))
	mseti.Add(new(types.StudentClassYpk33))
	mseti.Add(new(types.StudentClassYpkDkb))
	mseti.Add(new(types.StudentCustomer))
	mseti.Add(new(types.StudentCzIntegral))
	mseti.Add(new(types.StudentDeleteLog))
	mseti.Add(new(types.StudentGift))
	mseti.Add(new(types.StudentGrowthRecord))
	mseti.Add(new(types.StudentGrowthRecordFile))
	mseti.Add(new(types.StudentMarketActivity))
	mseti.Add(new(types.StudentOrgan))
	mseti.Add(new(types.StudentOrganSchool))
	mseti.Add(new(types.StudentReopened))
	mseti.Add(new(types.StudentRzIntegral))
	mseti.Add(new(types.StudentSale))
	mseti.Add(new(types.StudentTimeResource))
	mseti.Add(new(types.StudentTypeJournal))
	mseti.Add(new(types.StudentUseHourDetail))
	mseti.Add(new(types.StudentVisit))
	mseti.Add(new(types.StudentVisitLanguage))
	mseti.Add(new(types.StudyCard))
	mseti.Add(new(types.SystemMessage))
	mseti.Add(new(types.SystemNotice))
	mseti.Add(new(types.SystemOperaLogs))
	mseti.Add(new(types.SystemSet))
	mseti.Add(new(types.T0UserInfo))
	mseti.Add(new(types.T19LoginLog))
	mseti.Add(new(types.T2LoginLog))
	mseti.Add(new(types.T3Program))
	mseti.Add(new(types.T520LoginLog))
	mseti.Add(new(types.TaTargetPlan))
	mseti.Add(new(types.TeaLeaveManagement))
	mseti.Add(new(types.TeachLog))
	mseti.Add(new(types.TeacherHourDetail))
	mseti.Add(new(types.TeacherServiceProcess))
	mseti.Add(new(types.TeachingMaterial))
	mseti.Add(new(types.TeachingTime))
	mseti.Add(new(types.Temp20190330))
	mseti.Add(new(types.TempHuifu))
	mseti.Add(new(types.TempShaoerShow))
	mseti.Add(new(types.TempStorageCz))
	mseti.Add(new(types.TempStorageCzFlow))
	mseti.Add(new(types.TempStorageRz))
	mseti.Add(new(types.TempStu))
	mseti.Add(new(types.TempStuJifenAdd))
	mseti.Add(new(types.TestSites))
	mseti.Add(new(types.TextBook))
	mseti.Add(new(types.TextBookProductType))
	mseti.Add(new(types.TimeSet))
	mseti.Add(new(types.TimeTrigger))
	mseti.Add(new(types.TmLevel))
	mseti.Add(new(types.TmSi))
	mseti.Add(new(types.TmSiUnit))
	mseti.Add(new(types.TrainExternalPlace))
	mseti.Add(new(types.TrainOrgan))
	mseti.Add(new(types.TrainOrganSchool))
	mseti.Add(new(types.TrainerEvaluateInfo))
	mseti.Add(new(types.TrainerEvaluateParticipant))
	mseti.Add(new(types.TrainerTable))
	mseti.Add(new(types.TrialLesson))
	mseti.Add(new(types.TrialLessonProType))
	mseti.Add(new(types.TrialLessonStudent))
	mseti.Add(new(types.TrialLessonTea))
	mseti.Add(new(types.TrialLessonUser))
	mseti.Add(new(types.TscReforderFile))
	mseti.Add(new(types.Unit))
	mseti.Add(new(types.UserBussType))
	mseti.Add(new(types.UserCasting))
	mseti.Add(new(types.UserCourse))
	mseti.Add(new(types.UserDepartment))
	mseti.Add(new(types.UserGsSch))
	mseti.Add(new(types.UserMessageNotification))
	mseti.Add(new(types.UserOnoffRz))
	mseti.Add(new(types.UserOrgan))
	mseti.Add(new(types.UserOrganSchool))
	mseti.Add(new(types.UserProductType))
	mseti.Add(new(types.UserTimeResource))
	mseti.Add(new(types.UserTrainingTable))
	mseti.Add(new(types.UserTypeTable))
	mseti.Add(new(types.UserUploadFile))
	mseti.Add(new(types.ViolateRatioManage))
	mseti.Add(new(types.WastageFollowupRecord))
	mseti.Add(new(types.WechatMessage))
	mseti.Add(new(types.WholeCountry))
	mseti.Add(new(types.XlHc))
	mseti.Add(new(types.XyHourRecord))
	mseti.Add(new(types.YbPayFlowing))
	mseti.Add(new(types.YbPayOut))
	mseti.Add(new(types.YbzfPos))
	mseti.Add(new(types.YeepayReconciliationsRecord))
	mseti.Add(new(types.ZbxGiftrules))
	mseti.Add(new(types.ZbxzsgzValiddate))
	mseti.Add(new(types.ZbxzsyxCity))
	return nil
}

// SyncController defined
func (dol *Dolphin) SyncController() error {
	OrganRoutes(&dol.RouterGroup)
	return nil
}

// SyncSrv defined
func (dol *Dolphin) SyncSrv(svc svc.Svc) error {
	OrganInstance.Srv.Svc = svc
	return nil
}

// SyncService defined
func (dol *Dolphin) SyncService() error {
	return nil
}
