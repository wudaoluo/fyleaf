package mongodb

import (
	"go-mog/mog"
	"fmt"
	"time"
	"gopkg.in/mgo.v2/bson"
)

const DBNAME = "test"



type User struct {
	Id_  bson.ObjectId `bson:"_id"`
	Name string "bson:`name`"
	Age  int    "bson:`age`"
}

type a struct {
	Name string
}


var c *mog.DialContext

func main() {
	var err error
	c, err = mog.Dial("10.211.55.4", 30)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()


	// session
	//s := c.Ref()
	//defer c.UnRef(s)
	//ss := s.DB(DBNAME).C("name")
	//err = ss.Insert(&User{Name: "Tom", Age: 20})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//ss.Insert(&a{Name: "asassa"})
	//mog_select_one()
	//mog_select_limit()

	mog_insert()
	//mog_select_one()
	//mog_updateid()
	//mog_delete()
}


func mog_insert() {
	s := c.Ref()
	defer c.UnRef(s)
	ss := s.DB(DBNAME).C("name1")
	err := ss.Insert(&User{Id_:bson.NewObjectIdWithTime(time.Now()),Name: "Tom", Age: 20})
	if err != nil {
		fmt.Println(err)
	}
}



//查询单个，并删除
func mog_select_one() {
	c, err := mog.Dial("10.211.55.4", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	s := c.Ref()
	defer c.UnRef(s)
	ss := s.DB(DBNAME).C("name")


	result:=User{}
	err=ss.Find(bson.M{"name":"Tom"}).One(&result)

	fmt.Println(result.Id_.String())

	// this.ServeJson() // 显示
	ss.Remove(result)
}



func mog_select_limit() {
	c, err := mog.Dial("10.211.55.4", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	s := c.Ref()
	defer c.UnRef(s)
	ss := s.DB(DBNAME).C("name")

	var users []User
	err=ss.Find(nil).All(&users)
	if err != nil {
		panic(err)
	}
	fmt.Println(users)
}





//根据id删除
func mog_delete() {
	c, err := mog.Dial("10.211.55.4", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	s := c.Ref()
	defer c.UnRef(s)
	ss := s.DB(DBNAME).C("name")


	//bson.ObjectIdHex = "aaaaaaaaaaaaaaaa"
	err = ss.RemoveId(bson.ObjectIdHex("5ad06bfbce216c46453bdf1d"))
	fmt.Println(err)
}






//更新
func mog_update() {
	s := c.Ref()
	defer c.UnRef(s)
	ss := s.DB(DBNAME).C("name")
	ss.Update(bson.M{ "_id": bson.ObjectIdHex("5ad072780000000000000000") },bson.M{"age": 21,"name":"ds"})
	//err := ss.UpdateId(bson.ObjectIdHex("5ad06bfbce216c46453bdf1d"),&User{Name:"carlo"})

	//fmt.Println("err",err)
}


//通过id更新

func mog_updateid() {
	s := c.Ref()
	defer c.UnRef(s)
	ss := s.DB(DBNAME).C("name")
	ss.UpdateId(bson.ObjectIdHex("5ad072780000000000000000"),bson.M{"age": 212,"name":"dsaaa"})
}

