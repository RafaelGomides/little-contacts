package controllers

import (
	"encoding/json"
	"errors"
	"little-contacts/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

//  ContactController operations for Contact
type ContactController struct {
	beego.Controller
}

// URLMapping ...
func (c *ContactController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("SendEmail", c.SendEmail)
}

// SendEmail ...
// @Title SendEmail
// @Description Envia um email com as informações passadas
// @Param	body		body 	models.Contact	true		"body for Contact content"
// @Success 201 {int} models.Contact
// @Failure 500 alguma coisa aconteceu
// @router /email [post]
func (c *ContactController) SendEmail() {
	var v models.Contact
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := v.SendEmail(); err == nil {
		c.Ctx.Output.SetStatus(201)
	} else {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = err.Error()
	}
}

//
// Post ...
// @Title Post
// @Description create Contact
// @Param	body		body 	models.Contact	true		"body for Contact content"
// @Success 201 {int} models.Contact
// @Failure 403 body is empty
// @router / [post]
func (c *ContactController) Post() {
	var v models.Contact
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddContact(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Contact by ID
// @Param	ID		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Contact
// @Failure 403 :ID is empty
// @router /:ID [get]
func (c *ContactController) GetOne() {
	IDStr := c.Ctx.Input.Param(":ID")
	ID, _ := strconv.ParseInt(IDStr, 0, 64)
	v, err := models.GetContactByID(ID)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Contact
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Contact
// @Failure 403
// @router / [get]
func (c *ContactController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalID query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllContact(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Contact
// @Param	ID		path 	string	true		"The ID you want to update"
// @Param	body		body 	models.Contact	true		"body for Contact content"
// @Success 200 {object} models.Contact
// @Failure 403 :ID is not int
// @router /:ID [put]
func (c *ContactController) Put() {
	IDStr := c.Ctx.Input.Param(":ID")
	ID, _ := strconv.ParseInt(IDStr, 0, 64)
	v := models.Contact{ID: ID}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateContactByID(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Contact
// @Param	ID		path 	string	true		"The ID you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 ID is empty
// @router /:ID [delete]
func (c *ContactController) Delete() {
	IDStr := c.Ctx.Input.Param(":ID")
	ID, _ := strconv.ParseInt(IDStr, 0, 64)
	if err := models.DeleteContact(ID); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
