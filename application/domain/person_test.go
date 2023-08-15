package domain

import (
	"errors"
	"testing"
)

func TestCreatePerson(t *testing.T) {
	t.Run("Given valid data when try to create person should be ok", func(t *testing.T) {
		expectedNickName := "Zé"
		expectedName := "Jose Maria"
		expectedBirthDate := "1980-05-27"
		person, err := CreatePerson(expectedNickName, expectedName, expectedBirthDate)
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		if person.Id == "" {
			t.Errorf("expected id but not found")
		}
		if person.NickName != expectedNickName {
			t.Errorf("Unexpected value. Expected %v but found %v", expectedNickName, person.NickName)
		}
		if person.Name != expectedName {
			t.Errorf("Unexpected value. Expected %v but found %v", expectedName, person.Name)
		}
		if person.BirthDate.Format("2006-01-02") != expectedBirthDate {
			t.Errorf("Unexpected value. Expected %v but found %v", expectedBirthDate, person.BirthDate)
		}
	})

	t.Run("Given valid data with stack when try to create person should be ok", func(t *testing.T) {
		expectedNickName := "Zé"
		expectedName := "Jose Maria"
		expectedBirthDate := "1980-05-27"
		goLang, _ := NewStack("GoLang")
		csharp, _ := NewStack("C#")
		stackList := StackList{}
		stackList.AddStack(*goLang)
		stackList.AddStack(*csharp)
		person, err := CreatePerson(expectedNickName, expectedName, expectedBirthDate)
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		person.AddStackList(stackList.GetStacks())
		if len(person.StackList) != len(stackList.GetStacks()) {
			t.Errorf("Unexpected value. Expected %v but found %v", len(stackList.GetStacks()), len(person.StackList))
		}
	})

	t.Run("Given an invalid nickname when try to create person should return an error", func(t *testing.T) {
		expectedNickName := ""
		expectedName := "Jose Maria"
		expectedBirthDate := "1980-05-27"
		expectedError := errors.New("nickname is null")
		_, err := CreatePerson(expectedNickName, expectedName, expectedBirthDate)
		if err == nil {
			t.Errorf("an error was expected")
		}
		if err.Error() != expectedError.Error() {
			t.Errorf("value error. Expected %v, but received %v.", expectedError, err.Error())
		}
	})

	t.Run("Given an invalid nickname greater than 32 chars when try to create person should return an error", func(t *testing.T) {
		expectedNickName := "Este tipo é o Zé maria que mora na rua 102 e fuma crack"
		expectedName := "Jose Maria"
		expectedBirthDate := "1980-05-27"
		expectedError := errors.New("nickname is greater than 32 chars")
		_, err := CreatePerson(expectedNickName, expectedName, expectedBirthDate)
		if err == nil {
			t.Errorf("an error was expected")
		}
		if err.Error() != expectedError.Error() {
			t.Errorf("value error. Expected %v, but received %v.", expectedError, err.Error())
		}

	})

	t.Run("Given an invalid null name when try to create person should return an error", func(t *testing.T) {
		expectedError := errors.New("name is null")
		expectedNickName := "Zé"
		expectedName := ""
		expectedBirthDate := "1980-05-27"
		_, err := CreatePerson(expectedNickName, expectedName, expectedBirthDate)
		if err == nil {
			t.Errorf("an error was expected")
		}
		if err.Error() != expectedError.Error() {
			t.Errorf("value error. Expected %v, but received %v.", expectedError, err.Error())
		}
	})

	t.Run("Given an invalid name with more than 100 chars when try to create person should return an error", func(t *testing.T) {
		expectedError := errors.New("name is greater than 100 chars")
		expectedNickName := "Zé"
		expectedName := "José Maria Trindade dos Anjos Guerreiros de São Jorge da Capadócia do Norte da Italia na Epoca do Facismo"
		expectedBirthDate := "1980-05-27"
		_, err := CreatePerson(expectedNickName, expectedName, expectedBirthDate)
		if err == nil {
			t.Errorf("an error was expected")
		}
		if err.Error() != expectedError.Error() {
			t.Errorf("value error. Expected %v, but received %v.", expectedError, err.Error())
		}
	})

	t.Run("Given an invalid birthdate formate when try to create person should return an error", func(t *testing.T) {
		expectedNickName := "Zé"
		expectedName := "Jose Maria"
		expectedBirthDate := "10/05/10934"
		expectedError := errors.New("invalid birth date")
		_, err := CreatePerson(expectedNickName, expectedName, expectedBirthDate)
		if err == nil {
			t.Errorf("an error was expected")
		}
		if err.Error() != expectedError.Error() {
			t.Errorf("value error. Expected %v, but received %v.", expectedError, err.Error())
		}
	})
}
