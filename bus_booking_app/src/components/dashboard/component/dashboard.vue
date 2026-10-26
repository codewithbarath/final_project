<template>
    <div>

        <v-app-bar color="primary" dark class="px-4">
            <v-icon x-large>mdi-bus</v-icon>
            <v-toolbar-title class="font-weight-bold text-h6">
                BusBooking
            </v-toolbar-title>
            <v-spacer></v-spacer>

            <!-- <v-btn text class="mx-1" @click="home">Home</v-btn> -->
            <v-btn text class="mx-1" @click="signup">signup</v-btn>
            <v-btn text class="mx-1" @click="login">logout</v-btn>
            <v-btn text class="mx-1" @click="myticket">My Tickets</v-btn>
            <v-btn text class="mx-1" @click="setting">Settings</v-btn>
        </v-app-bar>
        <v-img class="" max-height="100%"
            src="https://assets.volvo.com/is/image/VolvoInformationTechnologyAB/blue-bus?qlt=82&wid=1024&ts=1660212095501&dpr=off&fit=constrain">
            <v-card style="position: relative;">


            </v-card>
            <!-- <h1>the data will show here{{ busarr }}</h1> -->
            <div class="d-flex justify-center"
                style="position: absolute; top: 400px; left: 500px; right: 0; background-color: white; width: 800px;">
                <div class="d-flex align-center gap-4 pa-4 rounded-lg elevation-2 bg-surface">
                    <v-autocomplete label="from" :items="fromvalue" v-model="from" outlined
                        class="mt-4"></v-autocomplete>
                    <v-autocomplete label="to" :items="tovalue" v-model="to" outlined
                        class="mx-4 mt-4"></v-autocomplete>
                    <v-menu ref="menu" v-model="menu" :close-on-content-click="false" :return-value.sync="date"
                        transition="scale-transition" offset-y min-width="auto">
                        <template v-slot:activator="{ on, attrs }">
                            <v-text-field v-model="date" label="Picker in menu" prepend-icon="mdi-calendar" readonly
                                v-bind="attrs" v-on="on"></v-text-field>
                        </template>
                        <v-date-picker v-model="date" no-title scrollable :min="todayDate">
                            <v-spacer></v-spacer>
                            <v-btn text color="primary" @click="menu = false">
                                Cancel
                            </v-btn>
                            <v-btn text color="primary" @click="$refs.menu.save(date)">
                                OK
                            </v-btn>
                        </v-date-picker>

                    </v-menu>
                    <v-overlay :value="isnavigating" opacity="0.7">
                        <div class="text-center">
                            <v-progress-circular indeterminate size="64" width="6"
                                color="primary"></v-progress-circular>
                            <h3 class="mt-4 white--text">Searching your buses ples wait....</h3>
                        </div>
                    </v-overlay>

                    <!-- <v-text-field label="no of passengers" outlined class="mx-4" v-model="passenger"></v-text-field> -->
                    <v-btn color="primary" elevation="2" large small x-large class="mb-3 mx-3" @click="search">
                        search</v-btn>






                </div>
            </div>
        </v-img>

    </div>
</template>


<script>
import eventservices from '../../../Services/eventservices';

export default {
    data() {
        return {
            // seebus:"",
            // passenger:"",
            from: "",
            to: "",
            fromvalue: ["chengalpattu", "tindivanam", "villupuram", "trichy", "Chennai", "Madurai", "tiruppur", "erode", "salem", "hosur", "Coimbatore", "Bengaluru "],
            tovalue: ["chengalpattu", "tindivanam", "villupuram", "trichy", "Chennai", "Madurai", "tiruppur", "erode", "salem", "hosur", "Coimbatore", "Bengaluru "],
            date: new Date().toISOString().substr(0, 10),
            // date:"",
            menu: false,
            isnavigating: false,
            busarr: []

        }
    },
    methods: {
        search() {
            eventservices.fetchbusdetail(this.from, this.to).then((pResponse) => {
                console.log(pResponse)
                console.log(pResponse.data.data)
                if (pResponse.data.status === "s") {
                    // this.$store.commit('backendbusdetail', pResponse.data.data);
                    // this.$router.push("/search");
                    this.$router.push({path: '/search', query: { buses: JSON.stringify(pResponse.data.data)}
                    });
                }
                else {
                    alert("no buses availabe")
                }
            }).catch((error) => {
                console.error(error);
            });
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
        signup() {
            this.$router.push("/")
        },
        login() {
            this.$store.commit("resetsempty")
            this.$router.push("/login")
        }
    },
    computed: {
        todayDate() {
            return new Date().toISOString().split('T')[0];
        }
    },



}


</script>