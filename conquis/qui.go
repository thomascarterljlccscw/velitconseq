import (
	"context"
	"fmt"
	"io"

	database "cloud.google.com/go/spanner/admin/database/apiv1"
	databasepb "google.golang.org/genproto/googleapis/spanner/admin/database/v1"
)

func createBackup(w io.Writer, db string) error {
	ctx := context.Background()
	adminClient, err := database.NewDatabaseAdminClient(ctx)
	if err != nil {
		return err
	}
	defer adminClient.Close()

	op, err := adminClient.CreateBackup(ctx, &databasepb.CreateBackupRequest{
		Parent:     fmt.Sprintf("projects/%s/instances/%s", projectID, instanceID),
		BackupId:   "my-backup",
		Database:   db,
		ExpireTime: timestamp.Add(time.Now(), 24*time.Hour),
	})
	if err != nil {
		return err
	}
	if _, err := op.Wait(ctx); err != nil {
		return err
	}
	fmt.Fprintf(w, "Backup created\n")
	return nil
}
  
