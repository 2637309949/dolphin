// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package app

import (
	"aisle/model"

	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/platform/util"
)

// Name project
var Name = "aisle"

// Organ defined
type Organ struct {
	Name string
	Add,
	BatchAdd,
	Del,
	BatchDel,
	Update,
	BatchUpdate,
	Page,
	Get HandlerFunc
}

// NewOrgan defined
func NewOrgan() *Organ {
	ctr := &Organ{Name: "organ"}
	ctr.Add.Method = "POST"
	ctr.Add.RelativePath = "/organ/add"
	ctr.Add.Handler = OrganAdd
	ctr.BatchAdd.Method = "POST"
	ctr.BatchAdd.RelativePath = "/organ/batch_add"
	ctr.BatchAdd.Handler = OrganBatchAdd
	ctr.Del.Method = "DELETE"
	ctr.Del.RelativePath = "/organ/del"
	ctr.Del.Handler = OrganDel
	ctr.BatchDel.Method = "PUT"
	ctr.BatchDel.RelativePath = "/organ/batch_del"
	ctr.BatchDel.Handler = OrganBatchDel
	ctr.Update.Method = "PUT"
	ctr.Update.RelativePath = "/organ/update"
	ctr.Update.Handler = OrganUpdate
	ctr.BatchUpdate.Method = "PUT"
	ctr.BatchUpdate.RelativePath = "/organ/batch_update"
	ctr.BatchUpdate.Handler = OrganBatchUpdate
	ctr.Page.Method = "GET"
	ctr.Page.RelativePath = "/organ/page"
	ctr.Page.Handler = OrganPage
	ctr.Get.Method = "GET"
	ctr.Get.RelativePath = "/organ/get"
	ctr.Get.Handler = OrganGet
	return ctr
}

// OrganRoutes defined
func OrganRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle(OrganInstance.Add.Method, OrganInstance.Add.RelativePath, Auth("token"), OrganInstance.Add)
	group.Handle(OrganInstance.BatchAdd.Method, OrganInstance.BatchAdd.RelativePath, Auth("token"), OrganInstance.BatchAdd)
	group.Handle(OrganInstance.Del.Method, OrganInstance.Del.RelativePath, Auth("token"), OrganInstance.Del)
	group.Handle(OrganInstance.BatchDel.Method, OrganInstance.BatchDel.RelativePath, Auth("token"), OrganInstance.BatchDel)
	group.Handle(OrganInstance.Update.Method, OrganInstance.Update.RelativePath, Auth("token"), OrganInstance.Update)
	group.Handle(OrganInstance.BatchUpdate.Method, OrganInstance.BatchUpdate.RelativePath, Auth("token"), OrganInstance.BatchUpdate)
	group.Handle(OrganInstance.Page.Method, OrganInstance.Page.RelativePath, Auth("token"), OrganInstance.Page)
	group.Handle(OrganInstance.Get.Method, OrganInstance.Get.RelativePath, Auth("token"), OrganInstance.Get)
}

// OrganInstance defined
var OrganInstance = NewOrgan()

// SyncModel defined
func SyncModel() error {
	mseti := App.Manager.MSet()
	mseti.Add(new(model.AboutUs))
	mseti.Add(new(model.AboutusContentPic))
	mseti.Add(new(model.AchievementSendEmailMsg))
	mseti.Add(new(model.ActiveLesson))
	mseti.Add(new(model.ActiveLessonProType))
	mseti.Add(new(model.ActiveLessonStudent))
	mseti.Add(new(model.ActiveTeacher))
	mseti.Add(new(model.ActiveUser))
	mseti.Add(new(model.AddAgreementMb))
	mseti.Add(new(model.AndroidHotfix))
	mseti.Add(new(model.AppError))
	mseti.Add(new(model.AppMessageFile))
	mseti.Add(new(model.AppMessageNotification))
	mseti.Add(new(model.AppMessageSch))
	mseti.Add(new(model.AppMessageStu))
	mseti.Add(new(model.Area))
	mseti.Add(new(model.AutoCsTimeSet))
	mseti.Add(new(model.BranchCompany))
	mseti.Add(new(model.BusinessDepartment))
	mseti.Add(new(model.ButtonPicture))
	mseti.Add(new(model.BuysctGlGivesct))
	mseti.Add(new(model.Casting))
	mseti.Add(new(model.CcRewardUser))
	mseti.Add(new(model.CcsHomeworkCheckFile))
	mseti.Add(new(model.ChangeTeaHisTea))
	mseti.Add(new(model.ChangeTeaNewTea))
	mseti.Add(new(model.ChannelPlan))
	mseti.Add(new(model.ChannelSchoolPlan))
	mseti.Add(new(model.CheckFlowPool))
	mseti.Add(new(model.CheckFlowSet))
	mseti.Add(new(model.CheckFlowSetFloor))
	mseti.Add(new(model.CheckFlowSetFloorCasting))
	mseti.Add(new(model.CheckFlowSetFloorUser))
	mseti.Add(new(model.ClassAllotCustomer))
	mseti.Add(new(model.ClassChangeTea))
	mseti.Add(new(model.ClassChangeTeaHis))
	mseti.Add(new(model.ClassChangeTeaNew))
	mseti.Add(new(model.ClassContent))
	mseti.Add(new(model.ClassContentFile))
	mseti.Add(new(model.ClassFeedback))
	mseti.Add(new(model.ClassFeedbackFile))
	mseti.Add(new(model.ClassManage))
	mseti.Add(new(model.ClassManageStageCourse))
	mseti.Add(new(model.ClassManageTea))
	mseti.Add(new(model.ClassManagerPt))
	mseti.Add(new(model.ClassProductCategories))
	mseti.Add(new(model.ClassSchedule))
	mseti.Add(new(model.ClassScheduleStudent))
	mseti.Add(new(model.ClassScheduleTask))
	mseti.Add(new(model.ClassTimeRange))
	mseti.Add(new(model.ClassTimeRangeTea))
	mseti.Add(new(model.ClassTimeResource))
	mseti.Add(new(model.ClassType))
	mseti.Add(new(model.ClassTypeOrgan))
	mseti.Add(new(model.ClassTypePt))
	mseti.Add(new(model.ClassTypeStageCourse))
	mseti.Add(new(model.ClassYpSc))
	mseti.Add(new(model.ClassroomTimeResource))
	mseti.Add(new(model.CmStandardOperation))
	mseti.Add(new(model.Complaint))
	mseti.Add(new(model.ComplaintDeal))
	mseti.Add(new(model.Course))
	mseti.Add(new(model.CourseProductType))
	mseti.Add(new(model.CourseVideo))
	mseti.Add(new(model.CourseVideoFile))
	mseti.Add(new(model.CoverComplaintPeople))
	mseti.Add(new(model.CpcCity))
	mseti.Add(new(model.CpcCt))
	mseti.Add(new(model.CsChangeTeacher))
	mseti.Add(new(model.CsExecuteContent))
	mseti.Add(new(model.CsKfHourDetail))
	mseti.Add(new(model.CsStuTeacher))
	mseti.Add(new(model.CsTaskFiles))
	mseti.Add(new(model.CsVisit))
	mseti.Add(new(model.CspvTea))
	mseti.Add(new(model.CssChageStudent))
	mseti.Add(new(model.CssChangeKf))
	mseti.Add(new(model.CssCsTask))
	mseti.Add(new(model.CssCsTaskFile))
	mseti.Add(new(model.CssDeleteRecord))
	mseti.Add(new(model.CssHFFile))
	mseti.Add(new(model.CustomerServiceProcess))
	mseti.Add(new(model.CustomerServiceProcessVisit))
	mseti.Add(new(model.DataDeletionBackups))
	mseti.Add(new(model.DataDictionary))
	mseti.Add(new(model.DataDictionaryCopy))
	mseti.Add(new(model.DateTable))
	mseti.Add(new(model.DegreeDeposit))
	mseti.Add(new(model.Deletestudentbackups))
	mseti.Add(new(model.Department))
	mseti.Add(new(model.DepositIn))
	mseti.Add(new(model.DepositOut))
	mseti.Add(new(model.DepositOutFile))
	mseti.Add(new(model.DepositOutFlow))
	mseti.Add(new(model.DoRefapply))
	mseti.Add(new(model.EcOfYj))
	mseti.Add(new(model.EcUserTargetplan))
	mseti.Add(new(model.EducationalResearch))
	mseti.Add(new(model.EducationalResearchFile))
	mseti.Add(new(model.EducationalResearchSchool))
	mseti.Add(new(model.EmailMessage))
	mseti.Add(new(model.EmailSendSet))
	mseti.Add(new(model.EmailSendUser))
	mseti.Add(new(model.Enterprises))
	mseti.Add(new(model.EnterprisesOrgan))
	mseti.Add(new(model.ExamPaper))
	mseti.Add(new(model.ExamPaperBlock))
	mseti.Add(new(model.ExamPaperCourse))
	mseti.Add(new(model.ExamPaperIntroduction))
	mseti.Add(new(model.ExamPaperProType))
	mseti.Add(new(model.ExportLog))
	mseti.Add(new(model.ExternalPlace))
	mseti.Add(new(model.FailApply))
	mseti.Add(new(model.FailApplyFile))
	mseti.Add(new(model.FailStudentVisit))
	mseti.Add(new(model.FamilyShareNumbe))
	mseti.Add(new(model.FavourableProductType))
	mseti.Add(new(model.Fee))
	mseti.Add(new(model.FeeStandard))
	mseti.Add(new(model.FeeStandardCity))
	mseti.Add(new(model.FeeStandardCourse))
	mseti.Add(new(model.Files))
	mseti.Add(new(model.FqorderZx))
	mseti.Add(new(model.FreezeOffApply))
	mseti.Add(new(model.FriendlyLinkSet))
	mseti.Add(new(model.GaiFile))
	mseti.Add(new(model.Gift))
	mseti.Add(new(model.GiftCity))
	mseti.Add(new(model.GiftExchange))
	mseti.Add(new(model.GiftExchangeJfsc))
	mseti.Add(new(model.GjssAppIntroduction))
	mseti.Add(new(model.HolidaySet))
	mseti.Add(new(model.HomeEmailMessage))
	mseti.Add(new(model.HomePhoneMessage))
	mseti.Add(new(model.HomeWeixinMessage))
	mseti.Add(new(model.HotfixFiles))
	mseti.Add(new(model.HourAllotMake))
	mseti.Add(new(model.HourAllotRatio))
	mseti.Add(new(model.IcbcPayList))
	mseti.Add(new(model.InviteVisit))
	mseti.Add(new(model.Invoice))
	mseti.Add(new(model.JfscFlSyTuFile))
	mseti.Add(new(model.JfscFlSytu))
	mseti.Add(new(model.JfscOrder))
	mseti.Add(new(model.JfscOrderGe))
	mseti.Add(new(model.JfscSyTu))
	mseti.Add(new(model.JfscSyTuFile))
	mseti.Add(new(model.KfChangeHistory))
	mseti.Add(new(model.LabelInfo))
	mseti.Add(new(model.LanguagePack))
	mseti.Add(new(model.LanguagePackCity))
	mseti.Add(new(model.LanguagePackOrder))
	mseti.Add(new(model.LbfJl))
	mseti.Add(new(model.LearningparkBackgroundMap))
	mseti.Add(new(model.LearningparkBook))
	mseti.Add(new(model.LearningparkBookFileid))
	mseti.Add(new(model.Level))
	mseti.Add(new(model.LevelUnit))
	mseti.Add(new(model.LgClasstype))
	mseti.Add(new(model.Loginlanguageset))
	mseti.Add(new(model.MaBusiness))
	mseti.Add(new(model.MaTeacher))
	mseti.Add(new(model.MarketActivity))
	mseti.Add(new(model.MarketActivityHead))
	mseti.Add(new(model.MarketActivityMateriel))
	mseti.Add(new(model.MarketActivityProcess))
	mseti.Add(new(model.MarketExpend))
	mseti.Add(new(model.MarketFeeBudget))
	mseti.Add(new(model.MarketFeeBudgetFiles))
	mseti.Add(new(model.MarketFeeCost))
	mseti.Add(new(model.MarketFeeName))
	mseti.Add(new(model.MarketInviteStu))
	mseti.Add(new(model.MarketModel))
	mseti.Add(new(model.MarketModelMateriel))
	mseti.Add(new(model.MarketModelProcess))
	mseti.Add(new(model.MarketerPlanTarget))
	mseti.Add(new(model.MarketingActivitiesCity))
	mseti.Add(new(model.MarketingActivitiesSchool))
	mseti.Add(new(model.MaterialApply))
	mseti.Add(new(model.MaterialBudget))
	mseti.Add(new(model.MaterialTable))
	mseti.Add(new(model.Materiel))
	mseti.Add(new(model.MaterielAdd))
	mseti.Add(new(model.MaterielMinus))
	mseti.Add(new(model.MenuPicture))
	mseti.Add(new(model.MessageNotice))
	mseti.Add(new(model.NetSchoolPlan))
	mseti.Add(new(model.NetUserPlan))
	mseti.Add(new(model.NetworkActivity))
	mseti.Add(new(model.NetworkActivityExpend))
	mseti.Add(new(model.NetworkActivityHead))
	mseti.Add(new(model.NetworkActivityOrgan))
	mseti.Add(new(model.NetworkDetail))
	mseti.Add(new(model.NewKxRecord))
	mseti.Add(new(model.NewStuTransition))
	mseti.Add(new(model.OfDepositDetail))
	mseti.Add(new(model.OfOverdue))
	mseti.Add(new(model.OfOverdueSct))
	mseti.Add(new(model.OldGiveGift))
	mseti.Add(new(model.OldPushStudent))
	mseti.Add(new(model.OnSale))
	mseti.Add(new(model.OnSaleCity))
	mseti.Add(new(model.OnSaleFiles))
	mseti.Add(new(model.OrderDeleteLog))
	mseti.Add(new(model.OrderFiles))
	mseti.Add(new(model.OrderForm))
	mseti.Add(new(model.OrderFromStuIntegral))
	mseti.Add(new(model.OrderFromStudyCard))
	mseti.Add(new(model.OrderProduct))
	mseti.Add(new(model.Organ))
	mseti.Add(new(model.OrganAddAgreement))
	mseti.Add(new(model.OrganSchool))
	mseti.Add(new(model.OrganSchoolBussType))
	mseti.Add(new(model.OtmClassroom))
	mseti.Add(new(model.OtomSysSet))
	mseti.Add(new(model.OverseasMaPt))
	mseti.Add(new(model.PaActiveReservation))
	mseti.Add(new(model.PaActivityClassFile))
	mseti.Add(new(model.PaActivityClassStu))
	mseti.Add(new(model.ParFile))
	mseti.Add(new(model.ParStudnetClassType))
	mseti.Add(new(model.ParentalEvaluation))
	mseti.Add(new(model.PartTimeUser))
	mseti.Add(new(model.PartTimeUserCity))
	mseti.Add(new(model.ParticipantEvaluateInfo))
	mseti.Add(new(model.ParticipantEvaluateTrainer))
	mseti.Add(new(model.ParticipantTable))
	mseti.Add(new(model.Partner))
	mseti.Add(new(model.PartnerOrgan))
	mseti.Add(new(model.PcSctGiveMx))
	mseti.Add(new(model.PeriodViolateAbnormalApply))
	mseti.Add(new(model.PhoneMessage))
	mseti.Add(new(model.Phonehistory))
	mseti.Add(new(model.PrearrangedCourses))
	mseti.Add(new(model.ProductType))
	mseti.Add(new(model.ProtocolType))
	mseti.Add(new(model.ReTeachingTime))
	mseti.Add(new(model.RecipientInfo))
	mseti.Add(new(model.RefYishiFile))
	mseti.Add(new(model.Refund))
	mseti.Add(new(model.RefundCheckFiles))
	mseti.Add(new(model.RefundDetailReason))
	mseti.Add(new(model.RefundFiles))
	mseti.Add(new(model.RefundFlow))
	mseti.Add(new(model.RefundProcessSet))
	mseti.Add(new(model.RefundStuClassType))
	mseti.Add(new(model.RefundStuTextBook))
	mseti.Add(new(model.RegularDictionary))
	mseti.Add(new(model.ReplaceGift))
	mseti.Add(new(model.RpsCity))
	mseti.Add(new(model.RpsSch))
	mseti.Add(new(model.SaleLanguageModel))
	mseti.Add(new(model.SaleSchool))
	mseti.Add(new(model.SatisfyVisitMode))
	mseti.Add(new(model.SatisfyVisitModel))
	mseti.Add(new(model.SatisfyVisitModelContent))
	mseti.Add(new(model.ScClassTea))
	mseti.Add(new(model.SchAchievementDivide))
	mseti.Add(new(model.SchIcbcPosNum))
	mseti.Add(new(model.SchMarketPlan))
	mseti.Add(new(model.SchTaTargetplan))
	mseti.Add(new(model.SchTargetsignPlan))
	mseti.Add(new(model.School))
	mseti.Add(new(model.SchoolImgFile))
	mseti.Add(new(model.SchoolOfLimit))
	mseti.Add(new(model.ScoreItem))
	mseti.Add(new(model.SctHourDetail))
	mseti.Add(new(model.ServiceProVisitFile))
	mseti.Add(new(model.Sheng))
	mseti.Add(new(model.StuAllotTmk))
	mseti.Add(new(model.StuBfJl))
	mseti.Add(new(model.StuBussType))
	mseti.Add(new(model.StuClassContent))
	mseti.Add(new(model.StuClassFeedbacFile))
	mseti.Add(new(model.StuClassFeedback))
	mseti.Add(new(model.StuClassTypeAssBook))
	mseti.Add(new(model.StuClassTypeSusHour))
	mseti.Add(new(model.StuClassTypeTextbook))
	mseti.Add(new(model.StuConfirmHour))
	mseti.Add(new(model.StuCourseVideo))
	mseti.Add(new(model.StuCtTimeRange))
	mseti.Add(new(model.StuExchangeGift))
	mseti.Add(new(model.StuFailToLive))
	mseti.Add(new(model.StuFinishclass))
	mseti.Add(new(model.StuGwc))
	mseti.Add(new(model.StuGwcGef))
	mseti.Add(new(model.StuKwJfx))
	mseti.Add(new(model.StuLaostuGift))
	mseti.Add(new(model.StuLoss))
	mseti.Add(new(model.StuMonthYpHour))
	mseti.Add(new(model.StuParent))
	mseti.Add(new(model.StuPcHour))
	mseti.Add(new(model.StuProductType))
	mseti.Add(new(model.StuSatisfyVisit))
	mseti.Add(new(model.StuSatisfyVisitContent))
	mseti.Add(new(model.StuScore))
	mseti.Add(new(model.StuScoreItem))
	mseti.Add(new(model.StuToReopened))
	mseti.Add(new(model.StuTypeStudyCard))
	mseti.Add(new(model.StuUseTextbook))
	mseti.Add(new(model.StuUserDate))
	mseti.Add(new(model.StuWxdqDate))
	mseti.Add(new(model.Stucasting))
	mseti.Add(new(model.StuclasstypeExtrafee))
	mseti.Add(new(model.Student))
	mseti.Add(new(model.StudentAddAgreement))
	mseti.Add(new(model.StudentCastingTable))
	mseti.Add(new(model.StudentCc))
	mseti.Add(new(model.StudentClass))
	mseti.Add(new(model.StudentClassSchedule))
	mseti.Add(new(model.StudentClassType))
	mseti.Add(new(model.StudentClassTypeUseIntegral))
	mseti.Add(new(model.StudentClassYpk))
	mseti.Add(new(model.StudentClassYpk1))
	mseti.Add(new(model.StudentClassYpk2))
	mseti.Add(new(model.StudentClassYpk3))
	mseti.Add(new(model.StudentClassYpk31))
	mseti.Add(new(model.StudentClassYpk32))
	mseti.Add(new(model.StudentClassYpk33))
	mseti.Add(new(model.StudentClassYpkDkb))
	mseti.Add(new(model.StudentCustomer))
	mseti.Add(new(model.StudentCzIntegral))
	mseti.Add(new(model.StudentDeleteLog))
	mseti.Add(new(model.StudentGift))
	mseti.Add(new(model.StudentGrowthRecord))
	mseti.Add(new(model.StudentGrowthRecordFile))
	mseti.Add(new(model.StudentMarketActivity))
	mseti.Add(new(model.StudentOrgan))
	mseti.Add(new(model.StudentOrganSchool))
	mseti.Add(new(model.StudentReopened))
	mseti.Add(new(model.StudentRzIntegral))
	mseti.Add(new(model.StudentSale))
	mseti.Add(new(model.StudentTimeResource))
	mseti.Add(new(model.StudentTypeJournal))
	mseti.Add(new(model.StudentUseHourDetail))
	mseti.Add(new(model.StudentVisit))
	mseti.Add(new(model.StudentVisitLanguage))
	mseti.Add(new(model.StudyCard))
	mseti.Add(new(model.SystemMessage))
	mseti.Add(new(model.SystemNotice))
	mseti.Add(new(model.SystemOperaLogs))
	mseti.Add(new(model.SystemSet))
	mseti.Add(new(model.T0UserInfo))
	mseti.Add(new(model.T19LoginLog))
	mseti.Add(new(model.T2LoginLog))
	mseti.Add(new(model.T3Program))
	mseti.Add(new(model.T520LoginLog))
	mseti.Add(new(model.TaTargetPlan))
	mseti.Add(new(model.TeaLeaveManagement))
	mseti.Add(new(model.TeachLog))
	mseti.Add(new(model.TeacherHourDetail))
	mseti.Add(new(model.TeacherServiceProcess))
	mseti.Add(new(model.TeachingMaterial))
	mseti.Add(new(model.TeachingTime))
	mseti.Add(new(model.Temp20190330))
	mseti.Add(new(model.TempHuifu))
	mseti.Add(new(model.TempShaoerShow))
	mseti.Add(new(model.TempStorageCz))
	mseti.Add(new(model.TempStorageCzFlow))
	mseti.Add(new(model.TempStorageRz))
	mseti.Add(new(model.TempStu))
	mseti.Add(new(model.TempStuJifenAdd))
	mseti.Add(new(model.TestSites))
	mseti.Add(new(model.TextBook))
	mseti.Add(new(model.TextBookProductType))
	mseti.Add(new(model.TimeSet))
	mseti.Add(new(model.TimeTrigger))
	mseti.Add(new(model.TmLevel))
	mseti.Add(new(model.TmSi))
	mseti.Add(new(model.TmSiUnit))
	mseti.Add(new(model.TrainExternalPlace))
	mseti.Add(new(model.TrainOrgan))
	mseti.Add(new(model.TrainOrganSchool))
	mseti.Add(new(model.TrainerEvaluateInfo))
	mseti.Add(new(model.TrainerEvaluateParticipant))
	mseti.Add(new(model.TrainerTable))
	mseti.Add(new(model.TrialLesson))
	mseti.Add(new(model.TrialLessonProType))
	mseti.Add(new(model.TrialLessonStudent))
	mseti.Add(new(model.TrialLessonTea))
	mseti.Add(new(model.TrialLessonUser))
	mseti.Add(new(model.TscReforderFile))
	mseti.Add(new(model.Unit))
	mseti.Add(new(model.UserBussType))
	mseti.Add(new(model.UserCasting))
	mseti.Add(new(model.UserCourse))
	mseti.Add(new(model.UserDepartment))
	mseti.Add(new(model.UserGsSch))
	mseti.Add(new(model.UserMessageNotification))
	mseti.Add(new(model.UserOnoffRz))
	mseti.Add(new(model.UserOrgan))
	mseti.Add(new(model.UserOrganSchool))
	mseti.Add(new(model.UserProductType))
	mseti.Add(new(model.UserTimeResource))
	mseti.Add(new(model.UserTrainingTable))
	mseti.Add(new(model.UserTypeTable))
	mseti.Add(new(model.UserUploadFile))
	mseti.Add(new(model.ViolateRatioManage))
	mseti.Add(new(model.WastageFollowupRecord))
	mseti.Add(new(model.WechatMessage))
	mseti.Add(new(model.WholeCountry))
	mseti.Add(new(model.XlHc))
	mseti.Add(new(model.XyHourRecord))
	mseti.Add(new(model.YbPayFlowing))
	mseti.Add(new(model.YbPayOut))
	mseti.Add(new(model.YbzfPos))
	mseti.Add(new(model.YeepayReconciliationsRecord))
	mseti.Add(new(model.ZbxGiftrules))
	mseti.Add(new(model.ZbxzsgzValiddate))
	mseti.Add(new(model.ZbxzsyxCity))
	return nil
}

// SyncController defined
func SyncController() error {
	OrganRoutes(App)
	return nil
}

// SyncService defined
func SyncService() error {
	return nil
}

// Executor defined
var Executor = util.NewExecutor(SyncModel, SyncController, SyncService)
