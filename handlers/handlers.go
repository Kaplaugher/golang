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

func CreateOpportunity(c echo.Context) error {
	// Parse the multipart form
	err := c.Request().ParseMultipartForm(10 << 20) // 10 MB max memory
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse form")
	}

	// Extract opportunity data
	opportunity := models.Opportunity{
		Name:        c.FormValue("name"),
		Description: c.FormValue("description"),
		Client:      c.FormValue("client"),
		Needs:       c.FormValue("needs"),
	}

	// Handle file uploads
	form, err := c.MultipartForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to get multipart form")
	}

	files := form.File["files"]
	for _, file := range files {
		// In a real application, you would save the file and store its path
		opportunity.Files = append(opportunity.Files, file.Filename)
	}

	opportunity.ID = 1 // Placeholder ID

	return c.JSON(http.StatusCreated, opportunity)
}

func GetOpportunities(c echo.Context) error {

	opportunities := []models.Opportunity{
		{
			ID:          1,
			Name:        "Sample Opportunity 1",
			Description: "This is a sample opportunity",
			Client:      "ACME Corp",
			Files:       []string{"file1.pdf", "file2.docx"},
			Needs:       "Requires immediate attention",
		},
		{
			ID:          2,
			Name:        "Sample Opportunity 2",
			Description: "Another sample opportunity",
			Client:      "XYZ Inc",
			Files:       []string{"document.pdf"},
			Needs:       "Looking for innovative solutions",
		},
	}

	return c.JSON(http.StatusOK, opportunities)
}
