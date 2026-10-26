package loginapi

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
)

var DB *sql.DB

type Login struct {
	Email    string `json:"mail" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type responsestruct struct {
	DataArr  []Login `json:"data"`
	Status   string  `json:"status"`
	ErrorMsg string  `json:"errorMsg"`
}

func GetAPI(writer http.ResponseWriter, request *http.Request) {
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
	log.Println("logincollectmethod(+)")
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
	log.Println("logincollectmethod(-)")
	return nil
}

func communicatemethod(loginsdata Login) (lData []Login, lErr error) {
	log.Println("logincommunicate(+)")
	lErr = InsertUserdetails(loginsdata)
	if lErr != nil {
		return nil, errors.New("Auvm:001" + lErr.Error())
	}
	defer log.Println("logincommuncate(-)")
	return nil, nil
}

func constructMethod(lRespec *responsestruct, lData []Login) {
	log.Println("loginconstruct method(+)")
	lRespec.DataArr = lData
	log.Println("loginconstruct method(-)")
	return
}

func complete(writer http.ResponseWriter, lRespRec responsestruct) {
	log.Println("logincompletemethod(+)")
	lRespByte, lErr := json.Marshal(lRespRec)
	if lErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(writer, `{"status":"e","errorMsg":"error in jsoning"}`)
		return
	}
	fmt.Fprint(writer, string(lRespByte))
	log.Println("logincompletemethod(-)")
}

func InsertUserdetails(lreqrec Login) (lErr error) {
	log.Println("Inerusrdetail9(+)")
	lDB, lErr := db.DBConnection()
	if lErr != nil {
		return errors.New("auid:001" + lErr.Error())
	}
	defer lDB.Close()
	var dbPassword string
	lErr = lDB.QueryRow("select password from stn013signupbus where email = ?", lreqrec.Email).Scan(&dbPassword)
	if lErr != nil {
		return errors.New("auid:002" + lErr.Error())
	}
	if dbPassword != lreqrec.Password {
		return errors.New("auid:003: incorrect password")
	}
	log.Println("Insertuserdetails(-)")
	return nil
}
