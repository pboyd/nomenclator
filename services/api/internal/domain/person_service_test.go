package domain

import (
	"context"
	"testing"

	"github.com/pboyd/nomenclator/api/internal/database"
	"github.com/stretchr/testify/assert"
)

func TestCreatePeople(t *testing.T) {
	cases := []struct {
		Person *Person
		Errors map[string]string
	}{
		{
			Person: &Person{
				Prefix:     "Mr",
				FirstName:  "John",
				MiddleName: "Q",
				LastName:   "Doe",
				Suffix:     "Jr",
			},
			Errors: nil,
		},
		{
			Person: &Person{
				FirstName:  "John",
				MiddleName: "Q",
				LastName:   "Doe",
			},
			Errors: nil,
		},
		{
			Person: &Person{
				Prefix:     "INVALID",
				FirstName:  "John",
				MiddleName: "Q",
				LastName:   "Doe",
				Suffix:     "INVALID",
			},
			Errors: map[string]string{
				"prefix": "invalid prefix",
				"suffix": "invalid suffix",
			},
		},
	}

	db := database.TestDB(t)
	assert := assert.New(t)

	personService := NewPersonService(database.New(db))

	for _, c := range cases {
		err := personService.Create(context.Background(), c.Person)
		if c.Errors == nil {
			if !assert.NoError(err) {
				continue
			}
		} else {
			validationErr, ok := err.(ErrValidationFailed)
			if !assert.True(ok) {
				continue
			}

			for k, v := range c.Errors {
				assert.Equal(v, validationErr[k])
			}
			continue
		}

		assert.NotZero(c.Person.ID)
		assert.NotZero(c.Person.CreatedAt)
		assert.NotZero(c.Person.UpdatedAt)
	}
}

func TestLoadPerson(t *testing.T) {
	db := database.TestDB(t, "minimal")
	assert := assert.New(t)

	personService := NewPersonService(database.New(db))

	person := &Person{
		Prefix:    "Ms",
		FirstName: "Jane",
		LastName:  "Jones",
	}
	assert.NoError(personService.Create(context.Background(), person))

	loadedPerson, err := personService.Load(context.Background(), person.ID)
	assert.NoError(err)
	assert.Equal(person, loadedPerson)

	loadedPerson, err = personService.Load(context.Background(), 0)
	assert.NoError(err)
	assert.Nil(loadedPerson)
}
