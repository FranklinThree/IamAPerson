package IamAPerson
import(
	"database/sql"
	//"fmt"
	_"github.com/go-sql-driver/mysql"

)
func main(){
	db,err := sql.Open("mysql","root:333333@(127.0.0.1:3306)/facedata?charset=utf8")
	defer db.Close()
	CheckErr(err)
	//stmt,err := db.Prepare("INSERT picture SET dapartmentID=?,iname=?")
	//
	//res,err:= stmt.Exec(1,"Franklin3")
	//CheckErr(err)
	//fmt.Println(res)

	//updates,err := db.Prepare("UPDATE picture SET iname= CONCAT(iname,'1') where dapartmentID = 1")
	//CheckErr(err)
	//res,err := updates.Exec()
	//CheckErr(err)
	//fmt.Println(res)

	//query,err := db.Query("SELECT iname FROM picture WHERE dapartmentID = 1 ")
	//for query.Next(){
	//	var res string
	//	query.Scan(&res)
	//	fmt.Println("res =",res)
	//}

}
