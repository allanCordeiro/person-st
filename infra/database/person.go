package database

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/AllanCordeiro/person-st/application/domain"
)

type PersonDB struct {
	DB *sql.DB
}

func NewPersonDB(db *sql.DB) *PersonDB {
	return &PersonDB{DB: db}
}

func (p *PersonDB) Save(person domain.Person) error {
	stmt, err := p.DB.Prepare(`INSERT INTO rinha.person (id, nickname, name, birth_date, stack)VALUES ($1, $2, $3, $4, $5)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	stacks, _ := json.Marshal(person.StackList)
	_, err = stmt.Exec(person.Id, person.NickName, person.Name, person.BirthDate, stacks)
	if err != nil {
		return err
	}
	return nil
}

func (p *PersonDB) GetByID(id string) (*domain.Person, error) {
	person := &domain.Person{}
	stackList := &domain.StackList{}
	stackJson := []byte{}
	stmt, err := p.DB.Prepare("SELECT id, nickname, name, birth_date, stack FROM rinha.person WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&person.Id,
		&person.NickName,
		&person.Name,
		&person.BirthDate,
		&stackJson,
	)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(stackJson, &stackList.Stacks)
	if err != nil {
		return nil, err
	}

	person.AddStackList(stackList.GetStacks())

	return person, nil
}

func (p *PersonDB) GetByTerms(term string) ([]domain.Person, error) {
	var personList []domain.Person
	rows, err := p.DB.Query("SELECT id, nickname, name, birth_date, stack FROM rinha.person WHERE full_search @@ PLAINTO_TSQUERY($1) LIMIT 50", term)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		person := &domain.Person{}
		stackList := &domain.StackList{}
		stackJson := []byte{}
		err = rows.Scan(&person.Id, &person.NickName, &person.Name, &person.BirthDate, &stackJson)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(stackJson, &stackList.Stacks)
		if err != nil {
			return nil, err
		}
		person.AddStackList(stackList.GetStacks())
		personList = append(personList, *person)
	}
	return personList, nil

	// if len(personList) > 0 {
	// 	return personList, nil
	// }
	// return p.fallback(term)

}

func (p *PersonDB) fallback(term string) ([]domain.Person, error) {
	log.Println("entrei aqui")
	var personList []domain.Person
	rows, err := p.DB.Query(`SELECT id, nickname, name, birth_date, stack FROM rinha.person WHERE
	 name ILIKE '%' || $1 || '%'`, term)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		person := &domain.Person{}
		stackList := &domain.StackList{}
		stackJson := []byte{}
		err = rows.Scan(&person.Id, &person.NickName, &person.Name, &person.BirthDate, &stackJson)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(stackJson, &stackList.Stacks)
		if err != nil {
			return nil, err
		}
		person.AddStackList(stackList.GetStacks())
		personList = append(personList, *person)
	}
	return personList, nil
}

func (p *PersonDB) GetTotal() (*int64, error) {
	var total int64
	stmt, err := p.DB.Prepare("SELECT count(id) FROM rinha.person")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow().Scan(&total)
	if err != nil {
		return nil, err
	}
	return &total, nil
}
