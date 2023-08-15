package tests

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"

	"github.com/AllanCordeiro/person-st/infra/database"
	"github.com/AllanCordeiro/person-st/infra/webserver/handlers"
)

type RequestInput struct {
	Nickname  string   `json:"apelido"`
	Name      string   `json:"nome"`
	BirthDate string   `json:"nascimento"`
	Stack     []string `json:"stack"`
}

type RequestOutput struct {
	ID        string   `json:"id"`
	NickName  string   `json:"apelido"`
	Name      string   `json:"nome"`
	BirthDate string   `json:"nascimento"`
	StackList []string `json:"stack"`
}

type OutputList struct {
	id       string
	nickName string
}

var endpointList []OutputList

func TestCreatePerson(t *testing.T) {
	tests := []struct {
		name          string
		input         RequestInput
		statusCode    int
		checkLocation bool
	}{
		{
			name: "Given valid person when call post should return status 201",
			input: RequestInput{
				Nickname:  "josé",
				Name:      "José Roberto",
				BirthDate: "2000-10-01",
				Stack:     []string{"C#", "Node", "Oracle"},
			},
			statusCode:    http.StatusCreated,
			checkLocation: true,
		},
		{
			name: "Given valid person without stack when call post should return status 201",
			input: RequestInput{
				Nickname:  "ana",
				Name:      "Ana Barbosa",
				BirthDate: "1985-09-23",
				Stack:     nil,
			},
			statusCode:    http.StatusCreated,
			checkLocation: true,
		},
		{
			name: "Given valid person without stack when call post should return status 201",
			input: RequestInput{
				Nickname:  "aninha",
				Name:      "Ana Barbosa",
				BirthDate: "1985-09-23",
				Stack:     []string{"Node", "Postgres"},
			},
			statusCode:    http.StatusCreated,
			checkLocation: true,
		},
		{
			name: "Given valid person with nickname already stored when call post should return status 422",
			input: RequestInput{
				Nickname:  "aninha",
				Name:      "Ana Barbosa",
				BirthDate: "1985-09-23",
				Stack:     []string{"Node", "Postgres"},
			},
			statusCode:    http.StatusUnprocessableEntity,
			checkLocation: false,
		},
		{
			name: "Given person without name when call post should return status 422",
			input: RequestInput{
				Nickname:  "ana",
				BirthDate: "1985-09-23",
				Stack:     []string{"Node", "Postgres"},
			},
			statusCode:    http.StatusUnprocessableEntity,
			checkLocation: false,
		},
		{
			name: "Given person without nickname when call post should return status 422",
			input: RequestInput{
				Name:      "Ana Barbosa",
				BirthDate: "1985-09-23",
				Stack:     []string{"Node", "Postgres"},
			},
			statusCode:    http.StatusUnprocessableEntity,
			checkLocation: false,
		},
	}
	db, err := dbConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	personDB := database.NewPersonDB(db)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body bytes.Buffer
			_ = json.NewEncoder(&body).Encode(tt.input)

			req, err := http.NewRequest(http.MethodPost, "/pessoas", &body)
			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}

			w := httptest.NewRecorder()
			personHander := handlers.NewPersonHandler(personDB)
			personHander.CreatePerson(w, req)
			res := w.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.statusCode {
				t.Errorf("unexpected status code. Expected %v, but received %v", tt.statusCode, res.StatusCode)
			}

			if tt.checkLocation {
				if res.Header.Get("Location") == "" {
					t.Errorf("expected header location, but not found")
				}
				endpointList = append(endpointList, OutputList{id: res.Header.Get("Location"), nickName: tt.input.Nickname})
			}
		})
	}
}

func TestGetPerson(t *testing.T) {
	tests := []struct {
		name       string
		input      RequestInput
		endpoint   string
		statusCode int
	}{
		{
			name: "Given a valid ID when call get person/id should return its data",
			input: RequestInput{
				Nickname:  "josé",
				Name:      "José Roberto",
				BirthDate: "2000-10-01",
				Stack:     []string{"C#", "Node", "Oracle"},
			},
			endpoint:   getEndpoint("josé"),
			statusCode: http.StatusOK,
		},
		{
			name: "Given a valid ID with no stack when call get person/id should return its data",
			input: RequestInput{
				Nickname:  "ana",
				Name:      "Ana Barbosa",
				BirthDate: "1985-09-23",
				Stack:     nil,
			},
			endpoint:   getEndpoint("ana"),
			statusCode: http.StatusOK,
		},
		{
			name:       "Given an invalid ID when call get person/id should return error 404",
			input:      RequestInput{},
			endpoint:   "/pessoas/an-invalid-id",
			statusCode: http.StatusNotFound,
		},
	}

	db, err := dbConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	personDB := database.NewPersonDB(db)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output RequestOutput
			personId := strings.TrimLeft(tt.endpoint, "/pessoas/")
			req, err := http.NewRequest(http.MethodGet, "/pessoas/{personID}", nil)
			rchi := chi.NewRouteContext()
			rchi.URLParams.Add("personID", personId)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rchi))

			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}

			w := httptest.NewRecorder()
			receivedStatusCode := w.Code
			personHander := handlers.NewPersonHandler(personDB)
			personHander.GetPersonById(w, req)
			if receivedStatusCode != tt.statusCode {
				t.Errorf("Unexpected status code. Expected %v but received %v", tt.statusCode, receivedStatusCode)
			}
			if receivedStatusCode == http.StatusOK {
				err = json.Unmarshal(w.Body.Bytes(), &output)
				if err != nil {
					t.Errorf("Unexpected error %v", err)
				}

				if output.NickName != tt.input.Nickname {
					t.Errorf("Unexpected value. Expected %v but received %v", tt.input.Nickname, output.NickName)
				}
				if output.Name != tt.input.Name {
					t.Errorf("Unexpected value. Expected %v but received %v", tt.input.Name, output.Name)
				}
				if output.BirthDate != tt.input.BirthDate {
					t.Errorf("Unexpected value. Expected %v but received %v", tt.input.BirthDate, output.BirthDate)
				}
				if len(output.StackList) != len(tt.input.Stack) {
					if output.NickName != tt.input.Nickname {
						t.Errorf("Unexpected value. Expected %v but received %v", len(tt.input.Stack), len(output.StackList))
					}
				}
			}
		})
	}
}

func getEndpoint(nickname string) string {
	for _, data := range endpointList {
		if data.nickName == nickname {
			return data.id
		}
	}
	return ""
}

func dbConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://rinha:rinha123@localhost/rinhadb?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
