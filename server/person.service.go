package main

// ListPeopleService service
func ListPeopleService() []Person {
	var people []Person
	DB.Limit(100).Find(&people)
	return people
}
