package tuya

import "github.com/Gefist/TuyaLanCLI/internal/commands"

type response struct {
	payload
	returnCode    int
	newSequenceNr int
	commandByte   commands.Type
}
