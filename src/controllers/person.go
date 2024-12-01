package controllers

import (
	"bff/src/models/out"
	"bff/src/services"
	"fmt"
	"log/slog"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetDescription(ctx *fiber.Ctx) error {
	queries := ctx.Queries()
	name := queries["name"]

	person, err := services.GetPersonByName(name)
	if err != nil {
		slog.Error("Problema ao realizar consulta")
		return err
	}

	slog.Info("Consulta Realizada")

	return ctx.JSON(provideDescriprion(person))
}

func provideDescriprion(person out.Person) out.Description {
	var films []string
	for _, film := range person.Films {
		films = append(films, film.Name)
	}
	filmsVerbose := strings.Join(films, ", ")
	desc := fmt.Sprintf("O persoagem %s, é um dos habitantes do planeta %s e participou dos filmes: %s", person.Name, person.Homeworld, filmsVerbose)
	var description out.Description
	description.Description = desc;
	return description
}
