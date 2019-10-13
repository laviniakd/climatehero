package backend_src

import (
	"climatehero/db"
	"climatehero/models"
)

func hasUser(email string) bool {
	return db.HasUser(email)
}

func AddUser(email, pin, name string) *models.Error {
	if hasUser := hasUser(email); hasUser {
		code := int32(400)
		msg := "User already exists"
		return &models.Error{&code, &msg}
	}

	if err := db.AddUser(email, pin, name); err != nil {
		code := int32(500)
		msg := "Could not add user"
		return &models.Error{&code, &msg}
	}

	return nil
}

func CheckItem(email, item string) (*int, *string, *models.Error) {
	if hasUser := hasUser(email); !hasUser {
		code := int32(500)
		msg := "User does not exist"
		return nil, nil, &models.Error{&code, &msg}
	}

	// update the list table, setting this item to checked
	if err := db.CheckItem(email, item); err != nil {
		code := int32(500)
		msg := "Database update error"
		return nil, nil, &models.Error{&code, &msg}
	}

	// update the last checked table
	if err := db.UpdateLastChecked(email); err != nil {
		code := int32(500)
		msg := "Database update error"
		return nil, nil, &models.Error{&code, &msg}
	}

	// update the points
	points, err := db.UpdatePoints(email)
	if err != nil {
		code := int32(500)
		msg := "Database update error"
		return nil, nil, &models.Error{&code, &msg}
	}

	checkMessage := "Congratulations, you've earned 10 points! Your polar bears are thriving!"

	return &points, &checkMessage, nil
}

func GetUserInfo(email string) (*int, []string, *string, *models.Error) {
	if hasUser := hasUser(email); !hasUser {
		code := int32(500)
		msg := "User does not exist"
		return nil, nil, nil, &models.Error{&code, &msg}
	}

	points, checkList, message, err := db.GetUserInfo(email)
	if err != nil {
		code := int32(500)
		msg := "Database error"
		return nil, nil, nil, &models.Error{&code, &msg}
	}

	return &points, checkList, &message, nil
}

func Login(email, pin string) (*string, *models.Error) {
	if hasUser := hasUser(email); !hasUser {
		code := int32(500)
		msg := "User does not exist"
		return nil, &models.Error{&code, &msg}
	}

	name, err := db.Login(email, pin)
	if err != nil {
		code := int32(500)
		msg := "Database error; possible invalid password"
		return nil, &models.Error{&code, &msg}
	}

	return name, nil
}

func UpdateList(email string, list []string) *models.Error {
	if hasUser := hasUser(email); !hasUser {
		code := int32(500)
		msg := "User does not exist"
		return &models.Error{&code, &msg}
	}

	return db.UpdateList(email, list)
}
