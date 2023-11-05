package domain

type ServiceRegistry struct {
	ClassServ           ClassService
	StudentClassService StudentClassService
}

type ViewRegistry struct {
	ClassVw        ClassView
	StudentClassVw StudentClassView
}