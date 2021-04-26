package ecode

// 之后的错误码规则为6位，如下
// 00 0000
// 前2位为错误码前缀，用于分类，从10开始累加
// 后4位为错误码主体，用于实例，从0001开始累加

// 普通错误码前缀从 10，开始，且须注释说明划分的服务，如：
//     100001     同级部门已存在相同名字（前缀10表示user服务）
//     110001     首页门户不可为空（前缀11表示new_portal服务）

var (
	ErrInvalidParam = New(100000, "参数错误")
	ErrNotFound     = New(100001, "暂无数据")
	ErrQuery        = New(100002, "查询失败")
	ErrInsert       = New(100003, "插入失败")
	ErrUpdate       = New(100004, "更新失败")
	ErrDelete       = New(100005, "删除失败")

	// user
	ErrDepartmentNameExist = New(110001, "同级部门已存在相同名字")

	//reward
	ErrPropNameTooLong = New(120001, "道具名称过长")
	ErrConfigureReward = New(120002, "打赏配置出错")
)
