package main

import (
	"os"
	"syscall"

	"github.com/gogo/protobuf/types"
	// This must be GOGO variant or the ENUMs won't work.
	"github.com/gogo/protobuf/jsonpb"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"pixielabs.ai/pixielabs/src/cloud/artifact_tracker/schema"
	"pixielabs.ai/pixielabs/src/cloud/shared/pgmigrate"
	vpb "pixielabs.ai/pixielabs/src/shared/artifacts/versionspb"
	"pixielabs.ai/pixielabs/src/shared/artifacts/versionspb/utils"
	"pixielabs.ai/pixielabs/src/shared/services/pg"
)

func init() {
	pflag.String("versions_file", "VERSIONS.json", "Path to the versions file")
	pflag.Bool("check_only", false, "Only run check")
}

func mustLoadDB() *sqlx.DB {
	db := pg.MustConnectDefaultPostgresDB()
	err := pgmigrate.PerformMigrationsUsingBindata(db, "artifacts_tracker_service_migrations", &pgmigrate.SchemaAssetFetcher{
		AssetNames: schema.AssetNames,
		Asset:      schema.Asset,
	})
	if err != nil {
		log.WithError(err).Fatal("Failed to apply migrations")
	}
	return db
}

func mustReadVersionFile() *vpb.ArtifactSet {
	versionsFilePath := viper.GetString("versions_file")
	log.WithField("file", versionsFilePath).Info("Reading file")

	r, err := os.Open(versionsFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Close()

	msg := &vpb.ArtifactSet{}
	u := jsonpb.Unmarshaler{}

	err = u.Unmarshal(r, msg)
	if err != nil {
		log.Fatalln(err)
	}
	return msg
}

func mustUpdateDatabase(db *sqlx.DB, artifacts *vpb.ArtifactSet) {
	name := artifacts.Name
	txn, err := db.Begin()
	if err != nil {
		log.Fatalln(err)
	}
	defer txn.Rollback()
	// This query does and update of the database. It will insert an artifact if it's missing.
	// Otherwise, it will update the change log if necessary.
	// Note: There is a `commit_hash=artifacts.commit_hash` update to cause the artifact id to get
	// returned. This is non-ideal, but works for our use case.
	query := `
    WITH ins AS (
      INSERT INTO artifacts (artifact_name, create_time, commit_hash, version_str, available_artifacts)
      VALUES($1, $2, $3, $4, $5)
      ON CONFLICT(artifact_name, version_str) DO UPDATE set commit_hash=artifacts.commit_hash
      RETURNING id as artifacts_id
    )
    INSERT INTO artifact_changelogs(artifacts_id, changelog) SELECT artifacts_id, $6 FROM ins
    ON CONFLICT (artifacts_id) DO UPDATE set changelog=EXCLUDED.changelog;
    `
	stmt, err := txn.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}

	for _, artifact := range artifacts.Artifact {
		t, _ := types.TimestampFromProto(artifact.Timestamp)
		_, err := stmt.Exec(name,
			t,
			artifact.CommitHash,
			artifact.VersionStr,
			utils.ToArtifactArray(artifact.AvailableArtifacts),
			artifact.Changelog)
		if err != nil {
			log.Fatalln(err)
		}
	}
	if err := txn.Commit(); err != nil {
		log.Fatalln(err)
	}
}

func checkVersionData(artifacts *vpb.ArtifactSet) {
	if len(artifacts.Name) == 0 {
		log.Fatal("Artifact name must be specfied in the versions file.")
	}
	for _, artifact := range artifacts.Artifact {
		if len(artifact.VersionStr) == 0 {
			log.WithField("entry", artifact.String()).Fatal("Version str must be specified")
		}
		if len(artifact.CommitHash) == 0 {
			log.WithField("version", artifact.VersionStr).Fatal("Must specify the commit hash")
		}
		if len(artifact.AvailableArtifacts) == 0 {
			log.WithField("version", artifact.VersionStr).Fatal("Must have atleast one available artifact")
		}
		if artifact.Timestamp == nil || (artifact.Timestamp.Seconds == 0) {
			log.WithField("version", artifact.VersionStr).Fatal("Timestamp should be specified.")
		}
	}
}

func main() {
	pflag.Parse()

	viper.AutomaticEnv()
	viper.SetEnvPrefix("PL")
	viper.BindPFlags(pflag.CommandLine)

	versionData := mustReadVersionFile()
	checkVersionData(versionData)

	// Print out latest artifact info to Stdout.
	latestArtifact := versionData.Artifact[0]
	m := jsonpb.Marshaler{}
	m.Marshal(os.Stdout, latestArtifact)

	if viper.GetBool("check_only") {
		syscall.Exit(0)
	}

	db := mustLoadDB()
	mustUpdateDatabase(db, versionData)
}
