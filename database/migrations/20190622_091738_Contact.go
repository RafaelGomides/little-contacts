package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Contact_20190622_091738 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Contact_20190622_091738{}
	m.Created = "20190622_091738"

	migration.Register("Contact_20190622_091738", m)
}

// Run the migrations
func (m *Contact_20190622_091738) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE Contact(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,`cellphone` varchar(128) NOT NULL,`facebook` varchar(128) NOT NULL,`githun` varchar(128) NOT NULL,`twitter` varchar(128) NOT NULL,`email` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Contact_20190622_091738) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `Contact`")
}
