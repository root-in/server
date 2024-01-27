package internal

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const spreadsheetId = "1fTYLgGZG8HJ25HtVg1y7as7EVG7eB347DZM6sVzw5MM"
const spreadsheetName = "rootin"

func saveUser(user *User, fileName string) error {
	secret, ok := os.LookupEnv("credentials.json")
	if !ok {
		return fmt.Errorf("unable to read client secret")
	}

	ctx := context.Background()
	sheetsService, err := sheets.NewService(ctx, option.WithCredentialsJSON([]byte(secret)))
	if err != nil {
		return fmt.Errorf("failed to create service with %v", err)
	}

	row := &sheets.ValueRange{
		Values: [][]interface{}{{user.Name, user.Surname, user.Email, user.Phone, getFileLink(fileName)}},
	}

	response, err := sheetsService.Spreadsheets.Values.Append(spreadsheetId, "Registrants", row).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Context(ctx).Do()
	if err != nil || response.HTTPStatusCode != 200 {
		return fmt.Errorf("failed to write user with %v", err)
	}

	return nil
}

func getFileLink(fileName string) string {
	u := url.URL{
		Scheme: "https",
		Host:   "storage.cloud.google.com",
		Path:   fmt.Sprintf("/%s-web/%s", spreadsheetName, fileName),
	}
	return u.String()
}
