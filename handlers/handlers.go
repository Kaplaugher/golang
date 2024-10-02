package handlers

import (
	"myapi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func GetOpportunity(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid opportunity ID")
	}

	// In a real application, you would fetch the opportunity from a database
	opportunity := models.Opportunity{
		ID:          id,
		Name:        "Sample Opportunity",
		Description: "This is a sample opportunity",
		Client:      "ACME Corp",
		Files:       []string{"file1.pdf", "file2.docx"},
		Needs:       "Requires immediate attention",
	}

	return c.JSON(http.StatusOK, opportunity)
}

func AddFilesToOpportunity(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid opportunity ID")
	}

	var files struct {
		Files []string `json:"files"`
	}

	if err := c.Bind(&files); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// In a real application, you would update the opportunity in the database
	// For this example, we'll just return the files that would be added
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "Files added successfully",
		"opportunityId": id,
		"addedFiles":    files.Files,
	})
}
