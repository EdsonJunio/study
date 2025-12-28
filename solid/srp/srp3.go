package srp

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type (
	Citizen struct {
		Name string
		Age  int
	}

	CSVLoader struct{}

	CitizenValidator struct{}

	GovAPIClient struct {
		Endpoint string
	}

	ETLProcessor struct {
		Loader    *CSVLoader
		Validator *CitizenValidator
		Client    *GovAPIClient
	}
)

func (*CSVLoader) Load(filename string) ([]Citizen, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	var citizens []Citizen
	for _, rec := range records {
		if len(rec) < 2 {
			continue
		}
		age, _ := strconv.Atoi(rec[1])
		citizens = append(citizens, Citizen{
			Name: strings.TrimSpace(rec[0]),
			Age:  age,
		})
	}

	return citizens, nil
}

func (*CitizenValidator) Validate(c Citizen) bool {
	return c.Age >= 18
}

func (g *GovAPIClient) Send(c Citizen) error {
	body, _ := json.Marshal(c)
	_, err := http.Post(g.Endpoint, "application/json", strings.NewReader(string(body)))
	return err
}

func (p *ETLProcessor) Run(filename string) error {
	citizens, err := p.Loader.Load(filename)
	if err != nil {
		return err
	}

	for _, c := range citizens {
		if !p.Validator.Validate(c) {
			continue
		}
		if err := p.Client.Send(c); err != nil {
			return err
		}
	}

	return nil
}
