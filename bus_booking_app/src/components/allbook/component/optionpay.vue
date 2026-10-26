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
        <!-- <h1>price{{ price }}</h1>
        <h2>noofpass{{ noofpassengers }}</h2>
        <h2>finalgstamount{{ finalgstamount }}</h2>
        <h2>finaltotalamount{{ finaltotalamount }}</h2> -->
        <v-btn @click="back" class="ms-4 mt-4">back</v-btn>
        <!-- <h1>data{{ mainarr }}</h1> -->
         <!-- <h2>{{ mainarr }}</h2> -->
         <!-- <h3>{{ mobileno }}</h3> -->
         <!-- <h3>{{ emailid }}</h3> -->
         <h3>the price{{ price }}</h3>
        <div class="d-flex justify-center">

            <div style="width: 700px; margin-top: 80px;">
                <v-card class="pa-8">
                    <div class="d-flex flex-column justify-center">
                        <h2>PAYMENT OPTION</h2>
                        <p>please choose a payment option</p>
                    </div>

                    <div>
                        <v-btn @click="showcoupen" color="purple">coution option</v-btn>
                    </div>
                    <h4>select coupten 10% offer</h4>
                    <div v-show="value">
                        <!-- <h4>10% offer</h4> -->
                        <h2>Xc@$A2</h2>
                        <v-text-field outlined v-model="coupen" label="enter the coupen"></v-text-field>
                    </div>
                    <div>
                        <v-radio-group v-model="selectedpayment">
                            <div class="d-flex justify-space-evenly mx-6">
                                <div>
                                    <v-img src="https://i.pinimg.com/564x/67/28/d8/6728d8f618ff531833c69bd830569376.jpg"
                                        style="width:100px"></v-img>
                                    <v-radio label="Google Pay (Gpay)" value="gpay"></v-radio>
                                </div>
                                <div>
                                    <v-img
                                        src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQeMOi7xSvLohppZSU3IJTGvUizZR470SgQvO9G_Ei573vwoE3uuZVZ1L1e&s=10"
                                        style="width:100px"></v-img>
                                    <v-radio label="Paytm" value="paytm"></v-radio>
                                </div>
                            </div>

                        </v-radio-group>

                    </div>
                    <v-btn block color="primary" elevation="8" @click="proceed">proceed</v-btn>
                </v-card>
            </div>
        </div>
        <v-dialog v-model="dialogstate" max-width="400">

            <v-card>
                <v-card-title class="text-h5">
                    CONFORM PAYMENT
                </v-card-title>

                <v-card-text>
                    <v-card class="pa-4 mt-4" elevation="3">
                        <div class="d-flex justify-space-between pa-2">
                            <span class=" font-weight-medium">no of passengers:</span>
                            <span class=" font-weight-bold">{{ noofpassengers }}</span>
                        </div>
                        <v-divider></v-divider>
                        <div class="d-flex justify-space-between pa-2">
                            <span class="font-weight-medium">price:</span>
                            <span class="font-weight-bold">₹{{ price }}</span>
                        </div>
                        <v-divider></v-divider>
                        <div class="d-flex justify-space-between pa-2">
                            <span class="font-weight-medium">gst amount:</span>
                            <span class="font-weight-bold">₹{{ finalgstamount }}</span>
                        </div>

                        <v-divider class="my-2"></v-divider>
                        <div class="d-flex justify-space-between pa-2 primary--text">
                            <span class="title font-weight-bold">total amount:</span>
                            <span class="title font-weight-bold">₹{{ finaltotalamount }}</span>
                        </div>

                        <v-btn block color="primary" class="mt-4" elevation="7" @click="conform">
                            confirm
                        </v-btn>

                    </v-card>
                </v-card-text>

                <v-card-actions>
                    <v-spacer></v-spacer> <!-- Pushes buttons to the right hand side -->

                    <!-- Close by setting state variable back to false -->
                    <!-- <v-btn color="grey darken-1" text @click="dialogstate = false">
                        Cancel
                    </v-btn> -->

                    <!-- <v-btn color="error" text @click="deleteItem">
                        Delete
                    </v-btn> -->
                </v-card-actions>
            </v-card>

        </v-dialog>
    </div>
</template>

<script>
export default {
    data() {
        return {
            coupen: "",
            selectedpayment: "",
            value: false,
            coupencode: "Xc@$A2",
            finaltotalamount: "",
            finalgstamount: "",
            // theprice: this.$store.state.selected.fare,
            thepass: this.$store.state.totalnofbookedseats.length,
            cardshow: false,
            dialogstate: false,
            mainarr:[],
            mobileno:"",
            emailid:""
            // mainarr:[]
        }
    },
  
    created() {

    this.mainarr = JSON.parse(this.$route.query.value);
    this.mobileno = this.$route.query.phone;
    this.emailid = this.$route.query.email;
    }
,
  
    computed: {
        price() {
            let total=0;
            for (let i=0;i<this.mainarr.length;i++){
                total=total+this.mainarr[i].fare
            }
            return total
        },
        noofpassengers() {
            return this.mainarr.length;
        }
    },
    methods: {
        showcoupen() {
            this.value = true;
        },
        proceed() {
            let originalamount = this.price
            let gstamount = originalamount * 0.10
            this.finalgstamount = gstamount
            if (this.selectedpayment === "") {
                alert("please select a payment option first")
            }
            if (this.coupen === "") {
                let totalamount = originalamount + gstamount
                this.finaltotalamount = totalamount
                this.cardshow = true;

            } else {
                if (this.coupen === this.coupencode) {
                    let discount = 100
                    let totalamount = (originalamount - discount) + gstamount
                    this.finaltotalamount = totalamount
                    this.cardshow = true

                } else {
                    alert("invalid coupon code")
                    this.cardshow = false
                }
            }
            this.dialogstate = true
            // this.cardshow=true
        },
        conform() {
            this.$store.commit("finalamountsignal", this.finaltotalamount)
            this.$router.push({path: "/phone",query: {value: JSON.stringify(this.mainarr),phone: this.mobileno,email: this.emailid,totalprice: this.finaltotalamount}})

        },
        back() {
            this.$router.push("/gpay")
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
