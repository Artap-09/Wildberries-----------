package pattern

import (
	"fmt"
)

/*
func main() {
	fmt.Println()
	walletFacade := newWalletFacade("abc", 1234) // Создаем аккаунт
	fmt.Println()

	err := walletFacade.addMoneyToWallet("abc", 1234, 10) // Добавляем деньги в кошелек
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5) // Уменьшаем баланс кошелька
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
*/

//Фасад
type walletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
	notification *notification
	ledger       *ledger
}

// newWalletFacade - фасад кошелька.
func newWalletFacade(accountID string, code int) *walletFacade {
	fmt.Println("Starting create account")
	walletFacade := &walletFacade{
		account:      newAccount(accountID), // Создаем аккаунт
		securityCode: newSecurityCode(code), // Создаем код безопасности
		wallet:       newWallet(),           // Создаем кошелек
		notification: &notification{},       // Уведомление
		ledger:       &ledger{},             // Бухгалерская книга
	}
	fmt.Println("Account created")
	return walletFacade
}

//Добавляет деньги в кошелек
func (w *walletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

//Уменьшает баланс кошелька
func (w *walletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}

	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}

	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

//Аккаунт
type account struct {
	name string
}

//Создает аккаунт
func newAccount(accountName string) *account {
	return &account{
		name: accountName,
	}
}

//Проверяет аккаунт
func (a *account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("Account Name is incorrect")
	}

	fmt.Println("Account Verified")
	return nil
}

//Код безопасности
type securityCode struct {
	code int
}

//Создает код безопасности
func newSecurityCode(code int) *securityCode {
	return &securityCode{
		code: code,
	}
}

//Проверяет код безопасности
func (s *securityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("Security Code is incorrect")
	}

	fmt.Println("SecurityCode Verified")
	return nil
}

//Кошелек
type wallet struct {
	balance int
}

//Создает кошелек
func newWallet() *wallet {
	return &wallet{
		balance: 0,
	}
}

//Выполняем операцию кредит с кошельком
func (w *wallet) creditBalance(amount int) {
	//изменяем баласн кошелька
	w.balance += amount
	fmt.Println("Wallet balance added successfully")
	return
}

//Выполняем операцию дебит с кошельком
func (w *wallet) debitBalance(amount int) error {
	//Проверяем баланс кошелька
	if w.balance < amount {
		return fmt.Errorf("Balance is not sufficient")
	}
	fmt.Println("Wallet balance is Sufficient")
	//Изменяем баланс кошелька
	w.balance = w.balance - amount
	return nil
}

//ledger
type ledger struct {
}

//Делаем запись в бух. книге
func (s *ledger) makeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
	return
}

//Уведомления
type notification struct {
}

//Отправляем сообщение о кредите кошелька
func (n *notification) sendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

//Отправляем сообщение о дебете кошелька
func (n *notification) sendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}
