package internal

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func SaveUser(user *User, fileLink string) error {
	secret, ok := os.LookupEnv("credentials.json")
	if !ok {
		return fmt.Errorf("unable to read client secret")
	}

	ctx := context.Background()
	sheetsService, err := sheets.NewService(ctx, option.WithCredentialsJSON([]byte(secret)))
	if err != nil {
		return fmt.Errorf("failed to create service with %v", err)
	}

	spreadsheetId := "1fTYLgGZG8HJ25HtVg1y7as7EVG7eB347DZM6sVzw5MM"

	row := &sheets.ValueRange{
		Values: [][]interface{}{{user.Name, user.Surname, user.Email, user.Phone, fileLink}},
	}

	response, err := sheetsService.Spreadsheets.Values.Append(spreadsheetId, "Registrants", row).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Context(ctx).Do()
	if err != nil || response.HTTPStatusCode != 200 {
		return fmt.Errorf("failed to write user with %v", err)
	}

	return nil
}
