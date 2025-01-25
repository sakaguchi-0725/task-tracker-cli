package model

type Status string

const (
	NotStarted Status = "not-started"
	InProgress Status = "in-progress"
	Done       Status = "done"
)

func NewStatus(s string) (Status, error) {
	if s == NotStarted.String() || s == InProgress.String() || s == Done.String() {
		return Status(s), nil
	}

	return "", nil
}

func GenerateStatus() Status {
	return NotStarted
}

func (s Status) String() string {
	return string(s)
}
