package cmd

import (
	"github.com/spf13/cobra"
)

type cmder interface {
	getCommand() *cobra.Command
}

type commandsBuilder struct {
	commands []cmder
}

func newCommandsBuilder() *commandsBuilder {
	return &commandsBuilder{}
}

func (b *commandsBuilder) addCommands(commands ...cmder) *commandsBuilder {
	b.commands = append(b.commands, commands...)
	return b
}

func (b *commandsBuilder) addAll() *commandsBuilder {
	b.addCommands(
		b.newAdjustCmd(),
		b.newBlurCmd(),
		b.newCropCmd(),
		b.newFitCmd(),
		b.newFlipCmd(),
		b.newGrayCmd(),
		b.newInvertCmd(),
		b.newResizeCmd(),
		b.newRotateCmd(),
		b.newServerCmd(),
		b.newSharpenCmd(),
		b.newThumbnailCmd(),
		b.newTransposeCmd(),
		b.newTransverseCmd(),
	)

	return b
}

func (b *commandsBuilder) build() *pictarCmd {
	pictarCmd := b.newPictarCmd()
	addCommands(pictarCmd.getCommand(), b.commands...)
	return pictarCmd
}

func addCommands(root *cobra.Command, commands ...cmder) {
	for _, command := range commands {
		cmd := command.getCommand()
		if cmd == nil {
			continue
		}
		root.AddCommand(cmd)
	}
}

type baseBuilderCmd struct {
	cmd     *cobra.Command
	builder *commandsBuilder
}

func (b *commandsBuilder) newBaseBuilderCmd(cmd *cobra.Command) *baseBuilderCmd {
	return &baseBuilderCmd{cmd: cmd, builder: b}
}

func (b *baseBuilderCmd) getCommand() *cobra.Command {
	return b.cmd
}

func (b *baseBuilderCmd) getCommandsBuilder() *commandsBuilder {
	return b.builder
}
