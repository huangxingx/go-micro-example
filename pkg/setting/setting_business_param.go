package setting

type businessParam struct {
	OssEndPoint    string
	OssBucketName  string
	OssDownloadUrl string
}

var businessParamDevelop = &businessParam{"oss-cn-hangzhou.aliyuncs.com", "cs-liaoliao",
	"https://cs-liaoliao.oss-cn-hangzhou.aliyuncs.com"}

var businessParamOnline = &businessParam{}

var businessParamOnlineTest = &businessParam{"oss-cn-hangzhou.aliyuncs.com", "cs-liaoliao",
	"https://cs-liaoliao.oss-cn-hangzhou.aliyuncs.com"}

var businessParamRelease = &businessParam{"oss-cn-hangzhou.aliyuncs.com", "cs-liaoliao",
	"https://cs-liaoliao.oss-cn-hangzhou.aliyuncs.com"}
