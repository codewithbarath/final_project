package seatsapi

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

	"github.com/go-playground/validator/v10"
)

var DB *sql.DB

type Login struct {
	// From string `json:"from" validate:"required"`
	// To   string `json:"to" validate:"required"`
	BusID string `json:"busid" validate:"required"`
}

type BusDetails struct {
	ID          int     `json:"id"`
	BussID      int     `json:"busid"`
	SeatNumber  string  `json:"seat_number"`
	RowIndex    int     `json:"row_index"`
	ColumnIndex int     `json:"column_index"`
	Seat        string  `json:"seat"`
	Fare        int     `json:"fare"`
	Gender      *string `json:"gender"`
	Prefer      *string `json:"prefer"`
	Name        *string `json:"name"`
	Age         *int    `json:"age"`
}

type responsestruct struct {
	DataArr  []BusDetails `json:"data"`
	Status   string       `json:"status"`
	ErrorMsg string       `json:"errorMsg"`
}

func SeatAPI(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "USER, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	var lRespRec responsestruct
	if strings.EqualFold(request.Method, "GET") {
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
	log.Println("seat_collectmethod(+)")
	data.BusID = r.URL.Query().Get("busid")
	if r.Method == "GET" {
		return nil
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	log.Println("seat_collectmethod(-)")
	return nil
}
func validatemethod(logins Login) error {
	log.Println("dahsboard_validatemethod(+)")
	validate := validator.New()
	err := validate.Struct(logins)
	if err != nil {
		return err
	}
	log.Println("dahsboard_validatemethod(-)")
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

func communicatemethod(loginsdata Login) (lData []BusDetails, lErr error) {
	log.Println("communicate(+)")

	lData, lErr = findbusdetails(loginsdata)
	if lErr != nil {
		return nil, errors.New("Auvm:001" + lErr.Error())
	}
	defer log.Println("communcate(-)")

	return lData, nil
}

func constructMethod(lRespec *responsestruct, lData []BusDetails) {
	log.Println("construct method(+)")
	lRespec.DataArr = lData
	log.Println("construct method(-)")
}

// func constructMethod(lRespec *responsestruct, lData []Login) {
// 	log.Println("construct method(+)")
// 	lRespec.DataArr = lData
// 	log.Println("construct method(-)")
// 	return
// }

func complete(writer http.ResponseWriter, lRespRec responsestruct) {
	log.Println("dahsboard_completemethod(+)")
	lRespByte, lErr := json.Marshal(lRespRec)
	if lErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(writer, `{"status":"e","errorMsg":"error in umnarshaling"}`)
		return
	}

	fmt.Fprint(writer, string(lRespByte))
	log.Println("dahsboard_completemethod(-)")
}
func findbusdetails(lreqrec Login) (lData []BusDetails, lErr error) {
	log.Println("findseatmethod(+)")

	lDB, lErr := db.DBConnection()
	if lErr != nil {
		return lData, errors.New("dAPIfindbusmehtod " + lErr.Error())
	}
	defer lDB.Close()
	// lRows, lErr := lDB.Query(`select id ,busid ,seat_number , row_index, column_index ,seat, gender,prefer from stn013seats where busid=?;`, lreqrec.BusID)
	// lRows, lErr := lDB.Query(`select  id , busid, seat_number, row_index, column_index, seat,fare, gender, prefer from stn013seats where busid=?;`, lreqrec.BusID)
	lRows, lErr := lDB.Query(`select id, busid, seat_number, row_index, column_index, seat, fare, gender, prefer, name, age from stn013seats where busid=?;`, lreqrec.BusID)

	if lErr != nil {
		return lData, errors.New("error in bus finding" + lErr.Error())
	}
	for lRows.Next() {
		var lDatas BusDetails
		lErr = lRows.Scan(
			&lDatas.ID,
			&lDatas.BussID,
			&lDatas.SeatNumber,
			&lDatas.RowIndex,
			&lDatas.ColumnIndex,
			&lDatas.Seat,
			&lDatas.Fare,
			&lDatas.Gender,
			&lDatas.Prefer,
			&lDatas.Name,
			&lDatas.Age,
		)
		if lErr != nil {
			return nil, lErr
		}
		lData = append(lData, lDatas)
	}
	log.Println("findseatmethod(-)")
	return lData, nil

}

// defer lDB.Close()
// var dbPassword string
// lErr = lDB.QueryRow("SELECT password FROM busstn013signup WHERER email = ?", lreqrec.Email).Scan(&dbPassword)
// if lErr != nil {
// 	return errors.New("auid:002" + lErr.Error())
// }
// if dbPassword != lreqrec.Password {
// 	return errors.New("auid:003: incorrect password")
// }
// log.Println("findbusmethod(-)")
// return nil
// func findbusdetails(lreqrec Login) ([]BusDetails, error) {
// 	log.Println("dahsboard_query(+)")

// 	lDB, lErr := db.DBConnection()
// 	if lErr != nil {
// 		return errors.New("auid:001" + lErr.Error())
// 	}
// 	defer lDB.Close()

// 	log.Println("dahsboard_qeryend(-)")
// 	return nil
// }
