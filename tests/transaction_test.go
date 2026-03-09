package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"transaction-service/dto"
	"transaction-service/handlers"
	"transaction-service/models"
	"transaction-service/services"
	"transaction-service/validator"
)

// MockTransactionRepository is a mock implementation of the TransactionRepository interface
type MockTransactionRepository struct {
	transactions map[int64]models.Transaction
	nextID       int64
}

func NewMockTransactionRepository() *MockTransactionRepository {
	return &MockTransactionRepository{
		transactions: make(map[int64]models.Transaction),
		nextID:       1,
	}
}

func (m *MockTransactionRepository) Insert(transaction models.Transaction) (models.Transaction, error) {
	transaction.TransactionId = m.nextID
	m.transactions[m.nextID] = transaction
	m.nextID++
	return transaction, nil
}

type MockTransactionTypeRepository struct {
	types map[int]string
}

func NewMockTransactionTypeRepository() *MockTransactionTypeRepository {
	return &MockTransactionTypeRepository{
		types: map[int]string{
			1: "Normal Purchase",
			2: "Purchase with Installments",
			3: "Withdrawal",
			4: "Credit Voucher",
		},
	}
}

func (m *MockTransactionTypeRepository) GetTypes() (map[int]string, error) {
	return m.types, nil
}

// TestCreateTransaction_NormalPurchase tests creating a normal purchase transaction
func TestCreateTransaction_NormalPurchase(t *testing.T) {
	// Arrange
	mockAccountRepo := NewMockAccountRepository()
	mockTransactionRepo := NewMockTransactionRepository()

	// Create an account
	account, _ := mockAccountRepo.Insert(models.Account{DocumentNumber: "123456789"})

	transactionService := services.NewTransactionService(
		mockTransactionRepo,
		mockAccountRepo,
	)

	transactionTypeRepo := NewMockTransactionTypeRepository()
	transactionValidator := validator.NewTransactionValidator(transactionTypeRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService, transactionValidator)

	requestBody := dto.TransactionCreateDTO{
		AccountID:       strconv.FormatInt(account.ID, 10),
		Amount:          100.50,
		TransactionType: "1", // Normal Purchase
	}
	body, _ := json.Marshal(requestBody)

	req := httptest.NewRequest("POST", "/transactions", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	transactionHandler.CreateTransaction(w, req)

	// Assert
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	var response dto.TransactionResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.Amount != -100.50 {
		t.Errorf("Expected amount 100.50, got %f", response.Amount)
	}

	if response.TransactionType != "1" {
		t.Errorf("Expected operation type ID 1, got %s", response.TransactionType)
	}
}

// TestCreateTransaction_PurchaseWithInstallments tests creating a purchase with installments
func TestCreateTransaction_PurchaseWithInstallments(t *testing.T) {
	// Arrange
	mockAccountRepo := NewMockAccountRepository()
	mockTransactionRepo := NewMockTransactionRepository()

	// Create an account
	account, _ := mockAccountRepo.Insert(models.Account{DocumentNumber: "123456789"})

	transactionService := services.NewTransactionService(
		mockTransactionRepo,
		mockAccountRepo,
	)

	transactionTypeRepo := NewMockTransactionTypeRepository()
	transactionValidator := validator.NewTransactionValidator(transactionTypeRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService, transactionValidator)

	requestBody := dto.TransactionCreateDTO{
		AccountID:       strconv.FormatInt(account.ID, 10),
		Amount:          1000.50,
		TransactionType: "2", // Purchase with Installments
	}
	body, _ := json.Marshal(requestBody)

	req := httptest.NewRequest("POST", "/transactions", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	transactionHandler.CreateTransaction(w, req)

	// Assert
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	var response dto.TransactionResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.TransactionType != "2" {
		t.Errorf("Expected operation type ID 2, got %s", response.TransactionType)
	}
}

// TestCreateTransaction_Withdrawal tests creating a withdrawal transaction
func TestCreateTransaction_Withdrawal(t *testing.T) {
	// Arrange
	mockAccountRepo := NewMockAccountRepository()
	mockTransactionRepo := NewMockTransactionRepository()

	// Create an account
	account, _ := mockAccountRepo.Insert(models.Account{DocumentNumber: "123456789"})

	transactionService := services.NewTransactionService(
		mockTransactionRepo,
		mockAccountRepo,
	)

	transactionTypeRepo := NewMockTransactionTypeRepository()
	transactionValidator := validator.NewTransactionValidator(transactionTypeRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService, transactionValidator)

	requestBody := dto.TransactionCreateDTO{
		AccountID:       strconv.FormatInt(account.ID, 10),
		Amount:          50.00,
		TransactionType: "3", // Withdrawal
	}
	body, _ := json.Marshal(requestBody)

	req := httptest.NewRequest("POST", "/transactions", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	transactionHandler.CreateTransaction(w, req)

	// Assert
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	var response dto.TransactionResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.TransactionType != "3" {
		t.Errorf("Expected operation type ID 3, got %s", response.TransactionType)
	}
}

// TestCreateTransaction_CreditVoucher tests creating a credit voucher transaction
func TestCreateTransaction_CreditVoucher(t *testing.T) {
	// Arrange
	mockAccountRepo := NewMockAccountRepository()
	mockTransactionRepo := NewMockTransactionRepository()

	// Create an account
	account, _ := mockAccountRepo.Insert(models.Account{DocumentNumber: "123456789"})

	transactionService := services.NewTransactionService(
		mockTransactionRepo,
		mockAccountRepo,
	)

	transactionTypeRepo := NewMockTransactionTypeRepository()
	transactionValidator := validator.NewTransactionValidator(transactionTypeRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService, transactionValidator)

	requestBody := dto.TransactionCreateDTO{
		AccountID:       strconv.FormatInt(account.ID, 10),
		Amount:          200.00,
		TransactionType: "4", // Credit Voucher
	}
	body, _ := json.Marshal(requestBody)

	req := httptest.NewRequest("POST", "/transactions", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	transactionHandler.CreateTransaction(w, req)

	// Assert
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	var response dto.TransactionResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.TransactionType != "4" {
		t.Errorf("Expected operation type ID 4, got %s", response.TransactionType)
	}
}

// TestCreateTransaction_InvalidAmount tests transaction with invalid amount
func TestCreateTransaction_InvalidAmount(t *testing.T) {
	// Arrange
	mockAccountRepo := NewMockAccountRepository()
	mockTransactionRepo := NewMockTransactionRepository()

	// Create an account
	account, _ := mockAccountRepo.Insert(models.Account{DocumentNumber: "123456789"})

	transactionService := services.NewTransactionService(
		mockTransactionRepo,
		mockAccountRepo,
	)

	transactionTypeRepo := NewMockTransactionTypeRepository()
	transactionValidator := validator.NewTransactionValidator(transactionTypeRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService, transactionValidator)

	requestBody := dto.TransactionCreateDTO{
		AccountID:       strconv.FormatInt(account.ID, 10),
		Amount:          -100.00, // Invalid Amount
		TransactionType: "1",
	}
	body, _ := json.Marshal(requestBody)

	req := httptest.NewRequest("POST", "/transactions", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	transactionHandler.CreateTransaction(w, req)

	// Assert
	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("Expected status code %d, got %d", http.StatusUnprocessableEntity, w.Code)
	}
}

// TestCreateTransaction_AccountNotFound tests transaction for non-existent account
func TestCreateTransaction_AccountNotFound(t *testing.T) {
	// Arrange
	mockAccountRepo := NewMockAccountRepository()
	mockTransactionRepo := NewMockTransactionRepository()

	transactionService := services.NewTransactionService(
		mockTransactionRepo,
		mockAccountRepo,
	)

	transactionTypeRepo := NewMockTransactionTypeRepository()
	transactionValidator := validator.NewTransactionValidator(transactionTypeRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService, transactionValidator)

	requestBody := dto.TransactionCreateDTO{
		AccountID:       "999", // Invalid Account
		Amount:          100.00,
		TransactionType: "1",
	}
	body, _ := json.Marshal(requestBody)

	req := httptest.NewRequest("POST", "/transactions", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	transactionHandler.CreateTransaction(w, req)

	// Assert
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
	}
}

// TestCreateTransaction_InvalidOperationType tests transaction with invalid operation type
func TestCreateTransaction_InvalidOperationType(t *testing.T) {
	// Arrange
	mockAccountRepo := NewMockAccountRepository()
	mockTransactionRepo := NewMockTransactionRepository()

	// Create an account
	account, _ := mockAccountRepo.Insert(models.Account{DocumentNumber: "123456789"})

	transactionService := services.NewTransactionService(
		mockTransactionRepo,
		mockAccountRepo,
	)

	transactionTypeRepo := NewMockTransactionTypeRepository()
	transactionValidator := validator.NewTransactionValidator(transactionTypeRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService, transactionValidator)

	requestBody := dto.TransactionCreateDTO{
		AccountID:       strconv.FormatInt(account.ID, 10),
		Amount:          100.00,
		TransactionType: "99", // Invalid Operation type
	}
	body, _ := json.Marshal(requestBody)

	req := httptest.NewRequest("POST", "/transactions", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	transactionHandler.CreateTransaction(w, req)

	// Assert
	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("Expected status code %d, got %d", http.StatusUnprocessableEntity, w.Code)
	}
}

// TestCreateTransaction_MalformedJSON tests transaction creation with malformed JSON
func TestCreateTransaction_MalformedJSON(t *testing.T) {
	// Arrange
	mockAccountRepo := NewMockAccountRepository()
	mockTransactionRepo := NewMockTransactionRepository()

	transactionService := services.NewTransactionService(
		mockTransactionRepo,
		mockAccountRepo,
	)

	transactionTypeRepo := NewMockTransactionTypeRepository()
	transactionValidator := validator.NewTransactionValidator(transactionTypeRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService, transactionValidator)

	malformedBody := []byte(`{invalid json}`)

	req := httptest.NewRequest("POST", "/transactions", bytes.NewReader(malformedBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	transactionHandler.CreateTransaction(w, req)

	// Assert
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}
}

// TestMultipleTransactions tests creating multiple transactions for the same account
func TestMultipleTransactions(t *testing.T) {
	// Arrange
	mockAccountRepo := NewMockAccountRepository()
	mockTransactionRepo := NewMockTransactionRepository()

	// Create an account
	account, _ := mockAccountRepo.Insert(models.Account{DocumentNumber: "123456789"})

	transactionService := services.NewTransactionService(
		mockTransactionRepo,
		mockAccountRepo,
	)

	// Act
	trans1, _ := transactionService.CreateTransaction(dto.TransactionCreateDTO{
		AccountID:       strconv.FormatInt(account.ID, 10),
		Amount:          100.00,
		TransactionType: "1",
	})
	trans2, _ := transactionService.CreateTransaction(dto.TransactionCreateDTO{
		AccountID:       strconv.FormatInt(account.ID, 10),
		Amount:          50.00,
		TransactionType: "3",
	})
	trans3, _ := transactionService.CreateTransaction(dto.TransactionCreateDTO{
		AccountID:       strconv.FormatInt(account.ID, 10),
		Amount:          200.00,
		TransactionType: "4",
	})

	// Assert
	if trans1.TransactionId != "1" {
		t.Errorf("Expected transaction 1 ID to be 1, got %s", trans1.TransactionId)
	}

	if trans2.TransactionId != "2" {
		t.Errorf("Expected transaction 2 ID to be 2, got %s", trans2.TransactionId)
	}

	if trans3.TransactionId != "3" {
		t.Errorf("Expected transaction 3 ID to be 3, got %s", trans3.TransactionId)
	}

	if trans1.Amount != -100.00 {
		t.Errorf("Expected transaction 1 amount 100.00, got %f", trans1.Amount)
	}

	if trans2.Amount != -50.00 {
		t.Errorf("Expected transaction 2 amount 50.00, got %f", trans2.Amount)
	}

	if trans3.Amount != 200.00 {
		t.Errorf("Expected transaction 3 amount 200.00, got %f", trans3.Amount)
	}
}

// TestTransactionZeroAmount tests transaction with zero amount
func TestTransactionZeroAmount(t *testing.T) {
	// Arrange
	mockAccountRepo := NewMockAccountRepository()
	mockTransactionRepo := NewMockTransactionRepository()

	// Create an account
	account, _ := mockAccountRepo.Insert(models.Account{DocumentNumber: "123456789"})

	transactionService := services.NewTransactionService(
		mockTransactionRepo,
		mockAccountRepo,
	)

	transactionTypeRepo := NewMockTransactionTypeRepository()
	transactionValidator := validator.NewTransactionValidator(transactionTypeRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService, transactionValidator)

	requestBody := dto.TransactionCreateDTO{
		AccountID:       strconv.FormatInt(account.ID, 10),
		Amount:          0.00,
		TransactionType: "1",
	}
	body, _ := json.Marshal(requestBody)

	req := httptest.NewRequest("POST", "/transactions", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	transactionHandler.CreateTransaction(w, req)

	// Assert
	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("Expected status code %d, got %d", http.StatusUnprocessableEntity, w.Code)
	}
}
