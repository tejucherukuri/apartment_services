package models

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gocql/gocql"
)

type AServiceOrm struct {
	orm.Ormer
}

func NewOrm() *AServiceOrm {
	db := orm.NewOrm()
	return &AServiceOrm{db}
}

func (this *AServiceOrm) InsertB(md interface{}) (i int64, e error) {
	curTime := time.Now()

	reflect.ValueOf(md).Elem().FieldByName("CreatedDate").Set(reflect.ValueOf(curTime))
	reflect.ValueOf(md).Elem().FieldByName("LastModified").Set(reflect.ValueOf(curTime))
	i, e = this.Insert(md)
	beego.Debug(e)
	return
}

func (this *AServiceOrm) UpdateB(md interface{}) (i int64, e error) {
	curTime := time.Now()
	reflect.ValueOf(md).Elem().FieldByName("LastModified").Set(reflect.ValueOf(curTime))
	i, e = this.Update(md)
	beego.Debug(e)
	return
}

var CassandraCluster *gocql.ClusterConfig
var CassandraSessions []*gocql.Session

func NewSession() (*gocql.Session, error) {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(5)
	return CassandraSessions[i], nil
}
