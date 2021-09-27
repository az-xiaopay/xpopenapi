package xpnotify

const NotifyType_ConsumeSuccess = "ConsumeSuccess"
const NotifyType_ReductionFail = "ReductionFail"
const NotifyType_Reacquire = "Reacquire"
const NotifyType_Datachange = "DataChange"

const (
	ChangeType_create_grade        = "create_grade"        //新增年级事件
	ChangeType_update_grade        = "update_grade"        //更新年级事件
	ChangeType_delete_grade        = "delete_grade"        //删除年级事件
	ChangeType_create_class        = "create_class"        //新增班级事件
	ChangeType_update_class        = "update_class"        //更新班级事件
	ChangeType_delete_class        = "delete_class"        //删除班级事件
	ChangeType_create_teacher_user = "create_teacher_user" //新增教职工事件
	ChangeType_update_teacher_user = "update_teacher_user" //更新教职工事件
	ChangeType_delete_teacher_user = "delete_teacher_user" //删除教职工事件
	ChangeType_create_student_user = "create_student_user" //新增学生事件
	ChangeType_update_student_user = "update_student_user" //更新学生事件
	ChangeType_delete_student_user = "delete_student_user" //删除学生事件
)
