package tests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gebv/acca/api"
)

func Test02_01SberbankStrategy(t *testing.T) {
	h := NewHelperData(t)

	t.Run("CreateCurrency", h.CreateCurrency("curr1"))

	t.Run("CreateAccount", h.CreateAccount("acc2.1.1", "curr1"))

	t.Run("CreateAccount", h.CreateAccount("acc2.1.2", "curr1"))

	t.Run("RechargeFromInvoice", func(t *testing.T) {

		t.Run("NewInvoice", h.NewInvoice("inv1", "simple", nil))

		Meta := map[string]string{
			"callback":    "https://ya.ru",
			"description": "",
			"email":       "test@mail.ru",
		}
		meta, err := json.Marshal(&Meta)
		require.NoError(t, err)

		t.Run("AddTransactionToInvoice", h.AddTransactionToInvoice(
			"inv1",
			"tx1",
			"sberbank",
			1000,
			&meta,
			[]*api.AddTransactionToInvoiceRequest_Oper{
				h.CreateOperation(
					"acc2.1.1",
					"acc2.1.2",
					"",
					true,
					api.OperStrategy_RECHARGE_OPS,
					1000,
				),
			},
		))

		t.Run("AuthInvoice", h.AuthInvoice("inv1"))

		// Так как авторизация транзакции проходит через 2 запроса.
		t.Run("WaitTransaction", h.WaitTransaction("tx1", api.TxStatus_AUTH_TX))
		t.Run("WaitTransaction", h.WaitTransaction("tx1", api.TxStatus_AUTH_TX))

		t.Run("CheckBalances", h.CheckBalances("acc2.1.1", "curr1"))

		t.Run("CheckBalances", h.CheckBalances("acc2.1.2", "curr1"))

		t.Run("CheckTransactionWithProvider", h.CheckTransactionWithProvider("tx1", "CREATED", api.TxStatus_AUTH_TX))

		t.Run("SendCardDataInSberbank", h.SendCardDataInSberbank("tx1"))

		t.Run("WaitInvoice", h.WaitInvoice("inv1", api.InvoiceStatus_WAIT_I))

		t.Run("CheckTransactionWithProvider", h.CheckTransactionWithProvider("tx1", "APPROVED", api.TxStatus_HOLD_TX))

		h.BalanceInc("acc2.1.1", 1000)
		h.AcceptedBalanceInc("acc2.1.1", 1000)
		h.BalanceInc("acc2.1.2", 1000)
		h.AcceptedBalanceInc("acc2.1.2", 1000)

		t.Run("AcceptInvoice", h.AcceptInvoice("inv1"))

		t.Run("WaitInvoice", h.WaitInvoice("inv1", api.InvoiceStatus_ACCEPTED_I))

		t.Run("CheckTransaction", h.CheckTransaction("tx1", api.TxStatus_ACCEPTED_TX))

		t.Run("CheckBalances", h.CheckBalances("acc2.1.1", "curr1"))

		t.Run("CheckBalances", h.CheckBalances("acc2.1.2", "curr1"))

	})
}

func Test02_02SberbankStrategy(t *testing.T) {
	h := NewHelperData(t)

	t.Run("CreateCurrency", h.CreateCurrency("curr1"))

	t.Run("CreateAccount", h.CreateAccount("acc2.2.1", "curr1"))

	t.Run("CreateAccount", h.CreateAccount("acc2.2.2", "curr1"))

	t.Run("RechargeFromInvoice", func(t *testing.T) {

		t.Run("NewInvoice", h.NewInvoice("inv1", "simple", nil))

		Meta := map[string]string{
			"callback":    "https://ya.ru",
			"description": "",
			"email":       "test@mail.ru",
		}
		meta, err := json.Marshal(&Meta)
		require.NoError(t, err)

		t.Run("AddTransactionToInvoice", h.AddTransactionToInvoice(
			"inv1",
			"tx1",
			"sberbank",
			1000,
			&meta,
			[]*api.AddTransactionToInvoiceRequest_Oper{
				h.CreateOperation(
					"acc2.2.1",
					"acc2.2.2",
					"",
					true,
					api.OperStrategy_RECHARGE_OPS,
					1000,
				),
			},
		))

		t.Run("AuthInvoice", h.AuthInvoice("inv1"))

		// Так как авторизация транзакции проходит через 2 запроса.
		t.Run("WaitTransaction", h.WaitTransaction("tx1", api.TxStatus_AUTH_TX))
		t.Run("WaitTransaction", h.WaitTransaction("tx1", api.TxStatus_AUTH_TX))

		t.Run("CheckBalances", h.CheckBalances("acc2.2.1", "curr1"))

		t.Run("CheckBalances", h.CheckBalances("acc2.2.2", "curr1"))

		t.Run("CheckTransactionWithProvider", h.CheckTransactionWithProvider("tx1", "CREATED", api.TxStatus_AUTH_TX))

		t.Run("SendCardDataInSberbank", h.SendCardDataInSberbank("tx1"))

		t.Run("WaitInvoice", h.WaitInvoice("inv1", api.InvoiceStatus_WAIT_I))

		t.Run("CheckTransactionWithProvider", h.CheckTransactionWithProvider("tx1", "APPROVED", api.TxStatus_HOLD_TX))

		t.Run("CheckBalances", h.CheckBalances("acc2.2.1", "curr1"))

		t.Run("CheckBalances", h.CheckBalances("acc2.2.2", "curr1"))

		t.Run("RejectInvoice", h.RejectInvoice("inv1"))

		// Так как отмена транзакции проходит через 2 запроса.
		t.Run("WaitTransaction", h.WaitTransaction("tx1", api.TxStatus_REJECTED_TX))
		t.Run("WaitTransaction", h.WaitTransaction("tx1", api.TxStatus_REJECTED_TX))

		t.Run("CheckTransactionWithProvider", h.CheckTransactionWithProvider("tx1", "REVERSED", api.TxStatus_REJECTED_TX))

		t.Run("CheckBalances", h.CheckBalances("acc2.2.1", "curr1"))

		t.Run("CheckBalances", h.CheckBalances("acc2.2.2", "curr1"))

	})
}

func Test02_03SberbankStrategy(t *testing.T) {
	h := NewHelperData(t)

	t.Run("CreateCurrency", h.CreateCurrency("curr1"))

	t.Run("CreateAccount", h.CreateAccount("acc2.3.1", "curr1"))

	t.Run("CreateAccount", h.CreateAccount("acc2.3.2", "curr1"))

	t.Run("RechargeFromInvoice", func(t *testing.T) {

		t.Run("NewInvoice", h.NewInvoice("inv1", "simple", nil))

		Meta := map[string]string{
			"callback":    "https://ya.ru",
			"description": "",
			"email":       "test@mail.ru",
		}
		meta, err := json.Marshal(&Meta)
		require.NoError(t, err)

		t.Run("AddTransactionToInvoice", h.AddTransactionToInvoice(
			"inv1",
			"tx1",
			"sberbank",
			1000,
			&meta,
			[]*api.AddTransactionToInvoiceRequest_Oper{
				h.CreateOperation(
					"acc2.3.1",
					"acc2.3.2",
					"",
					false,
					api.OperStrategy_RECHARGE_OPS,
					1000,
				),
				h.CreateOperation(
					"acc2.3.1",
					"acc2.3.2",
					"",
					false,
					api.OperStrategy_WITHDRAW_OPS,
					100,
				),
			},
		))

		t.Run("AuthInvoice", h.AuthInvoice("inv1"))

		// Так как авторизация транзакции проходит через 2 запроса.
		t.Run("WaitTransaction", h.WaitTransaction("tx1", api.TxStatus_AUTH_TX))
		t.Run("WaitTransaction", h.WaitTransaction("tx1", api.TxStatus_AUTH_TX))

		t.Run("CheckBalances", h.CheckBalances("acc2.3.1", "curr1"))

		t.Run("CheckBalances", h.CheckBalances("acc2.3.2", "curr1"))

		t.Run("CheckTransactionWithProvider", h.CheckTransactionWithProvider("tx1", "CREATED", api.TxStatus_AUTH_TX))

		h.BalanceInc("acc2.3.1", 1000)
		h.BalanceDec("acc2.3.1", 100)
		h.BalanceInc("acc2.3.2", 1000)
		h.BalanceDec("acc2.3.2", 100)
		h.AcceptedBalanceInc("acc2.3.1", 1000)
		h.AcceptedBalanceDec("acc2.3.1", 100)
		h.AcceptedBalanceInc("acc2.3.2", 1000)
		h.AcceptedBalanceDec("acc2.3.2", 100)

		t.Run("SendCardDataInSberbank", h.SendCardDataInSberbank("tx1"))

		t.Run("WaitInvoice", h.WaitInvoice("inv1", api.InvoiceStatus_ACCEPTED_I))

		t.Run("CheckTransactionWithProvider", h.CheckTransactionWithProvider("tx1", "DEPOSITED", api.TxStatus_ACCEPTED_TX))

		t.Run("CheckBalances", h.CheckBalances("acc2.3.1", "curr1"))

		t.Run("CheckBalances", h.CheckBalances("acc2.3.2", "curr1"))

	})
}