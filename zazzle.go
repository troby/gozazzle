package gozazzle

import "encoding/xml"

type Session struct {
	BaseUrl		string
	Id		string
	Secret		string
	Hash		string
	Response	*Response
}

type Response struct {
	XMLName			xml.Name	`xml:"Response"`
	Code			string		`xml:"Status>Code"`
	Info			string		`xml:"Status>Info"`
	Orders			[]Order		`xml:"Result>Orders>Order"`
	ShippingDocument	ZazzleDocument	`xml:"Result>ShippingDocuments>ShippingDocument"`
	Messages		[]Message	`xml:"Result>Messages>Message"`
}

type Order struct {
	OrderId		string		`xml:"OrderId"`
	OrderDate	string		`xml:"OrderDate"`
	OrderType	string		`xml:"OrderType"`
	DeliveryMethod	string		`xml:"DeliveryMethod"`
	Priority	string		`xml:"Priority"`
	Currency	string		`xml:"Currency"`
	Status		string		`xml:"Status"`
	Attributes	string		`xml:"Attributes"`
	ShippingAddress	Address		`xml:"ShippingAddress"`
	LineItems	[]LineItem	`xml:"LineItems>LineItem"`
	Products	[]Product	`xml:"Products>Product"`
	PackingSheet	[]ZazzleFile	`xml:"PackingSheet>Page>Front"`
}

type Address struct {
	Address1	string		`xml:"Address1"`
	Address2	string		`xml:"Address2"`
	Address3	string		`xml:"Address3"`
	Name		string		`xml:"Name"`
	Name2		string		`xml:"Name2"`
	City		string		`xml:"City"`
	State		string		`xml:"State"`
	Country		string		`xml:"Country"`
	CountryCode	string		`xml:"CountryCode"`
	Zip		string		`xml:"Zip"`
	Phone		string		`xml:"Phone"`
	Email		string		`xml:"Email"`
	Type		string		`xml:"Type"`
}

type LineItem struct {
	LineItemId		string		`xml:"LineItemId"`
	OrderId			string		`xml:"OrderId"`
	LineItemType		string		`xml:"LineItemType"`
	Quantity		string		`xml:"Quantity"`
	Description		string		`xml:"Description"`
	Attributes		string		`xml:"Attributes"`
	VendorAttributes	string		`xml:"VendorAttributes"`
	Price			string		`xml:"Price"`
	ProductId		string		`xml:"ProductId"`
	PrintFiles		[]ZazzleFile	`xml:"PrintFiles>PrintFile"`
	Previews		[]ZazzleFile	`xml:"Previews>PreviewFile"`
}

type ZazzleFile struct {
	Type		string	`xml:"Type"`
	Url		string	`xml:"Url"`
	Description	string	`xml:"Description"`
}

type ZazzleDocument struct {
	ZazzleFile
	Format		string	`xml:"Format"`
}

type Product struct {
	ProductId	string	`xml:"ProductId"`
	ProductInfo	string	`xml:"ProductInfo"`
}

type ShippingInfo struct {
	Carrier			string		`xml:"Carrier"`
	Method			string		`xml:"Method"`
	TrackingNumber		string		`xml:"TrackingNumber"`
	Weight			string		`xml:"Weight"`
	ShippingDocument	ZazzleDocument	`xml:"ShippingDocuments>ShippingDocument"`
}

type Message struct {
	OrderId		string	`xml:"OrderId"`
	Text		string	`xml:"Text"`
	MessageDate	string	`xml:"MessageDate"`
}

func NewSession(id string, secret string) (*Session, error) {
	session		:= new(Session)
	session.BaseUrl	= "https://vendor.zazzle.com/v100/api.aspx?"
	session.Id	= id
	session.Secret	= secret
	session.Hash	= hash(id + secret)
	return session, nil
}

func License() string {
	lic :=      "#########################################################################\n"
	lic = lic + "# Copyright (c) 2018, Ted Roby (git at sranet dot com)                  #\n"
	lic = lic + "# All rights reserved.                                                  #\n"
	lic = lic + "#                                                                       #\n"
	lic = lic + "# Redistribution and use in source and binary forms, with or without    #\n"
	lic = lic + "#       modification, are permitted provided that the following         #\n"
	lic = lic + "#       conditions are met:                                             #\n"
	lic = lic + "#     * Redistributions of source code must retain the above copyright  #\n"
	lic = lic + "#       notice, this list of conditions and the following disclaimer.   #\n"
	lic = lic + "#     * Redistributions in binary form must reproduce the above         #\n"
	lic = lic + "#       copyright notice, this list of conditions and the following     #\n"
	lic = lic + "#       disclaimer in the documentation and/or other materials provided #\n"
	lic = lic + "#       with the distribution.                                          #\n"
	lic = lic + "#     * The name of Ted Roby and the names of its contributors may not  #\n"
	lic = lic + "#       be used to endorse or promote products derived from this        #\n"
	lic = lic + "#       software without specific prior written permission.             #\n"
	lic = lic + "#                                                                       #\n"
	lic = lic + "# THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS   #\n"
	lic = lic + "# \"AS IS\" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT     #\n"
	lic = lic + "# LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR #\n"
	lic = lic + "# A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL TED ROBY BE    #\n"
	lic = lic + "# LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR   #\n"
	lic = lic + "# CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF  #\n"
	lic = lic + "# SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR       #\n"
	lic = lic + "# BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, #\n"
	lic = lic + "# WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE  #\n"
	lic = lic + "# OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE,     #\n"
	lic = lic + "# EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.                    #\n"
	lic = lic + "#########################################################################\n"
	return lic
}
