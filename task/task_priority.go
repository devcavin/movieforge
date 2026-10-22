package task

type TaskPriority string

const (
    PriorityHigh     TaskPriority = "high"
    PriorityModerate TaskPriority = "moderate"
    PriorityLow      TaskPriority = "low"
)