type accountInfo{
	address  := ""
	role := "sender"
	value  := 0
}

func makeNewAccount() *accountInfo{
	newAccount := new(accountInfo)
	//makeNewAddressは何も実装していない
	newAccount.address = makeNewAddress()
	newAccount.value = 0

	//作成されたアカウントを公開する処理
	return newAccount
}

func getAccountValue(){

}

