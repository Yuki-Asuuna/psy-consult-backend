package database

type Binding struct {
	BindingID    int64
	SupervisorID string
	CounsellorID string
}

func (Binding) TableName() string {
	return "binding"
}
