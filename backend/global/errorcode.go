package global

var (
	SUCCESS string = "success"
	FAILED  string = "failed"
	// user errcode
	USER_LOGIN_FAILED      string = "user.login.failed"
	USER_AUTH_NO_RIGHT     string = "user.auth.no.right"
	USER_NOT_EXIST         string = "user.no.exist"
	USER_USERNAME_EXIST    string = "user.username.exist"
	USER_PWD_VERIFY_FAILED string = "user.pwd.verify.failed"

	TOKEN_IS_INVALID string = "token.is.invalid"

	DETECT_START_FAILED string = "detect.start.failed"
	DETECT_EXCEPTION    string = "detect.exception"
	DETECT_IS_INRUNNING string = "detect.is.running"
	DETECT_NO_TASK      string = "detect.no.task"

	DETECT_PROGRESS_QUERY_FAILED string = "detect.progress.query.failed"

	// param error
	PARAM_ERROR string = "param.error"

	SYSTEM_EXCEPTION = "system.exception"

	SYS_INIT_FAILED string = "sys.init.failed"
)
