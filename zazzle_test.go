package gozazzle

import (
	"fmt"
	"testing"
)

func TestNewSession(t *testing.T){
	actual, err := NewSession(expected["Id"], expected["Secret"])
	if err != nil {
		t.Error(err.Error())
	}

	if actual.BaseUrl != expected["BaseUrl"] {
		t.Errorf("BaseUrl mismatch: actual: %s expected: %s", actual.BaseUrl, expected["BaseUrl"])
	}

	if actual.Id != expected["Id"] {
		t.Errorf("Id mismatch: actual: %s expected: %s", actual.Id, expected["Id"])
	}

	if actual.Secret != expected["Secret"] {
		t.Errorf("Secret mismatch: actual: %s expected: %s", actual.Secret, expected["Secret"])
	}

	if actual.Hash != expected["Hash"] {
		t.Errorf("Hash mismatch: actual: %s expected: %s", actual.Hash, expected["Hash"])
	}
}

func TestFetch(t *testing.T){
	session, err := NewSession(expected["Id"], expected["Secret"])
	if err != nil {
		t.Error(err.Error())
	}
	method := "listneworders"
	if err := session.fetch(method); err != nil {
		t.Error(err.Error())
	}
	resp := session.GetResponse()
	if resp.Code == "ERROR" && err == nil {
		msg := "ERROR api code detected but fetch returned nil error"
		msg = fmt.Sprintf("%s\n%#v\n", msg, resp)
	}
	if resp.Code != "SUCCESS" && resp.Code != "ERROR" {
		t.Errorf("Code mismatch: %v", resp.Code)
	}
}

func TestListNewOrders(t *testing.T){
	session, err := NewSession(expected["Id"], expected["Secret"])
	if err != nil {
		t.Error(err.Error())
	}
	if err := session.ListNewOrders(); err != nil {
		t.Error(err.Error())
	}
	resp := session.GetResponse()
	fmt.Printf("Orders found: %d\n", len(resp.Orders))
	if len(resp.Orders) > 0 {
		for _, order := range resp.Orders {
			fmt.Printf("%v\n", order.OrderId)
			fmt.Printf("%q\n", order.ShippingAddress.City)
		}
	}
}

func TestListUpdatedOrders(t *testing.T){
	session, err := NewSession(expected["Id"], expected["Secret"])
	if err != nil {
		t.Error(err.Error())
	}
	if err := session.ListUpdatedOrders(); err != nil {
		t.Error(err.Error())
	}
}

func TestGetOrder(t *testing.T){
	session, err := NewSession(expected["Id"], expected["Secret"])
	if err != nil {
		t.Error(err.Error())
	}
	order, err := session.GetOrder(testOrder.OrderId)
	if err != nil {
		t.Error(err.Error())
	}

	if order.OrderId != testOrder.OrderId {
		t.Errorf("OrderId mismatch: %s", order.OrderId)
	}
	if order.OrderDate != testOrder.OrderDate {
		t.Errorf("OrderDate mismatch: %s", order.OrderDate)
	}
	if order.OrderType != testOrder.OrderType {
		t.Errorf("OrderType mismatch: %s", order.OrderType)
	}
	if order.DeliveryMethod != testOrder.DeliveryMethod {
		t.Errorf("DeliveryMethod mismatch: %s", order.DeliveryMethod)
	}
	if order.Priority != testOrder.Priority {
		t.Errorf("Priority mismatch: %s", order.Priority)
	}
	if order.Currency != testOrder.Currency {
		t.Errorf("Currency mismatch: %s", order.Currency)
	}
	if order.Status != testOrder.Status {
		t.Errorf("Status mismatch: %s", order.Status)
	}
	if order.Attributes != testOrder.Attributes {
		t.Errorf("Attributes mismatch: %s", order.Attributes)
	}

	actualAddress := order.ShippingAddress
	expectedAddress := testOrder.ShippingAddress
	if actualAddress.Address1 != expectedAddress.Address1 {
		t.Errorf("Address1 mismatch: %s", actualAddress.Address1)
	}

	actualAddress = order.ShippingAddress
	expectedAddress = testOrder.ShippingAddress
	if actualAddress.Address2 != expectedAddress.Address2 {
		t.Errorf("Address2 mismatch: %s", actualAddress.Address2)
	}

	actualAddress = order.ShippingAddress
	expectedAddress = testOrder.ShippingAddress
	if actualAddress.Address3 != expectedAddress.Address3 {
		t.Errorf("Address3 mismatch: %s", actualAddress.Address3)
	}

	actualAddress = order.ShippingAddress
	expectedAddress = testOrder.ShippingAddress
	if actualAddress.Name != expectedAddress.Name {
		t.Errorf("Name mismatch: %s", actualAddress.Name)
	}

	actualAddress = order.ShippingAddress
	expectedAddress = testOrder.ShippingAddress
	if actualAddress.Name2 != expectedAddress.Name2 {
		t.Errorf("Name2 mismatch: %s", actualAddress.Name2)
	}

	actualAddress = order.ShippingAddress
	expectedAddress = testOrder.ShippingAddress
	if actualAddress.City != expectedAddress.City {
		t.Errorf("City mismatch: %s", actualAddress.City)
	}

	actualAddress = order.ShippingAddress
	expectedAddress = testOrder.ShippingAddress
	if actualAddress.State != expectedAddress.State {
		t.Errorf("State mismatch: %s", actualAddress.State)
	}

	actualAddress = order.ShippingAddress
	expectedAddress = testOrder.ShippingAddress
	if actualAddress.Country != expectedAddress.Country {
		t.Errorf("Country mismatch: %s", actualAddress.Country)
	}

	actualAddress = order.ShippingAddress
	expectedAddress = testOrder.ShippingAddress
	if actualAddress.CountryCode != expectedAddress.CountryCode {
		t.Errorf("CountryCode mismatch: %s", actualAddress.CountryCode)
	}

	actualAddress = order.ShippingAddress
	expectedAddress = testOrder.ShippingAddress
	if actualAddress.Zip != expectedAddress.Zip {
		t.Errorf("Zip mismatch: %s", actualAddress.Zip)
	}

	actualAddress = order.ShippingAddress
	expectedAddress = testOrder.ShippingAddress
	if actualAddress.Phone != expectedAddress.Phone {
		t.Errorf("Phone mismatch: %s", actualAddress.Phone)
	}

	actualAddress = order.ShippingAddress
	expectedAddress = testOrder.ShippingAddress
	if actualAddress.Email != expectedAddress.Email {
		t.Errorf("Email mismatch: %s", actualAddress.Email)
	}

	actualAddress = order.ShippingAddress
	expectedAddress = testOrder.ShippingAddress
	if actualAddress.Type != expectedAddress.Type {
		t.Errorf("Type mismatch: %s", actualAddress.Type)
	}
}

func TestListOrderMessages(t *testing.T) {
        session, err := NewSession(expected["Id"], expected["Secret"])
        if err != nil {
                t.Error(err.Error())
        }
        messages, err := session.ListOrderMessages()
	for _, msg := range messages {
		fmt.Printf("%#v\n", msg)
	}
}

func TestUrlEncode(t *testing.T) {
	raw := "this is raw text!"
	expectedResult := "this+is+raw+text%21"
	encoded, err := urlEncode(raw)
	if err != nil {
		t.Error(err.Error())
	}
	if encoded != expectedResult {
		t.Errorf("encoded does not match: %s\n", encoded)
	}
	// test 200 char limit
	tooLong := "aabbccddeeffgghhiijjkkllmmnnooppqqrrssttuuvvwwxxyyzz"
	tooLong = tooLong + "aabbccddeeffgghhiijjkkllmmnnooppqqrrssttuuvvwwxxyyzz"
	tooLong = tooLong + "aabbccddeeffgghhiijjkkllmmnnooppqqrrssttuuvvwwxxyyzz"
	tooLong = tooLong + "aabbccddeeffgghhiijjkkllmmnnooppqqrrssttuuvvwwxxyyzz"
	var errTooLong error
	encoded, errTooLong = urlEncode(tooLong)
	if errTooLong == nil {
		t.Errorf("errTooLong did not generate error")
	}
}

func TestLicense(t *testing.T) {
	msg := License()
	fmt.Printf("%s", msg)
}
