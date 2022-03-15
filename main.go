package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	faker "github.com/bxcodec/faker/v3"
)

type OktaUser struct {
	Login               string
	FirstName           string `faker:"first_name"`
	MiddleName          string `faker:"first_name"`
	LastName            string `faker:"last_name"`
	EmailDomain         string `faker:"domain_name"`
	Email               string
	LoyaltyIdTnf        string `faker:"uuid_digit"`
	LoyaltyIdVans       string `faker:"uuid_digit"`
	LoyaltyIdTimberland string `faker:"uuid_digit"`
	LoyaltyIdDickies    string `faker:"uuid_digit"`
}

func main() {
	err := writeRandomUsersCsv()
	if err != nil {
		fmt.Println(err)
	}
}

func writeRandomUsersCsv() error {
	limit := 100000
	fileName := "users.csv"
	csvFile, err := os.Create(fileName)
	defer csvFile.Close()

	if err != nil {
		fmt.Printf("failed creating file: %s", err)
	}

	w := csv.NewWriter(csvFile)

	header := []string{
		"userName",
		"login",
		"firstName",
		"middleName",
		"lastName",
		"email",
		"loyalty_id_tnf",
		"loyalty_id_vans",
		"loyalty_id_timberland",
		"loyalty_id_dickies",
	}
	w.Write(header)

	for i := 0; i < limit; i++ {
		var user OktaUser
		err := faker.FakeData(&user)

		if err != nil {
			return err
		}

		emailAddress := fmt.Sprintf("%s.%s.%s@%s", user.FirstName, user.MiddleName, user.LastName, user.EmailDomain)
		emailAddress = strings.Replace(emailAddress, "'", "", -1)
		user.Email = emailAddress
		user.Login = emailAddress

		row := []string{
			user.Login,
			user.Login,
			user.FirstName,
			user.MiddleName,
			user.LastName,
			user.Email,
			user.LoyaltyIdTnf,
			user.LoyaltyIdVans,
			user.LoyaltyIdTimberland,
			user.LoyaltyIdDickies,
		}
		err = w.Write(row)

		if err != nil {
			return err
		}
	}
	return nil
}
