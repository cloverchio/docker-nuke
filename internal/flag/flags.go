package flag

import (
    "flag"
    "github.com/cloverchio/docker-nuke/pkg"
)

var All *bool
var Containers *bool
var Images *bool
var AllImages *bool
var Volumes *bool
var Networks *bool

func NukeFlagSet() *flag.FlagSet {
    nuke := flag.NewFlagSet("nuke", flag.ContinueOnError)
    All = nuke.Bool("all", false, pkg.UsageMessage("unused", "containers, images, volumes, and networks"))
    Containers = nuke.Bool("containers", false, pkg.UsageMessage("stopped", "containers"))
    Images = nuke.Bool("images", false, pkg.UsageMessage("dangling", "images"))
    AllImages = nuke.Bool("all-images", false, pkg.UsageMessage("unused", "images"))
    Volumes = nuke.Bool("volumes", false, pkg.UsageMessage("unused", "volumes"))
    Networks = nuke.Bool("networks", false, pkg.UsageMessage("unused", "networks"))
    return nuke
}
