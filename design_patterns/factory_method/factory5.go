package main

import "fmt"

type CloudStorage interface {
	Upload(file string)
}

type AWS struct {
	Region string
}

type Google struct {
	ProjectID string
}

func (a *AWS) Upload(file string) {
	fmt.Printf("Send to Aws service (region %s)...\n", a.Region)
}

func (g *Google) Upload(file string) {
	fmt.Printf("Send to Google service (region %s)...\n", g.ProjectID)
}

func GetStorage(typ string) CloudStorage {
	switch typ {
	case "aws":
		return &AWS{Region: "us-east-1"}
	case "google":
		return &Google{ProjectID: "gcp-prod-001"}
	}

	return nil
}

func main() {
	service := GetStorage("aws")
	service.Upload("America")
}
