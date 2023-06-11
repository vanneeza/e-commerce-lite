package transactionservice

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/ksuid"
	"github.com/vanneeza/e-commerce-lite/internal/domain"
	"github.com/vanneeza/e-commerce-lite/internal/domain/entity"
	transactionweb "github.com/vanneeza/e-commerce-lite/internal/domain/web/transaction"
	"github.com/vanneeza/e-commerce-lite/utils/helper"
)

type TransactionServiceImpl struct {
	Db           *sql.DB
	TxRepository domain.TxRepository
}

func NewTransactionService(db *sql.DB, txRepository domain.TxRepository) TransactionService {
	return &TransactionServiceImpl{
		Db:           db,
		TxRepository: txRepository,
	}
}

// MakeOrder implements TransactionService.
func (ts *TransactionServiceImpl) MakeOrder(ctx context.Context, req transactionweb.CreateOrderRequest) (transactionweb.ResponseOrder, error) {
	tx, err := ts.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	product, err := ts.TxRepository.ProductRepository.FindById(ctx, tx, req.ProductId)
	helper.PanicError(err)

	customers, err := ts.TxRepository.CustomerRepository.FindByParams(ctx, tx, req.CustomerId, "")
	helper.PanicError(err)

	store, err := ts.TxRepository.StoreRepository.FindById(ctx, tx, product.Store.Id)
	helper.PanicError(err)

	totalPrice := float64(product.Price) * float64(req.Qty)
	makeOrderDetail := entity.OrderDetail{
		Id:         ksuid.New().String(),
		TotalPrice: totalPrice,
		Status:     "WAITING TO PAYMENT",
		BuyDate:    time.Now(),
	}

	orderDetail, err := ts.TxRepository.TxRepository.CreateOrderDetail(ctx, tx, makeOrderDetail)
	helper.PanicError(err)

	Makeorder := entity.Order{
		Id:         ksuid.New().String(),
		Qty:        req.Qty,
		ProductId:  product,
		CustomerId: customers,
		StoreId:    store,
		TxDetailId: orderDetail,
	}

	order, err := ts.TxRepository.TxRepository.CreateOrder(ctx, tx, Makeorder)
	helper.PanicError(err)

	orderResponse := transactionweb.ResponseOrder{
		Id:         order.Id,
		Qty:        req.Qty,
		ProductId:  product,
		CustomerId: customers,
		StoreId:    store,
		Detail:     orderDetail,
	}

	return orderResponse, nil
}

func (ts *TransactionServiceImpl) SaveCallBack(ctx context.Context, req transactionweb.CallbackResponse) error {
	tx, err := ts.Db.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	callback := entity.Callback{
		Id:                 req.Id,
		DetailId:           req.DetailId,
		UserId:             req.UserId,
		IsHigh:             req.IsHigh,
		PaymentMethod:      req.PaymentMethod,
		Status:             req.Status,
		MerchantName:       req.MerchantName,
		Amount:             req.Amount,
		PaidAmount:         req.PaidAmount,
		BankCode:           req.BankCode,
		PaidDate:           req.PaidDate,
		PayerEmail:         req.PayerEmail,
		Description:        req.Description,
		Created:            req.Created,
		Updated:            req.Updated,
		Currency:           req.Currency,
		PaymentChannel:     req.PaymentChannel,
		PaymentDestination: req.PaymentDestination,
	}
	_, errCallback := ts.TxRepository.TxRepository.CreateCallback(ctx, tx, callback)
	helper.PanicError(errCallback)
	log.Println(callback.Status, "callback status")
	fmt.Scanln()

	if callback.Status == "PAID" {
		err2 := ts.TxRepository.TxRepository.UpdateStatusDetail(ctx, tx, callback)
		helper.PanicError(err2)

		order, err4 := ts.TxRepository.TxRepository.FindOrderById(ctx, tx, callback.DetailId)
		helper.PanicError(err4)

		s, err5 := ts.TxRepository.StoreRepository.FindById(ctx, tx, order.StoreId.Id)
		helper.PanicError(err5)

		log.Println(order, "order")
		log.Println(s, "store")
		addBalance := s.Balance + int64(callback.Amount)

		store := entity.Store{
			Id:      s.Id,
			Name:    s.Name,
			Email:   s.Email,
			NoHp:    s.NoHp,
			Address: s.Address,
			Balance: addBalance,
		}
		log.Println(addBalance, "add balance")
		fmt.Scanln()
		_, err6 := ts.TxRepository.StoreRepository.Update(ctx, tx, store)
		helper.PanicError(err6)

		return nil
	} else {
		return errors.New("transaction was failed")
	}

}
