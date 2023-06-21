package database

import (
	"database/sql"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
)

// TestDB connects to the database and prepares it for use by a test, by
// running migrations and loading fixtures.
//
// It uses the following environment variables:
//
//   - TEST_DATABASE_HOST: the host to connect to (default: localhost)
//   - POSTGRES_PASSWORD_FILE: the file containing the postgres password
//     (default: <project-root>/secrets/postgres_password)
//
// fixtures is an optional list of additional fixture directories to load,
// relative to the "fixtures" directory in this package. The "default" fixture
// is always loaded.
//
// If there are any errors, the test will be aborted.
func TestDB(t *testing.T, fixtures ...string) *sql.DB {
	passwordFile := os.Getenv("POSTGRES_PASSWORD_FILE")
	if passwordFile == "" {
		passwordFile = filepath.Join(gitRoot(), "secrets/postgres_password.txt")
	}

	//nolint:gosec
	password, err := os.ReadFile(passwordFile)
	if err != nil {
		t.Fatalf("unable to read postgres password file: %v", err)
	}

	host := os.Getenv("TEST_DATABASE_HOST")
	if host == "" {
		host = "localhost"
	}

	dsn := (&url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword("postgres", string(password)),
		Host:     host,
		Path:     "nomenclator_test",
		RawQuery: "sslmode=disable",
	}).String()

	pkgPath := packagePath()
	cfg := &Config{
		DSN:           dsn,
		MigrationsDir: filepath.Join(pkgPath, "migrations"),
	}

	conn, err := Connect(cfg)
	if err != nil {
		t.Fatalf("unable to connect to database: %v", err)
	}

	// Run migrations.

	err = Migrate(cfg, conn)
	if err != nil {
		t.Fatalf("unable to migrate database: %v", err)
	}

	// Load fixtures.

	fixtureOptions := []func(*testfixtures.Loader) error{
		testfixtures.Database(conn),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory(filepath.Join(pkgPath, "fixtures", "default")),
	}
	for _, fixture := range fixtures {
		fixtureOptions = append(fixtureOptions, testfixtures.Directory(filepath.Join(pkgPath, "fixtures", fixture)))
	}

	loader, err := testfixtures.New(fixtureOptions...)
	if err != nil {
		t.Fatalf("unable to create fixtures loader: %v", err)
	}

	err = loader.Load()
	if err != nil {
		t.Fatalf("unable to load fixtures: %v", err)
	}

	return conn
}

// gitRoot returns the root of the git repository.
func gitRoot() string {
	output, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(output))
}

// packagePath returns the path to the current package.
func packagePath() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}
