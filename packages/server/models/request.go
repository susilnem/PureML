package models

import (
	"encoding/json"
	"mime/multipart"

	uuid "github.com/satori/go.uuid"
)

type Request struct {
	User        *UserClaims
	Org         *OrganizationHandleResponse
	Body        []byte
	ParsedBody  map[string]interface{}
	Headers     map[string]string
	PathParams  map[string]string
	QueryParams map[string]string
	FormFiles   map[string][]*multipart.FileHeader
	FormValues  map[string][]string
}

func (r *Request) GetUserMail() string {
	return r.User.Email
}

func (r *Request) GetOrgId() uuid.UUID {
	return r.Org.UUID
}

func (r *Request) GetPathParam(param string) string {
	return getValueFromMap(r.PathParams, param)
}

func (r *Request) GetQueryParam(param string) string {
	return getValueFromMap(r.QueryParams, param)
}

func (r *Request) ParseJsonBody() {
	body := map[string]interface{}{}
	err := json.Unmarshal(r.Body, &body)
	if err != nil {
		panic(err)
	}
	r.ParsedBody = body
}

func (r *Request) GetParsedBodyAttribute(attributeName string) interface{} {
	if len(r.ParsedBody) == 0 {
		return nil
	}
	if val, ok := r.ParsedBody[attributeName]; ok {
		return val
	}
	return nil
}

func (r *Request) GetFormFile(formKeyName string) *multipart.FileHeader {
	if val, ok := r.FormFiles[formKeyName]; !ok || len(val) == 0 {
		return nil
	} else {
		return val[0]
	}
}

func getValueFromMap(m map[string]string, param string) string {
	if val, ok := m[param]; ok {
		return val
	}
	return ""
}
