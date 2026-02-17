package srp

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type (
	LogEntry struct {
		Message    string
		CreditCard string
	}

	DataMasker struct{}

	JsonFormatter struct{}

	FileLogger struct {
		FilePath string
	}

	LogService struct {
		Masker    *DataMasker
		Formatter *JsonFormatter
		Writer    *FileLogger
	}
)

func (DataMasker) MaskCreditCard(card string) string {
	if len(card) <= 4 {
		return card
	}
	return strings.Repeat("*", len(card)-4) + card[len(card)-4:]
}

func (JsonFormatter) ToJSON(entry LogEntry) ([]byte, error) {
	return json.Marshal(entry)
}

func (f *FileLogger) Save(data []byte) error {
	file, err := os.OpenFile(f.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	_, err = file.Write([]byte("\n"))
	return err
}

func (l *LogService) Logss(entry LogEntry) error {
	entry.CreditCard = l.Masker.MaskCreditCard(entry.CreditCard)

	jsonData, err := l.Formatter.ToJSON(entry)
	if err != nil {
		return err
	}

	err = l.Writer.Save(jsonData)
	if err != nil {
		return err
	}

	fmt.Println("Log salvo!")
	return nil
}

func main() {
	service := LogService{
		Masker:    &DataMasker{},
		Formatter: &JsonFormatter{},
		Writer:    &FileLogger{FilePath: "system.log"},
	}

	entry := LogEntry{
		Message:    "Pagamento processado",
		CreditCard: "1234567812345678",
	}

	service.Logss(entry)
}
