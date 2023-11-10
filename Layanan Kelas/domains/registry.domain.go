package domain

type ServiceRegistry struct {
	ClassServ           ClassService
	ClassSubjectServ    ClassSubjectService
	StudentClassService StudentClassService
}

type ViewRegistry struct {
	ClassVw        ClassView
	StudentClassVw StudentClassView
	ClassSubjectVw ClassSubjectView
}