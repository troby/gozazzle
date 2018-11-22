package gozazzle

func (session *Session) ListNewOrders() error {
	sig := session.Id + session.Secret
	session.Hash = hash(sig)
	if err := session.fetch("listneworders"); err != nil {
		return err
	}
	return nil
}

func (session *Session) ListUpdatedOrders() error {
	sig := session.Id + session.Secret
	session.Hash = hash(sig)
	if err := session.fetch("listupdatedorders"); err != nil {
		return err
	}
	return nil
}

func (session *Session) GetResponse() *Response {
	resp := session.Response
	return resp
}

func (session *Session) GetOrder(orderId string) (Order, error) {
	sig := session.Id + orderId + session.Secret
	session.Hash = hash(sig)
	url := "getorder&orderid=" + orderId
	if err := session.fetch(url); err != nil {
		return Order{}, err
	}
	order := session.Response.Orders[0]
	return order, nil
}

func (session *Session) GetShippingLabel(orderId string, weight string, format string) (ZazzleDocument, error) {
	switch format {
	case "PDF":
	case "ZPL":
	case "PNG":
	default:
		format = "PDF"
	}
	sig := session.Id + orderId + weight + format + session.Secret
	session.Hash = hash(sig)
	url := "getshippinglabel&orderid=" + orderId + "&weight=" + weight + "&format=" + format
	if err := session.fetch(url); err != nil {
		return ZazzleDocument{}, err
	}
	shippingDoc := session.Response.ShippingDocument
	return shippingDoc, nil
}

func (session *Session) AckNewOrder(orderId string) error {
	sig := session.Id + orderId + "new" + session.Secret
	session.Hash = hash(sig)
	url := "ackorder&orderid=" + orderId + "&type=new&action=accept"
	if err := session.fetch(url); err != nil {
		return err
	}
	return nil
}

func (session *Session) ListOrderMessages() ([]Message, error) {
	sig := session.Id + session.Secret
	session.Hash = hash(sig)
	url := "listordermessages"
	if err := session.fetch(url); err != nil {
		return []Message{}, err
	}
	messages := session.Response.Messages
	return messages, nil
}

func (session *Session) AddOrderActivity(orderId string, activity string) error {
	encoded, err := urlEncode(activity)
	if err != nil {
		return err
	}
	sig := session.Id + orderId + encoded + session.Secret
	session.Hash = hash(sig)
	url := "addorderactivity&orderid=" + orderId + "&activity=" + encoded
	if err := session.fetch(url); err != nil {
		return err
	}
	return nil
}
