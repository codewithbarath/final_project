<template>
    <div>
        <!-- <h5>the val {{ themaindata }}</h5> -->
        <div>

            <!-- <h2>{{ arr }}</h2> -->


        </div>
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
        <v-btn @click="back" class="mt-4 ms-4">back</v-btn>
        <!-- <h3 class="text-center text-uppercase">{{ bus.busName }}</h3> -->
        <div class="d-flex justify-center">
            <v-card class="d-flex pa-9">
                <div class="d-flex">
                    <v-icon x-large>mdi-human-female</v-icon>
                    <h4 class="mt-4 ms-4">FEMALE RESERVED SEATS</h4>
                </div>
                <div class="d-flex">
                    <v-icon x-large>mdi-human-male</v-icon>
                    <h4 class="mt-4 ms-4">MALE RESERVED SEATS</h4>
                </div>
                <div class="d-flex ms-5">
                    <v-card height="50" width="30"></v-card>
                    <h4 class="mt-4 ms-4">AVAILABEL SEATS</h4>
                </div>
            </v-card>
        </div>
        <!-- <h1>{{ arr }}</h1> -->
        <div class="text-center">
            <!-- Amount per seat:<h3>₹{{ totalfare }}</h3> -->
        </div>

        <!-- <v-card width="500">
            <h2 class="text-center mt-4">BUS ID:{{ val.busId }}</h2>
            <h1 class="text-center mt-4">BUS NAME:{{ val.busName }}</h1>


        </v-card> -->
        <div class="d-flex justify-center">
            <h1>₹{{ fairs * arr.length }}</h1>
        </div>
        <v-img style="width: 100px; margin-left: 950px; margin-top: 40px;"
            src="https://static.vecteezy.com/system/resources/previews/006/693/284/non_2x/steering-wheel-icon-template-black-color-editable-steering-wheel-icon-symbol-flat-illustration-for-graphic-and-web-design-free-vector.jpg"></v-img>






        <div v-for="(row, rowindex) in groupedSeats" :key="rowindex" class="d-flex justify-center align-start">

            <v-btn v-for="(seat, columnindex) in row" :key="seat.id" :class="['ma-2', columnindex == 0 ? 'mr-8' : '']"
                height="90" width="60" outlined :color="isSeatSelectedinarr(seat) ? 'green' : 'red'" elevation="2"
                :disabled="isalreadypresentinlocal(seat.busid, seat.seat_number) || seat.seat === 'booked'"
                @click="select(seat); again(seat.fare)">
                <v-icon v-if="seat.gender === 'male'" class="full-btn-icon" color="#2196F3" x-large>
                    mdi-human-male
                </v-icon>
                <v-icon v-else-if="seat.gender === 'female'" class="full-btn-icon" color="#D50000" x-large>
                    mdi-human-female
                </v-icon>
                <div v-else class="d-flex flex-column align-center">
                    <span class="text-subtitle-1 font-weight-bold">{{ seat.seat_number }}</span>
                    <span class="text-caption text--secondary">₹{{ seat.fare }}</span>
                </div>
            </v-btn>

        </div>
        <div class="d-flex justify-center">
            <v-btn @click="proceed">proceed</v-btn>
        </div>




    </div>
</template>


<script>

export default {
    data() {
        return {
            busid: "",
            name: "",
            gender: "",
            seatno: "",
            show: false,
            dialog: false,
            age: "",
            noofpassengers: this.$store.state.passno,
            arr: [],
            bus: this.$store.state.selected,
            yes: false,
            text: "",
            snack: false,
            themaindata: [],
            timeout: 1000,
            fair: null,
            // disable:"isa"

            // fares:this.val.fare

            // change:this.$store.state.thechanevalue
            //    noofpass:this.$store.state.passno
        }
    },

    created() {
        this.themaindata = JSON.parse(this.$route.query.values);
        console.log("seats", this.themaindata);
    },
    methods: {
        isSeatSelectedinarr(seat) {
            for (let i = 0; i < this.arr.length; i++) {
                if (this.arr[i].busid === seat.busid && this.arr[i].seatno === seat.seat_number) {
                    return true;
                }
            }
            return false;
        },




        // select(seat) {
        //     this.text = seat.seat_number;
        //     this.snack = true;
        //     let index = -1;
        //     for (let i = 0; i < this.arr.length; i++) {
        //         if (this.arr[i].busid === seat.busid && this.arr[i].seatno === seat.seat_number) {
        //             index = i;
        //             break;
        //         }
        //     }
        //     if (index > -1) {
        //         this.arr.splice(index, 1);
        //     } else {
        //         this.arr.push({ busid: seat.busid, seatno: seat.seat_number, prefer: seat.prefer, id: seat.id, row: seat.row_index, column: seat.column_index, seat: seat.seat, fare: seat.fare, gender: seat.gender,name:seat.name , age:seat.age ,from:seat.from,to:seat.to,busname:seat.busname});
        //     }
        // },
        // again(seat) {
        //     this.fair=seat

        // },
        select(seat) {
            this.text = seat.seat_number;
            this.snack = true;
            if (seat.column_index === 1 || seat.column_index === 2) {
                  let adjacentcol;
                  if (seat.column_index === 1) {
                      adjacentcol = 2;
                  } else {
                      adjacentcol = 1;
                  }
                  let adjacentseat = this.groupedSeats[seat.row_index][adjacentcol];
                  if (adjacentseat != null) {
                       if (adjacentseat.gender === "female") {
                          seat.gender = "female";
                       }
                  }
               
            }
            let index = -1;
            for (let i = 0; i < this.arr.length; i++) {
                if (this.arr[i].busid === seat.busid && this.arr[i].seatno === seat.seat_number) {
                    index = i;
                    break;
                }
            }
            if (index > -1) {
                this.arr.splice(index, 1);
            } else {
                this.arr.push({ busid: seat.busid, seatno: seat.seat_number, prefer: seat.prefer, id: seat.id, row: seat.row_index, column: seat.column_index, seat: seat.seat, fare: seat.fare,gender: seat.gender,name: seat.name,age: seat.age,from: seat.from,to: seat.to,busname: seat.busname});
            }
        },
        again(amount) {
            this.fair = amount
            console.log(this.fair)
        },
        isSelected(valb) {
            for (let i = 0; i < this.arr.length; i++) {
                if (this.arr[i].seatno == valb) {
                    return true
                }
            }
            return false
        },
        book() {
            this.show = true
        },
        proceed() {
            if (this.arr.length == 0) {
                alert("First choose the seats")
            }
            else {
                this.$store.commit("bookingvalue", this.arr)
                // this.$router.push("/allbook")
                this.$router.push({ path: '/allbook', query: { values: JSON.stringify(this.arr) } });
            }


        },
        back() {
            localStorage.clear()
            this.$router.push("/search")
        },
        booked(value) {
            if (value.seat == "booked") {
                return true
            } else {
                return false
            }
        },
        home() {
            this.$router.push("/dash")
        },
        myticket() {
            this.$router.push("/ticket")
        },
        setting() {
            this.$router.push("/setting")
        },
        isalreadypresentinlocal(id, no) {
            let seatsval = JSON.parse(localStorage.getItem("seatno"));
            console.log("localstoreval", seatsval);
            if (!seatsval) {
                return false;
            }
            for (let i = 0; i < seatsval.length; i++) {
                let seperateseat = seatsval[i];
                if (seperateseat.busid == id && seperateseat.seatno == no) {
                    return true;
                }
            }
            return false;
        },
        // selectseat(seat){
        //     if(seat.column_index===1 || seat.column_index===2){
        //         let ajacentcol;
        //         if(seat.column_index===1){
        //             adjacentcol=2
        //         } else {
        //             adjacentcol=1
        //         }
        //         let adjacentseat=this.groupedSeats[seat.row_index][adjacentcol]
        //         if(adjacentseat.gender==="female"){
        //             seat.gender="female"
        //         }
        //     }
        // }


    },


    computed: {
        val() {
            return this.$store.state.selected
        },
        nopass() {
            return this.$store.state.passno;

        },
        fare() {
            return this.$store.state.selected.fare
        },
        totalfare() {
            let basefare = this.$store.state.selected.fare || 0;
            return basefare * this.arr.length;
        },
        groupedSeats() {
            let rows = []
            for (let i = 0; i < this.themaindata.length; i++) {
                let seat = this.themaindata[i]
                let rowindex = seat.row_index
                let columnindex = seat.column_index
                if (!rows[rowindex]) {
                    rows[rowindex] = []
                }
                rows[rowindex][columnindex] = seat
            }
            return rows
        },
        fairs() {
            return this.themaindata[0].fare;
        }


    },

}




</script>



<style>
.full {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    font-size: 75px;
    opacity: 0, 25;
    z-index: 1;
}

.out {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    font-size: 75px;
    opacity: 0, 25;
    z-index: 1;

}

.blurred-female-icon {
    filter: blur(3px);
    opacity: 0.45;
}
</style>



<!-- <v-icon v-else-if="seat.prefer === 'female'" class="blurred-female-icon" color="#D50000" x-large>
                    mdi-human-female
                </v-icon> -->



















/*
<!-- <div v-for="(row, rowindex) in val.themaindata" :key="rowindex" class="d-flex justify-center align-start">

            <v-btn v-for="(seat, columnindex) in row" :key="seat.number"
                :class="['ma-2', columnindex == 0 ? 'mr-8' : '']" height="90" width="60" outlined 
                color="red" elevation="2"
                :disabled="isalreadypresentinlocal(val.busId, seat.number) || seat.seat === 'booked'"

            
                @click="select(val.busId, seat.number, seat.prefer); again(seat.seat)">
                <v-icon v-if="seat.gender === 'male'" class="full-btn-icon" color="#2196F3" x-large>
                    mdi-human-male
                </v-icon>
                <v-icon v-else-if="seat.gender === 'female'" class="full-btn-icon" color="#D50000" x-large>
                    mdi-human-female
                </v-icon> -->
<!-- <v-icon v-else-if="seat.prefer === 'female'" class="blurred-female-icon" color="#D50000" x-large>
                    mdi-human-female
                </v-icon> -->
<!-- <v-icon v-else class="full-btn-icon" color="#D500F9" x-large>

                </v-icon> -->

<!-- <span class="text-h6" v-if="!seat.gender"> -->
<!-- {{ seat.number }} -->
<!-- {{ seat.prefer }} -->
<!-- {{ seat.seat }} -->
<!-- {{ seat.gender }} -->
<!-- </span> -->
<!-- <p>{{ rowindex }}-{{ colIndex }}</p> -->
<!-- <v-snackbar v-model="snack" :timeout="timeout" color="success"> -->
<!-- <h3>you selected the seat no {{ text }}</h3> -->
<!-- </v-snackbar> -->
<!-- </v-btn> -->

<!-- </div> -->


<!-- <div class="d-flex justify-center mt-5" style="margin-bottom: 100px;">
            <v-btn @click="proceed" color="blue"> check seats</v-btn>
        </div> -->