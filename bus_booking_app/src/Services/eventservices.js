import axios from "axios";

const baseclientapiu = axios.create({
    withCredentials: false,
    headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
    },
    baseURL: "http://localhost:9091",
});

export default {
    FetchApiDetail(username, mail, phoneno, password) {
        return baseclientapiu.post("/register", {
            username: username,
            mail: mail,
            phoneno: phoneno,
            password: password
        });
    },
    fetchlogin(loginmail, loginpassword) {
        return baseclientapiu.post("/login", {
            mail: loginmail,
            password: loginpassword
        });
    },
    fetchbusdetail(from, to) {
        return baseclientapiu.post("/dash", {
            from: from,
            to: to
        })
    },
    fetchseats(busid) {
        return baseclientapiu.get("/search", {
            params: { busid: busid}
        });
    },
    updatesValues(mainarr,mobileno,emailid,finalamounts){
        return baseclientapiu.post("/book",{
            seats:mainarr,
            mobileno:mobileno,
            emailid:emailid,
            finalamounts:finalamounts


        })
    },
    getTicketmethod(mail){
        return baseclientapiu.get("/ticket",{
            params:{mail:mail}
        })
    },
    getDetails(mail){
        return baseclientapiu.get("/details",{
            params:{mail:mail}
        })
    }


};