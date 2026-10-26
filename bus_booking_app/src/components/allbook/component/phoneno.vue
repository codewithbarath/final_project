<template>
    <div>
        <v-app-bar color="primary" dark class="px-4">
              <v-icon x-large>mdi-bus</v-icon>
            <v-toolbar-title class="font-weight-bold text-h6">
                BusBooking
            </v-toolbar-title>
            <v-spacer></v-spacer>

            <v-btn text class="mx-1" @click="home">Home</v-btn>
            <v-btn text class="mx-1" @click="myticket">My Tickets</v-btn>
            <v-btn text class="mx-1" @click="setting">Settings</v-btn>
        </v-app-bar>
        <v-btn @click="back" class="mt-4 ms-4">
            back
        </v-btn>
         <!-- <h2>{{ mainarr }}</h2> -->
         <!-- <h3>{{ mobileno }}</h3> -->
         <!-- <h3>{{ emailid }}</h3> -->
         <!-- <h3>{{ finalamounts }}</h3> -->
         <!-- <h3>{{ frontenddata }}</h3> -->
        <!-- <v-overlay :value="isnavigating" opacity="0.7">
            <div class="text-center">
                <v-progress-circular indeterminate size="64" width="6" color="primary"></v-progress-circular>
                <h3 class="mt-4 white--text">paymentprocessing....</h3>
            </div>
        </v-overlay> -->
        <v-card style="width: 400px; margin-left: 700px; margin-top: 100px;" class="pa-9">
            <v-img src="https://i.pinimg.com/736x/b9/06/55/b906559125f9dbb97573503f379de4e8.jpg" style="width: 400px;"
                class="d-flex flex-column align-center"></v-img>
            <h1 class="text-center">ENTER THE UPI PIN</h1>
            <!-- <v-text-field label="enter the pin" class="mt-4" v-model="userinput"></v-text-field> -->
            <v-otp-input v-model="userinput" :disabled="loading" @finish="onFinish"></v-otp-input>
            <v-overlay absolute :value="loading">
                <v-progress-circular indeterminate color="primary"></v-progress-circular>
            </v-overlay>
            <v-snackbar v-model="snackbar" :color="snackbarColor" :timeout="2000">
                {{ text }}
            </v-snackbar>




        </v-card>
        <!-- <h1>changes{{ books }}</h1> -->
    </div>
</template>

<script>

import eventservices from '../../../Services/eventservices';

export default {
    data() {
        return {
            pin: 421622,
            userinput: "",
            // books: this.$store.state.totalnofbookedseats,
            // from: this.$store.state.from,
            // to: this.$store.state.to,
            // date: this.$store.state.date,
            // details: this.$store.state.selected.busId,//busid
            // detailss: this.$store.state.selected.busName,
            // detailsss: this.$store.state.selected.routes,
            // detailssss: this.$store.state.selected.timings.departure,
            // arrival: this.$store.state.selected.timings.arrival,
            // bookingdetails: this.$store.state.totalnofbookedseats,
            // no: this.$store.state.totalnofbookedseats.length,
            // email: this.$store.state.emailphoneno,
            // amount: this.$store.state.finalamounts,
            loading:false,
            snackbar:false,
            text:"",
            snackbarColor:"",
            mainarr:[],
            mobileno:"",
            emailid:"",
            finalamounts:"",
            frontenddata:""

            // isnavigating:false
        }
    },
     created() {

    this.mainarr = JSON.parse(this.$route.query.value);
    this.mobileno = this.$route.query.phone;
    this.emailid = this.$route.query.email;
    // this.finalamounts=this.$route.query.finaltotalamount
     this.finalamounts = this.$route.query.totalprice;
    },
    methods: {
        // verify() {
        //     if (this.userinput === "") {
        //         alert("Enter the pin first");
        //     } else if (parseInt(this.userinput) !== this.pin) {
        //         alert("Invalid pin no");
        //         this.userinput = "";
        //     } else {
        //         this.$store.commit("booksignal", { changes: this.books })
        //         this.$store.commit("value", {
        //             tfrom: this.from,
        //             tto: this.to,
        //             tdate: this.date,
        //             busid: this.details,
        //             busname: this.detailss,
        //             routes: this.detailsss,
        //             departure: this.detailssss,
        //             arrival: this.arrival,
        //             amount: this.amount,
        //             noofpass: this.no,
        //             id: Date.now(),
        //             names: this.$store.state.totalnofbookedseats

        //         });
        //         localStorage.clear()
        //         // this.$router.push("/ticket");
        //         this.isnavigating = true
        //         setTimeout(() => {
        //             this.isnavigating = false
        //             this.$router.push("/ticket")
        //         }, 4000)
        //     }
        // },
        onFinish(enteredValue){
               this.loading=true
            
               setTimeout( ()=>{
                this.loading=false
                this.snackbar=true
                if(parseInt(enteredValue) === this.pin){
                    eventservices.updatesValues(this.mainarr,this.mobileno,this.emailid,this.finalamounts).then((pResponse)=>{
                       console.log(pResponse.data);
                       if (pResponse.data.status === "s"){
                        console.log("success fully fetched")
                        console.log(pResponse.data.data)
                        this.text="payment success"
                        this.frontenddata=pResponse.data.data
                        // this.$router.push({path:"/ticket",query:{value:JSON.stringify(this.frontenddata)}})
                        this.$router.push("/ticket")

                       }
                    })
                    .catch((error)=>{
                        console.error(error);
                    })

                    // this.$store.commit("booksignal", { changes: this.books })
                    // this.$store.commit("value", {tfrom: this.from, tto: this.to, tdate: this.date,  busid: this.details, busname: this.detailss, routes: this.detailsss, departure: this.detailssss, arrival: this.arrival, amount: this.amount,noofpass: this.no, id: Date.now(), names: this.$store.state.totalnofbookedseats});
                    localStorage.clear()
                    // this.$router.push("/ticket")

                } else {
                    this.snackbarColor="error",
                    this.text="Wrong pin no",
                    this.userinput=""
                }
               },3000)
        },

        back() {
            // this.$store.commit("booksignal",{changes:this.books})
            this.$router.push("/payoption")
        },
        home() {
            this.$router.push("/dash")
        },
        myticket() {
            this.$router.push("/allticket")
        },
        setting() {
            this.$router.push("/setting")
        }
    }
}
</script>
