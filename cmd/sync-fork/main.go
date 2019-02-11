package main;

import (
    "fmt"
    "log"
    "os/exec"
    "strings"
)

func failOnError(err error) {
    if err == nil {
        return
    }
    log.Fatal(err)
}

func fetchUpstream() error {
    git := exec.Command("git", "remote")
    output, err := git.CombinedOutput()

    if err != nil {
        return err
    }

    outputStr := fmt.Sprintf("%s", output)
    remotes := strings.Split(outputStr, "\n")
    hasOrigin := false
    for _, remote := range remotes {
        if remote == "upstream" {
            hasOrigin = true
            break
        }
    }

    if !hasOrigin {
        return fmt.Errorf("No upstream repository found!")
    }

    git = exec.Command("git", "fetch", "upstream")
    output, err = git.CombinedOutput()
    if err != nil {
        return err
    }

    return nil
}

func checkoutMaster() error {
    git := exec.Command("git", "checkout", "master")
    output, err := git.CombinedOutput()
    log.Printf("%s\n", output)
    if err != nil {
        return err
    }
    return nil
}

func mergeUpstreamMaster() error {
    git := exec.Command("git", "merge", "upstream/master")
    output, err := git.CombinedOutput()
    log.Printf("%s\n", output)
    if err != nil {
        return err
    }
    return nil
}

func pushOrigin() error {
    git := exec.Command("git", "push")
    output, err := git.CombinedOutput()
    log.Printf("%s\n", output)
    if err != nil {
        return err
    }
    return nil
}

func main() {
    log.Println("Syncing with upstream...")
    // TODO allow configuration of folder to run in
    // TODO allow configuration of upstream name
    // TODO allow configuration of origin name
    // TODO allow configuration of master branch name
    // TODO add option to recursively search for git repositories in subfolders
    // TODO make push operation optional

    err := fetchUpstream()
    failOnError(err)
    err = checkoutMaster()
    failOnError(err)
    err = mergeUpstreamMaster()
    failOnError(err)
    err = pushOrigin()
    failOnError(err)
}
