<template>
    <div>
        <v-img contain class="d-flex align-center justify-center pa-4" style="background-color: #f5f5f5;"
            src="../../../assets/signupimagseofbus.jpg">

            <v-container style="margin-top: 100px; width: 500px">
                <v-alert v-if="showAlert" color="green" elevation="12" type="success" class="ml-4">
                    Successfully registered
                </v-alert>
                <v-alert type="error" v-if="anyerror">{{ errormessage }}</v-alert>
                <v-card class="pa-6" elevation="3 d-flex  flex-column justify-center" flat
                    style="background-color: rgb(255, 255, 255); width: 100%; max-width: 420px; margin: auto;">
                    <h1 class=" text-center text-h4 font-weight-bold mb-6">Login</h1>
                    <v-text-field label="Enter your mail" outlined prepend-icon="mdi-account"
                        v-model="loginmail"></v-text-field>
                    <v-text-field label="Enter password" :type="showPassword ? 'text' : 'password'" outlined
                        :prepend-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
                        @click:prepend="showPassword = !showPassword" v-model="loginpassword"></v-text-field>
                    <v-card-actions class="justify-center">
                        <v-btn color="primary" elevation="2" large block @click="fetchlogindetails">
                            Login
                        </v-btn>
                    </v-card-actions>
                    <div class="d-flex justify-center">
                        <h4> Did not have accoun ?</h4>
                        <router-link to="/">signup</router-link>
                    </div>
                </v-card>
            </v-container>
        </v-img>
    </div>
</template>

<script>
import eventservices from '../../../Services/eventservices';

export default {
    data() {
        return {
            showAlert: false,
            anyerror: false,
            errormessage: "",
            loginmail: "",
            loginpassword: "",
            getsignupdetail: this.$store.state.signupdetail,
            showPassword: false
        };
    },
    methods: {
        fetchlogindetails() {
            this.anyerror = false;
            this.showAlert = false;
            eventservices.fetchlogin(this.loginmail, this.loginpassword).then((pResponse) => {
                    console.log(pResponse);
                    console.log(pResponse.data);
                    if (pResponse.data.status === "s") {
                        console.log("success fully get data")
                        this.showAlert = true;
                        this.$store.commit('setusermail', this.loginmail);

                        setTimeout(() => {
                            this.showAlert = false;
                            this.$router.push("/dash");
                        }, 3000);
                    } else {
                        this.errormessage = "invalid email or a password";
                        console.log("the eror ocurred inthe data")
                        this.anyerror = true;
                        setTimeout(() => {
                            this.anyerror = false;
                            this.errormessage = "";
                        }, 3000);
                    }
                })
                .catch((error) => {
                    console.error(error);
                });
                
        }

    }
}
</script>
