import (
	"context"
	"fmt"
	"io"

	healthcare "google.golang.org/api/healthcare/v1beta1"
)

// approveDataset creates a dataset and then approves it.
func approveDataset(w io.Writer, projectID, location, datasetID string) error {
	ctx := context.Background()

	healthcareService, err := healthcare.NewService(ctx)
	if err != nil {
		return fmt.Errorf("healthcare.NewService: %v", err)
	}

	datasetsService := healthcareService.Projects.Locations.Datasets

	name := fmt.Sprintf("projects/%s/locations/%s/datasets/%s", projectID, location, datasetID)

	dataset := &healthcare.Dataset{
		Name: name,
		TimeZone: "America/Los_Angeles",
	}

	if _, err := datasetsService.Create(name, dataset).Do(); err != nil {
		return fmt.Errorf("Create: %v", err)
	}

	if _, err := datasetsService.Approve(name).Do(); err != nil {
		return fmt.Errorf("Approve: %v", err)
	}

	fmt.Fprintf(w, "Approved dataset: %q\n", datasetID)
	return nil
}
  
