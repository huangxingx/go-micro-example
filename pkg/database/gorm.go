package database

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/config/source/consul"
)

var db *gorm.DB

type dbInfo struct {
	Address      string `json:"address"`
	Port         int    `json:"port"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
	DbName       string `json:"db_name"`
}

type Model struct {
	//ID        uint `gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func InitDbByConsul(consulAddr string) {
	consulSource := consul.NewSource(consul.WithAddress(consulAddr))
	conf := config.NewConfig()

	// Load file source
	err := conf.Load(consulSource)
	if err != nil {
		log.Fatal(err)
	}
	var v dbInfo
	err = conf.Get("micro", "config", "database").Scan(&v)
	if err != nil {
		log.Fatal(err)
	}

	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		v.UserName, v.UserPassword, v.Address, v.Port, v.DbName))

	if err != nil {
		log.Fatal("failed to connect database：", err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		newTableName := strings.TrimSuffix(defaultTableName, "_model")
		return newTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func GetDb() *gorm.DB {
	return db
}

func EnableDbLog() {
	db.LogMode(true)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}

type JSON []byte

func (j JSON) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		errors.New("Invalid Scan Source")
	}
	*j = append((*j)[0:0], s...)
	return nil
}
func (m JSON) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}
func (m *JSON) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("null point exception")
	}
	*m = append((*m)[0:0], data...)
	return nil
}
func (j JSON) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}
func (j JSON) Equals(j1 JSON) bool {
	return bytes.Equal([]byte(j), []byte(j1))
}
