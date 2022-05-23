package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/authgear/authgear-server/cmd/portal/internal"
	"github.com/authgear/authgear-server/pkg/lib/config/configsource"
	"github.com/authgear/authgear-server/pkg/lib/theme"
	"github.com/spf13/cobra"
)

var cmdInternalMigrateMediaQueryDarkTheme = &cobra.Command{
	Use:   "migrate-media-query-dark-theme",
	Short: "Migrate media query dark theme",
	Run: func(cmd *cobra.Command, args []string) {
		binder := getBinder()
		dbURL, err := binder.GetRequiredString(cmd, ArgDatabaseURL)
		if err != nil {
			log.Fatalf(err.Error())
		}
		dbSchema, err := binder.GetRequiredString(cmd, ArgDatabaseSchema)
		if err != nil {
			log.Fatalf(err.Error())
		}

		internal.MigrateResources(&internal.MigrateResourcesOptions{
			DatabaseURL:            dbURL,
			DatabaseSchema:         dbSchema,
			UpdateConfigSourceFunc: migrateMediaQueryDarkTheme,
			DryRun:                 &MigrateResourcesDryRun,
		})
	},
}

func migrateMediaQueryDarkTheme(appID string, configSourceData map[string]string, dryRun bool) error {
	originalPath := "static/authgear-dark-theme.css"
	escapedPath := configsource.EscapePath(originalPath)
	encodedData, ok := configSourceData[escapedPath]
	if !ok {
		return nil
	}

	decoded, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return fmt.Errorf("failed to decode %v: %w", originalPath, err)
	}

	if dryRun {
		log.Printf("Converting app secret (%s)", appID)
		log.Printf("Before updated:")
		log.Printf("\n%s\n", string(decoded))
	}

	r := bytes.NewReader(decoded)
	migrated, err := theme.MigrateMediaQueryToClassBased(r)
	if err != nil {
		return fmt.Errorf("failed to migrate %v: %w", originalPath, err)
	}

	if dryRun {
		log.Printf("After updated:")
		log.Printf("\n%s\n", string(migrated))
	}

	configSourceData[escapedPath] = base64.StdEncoding.EncodeToString(migrated)
	return nil
}

func init() {
	cmdInternalBreakingChangeMigrateResources.AddCommand(cmdInternalMigrateMediaQueryDarkTheme)
}
