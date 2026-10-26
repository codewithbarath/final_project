package signup

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

	// _ "://github.com"

	"github.com/go-playground/validator/v10"
)

var DB *sql.DB

type Login struct {
	Username       string    `json:"username" validate:"required"`
	Phone          string    `json:"phoneno" validate:"required"`
	Email          string    `json:"mail" validate:"required,email"`
	Password       string    `json:"password" validate:"required"`
	ID             int       `json:"id"`
	CreatedBy      string    `json:"createdBy"`
	CreatedDate    time.Time `json:"createdDate"`
	UpdatedBy      string    `json:"updatedBy"`
	UpdatedDate    time.Time `json:"updatedDate"`
	CreatedProgram string    `json:"createdProgram"`
	UpdatedProgram string    `json:"updatedProgram"`
	IsActive       int       `json:"isActive"`
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
	log.Println("collectmethod(+)")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	err = json.Unmarshal(body, &data)

	if err != nil {
		return err
	}
	log.Println("collectmethod(-)")
	return nil
}
func validatemethod(logins Login) error {
	log.Println("validatemethod(+)")
	validate := validator.New()
	err := validate.Struct(logins)
	if err != nil {
		return err
	}
	log.Println("validatemethod(-)")
	return nil
}
func communicatemethod(loginsdata Login) (lData []Login, lErr error) {
	log.Println("communicate(+)")
	lErr = InsertUserdetails(loginsdata)
	if lErr != nil {
		return nil, errors.New("Auvm:001" + lErr.Error())
	}
	defer log.Println("communcate(-)")
	return nil, nil
}

func constructMethod(lRespec *responsestruct, lData []Login) {
	log.Println("construct method(+)")
	lRespec.DataArr = lData
	log.Println("construct method(-)")
	return
}

func complete(writer http.ResponseWriter, lRespRec responsestruct) {
	log.Println("completemethod(+)")

	writer.Header().Set("Content-Type", "application/json")
	lRespByte, lErr := json.Marshal(lRespRec)
	if lErr != nil {
		fmt.Fprint(writer, `{"status":"e","errorMsg":"error while jsoning"}`)
		return
	}

	fmt.Fprint(writer, string(lRespByte))
	log.Println("completemethod(-)")
}

func InsertUserdetails(lreqrec Login) (lErr error) {
	log.Println("Inerusrdetail9(+)")
	// lDB
	lDB, err := db.DBConnection()
	if err != nil {
		return errors.New("auid:001" + err.Error())
	}
	defer lDB.Close()
	currentTime := time.Now()
	programName := "flattrade"
	isActiveValue := 1
	_, lErr = lDB.Exec(`insert into stn013signupbus (email, username, phone, password, created_by, created_date, updated_by, updated_date, created_program, updated_program, isactive) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, lreqrec.Email, lreqrec.Username, lreqrec.Phone, lreqrec.Password, lreqrec.Username, currentTime, lreqrec.Username, currentTime, programName, programName, isActiveValue)
	if lErr != nil {
		return errors.New("auid:002" + lErr.Error())
	}
	log.Println("Insertuserdetails(-)")
	return nil
}
