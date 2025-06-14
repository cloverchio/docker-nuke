package pkg

import "fmt"

func Usage() func() {
	return func() {
		fmt.Println("Usage: docker-nuke nuke [options]")
		fmt.Println("Subcommands:")
		fmt.Println(SubCommandUsageMessage("all    ", "unused", "containers, images, volumes, and networks"))
		fmt.Println(SubCommandUsageMessage("containers", "stopped", "containers"))
		fmt.Println(SubCommandUsageMessage("images", "dangling", "images"))
		fmt.Println(SubCommandUsageMessage("all-images", "unused", "images"))
		fmt.Println(SubCommandUsageMessage("volumes", "unused", "volumes"))
		fmt.Println(SubCommandUsageMessage("networks", "unused", "networks"))
		fmt.Println("Example: docker-nuke nuke --containers --images")
	}
}

func SubCommandMessage(verb string, object string) string {
	return fmt.Sprintf("Nuking all %s Docker %s...", verb, object)
}

func SubCommandUsageMessage(command string, verb string, object string) string {
	return fmt.Sprintf("  --%s	%s", command, UsageMessage(verb, object))
}

func UsageMessage(verb string, object string) string {
	return fmt.Sprintf("Nuke all %s Docker %s", verb, object)
}
