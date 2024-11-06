package chainresponsabilty

func ProcessorUserSignUpFactory() IProcessorUserSignUp {
	validHandler := &ValidUserHandler{}
	aggregaterHandler := &AggregaterHandler{}
	saveHandler := &SaveUserHandler{}

	validHandler.LinkChain(aggregaterHandler)
	aggregaterHandler.LinkChain(saveHandler)

	return validHandler
}
