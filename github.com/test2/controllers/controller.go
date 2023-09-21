package controller

// func RegistrationController(w http.ResponseWriter, r *http.Request) error {
// 	var reg handlers.Registration

// 	err := json.NewDecoder(r.Body).Decode(&reg)
// 	if err != nil {
// 		return err
// 	}

// 	if reg.Name == "" || reg.SecondName == "" || reg.Email == "" || reg.Password == "" {
// 		return err
// 	}

// 	err = repository.RegistrtionRepository(reg)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
