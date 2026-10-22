package task

type TaskStatus string

const (
    StatusCreated    TaskStatus = "created"
    StatusInProgress TaskStatus = "in_progress"
    StatusCancelled  TaskStatus = "cancelled"
    StatusCompleted  TaskStatus = "completed"
)