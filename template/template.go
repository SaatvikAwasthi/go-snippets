package template

import (
	"fmt"
	"go-snips/models"
	"html/template"
	"io"
	"log"
)

type templates struct {
	greetTemplate *template.Template
}

type Template interface {
	GetHTML(message string, writer io.Writer) error
}

func New() Template {
	templateString := "<B><Font color='blue' size='16'> {{ .Message }} </B>"
	greetTemplate := template.New("GreetingTemplate")
	tmp, err := greetTemplate.Parse(templateString)
	if err != nil {
		log.Fatalf("error occurred in generating template. %e", err)
	}

	return &templates{
		greetTemplate: tmp,
	}
}

func (t *templates) GetHTML(message string, writer io.Writer) error {
	err := t.greetTemplate.Execute(writer, models.Greeting{
		Message: message,
	})
	if err != nil {
		return fmt.Errorf("error occurred in generating template. %w", err)
	}
	return nil
}
