package chainresponsabilty

type IProcessorUserSignUp interface {
	LinkChain(IProcessorUserSignUp)
	Execute(*User) Status
}
