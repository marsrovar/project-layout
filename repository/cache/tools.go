package cache

func bindLastLoginTokenKey(account string) string {
	return "login:" + account
}

func bindUserInfoKey(account string) string {
	return "userInfo:" + account
}
