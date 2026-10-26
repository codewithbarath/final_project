<template>
    <div>
        <!-- <h1>array{{ mainarr }}</h1> -->
        <v-app-bar color="primary" dark class="px-4">
            <v-toolbar-title class="font-weight-bold text-h6">
                BusBooking
            </v-toolbar-title>
            <v-spacer></v-spacer>

            <v-btn text class="mx-1" @click="home">Home</v-btn>
            <v-btn text class="mx-1" @click="myticket">My Tickets</v-btn>
            <v-btn text class="mx-1" @click="setting">Settings</v-btn>
        </v-app-bar>
        <div style="margin-left: 500px;" class="mt-10">
            <v-card style="width:600px" class="pa-10">
                <h1>Booking information</h1>
                <v-text-field label="phone no" v-model="phno"></v-text-field>
                <v-text-field label="email"></v-text-field>
                <v-expansion-panels accordion>
                    <v-expansion-panel v-for="(value, index) in getval" :key="index">
                        <v-expansion-panel-header>passenger detail</v-expansion-panel-header>
                        <v-expansion-panel-content>
                            <h2>Passenger Name :{{ value.passname }}</h2>
                            <h2>Gender : {{ value.passgender }}</h2>
                            <h2>age: {{ value.passage }}</h2>
                        </v-expansion-panel-content>
                    </v-expansion-panel>
                </v-expansion-panels>
                <v-btn block color="primary" elevation="2" x-large class="mt-4">proceed</v-btn>


            </v-card>
        </div>
    </div>
</template>

<script>


export default {
    data() {
        return {
            phno: "",
            email: "",
            mainarr:""
        }
    },
    methods: {
        home() {
            this.$router.push("/dash")
        },
        myticket() {
            this.$router.push("/allticket")
        },
        setting() {
            this.$router.push("/setting")
        }
    },
    //  created() {
    //    this.mainarr=JSON.parse(this.$route.query.values)
    // },
    computed: {
        getval() {
            return this.$store.state.bookingdetail
        }
    }

}




</script>