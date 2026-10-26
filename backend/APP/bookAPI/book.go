package bookapi

import (
	"backend/db"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

var DB *sql.DB

type Seat struct {
	BusID       int       `json:"busid"`
	SeatNo      string    `json:"seatno"`
	Prefer      *string   `json:"prefer"`
	ID          int       `json:"id"`
	Row         int       `json:"row"`
	Column      int       `json:"column"`
	Seat        string    `json:"seat"`
	Fare        int       `json:"fare"`
	Gender      string    `json:"gender"`
	Name        string    `json:"name"`
	Age         string    `json:"age"`
	CreatedBy   string    `json:"createdBy"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedBy   string    `json:"updatedBy"`
	UpdatedDate time.Time `json:"updatedDate"`
	IsActive    int       `json:"isActive"`
}

type Login struct {
	Seats        []Seat `json:"seats"`
	MobileNo     string `json:"mobileno"`
	MailID       string `json:"emailid"`
	FinalAmounts int    `json:"finalamounts,string"`
}

// type BusDetails struct {
// 	BusID           string `json:"busId"`
// 	BusName         string `json:"busName"`
// 	FromLocation    string `json:"from"`
// 	ToLocation      string `json:"to"`
// 	DepartureTime   string `json:"departureTime"`
// 	ArrivalTime     string `json:"arrivalTime"`
// 	SeatNo          string `json:"seatNo"`
// 	Gender          string `json:"gender"`
// 	Name            string `json:"name"`
// 	Age             string `json:"age"`
// 	TotalFare       int    `json:"totalFare"`
// 	MobileNumber    string `json:"mobileNo"`
// 	TotalPaidAmount int    `json:"totalPaidAmount"`
// }

type responsestruct struct {
	DataArr  []Login `json:"data"`
	Status   string  `json:"status"`
	ErrorMsg string  `json:"errorMsg"`
}

//	func main() {
//		fmt.Println("server started........")
//		var err error
//		DB, err = DBConnection()
//		if err != nil {
//			log.Fatal("Could not connect to database:", err)
//		}
//		defer DB.Close()
//		fmt.Println("Connected smoothly!")
//		http.HandleFunc("/register", getAPI)
//		err = http.ListenAndServe(":9091", nil)
//		if err != nil {
//			fmt.Println("the error is", err)
//		}
//	}
func BookAPI(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "USER, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	var lRespRec responsestruct
	if strings.EqualFold(request.Method, "POST") {
		var logins Login
		lRespRec.Status = "s"
		lErr := collectMehtod(request, &logins)
		if lErr != nil {
			lRespRec.Status = "e"
			lRespRec.ErrorMsg = "AgA:001: " + lErr.Error()
			log.Println(lRespRec.ErrorMsg)
			complete(writer, lRespRec)
			return
		}
		lErr = validatemethod(logins)
		if lErr != nil {
			lRespRec.Status = "e"
			lRespRec.ErrorMsg = "Avd:001: " + lErr.Error()
			log.Println(lRespRec.ErrorMsg)
			complete(writer, lRespRec)
			return
		}
		lData, lErr := communicatemethod(logins)
		if lErr != nil {
			lRespRec.Status = "e"
			lRespRec.ErrorMsg = "AgC:001: " + lErr.Error()
			log.Println(lRespRec.ErrorMsg)
			complete(writer, lRespRec)
			return
		}
		constructMethod(&lRespRec, lData)

		complete(writer, lRespRec)

	}
}
func collectMehtod(r *http.Request, data *Login) error {
	log.Println("seatbook_collectmethod(+)")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	err = json.Unmarshal(body, &data)
	// fmt.Println("the value is ", data.Dob)
	// fmt.Println("the va;ues ois ", data.Income)
	if err != nil {
		return err
	}
	log.Println("seatbook_collectmethod(-)")
	return nil
}
func validatemethod(logins Login) error {
	log.Println("seatbook_validatemethod(+)")
	validate := validator.New()
	err := validate.Struct(logins)
	if err != nil {
		return err
	}
	log.Println("seatbook_validatemethod(-)")
	return nil
}

// func communicatemethod(loginsdata Login) (lData []Login, lErr error) {
// 	log.Println("dahsboard_communicate(+)")

// 	lErr = findbusdetails(loginsdata)
// 	if lErr != nil {
// 		return nil, errors.New("Auvm:001" + lErr.Error())
// 	}
// 	defer log.Println("dahsboard_communcate(-)")

// 	return , nil
// }

func communicatemethod(loginsdata Login) (lData []Login, lErr error) {
	log.Println("seatbook_communicate(+)")
	lData, lErr = findbusdetails(loginsdata)
	if lErr != nil {
		return nil, errors.New("Auvm:001" + lErr.Error())
	}
	defer log.Println("seatbook_communcate(-)")
	return lData, nil
}

func constructMethod(lRespec *responsestruct, lData []Login) {
	log.Println("seat_book construct method(+)")
	lRespec.DataArr = lData
	log.Println("seatbook_construct method(-)")
}

// func constructMethod(lRespec *responsestruct, lData []Login) {
// 	log.Println("construct method(+)")
// 	lRespec.DataArr = lData
// 	log.Println("construct method(-)")
// 	return
// }

func complete(writer http.ResponseWriter, lRespRec responsestruct) {
	log.Println("seatbook_completemethod(+)")
	lRespByte, lErr := json.Marshal(lRespRec)
	if lErr != nil {
		fmt.Fprint(writer, `{"status":"e","errorMsg":"error in jsoning"}`)
		return
	}
	fmt.Fprint(writer, string(lRespByte))
	log.Println("seatbook_completemethod(-)")
}
func findbusdetails(lreqrec Login) (lData []Login, lErr error) {
	log.Println("seatbook_method(+)")
	lDB, lErr := db.DBConnection()
	if lErr != nil {
		return lData, errors.New("dAPIfindbusmehtod " + lErr.Error())
	}
	defer lDB.Close()
	for _, seat := range lreqrec.Seats {
		_, lErr := lDB.Exec(`update stn013seats set gender=? , name=?,age =? ,seat="booked" where busid=? and seat_number=?`, seat.Gender, seat.Name, seat.Age, seat.BusID, seat.SeatNo)
		if lErr != nil {
			log.Println("error in finding values")
			return lData, errors.New("error in finding bus details" + lErr.Error())
		}
		currentTime := time.Now()
		isActive := 1
		createdywho := "admin"

		// _, lErr = lDB.Exec(`insert into stn013bookings (busid, seat_number, passenger_name, passenger_age, passenger_gender, fare, mobile_number, email_id, total_paid_amount ,created_by, created_date, updated_by, updated_date,isactive) values (?, ?, ?, ?, ?, ?, ?, ?, ?,?,?,?,?,?)`, seat.BusID, seat.SeatNo, seat.Name, seat.Age, seat.Gender, seat.Fare, lreqrec.MobileNo, lreqrec.MailID, lreqrec.FinalAmounts, createdywho, currentTime, seat.Name, currentTime, isActive)
		_, lErr = lDB.Exec(`insert into stn013bookings (busid, seat_number, passenger_name, passenger_age, passenger_gender, fare, mobile_number, email_id, total_paid_amount, created_by, created_date, updated_by, updated_date, isactive) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, seat.BusID, seat.SeatNo, seat.Name, seat.Age, seat.Gender, seat.Fare, lreqrec.MobileNo, lreqrec.MailID, lreqrec.FinalAmounts, createdywho, currentTime, seat.Name, currentTime, isActive)

		if lErr != nil {
			log.Println("error in finding values")
			return lData, errors.New("error in finding bus details" + lErr.Error())
		}
		// 		lRows.lErr:=lDB.Query(`select id, busid, seat_number, passenger_name, passenger_age, passenger_gender, fare, mobile_number, total_paid_amount
		// from stn013bookings
		// where mobile_number = ?`,lreqrec.MobileNo)

		// lData=append(lData, lRow)
	}

	// lRows, lErr := lDB.Query(`select  b.busid,b.busname,b.fromlocation, b.tolocation, b.departuretime, b.arrivaltime,s.seat_number,s.passenger_name,s.passenger_age,s.passenger_gender,s.fare,s.mobile_number,s.total_paid_amount
	// from busdetailsstn013 b
	// inner join stn013bookings s ON b.busid = s.busid
	// where s.mobile_number = ?`, lreqrec.MobileNo)
	// if lErr != nil {
	// 	return lData, errors.New("error in finding bus details" + lErr.Error())
	// }
	// for lRows.Next() {
	// 	var lDatas Login
	// 	lErr = lRows.Scan(
	// 		&lDatas.BusID,
	// 		&lDatas.BusName,
	// 		&lDatas.FromLocation,
	// 		&lDatas.ToLocation,
	// 		&lDatas.DepartureTime,
	// 		&lDatas.ArrivalTime,
	// 		&lDatas.SeatNo,
	// 		&lDatas.Gender,
	// 		&lDatas.Name,
	// 		&lDatas.Age,
	// 		&lDatas.TotalFare,
	// 		&lDatas.MobileNumber,
	// 		&lDatas.TotalPaidAmount,
	// 	)
	// 	if lErr != nil {
	// 		return nil, lErr
	// 	}
	// 	lData = append(lData, lDatas)
	// 	fmt.Println("success", lDatas.Age)
	// }

	log.Println("seatbook_method(-)")
	return lData, nil

}
