package gotag

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/manifoldco/promptui"
)

const (
	SvIncMajor = 1
	SvIncMinor = 2
	SvIncPatch = 3
)

func Versioning(currentVersion string, svOption uint32) (string, error) {
	versionStr := strings.TrimSpace(currentVersion)
	if versionStr == "" {
		versionStr = "v0.0.0"
	}
	curVersion, err := semver.NewVersion(versionStr)
	if err != nil {
		return "", fmt.Errorf("err: version %s: %s", currentVersion, err)
	}
	switch svOption {
	case SvIncMajor:
		nv := curVersion.IncMajor()
		return nv.Original(), nil
	case SvIncMinor:
		nv := curVersion.IncMinor()
		return nv.Original(), nil
	case SvIncPatch:
		nv := curVersion.IncPatch()
		return nv.Original(), nil
	}
	return "", fmt.Errorf("err: wrong sv option major,minor,patch")
}

func GetLatestVersion() string {
	fetchTagCmd := exec.Command("git", "fetch", "--all", "--tags")
	err := fetchTagCmd.Run()
	if err != nil {
		fmt.Println("\033[1;31m✗\033[0m Command exit err: ", err)
		os.Exit(1)
	}
	cmd := exec.Command("git", "tag", "--sort=-v:refname")
	out, _ := cmd.Output()
	tagList := strings.Split(string(out), "\n")
	if len(strings.Join(tagList, ",")) > 0 {
		return tagList[0]
	}
	return "v0.0.0"
}

func UpdateVersion(svOption uint32) string {
	fmt.Printf("\033[1;32m✔\033[0m [1/5] RUN: fetch tag from repository...\n")
	fetchTagCmd := exec.Command("git", "fetch", "--all", "--tags")
	err := fetchTagCmd.Run()
	if err != nil {
		fmt.Println("\033[1;31m✗\033[0m Command exit err: ", err)
		os.Exit(1)
	}
	fmt.Printf("\033[1;32m✔\033[0m [2/5] RUN: get latest tag from git...\n")
	cmd := exec.Command("git", "tag", "--sort=-v:refname")
	out, _ := cmd.Output()
	tagList := strings.Split(string(out), "\n")
	curVersion := ""
	if len(strings.Join(tagList, ",")) > 0 {
		curVersion = tagList[0]
	}
	fmt.Println("\033[1;32m✔\033[0m [3/5] RUN: generate new version")
	newVersion, err := Versioning(curVersion, svOption)
	if err != nil {
		fmt.Println("\033[1;31m✗\033[0m Command exit err: ", err)
		os.Exit(1)
	}
	promptLabel := fmt.Sprintf("New version %s N for exit and y for run git tag", newVersion)
	prompt := promptui.Prompt{
		Label:     promptLabel,
		IsConfirm: true,
	}
	_, err = prompt.Run()
	if err != nil {
		fmt.Printf("\033[1;31m✗\033[0m Command exit %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("\033[1;32m✔\033[0m [4/5] RUN: git tag %s\n", newVersion)
	err = exec.Command("git", "tag", newVersion).Run()
	if err != nil {
		fmt.Println("\033[1;31m✗\033[0m Command exit err: ", err)
		os.Exit(1)
	}
	fmt.Println("\033[1;32m✔\033[0m [5/5] Completed version:", newVersion)
	fmt.Printf("🎉 !!!Please run command for push new tag: $ git push origin %s\n", newVersion)
	return newVersion
}
