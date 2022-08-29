package handlers

import (
	"encoding/json"
	"net/http"
	"math/rand"
	"strconv"
	"os"
	"log"
	"fmt"
	dto "waysbuck/dto/result"
	transactionsdto "waysbuck/dto/transaction"
	"waysbuck/models"
	"waysbuck/repositories"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"

	"gopkg.in/gomail.v2"
)

var c = coreapi.Client{
	ServerKey: os.Getenv("SERVER_KEY"),
	ClientKey:  os.Getenv("CLIENT_KEY"),
  }

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) FindTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transactions, err := h.TransactionRepository.FindTransactions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transactions}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) GetTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(transaction)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request  transactionsdto.CreateTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// request := new(transactionsdto.CreateTransactionRequest)
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	// validate := validator.New()
	// err := validate.Struct(request)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// }

	// data form pattern submit to pattern entity db user
	transaction := models.Transaction{
		// Name: request.Name,
		// Desc: request.Desc,
		BuyerID: request.BuyerID,
		// Price: request.Price,
		// Qty:   request.Qty,
		// Image: request.Image,
		//Category: request.Category,
		//ProductID: request.ProductID,
		Price: request.Price,
	}

    // data, err := h.TransactionRepository.CreateTransaction(transaction)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	json.NewEncoder(w).Encode(err.Error())
	// }

	var TransIdIsMatch = false
	var TransactionId int
	for !TransIdIsMatch {
	TransactionId = transaction.BuyerID + request.ProductID + rand.Intn(10000) - rand.Intn(100)
	transactionData, _ := h.TransactionRepository.GetTransaction(TransactionId)
	if transactionData.ID == 0 {
		TransIdIsMatch = true
	}
	}

	newTransaction, err := h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	dataTransactions, err := h.TransactionRepository.GetTransaction(newTransaction.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}


			// 1. Initiate Snap client
		var s = snap.Client{}
		s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)
		// Use to midtrans.Production if you want Production Environment (accept real transaction).

		// 2. Initiate Snap request param
		req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(dataTransactions.ID),
			GrossAmt: int64(dataTransactions.Price),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: dataTransactions.Product.User.Name,
			Email: dataTransactions.Product.User.Email,
		},
		}

		// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: snapResp}
	json.NewEncoder(w).Encode(response)


	// w.WriteHeader(http.StatusOK)
	// response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(data)}
	// json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	idTrans := int(userInfo["time"].(float64))

	request := new(transactionsdto.UpdateTransaction)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	transaction, err := h.TransactionRepository.GetTransaction(idTrans)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	if request.BuyerID != 0 {
		transaction.BuyerID = request.BuyerID
	}

	if request.Total != 0 {
		transaction.Total = request.Total
	}

	if request.Status != "" {
		transaction.Status = request.Status
	}

	_, err = h.TransactionRepository.UpdateTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	//----------------

	dataTransactions, err := h.TransactionRepository.GetTransaction(idTrans)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaction).

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(idTrans),
			GrossAmt: int64(dataTransactions.Total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: dataTransactions.Buyer.Name,
			Email: dataTransactions.Buyer.Email,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: snapResp}
	json.NewEncoder(w).Encode(response)

}

func (h handlerTransaction) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contetmt-type", "application/json")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	data, err := h.TransactionRepository.DeleteTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) Notification(w http.ResponseWriter, r *http.Request) {
	var notificationPayload map[string]interface{}
  
	err := json.NewDecoder(r.Body).Decode(&notificationPayload)
	if err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)

	// // Get One Transaction from repository GetOneTransaction using orderId parameter here ...
	// transaction, _ := h.TransactionRepository.GetOneTransaction(orderId)
  
	if transactionStatus == "capture" {
	  if fraudStatus == "challenge" {
		// TODO set transaction status on your database to 'challenge'
		// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
		h.TransactionRepository.UpdateTransactions("pending",  orderId)
	  } else if fraudStatus == "accept" {
		// TODO set transaction status on your database to 'success'
		h.TransactionRepository.UpdateTransactions("success",  orderId)
	  }
	} else if transactionStatus == "settlement" {
	  // TODO set transaction status on your databaase to 'success'
	  h.TransactionRepository.UpdateTransactions("success",  orderId)
	} else if transactionStatus == "deny" {
	  // TODO you can ignore 'deny', because most of the time it allows payment retries
	  // and later can become success
	  h.TransactionRepository.UpdateTransactions("failed",  orderId)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
	  // TODO set transaction status on your databaase to 'failure'
	  h.TransactionRepository.UpdateTransactions("failed",  orderId)
	} else if transactionStatus == "pending" {
	  // TODO set transaction status on your databaase to 'pending' / waiting payment
	  h.TransactionRepository.UpdateTransactions("pending",  orderId)
	}
  
	w.WriteHeader(http.StatusOK)
  }

  func SendMail(status string, transaction models.Transaction) {

	if status != transaction.Status && (status == "success") {
	  var CONFIG_SMTP_HOST = "smtp.gmail.com"
	  var CONFIG_SMTP_PORT = 587
	  var CONFIG_SENDER_NAME = "waysbucks <mu.rizzki2000@gmail.com>"
	  var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL_SYSTEM")
	  var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD_SYSTEM")
  
	  var productName = transaction.Product.Name
	  var price = strconv.Itoa(transaction.Product.Price)
  
	  mailer := gomail.NewMessage()
	  mailer.SetHeader("From", CONFIG_SENDER_NAME)
	  mailer.SetHeader("To", transaction.Buyer.Email)
	  mailer.SetHeader("Subject", "Transaction Status")
	  mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
	  <html lang="en">
		<head>
		<meta charset="UTF-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>Document</title>
		<style>
		  h1 {
		  color: brown;1
		  }
		</style>
		</head>
		<body>
		<h2>Product payment :</h2>
		<ul style="list-style-type:none;">
		  <li>Name : %s</li>
		  <li>Total payment: Rp.%s</li>
		  <li>Status : <b>%s</b></li>
		</ul>
		</body>
	  </html>`, productName, price, status))
  
	  dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	  )
  
	  err := dialer.DialAndSend(mailer)
	  if err != nil {
		log.Fatal(err.Error())
	  }
  
	  log.Println("Mail sent! to " + transaction.Buyer.Email)
	}
  }

func convertResponseTransaction(u models.Transaction) transactionsdto.TransactionResponse {
	return transactionsdto.TransactionResponse{
		BuyerID: u.BuyerID,
		Price:   u.Price,
		// Product: u.Product,
	}
}
