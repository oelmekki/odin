package cli

func (cmd *CLI) parseSubCommands(args []string) ([]string, bool) {
	if len(args) == 0 || len(cmd.subCommands) == 0 {
		return args, false
	}
	name := args[0]
	subcmd, ok := cmd.subCommands[name]
	if !ok {
		subcmd.errf("invalid command: %s", name)
	}
	subcmd.Start(args...)

	return []string{}, true
}
