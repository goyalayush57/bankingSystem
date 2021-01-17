package controller

import "src/src/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Employee routes
	s.Router.HandleFunc("/employees", middlewares.SetMiddlewareJSON(middlewares.SetAdminMiddlewareAuthentication(s.CreateEmployee))).Methods("POST")
	s.Router.HandleFunc("/employees", middlewares.SetMiddlewareJSON(middlewares.SetAdminMiddlewareAuthentication(s.GetEmployees))).Methods("GET")
	s.Router.HandleFunc("/employees/{id}", middlewares.SetMiddlewareJSON(middlewares.SetAdminMiddlewareAuthentication(s.FindEmployeeById))).Methods("GET")
	s.Router.HandleFunc("/employees/{id}", middlewares.SetMiddlewareJSON(middlewares.SetAdminMiddlewareAuthentication(s.UpdateEmployee))).Methods("PUT")
	s.Router.HandleFunc("/employees/{id}", middlewares.SetMiddlewareJSON(middlewares.SetAdminMiddlewareAuthentication(s.DeleteEmployee))).Methods("DELETE")
	s.Router.HandleFunc("/employees/{id}/kycs", middlewares.SetMiddlewareJSON(middlewares.SetAdminMiddlewareAuthentication(s.LinkEmployeeKYC))).Methods("POST")
	s.Router.HandleFunc("/employees/{id}/kycs/{kyc_id}", middlewares.SetMiddlewareJSON(middlewares.SetAdminMiddlewareAuthentication(s.UnLinkEmployeeKYC))).Methods("UNLINK")
	s.Router.HandleFunc("/employees/{id}/accounts", middlewares.SetMiddlewareJSON(middlewares.SetAdminMiddlewareAuthentication(s.LinkEmployeeAccount))).Methods("POST")
	s.Router.HandleFunc("/employees/{id}/accounts/{acc_id}", middlewares.SetMiddlewareJSON(middlewares.SetAdminMiddlewareAuthentication(s.UnLinkEmployeeAccount))).Methods("UNLINK")

	//Customer routes
	s.Router.HandleFunc("/customers", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.CreateCustomer))).Methods("POST")
	s.Router.HandleFunc("/customers", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetCustomers))).Methods("GET")
	s.Router.HandleFunc("/customers/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.FindCustomerById))).Methods("GET")
	s.Router.HandleFunc("/customers/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateCustomer))).Methods("PUT")
	s.Router.HandleFunc("/customers/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.DeleteCustomer))).Methods("DELETE")
	s.Router.HandleFunc("/customers/{id}/kycs", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.LinkCustomerKYC))).Methods("POST")
	s.Router.HandleFunc("/customers/{id}/kycs/{kyc_id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UnLinkCustomerKYC))).Methods("UNLINK")
	s.Router.HandleFunc("/customers/{id}/accounts", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.LinkCustomerAccount))).Methods("POST")
	s.Router.HandleFunc("/customers/{id}/accounts/{acc_id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UnLinkCustomerAccount))).Methods("UNLINK")
	// 	//Users routes
	// 	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	// 	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	// 	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	// 	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	// 	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	// 	//Posts routes
	// 	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	// 	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	// 	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	// 	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	// 	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")
}
