package utils

var (
	LoginVerify          = Rules{"CaptchaId": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify       = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
	ChangePasswordVerify = Rules{"Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	ArticleCreate        = Rules{"Title": {NotEmpty()}, "Content": {NotEmpty()}, "CategoryId": {NotEmpty()}}
	ArticleUpdate        = Rules{"Title": {NotEmpty()}, "Content": {NotEmpty()}}
	ArticleSave          = Rules{"Title": {NotEmpty()}}
	PageInfoVerify       = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
)
