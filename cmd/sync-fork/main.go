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
    masterName string
    upstreamName string
}

func failOnError(err error) {
    if err == nil {
        return
    }
    log.Fatal(err)
}

func fetchUpstream(settings Settings) error {
    git := exec.Command("git", "remote")
    output, err := git.CombinedOutput()

    if err != nil {
        return err
    }

    outputStr := fmt.Sprintf("%s", output)
    remotes := strings.Split(outputStr, "\n")
    hasOrigin := false
    for _, remote := range remotes {
        if remote == settings.upstreamName {
            hasOrigin = true
            break
        }
    }

    if !hasOrigin {
        return fmt.Errorf("Could not find upstream repository with name %s", settings.upstreamName)
    }

    git = exec.Command("git", "fetch", settings.upstreamName)
    output, err = git.CombinedOutput()
    if err != nil {
        return err
    }

    return nil
}

func checkoutMaster(settings Settings) error {
    git := exec.Command("git", "checkout", settings.masterName)
    output, err := git.CombinedOutput()
    log.Printf("%s\n", output)
    if err != nil {
        return err
    }
    return nil
}

func mergeUpstreamMaster(settings Settings) error {
    upstreamMaster := fmt.Sprintf("%s/%s", settings.upstreamName, settings.masterName)
    git := exec.Command("git", "merge", upstreamMaster)
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
    flag.StringVar(&settings.masterName, "master", "master", "Name of the master branch, default: 'master'")
    flag.StringVar(&settings.masterName, "m", "master", "Name of the master branch, default: 'master'")
    flag.StringVar(&settings.upstreamName, "upstream", "upstream", "Name of the upstream remote entry, default: 'upstream'")
    flag.StringVar(&settings.upstreamName, "u", "upstream", "Name of the upstream remote entry, default: 'upstream'")
    flag.Parse()
    return settings
}

func main() {
    settings := parseCommandLine()
    log.Printf("Syncing with fork:\n\tUpstream: %s\n\tMaster: %s\n\tPush: %t\n", settings.upstreamName, settings.masterName, !settings.doNotPush)

    err := fetchUpstream(settings)
    failOnError(err)
    err = checkoutMaster(settings)
    failOnError(err)
    err = mergeUpstreamMaster(settings)
    failOnError(err)
    if !settings.doNotPush {
        log.Println("Pushing to origin.")
        err = pushOrigin()
        failOnError(err)
    } else {
        log.Println("Not pushing to origin because no-push was specified")
    }
}
