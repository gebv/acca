package engine

import (
	"github.com/pkg/errors"
	"gopkg.in/reform.v1"
)

func NewSimpleService(db *reform.DB) *SimpleService {
	return &SimpleService{
		db: db,
	}
}

type SimpleService struct {
	db *reform.DB
}

func (s *SimpleService) InternalTransfer(srcAccID, dstAccID, amount int64) (int64, error) {
	var newInvoiceID int64

	err := s.db.InTransaction(func(tx *reform.TX) error {
		newInvoice := &Invoice{
			Key:         "simple1",
			Strategy:    "simple",
			TotalAmount: amount,
		}
		if err := s.db.Insert(newInvoice); err != nil {
			return errors.Wrap(err, "failed insert new invoice")
		}

		newInvoiceID = newInvoice.InvoiceID
		newTransaction := &Transaction{
			Provider:  "internal",
			InvoiceID: newInvoiceID,
		}
		if err := s.db.Insert(newTransaction); err != nil {
			return errors.Wrap(err, "failed insert new transaction")
		}

		opers := []*Operation{
			{
				SrcAccID: srcAccID,
				DstAccID: dstAccID,
				Amount:   amount,
				Strategy: SIMPLE_OPS,
			},
		}
		for _, oper := range opers {
			oper.TransactionID = newTransaction.TransactionID
			oper.InvoiceID = newTransaction.InvoiceID
			if err := s.db.Insert(oper); err != nil {
				return errors.Wrap(err, "failed insert new operation")
			}
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	return newInvoiceID, nil
}
