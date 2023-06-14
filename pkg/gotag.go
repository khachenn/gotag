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

const CheckedEmoji = "\033[1;32m✔\033[0m"
const CloseEmoji = "\033[1;31m✗\033[0m"

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
		fmt.Printf("%s Command exit err: %s \n", CloseEmoji, err)
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
	fmt.Printf("%s [1/5] RUN: fetch tag from repository...\n", CheckedEmoji)
	fetchTagCmd := exec.Command("git", "fetch", "--all", "--tags")
	err := fetchTagCmd.Run()
	if err != nil {
		fmt.Printf("%s Command exit err: %s \n", CloseEmoji, err)
		os.Exit(1)
	}
	fmt.Printf("%s [2/5] RUN: get latest tag from git...\n", CheckedEmoji)
	cmd := exec.Command("git", "tag", "--sort=-v:refname")
	out, _ := cmd.Output()
	tagList := strings.Split(string(out), "\n")
	curVersion := ""
	if len(strings.Join(tagList, ",")) > 0 {
		curVersion = tagList[0]
	}
	newVersion, err := Versioning(curVersion, svOption)
	if err != nil {
		fmt.Printf("%s Command exit err: %s \n", CloseEmoji, err)
		os.Exit(1)
	}
	fmt.Printf("%s [3/5] RUN: update version from \033[1;30m%s\033[0m to \033[1;33m%s\033[0m\n", CheckedEmoji, curVersion, newVersion)
	promptLabel := fmt.Sprintf("New version \033[1;33m%s\033[0m choose \033[1;37mN\033[0m for exit or \033[1;37my\033[0m for run git tag", newVersion)
	prompt := promptui.Prompt{
		Label:     promptLabel,
		IsConfirm: true,
	}
	_, err = prompt.Run()
	if err != nil {
		fmt.Printf("%s Command exit err: %s \n", CloseEmoji, err)
		os.Exit(1)
	}
	fmt.Printf("%s [4/5] RUN: git tag %s\n", CheckedEmoji, newVersion)
	err = exec.Command("git", "tag", newVersion).Run()
	if err != nil {
		fmt.Printf("%s Command exit err: %s\n", CloseEmoji, err)
		os.Exit(1)
	}
	fmt.Printf("%s [5/5] Completed version: \033[1;32m%s\033[0m\n", CheckedEmoji, newVersion)
	fmt.Printf("🎉 !!!Please run command for push new tag: $ git push origin %s\n", newVersion)
	return newVersion
}
