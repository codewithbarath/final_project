import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    from: "",
    to: "",
    amount: "",
    date: "",
    gender: "",
    name: "",
    val: "",
    passno: "",
    selected: null,//i am having the selected wich bus
    signupdetail: [],
    logineduser: [],
    searchbus: [],
    thechanevalue: "",
    bookingdetail: [],//here i am having the booked detail
    bookvalue: [],
    totalnofbookedseats: [],// no of seats and the male female here  very important,
    emailphoneno: [],//uder booking email phoneno
    tempbooking: [],//store one ticket,
    allbooking: [],// all tickets
    finalamounts: "",
    backendbus:[],
    themail:"",
    theloginmail:"",
    buses: [
      {
        busId: "1001",
        busName: "SETC Ultra Deluxe",
        from: "Chennai",
        to: "Madurai",
        fare: 650,
        timings: { departure: "21:00 pm", arrival: "05:30 am" },
        routes: ["chengalpattu", "tindivanam", "villupuram", "trichy"],
        seats: [
          [{ number: 1, seat: "available", gender: null }, { number: 2, seat: "available", gender: null, prefer: "" }, { number: 3, seat: "available", gender: null, prefer: "" }],
          [{ number: 4, seat: "available", gender: null }, { number: 5, seat: "available", gender: null, prefer: "" }, { number: 6, seat: "available", gender: null, prefer: "" }],
          [{ number: 7, seat: "available", gender: null }, { number: 8, seat: "available", gender: null, prefer: "" }, { number: 9, seat: "available", gender: null, prefer: "" }],
          [{ number: 10, seat: "available", gender: null }, { number: 11, seat: "available", gender: null, prefer: "" }, { number: 12, seat: "available", gender: null, prefer: "" }],
          [{ number: 13, seat: "available", gender: null }, { number: 14, seat: "available", gender: null, prefer: "" }, { number: 15, seat: "available", gender: null, prefer: "" }],
          [{ number: 16, seat: "available", gender: null }, { number: 17, seat: "available", gender: null, prefer: "" }, { number: 18, seat: "available", gender: null, prefer: "" }]
        ]
      },
      {
        busId: "1002",
        busName: "KPN Travels AC Sleeper",
        from: "Coimbatore",
        to: "Bengaluru ",
        fare: 1100,
        timings: { departure: "22:30 pm", arrival: "06:15 am" },
        routes: ["tiruppur", "erode", "salem", "hosur"],
        seats: [
          [{ number: 1, seat: "available", gender: null }, { number: 2, seat: "available", gender: null, prefer: "" }, { number: 3, seat: "available", gender: null, prefer: "" }],
          [{ number: 4, seat: "available", gender: null }, { number: 5, seat: "booked", gender: null, prefer: "" }, { number: 6, seat: "available", gender: null, prefer: "" }],
          [{ number: 7, seat: "available", gender: null }, { number: 8, seat: "available", gender: null, prefer: "" }, { number: 9, seat: "available", gender: null, prefer: "" }],
          [{ number: 10, seat: "booked", gender: "null" }, { number: 11, seat: "available", gender: null, prefer: "" }, { number: 12, seat: "available", gender: null, prefer: "" }],
          [{ number: 13, seat: "available", gender: null }, { number: 14, seat: "available", gender: null, prefer: "" }, { number: 15, seat: "available", gender: null, prefer: "" }],
          [{ number: 16, seat: "available", gender: null }, { number: 17, seat: "available", gender: null, prefer: "" }, { number: 18, seat: "available", gender: null, prefer: "" }]
        ]
      },
      {
        busId: "1003",
        busName: "TN Ultra Deluxe",
        from: "Chennai",
        to: "Madurai",
        fare: 620,
        tempbooking: null,
        timings: { departure: "08:15 am", arrival: "16:45 pm" },
        routes: ["chengalpattu", "tindivanam", "villupuram", "trichy"],
        seats: [
          [{ number: 1, seat: "available", gender: null }, { number: 2, seat: "available", gender: null, prefer: "" }, { number: 3, seat: "available", gender: null, prefer: "" }],
          [{ number: 4, seat: "available", gender: null }, { number: 5, seat: "available", gender: null, prefer: "" }, { number: 6, seat: "available", gender: null, prefer: "" }],
          [{ number: 7, seat: "available", gender: null }, { number: 8, seat: "available", gender: null, prefer: "" }, { number: 9, seat: "available", gender: null, prefer: "" }],
          [{ number: 10, seat: "available", gender: null }, { number: 11, seat: "available", gender: null, prefer: "" }, { number: 12, seat: "available", gender: null, prefer: "" }],
          [{ number: 13, seat: "available", gender: null }, { number: 14, seat: "available", gender: null, prefer: "" }, { number: 15, seat: "available", gender: null, prefer: "" }],
          [{ number: 16, seat: "available", gender: null }, { number: 17, seat: "available", gender: null, prefer: "" }, { number: 18, seat: "available", gender: null, prefer: "" }]
        ]
      }
    ]




  },
  mutations: {
    signal(state, value) {
      state.signupdetail.push(value)
    },
    semisignal(state, value) {
      state.logineduser.push(value)
    },
    searchedbussignal(state, value) {
      state.searchbus = value

    },
    setusermail(state, email) {
      state.theloginmail = email;
      // alert("the mail is saved in the global "+state.theloginmail)
    },
    resetsempty(state){
      state.theloginmail="";
      alert("the logout mail is "+state.theloginmail)
    },
    themail(state, payload) {
      state.themail = payload; 
    },
    noofpass(state, value) {
      state.passno = value
    },
    selectarray(state, value) {
      state.selected = value
      alert(state.selected)
    },
    booksignal(state, value) {
      let tickets = value.changes;
      for (let t = 0; t < tickets.length; t++) {
      let currentticket = tickets[t];
        for (let i = 0; i < state.buses.length; i++) {
        if (state.buses[i].busId == currentticket.busid) {
            for (let j = 0; j < state.buses[i].seats.length; j++) {
                for (let k = 0; k < state.buses[i].seats[j].length; k++) {
                  if (state.buses[i].seats[j][k].number == currentticket.seatno) {
                    state.buses[i].seats[j][k].seat = "booked";
                    state.buses[i].seats[j][k].gender = currentticket.gender;
                     if (currentticket.gender === "female") {
                       if (k === 1 && state.buses[i].seats[j][k + 1]) {
                          state.buses[i].seats[j][k + 1].prefer = "female";
                       }
                       if (k === 2 && state.buses[i].seats[j][k - 1]) {
                          state.buses[i].seats[j][k - 1].prefer = "female";
                    }
                  }
                  break;
                }

              }
            }
            break;
          }
        }
      }
      alert("seat is bookeed");
    },
    backendbusdetail(state,value){
         state.backendbus = value;
    },


    value(state, value) {
      state.allbooking.push(value)
      console.log("allbooking is", state.allbooking.length)
      // alert("the length is", state.allbooking.length)
    },

    bookingdetails(state, value) {
      state.bookingdetail.push(value)

    },
    allval(state, value) {

      state.from = value.thefrom;
      state.to = value.theto;
      state.date = value.thedate;
    },
    saveTempbooking(state, value) {
      state.tempbooking = value;
    },
    bookingvalue(state, value) {
      state.bookvalue = value
      // alert("theee", state.bookvalue.length)
    },
    seatdetail(state, value) {
      state.totalnofbookedseats = value
    },
    emailphone(state, value) {
      state.emailphoneno.push(value)
    },
    finalamountsignal(state, value) {
      state.finalamounts = value
    }

  },
  actions: {},
  modules: {},
});
