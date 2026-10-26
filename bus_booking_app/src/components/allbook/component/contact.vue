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
        <!-- <h1>{{ logindetail }}</h1> -->
        
        <div class="d-flex justify-center flex-column align-center">
        <v-card v-for="(value,index) in logindetail" :key="index" style="width: 700px;" class="pa-6 mt-10" color="yellow">
            <div class="d-flex justify-center">
            <h1 text-bold>{{ value.username}} here!</h1>
            </div>
            <h3>Mail id: {{ value.email }}</h3>
            <h3>Phone no:{{ value.phone }}</h3>
            <!-- <h3>User id: {{ value.id }}</h3> -->
            <v-btn @click="logout" class="mt-4 pa-3" color="red">logout</v-btn>

        </v-card>
        </div>
    </div>
</template>


<script>
import eventservices from '../../../Services/eventservices';


export default {
    data() {
        return {
            // logindetail: this.$store.state.logineduser,
            // signup: this.$store.state.signupdetail,
            logindetail:[]

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
        },
        logout() {
            this.$store.commit("resetsempty")
            this.$router.push("/login")
        }
       
    },
    mounted(){
        var mail=this.$store.state.theloginmail
        if(!mail){
            alert("/login first")
            this.$router.push("/login")
            return
        }
        eventservices.getDetails(mail).then((pResponse)=>{
            console.log(pResponse.data)
            this.logindetail=pResponse.data.data

        }).catch((error)=>{
            console.log(error)
        })
    }
























    // computed: {
    //     results() {
    //         let arr = [];
    //         for (let i = 0; i < this.logindetail.length; i++) {
    //             let person = this.logindetail[i];
    //             for (let j = 0; j < this.signup.length; j++) {
    //                 let thename = this.signup[j];
    //                 if (person.username == thename.name) {
    //                     arr.push(thename);
    //                 }
    //             }
    //         }
    //         return arr;
    //     }

    // }
}


</script>