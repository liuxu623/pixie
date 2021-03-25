package main

// TODO(zasgar/michelle/philkuz): Remove this when we have automated cert provisioning.
// This is a throwaway script to load certs into the database until we setup automated provisioning.

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"

	"pixielabs.ai/pixielabs/src/cloud/dnsmgr/schema"
	"pixielabs.ai/pixielabs/src/cloud/shared/pgmigrate"
	"pixielabs.ai/pixielabs/src/shared/services/pg"
)

type certInfo struct {
	Cert string `yaml:"crt"`
	Key  string `yaml:"key"`
}

// SSLCert is the record type in the database.
type SSLCert struct {
	Cname     string    `db:"cname"`
	ClusterID uuid.UUID `db:"cluster_id"`
	Cert      string    `db:"cert"`
	Key       string    `db:"key"`
}

func init() {
	pflag.String("certs_path", "../../../../credentials/certs", "The path to the certs")
	pflag.String("env_type", "dev", "The env type (dev, testing, nightly, staging, prod")
	pflag.String("domain_name_suffix", "clusters.dev.withpixie.dev", "The suffix of the domain name to strip out")
	pflag.Bool("update_only", false, "Whether the script should update_only")
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func getCertFilePath() string {
	certsPath := viper.GetString("certs_path")
	envType := viper.GetString("env_type")
	absPath, err := filepath.Abs(certsPath)
	if err != nil {
		log.WithError(err).Fatal("Failed to get cert path")
	}

	return filepath.Join(absPath, envType, "certs.yaml")
}

func getQuery(updateOnly bool) string {
	if updateOnly {
		return `UPDATE ssl_certs SET cert=:cert, key=:key WHERE cname=:cname`
	}
	return `INSERT INTO ssl_certs(cname, cert, key) VALUES (:cname, :cert, :key)`
}

func loadCerts(db *sqlx.DB) {
	certFilePath := getCertFilePath()
	if !fileExists(certFilePath) {
		log.WithField("certFile", certFilePath).
			Fatal("File does not exist")
	}

	log.WithField("certFile", certFilePath).WithField("update_only", viper.GetBool("update_only")).
		Info("Deploying certs")

	out, err := exec.Command("sops", "--decrypt", certFilePath).Output()
	if err != nil {
		log.WithError(err).Fatal("Failed to decrypt certs")
	}

	certs := map[string]certInfo{}
	err = yaml.Unmarshal(out, &certs)
	if err != nil {
		log.WithError(err).Fatal("Failed to unmarshal cert YAMLs")
	}

	domainSuffix := viper.GetString("domain_name_suffix")
	domainSuffix = "." + domainSuffix
	for fullCname, certInfo := range certs {
		log.Infof("inserting %s", fullCname)
		if !strings.HasSuffix(fullCname, domainSuffix) {
			log.Fatal("certificate suffix does not match the supplied domain")
		}

		query := getQuery(viper.GetBool("update_only"))
		cname := strings.TrimSuffix(fullCname, domainSuffix)
		// Remove the wildcard char.
		cname = strings.TrimPrefix(cname, "_.")
		sc := &SSLCert{
			Cname: cname,
			Cert:  certInfo.Cert,
			Key:   certInfo.Key,
		}
		_, err := db.NamedExec(query, &sc)
		if err != nil {
			log.WithError(err).Fatal("Failed to insert certs into DB")
		}
	}
}

// InitDBAndLoadCerts initializes the database then loads or updates the certs.
func InitDBAndLoadCerts() {
	db := pg.MustConnectDefaultPostgresDB()
	err := pgmigrate.PerformMigrationsUsingBindata(db, "dnsmgr_service_migrations", &pgmigrate.SchemaAssetFetcher{
		AssetNames: schema.AssetNames,
		Asset:      schema.Asset,
	})
	if err != nil {
		log.WithError(err).Fatal("Failed to apply migrations")
	}
	loadCerts(db)
}

func main() {
	log.WithField("exec", "load_certs").Info("Starting load_certs...")
	pflag.Parse()

	viper.AutomaticEnv()
	viper.SetEnvPrefix("PL")
	viper.BindPFlags(pflag.CommandLine)

	InitDBAndLoadCerts()

}
