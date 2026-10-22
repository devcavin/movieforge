package task

type Task struct {
    id          int
    title       string
    description string
    status      TaskStatus
    priority    TaskPriority
}