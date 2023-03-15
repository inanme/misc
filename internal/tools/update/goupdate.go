package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	log "github.com/sirupsen/logrus"
)

//nolint:tagliatelle
type ModuleUpdate struct {
	Path      string    `json:"Path"`
	Version   string    `json:"Version"`
	Time      time.Time `json:"Time"`
	Update    *Update   `json:"Update"`
	Indirect  *bool     `json:"Indirect"`
	GoMod     string    `json:"GoMod"`
	GoVersion string    `json:"GoVersion"`
}

//nolint:tagliatelle
type Update struct {
	Path    string    `json:"Path"`
	Version string    `json:"Version"`
	Time    time.Time `json:"Time"`
}

var ErrEmptyCommand = errors.New("empty command string")

func execCommandString(command string) ([]byte, error) {
	if command == "" || strings.TrimSpace(command) == "" {
		return nil, ErrEmptyCommand
	}
	commandAndArgs := regexp.MustCompile(`\s+`).Split(command, -1)
	return exec.Command(commandAndArgs[0], commandAndArgs[1:]...).Output() //nolint:gosec,wrapcheck
}

func execSimple(command string) (string, error) {
	if command == "" || strings.TrimSpace(command) == "" {
		return "", ErrEmptyCommand
	}
	commandAndArgs := regexp.MustCompile(`\s+`).Split(command, -1)
	out, err := exec.Command(commandAndArgs[0], commandAndArgs[1:]...).CombinedOutput() //nolint:gosec
	return string(out), err
}

const (
	gomod    = "go.mod"
	fileMode = 0o600
)

//nolint:main
func main() {
	gomodContents := update()
	if gomodContents == "" {
		log.Info("up-to-date")
		return
	}
	err := os.WriteFile(gomod, []byte(gomodContents), os.FileMode(fileMode))
	if err != nil {
		log.WithError(err).Fatal("failed to write " + gomod)
	}
	tidyOut, err := execSimple("go mod tidy")
	if err != nil {
		log.WithError(err).Fatal(tidyOut)
	}
	if tidyOut != "" {
		log.Println(tidyOut)
	}

	downloadOut, err := execSimple("go mod download")
	if err != nil {
		log.WithError(err).Fatal(downloadOut)
	}
	if downloadOut != "" {
		log.Println(downloadOut)
	}
}

//nolint:cyclop
func update() string {
	gomodByte, err := os.ReadFile(gomod)
	if err != nil {
		log.WithError(err).Fatal("failed to read go.mod")
	}
	gomodContents := string(gomodByte)
	out, err := execCommandString("go list -json -u -m all")
	if err != nil {
		log.WithError(err).Fatal(string(out))
	}

	thereIsNoUpdate := true
	dec := json.NewDecoder(bytes.NewReader(out))
	for {
		var update ModuleUpdate
		if err := dec.Decode(&update); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.WithError(err).Fatalf("reading go list output failed")
		}

		if update.Update == nil || (update.Indirect != nil && *update.Indirect) {
			continue
		}
		thereIsNoUpdate = false
		currentVersion := semver.MustParse(update.Version)
		updatedVersion := semver.MustParse(update.Update.Version)

		constraints, err := semver.NewConstraint("^" + currentVersion.String())
		if err != nil {
			log.Fatal(err)
		}
		success, errs := constraints.Validate(updatedVersion)
		if !success {
			for _, err := range errs {
				log.WithError(err).Warnf("validation failed for %s ", updatedVersion.String())
			}
		}

		regex := regexp.MustCompile(fmt.Sprintf(`%s\s+%s`, update.Path, update.Version))
		gomodContents = regex.ReplaceAllLiteralString(gomodContents, fmt.Sprintf("%s %s", update.Path, update.Update.Version))

		log.Infof("%s %s %+v", update.Path, update.Version, update.Update.Version)
	}
	if thereIsNoUpdate {
		return ""
	}
	return gomodContents
}
