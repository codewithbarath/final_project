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
        <v-btn @click="back" class="mt-4 ms-4">back</v-btn>
        <!-- <h1>data{{ mainarr}}</h1> -->

        <div class="d-flex justify-center">
            <v-card elevation="3" class="mt-6 pa-8 " width="1000px ">
                <h1>additionl details</h1>
                <v-text-field placeholder="mobile Number" v-model="mobileno" outlined
                    prepend-inner-icon="mdi-phone"></v-text-field>
                <v-text-field placeholder="email id" v-model="emailid" outlined
                    prepend-inner-icon="mdi-email"></v-text-field>
                <v-btn block color="primary" elevation="11" large @click="updates"> book now</v-btn>
            </v-card>
        </div>
        <div class="d-flex flex-column" style="width: 800px; margin-top: 50px; margin-left: 500px;">
            <v-expansion-panels color="#64B5F6">
                <v-expansion-panel v-for="(value, index) in mainarr" :key="index" elevation="5" color="#64B5F6">
                    <v-expansion-panel-header>

                        <v-icon x-large>mdi-account-outline</v-icon>

                        <h2 class=" font-weight-black text-h5 primary--text text--darken-2 my-6 ">Passenger
                            {{ index + 1 }} detail</h2>
                    </v-expansion-panel-header>
                    <div class="bg-light-blue-lighten-4">
                        <v-expansion-panel-content class="bg-light-blue-lighten-4">
                            <v-simple-table id="one">
                                <template v-slot:default>
                                    <thead>
                                        <tr>
                                            <th class="text-left">
                                                Name
                                            </th>
                                            <th class="text-left">
                                                age
                                            </th>
                                            <th class="text-left">
                                                gender
                                            </th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <tr >
                                            <td>{{ value.name }}</td>
                                            <td>{{ value.age }}</td>
                                            <td>{{ value.gender }}</td>
                                        </tr>
                                    </tbody>
                                </template>
                            </v-simple-table>
                        </v-expansion-panel-content>
                    </div>
                </v-expansion-panel>
            </v-expansion-panels>
        </div>




    </div>
</template>


<script>

export default {
    data() {
        return {
            rawdata: "",
            mobileno: "",
            emailid: "",
            details: this.$store.state.totalnofbookedseats,
            mainarr: []
        }
    },
    created() {
        
            this.mainarr = JSON.parse(this.$route.query.values);
        
    },



    methods: {
        back() {
            this.$router.push("/allbook")
        },

        updates() {
            if (!this.mobileno.trim() || !this.emailid.trim()) {
                alert("Enter the details first")
                return
            }

            this.$store.commit("emailphone", {
                email: this.emailid,
                mobileno: this.mobileno
            })
            this.$router.push({path: "/payoption", query: { value: JSON.stringify(this.mainarr),phone: this.mobileno, email: this.emailid }})
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


<style>
#one {
    background-color: #4FC3F7;
}
</style>