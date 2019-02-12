package main;

import (
    "flag"
    "fmt"
    "log"
    "os/exec"
    "strings"
)

type Settings struct {
    doNotPush bool
}

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

func parseCommandLine() Settings {
    settings := Settings{}
    flag.BoolVar(&settings.doNotPush, "no-push", false, "Automatically push a succesful sync to origin branch.")
    flag.BoolVar(&settings.doNotPush, "np", false, "Automatically push a succesful sync to origin branch.")
    flag.Parse()
    return settings
}

func main() {
    log.Println("Syncing with upstream...")

    settings := parseCommandLine()

    err := fetchUpstream()
    failOnError(err)
    err = checkoutMaster()
    failOnError(err)
    err = mergeUpstreamMaster()
    failOnError(err)
    if !settings.doNotPush {
        log.Println("Pushing to origin.")
        err = pushOrigin()
        failOnError(err)
    } else {
        log.Println("Not pushing to origin because no-push was specified")
    }
}
