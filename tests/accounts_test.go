package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"transaction-service/dto"
	"transaction-service/handlers"
	"transaction-service/models"
	"transaction-service/services"
)

func newReqWithParam(method, url, key, val string, body io.Reader) *http.Request {
	req := httptest.NewRequest(method, url, body)
	req.SetPathValue(key, val)
	return req
}

// MockAccountRepository is a mock implementation of the AccountRepository interface
type MockAccountRepository struct {
	accounts map[int64]models.Account
	nextID   int64
}

func NewMockAccountRepository() *MockAccountRepository {
	return &MockAccountRepository{
		accounts: make(map[int64]models.Account),
		nextID:   1,
	}
}

func (m *MockAccountRepository) Insert(account models.Account) (models.Account, error) {
	account.ID = m.nextID
	account.CreatedAt = time.Now()
	m.accounts[m.nextID] = account
	m.nextID++
	return account, nil
}

func (m *MockAccountRepository) FindByDocumentNumber(documentNumber string) (models.Account, error) {
	for _, account := range m.accounts {
		if account.DocumentNumber == documentNumber {
			return account, nil
		}
	}
	return models.Account{}, errors.New("account not found")
}

func (m *MockAccountRepository) FindByID(id string) (models.Account, error) {
	accountID, _ := strconv.ParseInt(id, 10, 64)
	account, exists := m.accounts[accountID]
	if !exists {
		return models.Account{}, errors.New("account not found")
	}
	return account, nil
}

// TestCreateAccount_Success tests successful account creation
func TestCreateAccount_Success(t *testing.T) {
	// Arrange
	mockRepo := NewMockAccountRepository()
	accountService := services.NewAccountService(mockRepo)
	accountHandler := handlers.NewAccountHandler(accountService)

	requestBody := dto.CreateAccountRequest{
		DocumentNumber: "123456789",
	}
	body, _ := json.Marshal(requestBody)

	req := httptest.NewRequest("POST", "/accounts", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	accountHandler.CreateAccount(w, req)

	// Assert
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	var response dto.AccountResponse
	json.Unmarshal(w.Body.Bytes(), &response)
	if response.DocumentNumber != "123456789" {
		t.Errorf("Expected document number '123456789', got '%s'", response.DocumentNumber)
	}

	if response.AccountID == "0" {
		t.Errorf("Expected non-zero account ID, got %s", response.AccountID)
	}
}

// TestCreateAccount_InvalidInput tests account creation with invalid input
func TestCreateAccount_InvalidInput(t *testing.T) {
	// Arrange
	mockRepo := NewMockAccountRepository()
	accountService := services.NewAccountService(mockRepo)
	accountHandler := handlers.NewAccountHandler(accountService)

	requestBody := dto.CreateAccountRequest{
		DocumentNumber: "", // Empty document number
	}
	body, _ := json.Marshal(requestBody)

	req := httptest.NewRequest("POST", "/accounts", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	accountHandler.CreateAccount(w, req)

	// Assert
	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("Expected status code %d, got %d", http.StatusUnprocessableEntity, w.Code)
	}
}

// TestGetAccount_Success tests successful account retrieval
func TestGetAccount_Success(t *testing.T) {
	// Arrange
	mockRepo := NewMockAccountRepository()
	accountService := services.NewAccountService(mockRepo)
	accountHandler := handlers.NewAccountHandler(accountService)

	// Create an account first
	testAccount := models.Account{
		DocumentNumber: "987654321",
	}
	createdAccount, _ := mockRepo.Insert(testAccount)

	req := newReqWithParam("GET", "/accounts/1", "accountId", "1", nil)
	w := httptest.NewRecorder()

	// Act
	accountHandler.GetAccountByID(w, req)

	// Assert
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response dto.AccountResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	if response.AccountID != strconv.FormatInt(createdAccount.ID, 10) {
		t.Errorf("Expected account ID %d, got %s", createdAccount.ID, response.AccountID)
	}

	if response.DocumentNumber != createdAccount.DocumentNumber {
		t.Errorf("Expected document number '%s', got '%s'", createdAccount.DocumentNumber, response.DocumentNumber)
	}
}

// TestGetAccount_NotFound tests account retrieval when account doesn't exist
func TestGetAccount_NotFound(t *testing.T) {
	// Arrange
	mockRepo := NewMockAccountRepository()
	accountService := services.NewAccountService(mockRepo)
	accountHandler := handlers.NewAccountHandler(accountService)

	//req := httptest.NewRequest("GET", "/accounts/999", nil)
	req := newReqWithParam("GET", "/accounts/999", "accountId", "999", nil)
	w := httptest.NewRecorder()

	// Act
	accountHandler.GetAccountByID(w, req)

	// Assert
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
	}
}

// TestCreateAccount_MalformedJSON tests account creation with malformed JSON
func TestCreateAccount_MalformedJSON(t *testing.T) {
	// Arrange
	mockRepo := NewMockAccountRepository()
	accountService := services.NewAccountService(mockRepo)
	accountHandler := handlers.NewAccountHandler(accountService)

	malformedBody := []byte(`{invalid json}`)

	req := httptest.NewRequest("POST", "/accounts", bytes.NewReader(malformedBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	accountHandler.CreateAccount(w, req)

	// Assert
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}
}

// TestMultipleAccountCreation tests creating multiple accounts
func TestMultipleAccountCreation(t *testing.T) {
	// Arrange
	mockRepo := NewMockAccountRepository()
	accountService := services.NewAccountService(mockRepo)

	// Act
	account1, _ := accountService.CreateAccount(dto.CreateAccountRequest{
		DocumentNumber: "111111111",
	})
	account2, _ := accountService.CreateAccount(dto.CreateAccountRequest{
		DocumentNumber: "222222222",
	})
	account3, _ := accountService.CreateAccount(dto.CreateAccountRequest{
		DocumentNumber: "333333333",
	})

	// Assert
	if account1.AccountID != "1" {
		t.Errorf("Expected account 1 ID to be 1, got %s", account1.AccountID)
	}

	if account2.AccountID != "2" {
		t.Errorf("Expected account 2 ID to be 2, got %s", account2.AccountID)
	}

	if account3.AccountID != "3" {
		t.Errorf("Expected account 3 ID to be 3, got %s", account3.AccountID)
	}

	if account1.DocumentNumber != "111111111" {
		t.Errorf("Expected document number '111111111', got '%s'", account1.DocumentNumber)
	}
}
