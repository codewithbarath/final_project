
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
        <!-- <h1>{{ names }}</h1> -->
        <!-- <h1>{{ allvalues }}</h1> -->
         <v-btn clas="mt-4 ml-4" @click="back">back</v-btn>
        <div class="d-flex justify-center">
            <h1>YOUR TICKETS</h1>
        </div>
        <div class="d-flex flex-column align-center justify-center">
            <v-card color="red"  v-for="(value, index) in allvalues" :key="index" style="width: 900px;"
                class="pt-3 mt-10">
                <div class="d-flex justify-space-around">
                    <h1 class="pa-4">{{ value.tfrom }}</h1>
                    <v-divider class="mt-10"></v-divider>
                    <h1 class="pt-4">to</h1>
                    <v-divider class="mt-10"></v-divider>
                    <h1 class="pa-4">{{ value.tto }}</h1>

                </div>
                <v-card>
                    <h2 class="text-center">DATE:{{ value.tdate }}</h2>

                </v-card>
                <v-card class="d-flex justify-center">
                    <h3>{{ value.departure }}--to--</h3>
                    <h3>{{ value.arrival }}</h3>
                </v-card>
                <v-card class="pa-4">
                    <h1 color="blue" class="text-center">{{ value.busname }}</h1>
                    <div v-for="(passenger, pIndex) in value.names" :key="pIndex">
                        <h3 class="pa-4  text-center">Passenger Name: {{ passenger.name }} (Age: {{ passenger.age }})
                        </h3>
                        <h3 class="pa-4 text-center">Seat No: {{ passenger.seatno }} | Gender: {{ passenger.gender }}
                        </h3>
                    </div>
                </v-card>
                <v-card class="d-flex justify-space-around pa-5"  >
                    <h3 class="green--text">₹{{ value.amount }}</h3>
                    <h3 class="green--text">NO OF PASS:{{ value.noofpass }}</h3>
                    <h3 class="green--text">TICKET ID:{{ value.id }}</h3>
                </v-card>
            </v-card>
        </div>
    </div>
</template>

<script>


export default {
    data() {
        return {
            allvalues: this.$store.state.allbooking,
            names: this.$store.state.totalnofbookedseats
        }
    },
    methods: {
      
        back() {
            this.$router.push("/phone")


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