package api

type Task interface {
	Run()
	CronString() string
	Name() string
}
