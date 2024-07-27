package core

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var marmotVersion string

func ExecutableVersion() (string, error) {
	if versionPath, pathErr := versionFilePath(); pathErr != nil {
		return "", fmt.Errorf("failed to locate version file; %w", pathErr)
	} else if rawVersion, readErr := readVersion(versionPath); readErr != nil {
		return "", fmt.Errorf("failed to read version from %s; %w", versionPath, readErr)
	} else if version, parseErr := parseVersion(rawVersion); parseErr != nil {
		return "", fmt.Errorf("failed to parse version from %s; %w", versionPath, parseErr)
	} else {
		return version, nil
	}
}

func InitMarmotVersion(version string) error {
	if marmotVersion != "" {
		return fmt.Errorf("marmot version already set to <%s> but given <%s>", marmotVersion, version)
	} else {
		marmotVersion = version
		return nil
	}
}

func MarmotVersion() (string, error) {
	if marmotVersion == "" {
		return "", errors.New("marmot version not set")
	} else {
		return marmotVersion, nil
	}
}

func parseVersion(versionRaw string) (string, error) {
	if version := strings.TrimSpace(versionRaw); version == "" {
		return "", fmt.Errorf("<%s> is effectively empty", versionRaw)
	} else {
		return version, nil
	}
}

func readVersion(versionFilename string) (string, error) {
	if versionBytes, readErr := os.ReadFile(versionFilename); readErr != nil {
		return "", readErr
	} else {
		return string(versionBytes), nil
	}
}

func versionFilePath() (string, error) {
	if executablePath, executableErr := os.Executable(); executableErr != nil {
		return "", executableErr
	} else if linkTarget, err := filepath.EvalSymlinks(executablePath); err != nil {
		return "", err
	} else {
		fmt.Printf("executablePath=%s\n", executablePath)
		fmt.Printf("linkTarget=%s\n", linkTarget)

		programDir := filepath.Dir(linkTarget)
		versionPath := filepath.Join(programDir, "version")
		fmt.Printf("versionPath=%s\n", versionPath)
		return versionPath, nil
	}
}
