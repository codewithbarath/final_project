<template>
    <div>
        <!-- <h1>{{ seats}}</h1> -->
         

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
        <v-img src="../../../assets/signupbuspic.jpg">
        <v-btn class="mt-4 ms-4" @click="back">back</v-btn>
        <h1 class="text-center justify-space-around white--text">AVAILABEL BUSES</h1>

        <v-container>
            <div class="d-flex  flex-column">



                <v-card v-for="value in theavailabebus" :key="value.busId" class="pa-4 mb-4 mt-10">


                    <div class="headline font-weight-bold d-flex justify-center "
                        style="background-color: orange; color: white;">{{ value.busName }}</div>
                    <div class="caption black--text mb-2">bus ID: {{ value.busId }}</div>
                    <div class="d-flex justift-space-around">

                        <div class="subtitle-1 mb-1">
                            <h4>from:</h4> {{ value.from }}
                        </div>
                        <div>
                            <h2>----------------- </h2>
                        </div>
                        <div class="subtitle-1 mb-1">
                            <h4>to:</h4> {{ value.to }}
                        </div>
                    </div>
                    <div class="subtitle-1 mb-1">
                        <h4>time:</h4> {{ value.departureTime }} - {{ value.arrivalTime }}
                    </div>
                    <v-divider></v-divider>
                    <v-divider></v-divider>


                    <div class="d-flex align-center justify-space-between">
                        <span class="title font-weight-bold green--
                        text">₹{{ value.fare }}</span>
                        <div class="d-flex ">
                            <v-icon color="yellow">mdi-star</v-icon>
                             <v-icon color="yellow">mdi-star</v-icon>
                              <v-icon color="yellow">mdi-star</v-icon>
                               <v-icon color="yellow">mdi-star</v-icon>
                                <v-icon color="yellow">mdi-star</v-icon>
                                 <v-icon color="grey">mdi-star-outline</v-icon>
                        </div>
                        <v-btn color="primary" @click="book(value)">View Seats</v-btn>
                    </div>
                    <v-progress-linear v-if="matchbusid == value.busId " indeterminate color="blue" class="my-4"></v-progress-linear>
                     <p v-if="matchbusid == value.busId" class="blue--text text-center">Please wait</p>

                    

                </v-card>




            </div>
        </v-container>
        </v-img>


    </div>
</template>


<script>

import eventservices from '../../../Services/eventservices';
export default {
    data() {
        return {
            matchbusid: null,
            theavailabebus: [] ,
            seats:[]
        }
    },
    created() {
       this.theavailabebus=JSON.parse(this.$route.query.buses)
    },
    methods: {
        book(value){
            var busid=value.busId
             eventservices.fetchseats(busid ).then((pResponse) => {
                    if (pResponse.data.status === "s") {
                        this.seats=pResponse.data.data
                        // console.log(this.seats)
                        // this.$router.push({path: '/seatbook', query: { values: JSON.stringify(pResponse.data.data)} });
                        this.$router.push({ path: '/seatbook', query: { values: JSON.stringify(pResponse.data.data) } });

                    }
                })
                .catch((error) => {
                    console.error(error);
                });
        },
        // book(value) {
        //     this.matchbusid = value.busId
        //     this.$store.commit("selectarray", value)
        //     setTimeout(() => {
        //         this.matchbusid = null
        //         this.$router.push("/seatbook")
        //     }, 4000)
        // },
        home() {
            this.$router.push("/dash")
        },
        myticket() {
            this.$router.push("/ticket")
        },
        setting() {
            this.$router.push("/setting")
        },
        back() {
            this.$router.push("/dash")
        }
    }
}
</script>
