package enums

//自定义枚举
type ResultStatus int

const (
	OK       ResultStatus = 200
	NotFound              = 404
	Error                 = 500
	NoLogin               = 401
	JumpPage              = 302
)

type BackEndStatus int

const (
	Deleted   BackEndStatus = 1
	NoDeleted               = 0
	Active                  = 1
	NoActive                = 0
)
