package main

// ListCompaniesService service
func ListCompaniesService() []Company {
	var companies []Company
	DB.Find(&companies)
	return companies
}
