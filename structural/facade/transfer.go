package facade

import "errors"

type Transfer struct {
	accounts map[string]float64
}

func (t *Transfer) Exec(acc1 string, acc2 string, val float64) error {
	curr1, ok := t.accounts[acc1]
	if !ok {
		return errors.New("account 1 not found")
	} else if curr1 <= 0 || val > curr1 {
		return errors.New("impossible transfer from this account")
	}

	_, ok = t.accounts[acc2]
	if !ok {
		return errors.New("account 2 not found")
	}

	t.accounts[acc1] -= val
	t.accounts[acc2] += val
	return nil
}

type TransferFacade struct {
	account1 string
	account2 string
	transfer *Transfer
}

func NewTransferFacade(acc1, acc2 string) TransferFacade {
	return TransferFacade{
		account1: acc1,
		account2: acc2,
		transfer: &Transfer{
			accounts: make(map[string]float64),
		},
	}
}

func (t *TransferFacade) Exec(val float64) error {
	t.transfer.Exec(t.account1, t.account2, val)
	return nil
}
