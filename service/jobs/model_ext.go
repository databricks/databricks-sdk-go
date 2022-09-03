package jobs

func (r *Run) FirstTask() *RunTask {
	if len(r.Tasks) == 0 {
		return nil
	}
	return &r.Tasks[0]
}
