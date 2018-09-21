package walletdmanager

const (
	// DefaultTransferFee is the default fee. It is expressed in TRTL
	DefaultTransferFee float64 = 1
	// DefaultTransferMixin is the default mixin
	DefaultTransferMixin = 7

	logWalletdCurrentSessionFilename     = "roto2-service-session.log"
	logWalletdAllSessionsFilename        = "roto2-service.log"
	logTurtleCoindCurrentSessionFilename = "Roto2Coind-session.log"
	logTurtleCoindAllSessionsFilename    = "Roto2Coind.log"
	walletdLogLevel                      = "3" // should be at least 3 as I use some logs messages to confirm creation of wallet
	walletdCommandName                   = "roto2-service"
	turtlecoindCommandName               = "Roto2Coind"
)
