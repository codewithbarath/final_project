package contactapi

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
	Mailid string `json:"mail" validate:"required"`
}

type BusDetails struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type responsestruct struct {
	DataArr  []BusDetails `json:"data"`
	Status   string       `json:"status"`
	ErrorMsg string       `json:"errorMsg"`
}

func DetailAPI(writer http.ResponseWriter, request *http.Request) {
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
	log.Println("contact_collectmethod(+)")
	data.Mailid = r.URL.Query().Get("mail")
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
	log.Println("contact_collectmethod(-)")
	return nil
}
func validatemethod(logins Login) error {
	log.Println("contact_validatemethod(+)")
	validate := validator.New()
	err := validate.Struct(logins)
	if err != nil {
		return err
	}
	log.Println("contact_validatemethod(-)")
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
	log.Println("contact_communicate(+)")

	lData, lErr = findbusdetails(loginsdata)
	if lErr != nil {
		return nil, errors.New("Auvm:001" + lErr.Error())
	}
	defer log.Println("contact_communcate(-)")

	return lData, nil
}

func constructMethod(lRespec *responsestruct, lData []BusDetails) {
	log.Println("contact_construct method(+)")
	lRespec.DataArr = lData
	log.Println("contact_construct method(-)")
}

// func constructMethod(lRespec *responsestruct, lData []Login) {
// 	log.Println("construct method(+)")
// 	lRespec.DataArr = lData
// 	log.Println("construct method(-)")
// 	return
// }

func complete(writer http.ResponseWriter, lRespRec responsestruct) {
	log.Println("contact_completemethod(+)")
	lRespByte, lErr := json.Marshal(lRespRec)
	if lErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(writer, `{"status":"e","errorMsg":"error in umnarshaling"}`)
		return
	}

	fmt.Fprint(writer, string(lRespByte))
	log.Println("contact_completemethod(-)")
}
func findbusdetails(lreqrec Login) (lData []BusDetails, lErr error) {
	log.Println("contactmethod(+)")

	lDB, lErr := db.DBConnection()
	if lErr != nil {
		return lData, errors.New("contact_mehtod " + lErr.Error())
	}
	defer lDB.Close()
	// lRows, lErr := lDB.Query(`select id ,busid ,seat_number , row_index, column_index ,seat, gender,prefer from stn013seats where busid=?;`, lreqrec.BusID)
	// lRows, lErr := lDB.Query(`select  id , busid, seat_number, row_index, column_index, seat,fare, gender, prefer from stn013seats where busid=?;`, lreqrec.BusID)
	lRows, lErr := lDB.Query(`select email,username , password, phone from stn013signupbus  where email=?`, lreqrec.Mailid)
	if lErr != nil {
		return lData, errors.New("error in finding bus details" + lErr.Error())
	}
	for lRows.Next() {
		var lDatas BusDetails
		lErr = lRows.Scan(
			&lDatas.Email,
			&lDatas.Username,
			&lDatas.Password,
			&lDatas.Phone,
		)

		if lErr != nil {
			return nil, lErr
		}
		lData = append(lData, lDatas)

	}

	log.Println("contactmethod(-)")
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
