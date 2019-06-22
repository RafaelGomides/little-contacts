package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"reflect"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Contact struct {
	ID        int64  `orm:"auto"`
	Name      string `orm:"size(128)"`
	Cellphone string `orm:"size(128)"`
	Facebook  string `orm:"size(128)"`
	Github    string `orm:"size(128)"`
	Twitter   string `orm:"size(128)"`
	Email     string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Contact))
}

// AddContact insert a new Contact into database and returns
// last inserted ID on success.
func AddContact(m *Contact) (ID int64, err error) {
	o := orm.NewOrm()
	ID, err = o.Insert(m)
	return
}

// GetContactByID retrieves Contact by ID. Returns error if
// ID doesn't exist
func GetContactByID(ID int64) (v *Contact, err error) {
	o := orm.NewOrm()
	v = &Contact{ID: ID}
	if err = o.QueryTable(new(Contact)).Filter("ID", ID).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllContact retrieves all Contact matches certain condition. Returns empty list if
// no records exist
func GetAllContact(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Contact))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: InvalID order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: InvalID order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Contact
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateContact updates Contact by ID and returns error if
// the record to be updated doesn't exist
func UpdateContactByID(m *Contact) (err error) {
	o := orm.NewOrm()
	v := Contact{ID: m.ID}
	// ascertain ID exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteContact deletes Contact by ID and returns error if
// the record to be deleted doesn't exist
func DeleteContact(ID int64) (err error) {
	o := orm.NewOrm()
	v := Contact{ID: ID}
	// ascertain ID exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Contact{ID: ID}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//SendEmail envia um email para o endereço registrado
func (c *Contact) SendEmail() (val bool, err error) {
	from := beego.AppConfig.String("emailUser")
	pass := beego.AppConfig.String("emailPass")
	to := beego.AppConfig.String("emailUser")
	cc := c.Email

	body, _ := json.Marshal(c)

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		string(body)

	err = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to, cc}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return false, err
	}

	return true, nil
}
