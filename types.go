package main

import "encoding/json"

func UnmarshalHar(data []byte) (Har, error) {
	var r Har
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Har) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Har struct {
	Log Log `json:"log"`
}

type Log struct {
	Version string  `json:"version"`
	Creator Creator `json:"creator"`
	Pages   []Page  `json:"pages"`
	Entries []Entry `json:"entries"`
}

type Creator struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Entry struct {
	Initiator         Initiator           `json:"_initiator"`
	Priority          *Priority           `json:"_priority"`
	ResourceType      ResourceType        `json:"_resourceType"`
	Cache             Cache               `json:"cache"`
	Connection        *string             `json:"connection,omitempty"`
	Pageref           Pageref             `json:"pageref"`
	Request           Request             `json:"request"`
	Response          Response            `json:"response"`
	ServerIPAddress   ServerIPAddressEnum `json:"serverIPAddress"`
	StartedDateTime   string              `json:"startedDateTime"`
	Time              float64             `json:"time"`
	Timings           Timings             `json:"timings"`
	FromCache         *FromCache          `json:"_fromCache,omitempty"`
	WebSocketMessages []interface{}       `json:"_webSocketMessages,omitempty"`
}

type Cache struct {
}

type Initiator struct {
	Type       Type    `json:"type"`
	URL        *string `json:"url,omitempty"`
	LineNumber *int64  `json:"lineNumber,omitempty"`
	Stack      *Stack  `json:"stack,omitempty"`
}

type Stack struct {
	CallFrames []CallFrame  `json:"callFrames"`
	Parent     *StackParent `json:"parent,omitempty"`
	ParentID   *ParentID    `json:"parentId,omitempty"`
}

type CallFrame struct {
	FunctionName string `json:"functionName"`
	ScriptID     string `json:"scriptId"`
	URL          string `json:"url"`
	LineNumber   int64  `json:"lineNumber"`
	ColumnNumber int64  `json:"columnNumber"`
}

type StackParent struct {
	Description Description   `json:"description"`
	CallFrames  []CallFrame   `json:"callFrames"`
	Parent      *PurpleParent `json:"parent,omitempty"`
	ParentID    *ParentID     `json:"parentId,omitempty"`
}

type PurpleParent struct {
	Description Description   `json:"description"`
	CallFrames  []CallFrame   `json:"callFrames"`
	Parent      *FluffyParent `json:"parent,omitempty"`
	ParentID    *ParentID     `json:"parentId,omitempty"`
}

type FluffyParent struct {
	Description Description      `json:"description"`
	CallFrames  []CallFrame      `json:"callFrames"`
	Parent      *TentacledParent `json:"parent,omitempty"`
	ParentID    *ParentID        `json:"parentId,omitempty"`
}

type TentacledParent struct {
	Description Description   `json:"description"`
	CallFrames  []CallFrame   `json:"callFrames"`
	Parent      *StickyParent `json:"parent,omitempty"`
	ParentID    *ParentID     `json:"parentId,omitempty"`
}

type StickyParent struct {
	Description Description   `json:"description"`
	CallFrames  []CallFrame   `json:"callFrames"`
	ParentID    *ParentID     `json:"parentId,omitempty"`
	Parent      *IndigoParent `json:"parent,omitempty"`
}

type IndigoParent struct {
	Description Description `json:"description"`
	CallFrames  []CallFrame `json:"callFrames"`
}

type ParentID struct {
	ID         string `json:"id"`
	DebuggerID string `json:"debuggerId"`
}

type Request struct {
	Method      Method          `json:"method"`
	URL         string          `json:"url"`
	HTTPVersion HTTPVersionEnum `json:"httpVersion"`
	Headers     []Header        `json:"headers"`
	QueryString []Header        `json:"queryString"`
	Cookies     []Cooky         `json:"cookies"`
	HeadersSize int64           `json:"headersSize"`
	BodySize    int64           `json:"bodySize"`
	PostData    *PostData       `json:"postData,omitempty"`
}

type Cooky struct {
	Name     Name      `json:"name"`
	Value    string    `json:"value"`
	Path     Path      `json:"path"`
	Domain   Domain    `json:"domain"`
	Expires  string    `json:"expires"`
	HTTPOnly bool      `json:"httpOnly"`
	Secure   bool      `json:"secure"`
	SameSite *SameSite `json:"sameSite,omitempty"`
}

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type PostData struct {
	MIMEType PostDataMIMEType `json:"mimeType"`
	Text     string           `json:"text"`
}

type Response struct {
	Status       int64           `json:"status"`
	StatusText   StatusTextEnum  `json:"statusText"`
	HTTPVersion  HTTPVersionEnum `json:"httpVersion"`
	Headers      []Header        `json:"headers"`
	Cookies      []interface{}   `json:"cookies"`
	Content      Content         `json:"content"`
	RedirectURL  string          `json:"redirectURL"`
	HeadersSize  int64           `json:"headersSize"`
	BodySize     int64           `json:"bodySize"`
	TransferSize int64           `json:"_transferSize"`
	Error        *Error          `json:"_error"`
}

type Content struct {
	Size        int64           `json:"size"`
	MIMEType    ContentMIMEType `json:"mimeType"`
	Text        *string         `json:"text,omitempty"`
	Encoding    *string         `json:"encoding,omitempty"`
	Compression *int64          `json:"compression,omitempty"`
}

type Timings struct {
	Blocked         float64 `json:"blocked"`
	DNS             float64 `json:"dns"`
	SSL             float64 `json:"ssl"`
	Connect         float64 `json:"connect"`
	Send            float64 `json:"send"`
	Wait            float64 `json:"wait"`
	Receive         float64 `json:"receive"`
	BlockedQueueing float64 `json:"_blocked_queueing"`
}

type Page struct {
	StartedDateTime string      `json:"startedDateTime"`
	ID              Pageref     `json:"id"`
	Title           string      `json:"title"`
	PageTimings     PageTimings `json:"pageTimings"`
}

type PageTimings struct {
	OnContentLoad float64 `json:"onContentLoad"`
	OnLoad        float64 `json:"onLoad"`
}

type FromCache string

const (
	Disk   FromCache = "disk"
	Memory FromCache = "memory"
)

type Description string

const (
	Await       Description = "await"
	Image       Description = "Image"
	Load        Description = "load"
	PromiseThen Description = "Promise.then"
	SetInterval Description = "setInterval"
	SetTimeout  Description = "setTimeout"
)

type Type string

const (
	Other      Type = "other"
	Parser     Type = "parser"
	TypeScript Type = "script"
)

type Pageref string

const (
	Page4 Pageref = "page_4"
)

type Priority string

const (
	High     Priority = "High"
	Low      Priority = "Low"
	VeryHigh Priority = "VeryHigh"
	VeryLow  Priority = "VeryLow"
)

type Domain string

const (
	BinanceCOM    Domain = ".binance.com"
	P2PBinanceCOM Domain = "p2p.binance.com"
)

type Name string

const (
	BNCFvKey                  Name = "BNC_FV_KEY"
	BNCFvKeyExpire            Name = "BNC_FV_KEY_EXPIRE"
	BNCFvKeyT                 Name = "BNC_FV_KEY_T"
	BNCLocation               Name = "BNC-Location"
	BNCUUID                   Name = "bnc-uuid"
	Campaign                  Name = "campaign"
	Cr00                      Name = "cr00"
	D1Og                      Name = "d1og"
	F30L                      Name = "f30l"
	Lang                      Name = "lang"
	Logined                   Name = "logined"
	OptanonAlertBoxClosed     Name = "OptanonAlertBoxClosed"
	OptanonConsent            Name = "OptanonConsent"
	P20T                      Name = "p20t"
	R2O1                      Name = "r2o1"
	SEGd                      Name = "se_gd"
	SEGsd                     Name = "se_gsd"
	SESD                      Name = "se_sd"
	Sensorsdata2015Jssdkcross Name = "sensorsdata2015jssdkcross"
	Source                    Name = "source"
	Theme                     Name = "theme"
	UserPreferredCurrency     Name = "userPreferredCurrency"
)

type Path string

const (
	Empty Path = "/"
)

type SameSite string

const (
	Lax SameSite = "Lax"
)

type HTTPVersionEnum string

const (
	H3          HTTPVersionEnum = "h3"
	HTTP11      HTTPVersionEnum = "HTTP/1.1"
	HTTP20      HTTPVersionEnum = "http/2.0"
	HTTPVersion HTTPVersionEnum = ""
)

type Method string

const (
	Get  Method = "GET"
	Post Method = "POST"
)

type PostDataMIMEType string

const (
	PurpleApplicationJSON PostDataMIMEType = "application/json"
	TextPlainCharsetUTF8  PostDataMIMEType = "text/plain;charset=UTF-8"
)

type ResourceType string

const (
	Document           ResourceType = "document"
	Fetch              ResourceType = "fetch"
	Font               ResourceType = "font"
	Ping               ResourceType = "ping"
	ResourceTypeImage  ResourceType = "image"
	ResourceTypeScript ResourceType = "script"
	Stylesheet         ResourceType = "stylesheet"
	Websocket          ResourceType = "websocket"
	Xhr                ResourceType = "xhr"
)

type ContentMIMEType string

const (
	ApplicationJavascript  ContentMIMEType = "application/javascript"
	ApplicationXJavascript ContentMIMEType = "application/x-javascript"
	BinaryOctetStream      ContentMIMEType = "binary/octet-stream"
	FluffyApplicationJSON  ContentMIMEType = "application/json"
	ImagePNG               ContentMIMEType = "image/png"
	ImageSVGXML            ContentMIMEType = "image/svg+xml"
	TextCSS                ContentMIMEType = "text/css"
	TextHTML               ContentMIMEType = "text/html"
	TextJavascript         ContentMIMEType = "text/javascript"
	XUnknown               ContentMIMEType = "x-unknown"
)

type Error string

const (
	NetERRBLOCKEDBYCLIENT Error = "net::ERR_BLOCKED_BY_CLIENT"
)

type StatusTextEnum string

const (
	StatusText         StatusTextEnum = ""
	SwitchingProtocols StatusTextEnum = "Switching Protocols"
)

type ServerIPAddressEnum string

const (
	ServerIPAddress ServerIPAddressEnum = ""
	The10418130236  ServerIPAddressEnum = "104.18.130.236"
	The10815810421  ServerIPAddressEnum = "108.158.104.21"
	The10815810433  ServerIPAddressEnum = "108.158.104.33"
	The10815810461  ServerIPAddressEnum = "108.158.104.61"
	The1816413107   ServerIPAddressEnum = "18.164.13.107"
	The1816413123   ServerIPAddressEnum = "18.164.13.123"
	The181641378    ServerIPAddressEnum = "18.164.13.78"
)
