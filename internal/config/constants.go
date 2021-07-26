package config

const (
	ServiceName = "MyDesk CLI"
	InfoUsage   = `
		Use search options using commands below.
		
		Examples:
		
		user --id 3
		user --name "Francisca Rasmussen"
		ticket --id "2614576f-98fb-4031-9e13-beca7a6a73ee" 
		ticket --type "question" 
		ticket --tag "Wisconsin" 
	`
	ErrorUseHelp = "Use 'help' to see list of all available commands"
)
