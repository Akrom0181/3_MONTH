package config

const (
	ERR_INFORMATION     = "The server has received the request and is continuing the process"
	SUCCESS             = "The request was successful"
	ERR_REDIRECTION     = "You have been redirected and the completion of the request requires further action"
	ERR_BADREQUEST      = "Bad request"
	ERR_INTERNAL_SERVER = "While the request appears to be valid, the server could not complete the request"
	User_ROLE           = "User"
	ADMIN_ROLE          = "admin"


	SmtpServer   = "smtp.gmail.com"
	SmtpPort     = "587"
	SmtpUsername = "akromjonotaboyev@gmail.com"
	SmtpPassword = "ndck xoka brsb iasx"
)

var SignedKey = []byte("MGJd@Ro]yKoCc)mVY1^c:upz~4rn9Pt!hYd]>c8dt#+%")
