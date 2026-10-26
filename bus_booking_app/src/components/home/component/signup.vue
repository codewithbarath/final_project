<template>

    <div>
        <v-container style="margin-top: 100px; width: 500px;">

            <v-alert v-if="showAlert" color="green" elevation="12" type="success" class="ml-4">
                Successfully registered
            </v-alert>
            <v-alert type="error" v-if="anyerror">{{ error }}</v-alert>
            <!-- <v-img src="../../../assets/"></v-img> -->

            <v-card class="pa-6" elevation="3">
                <h1 class="text-center text-h4 font-weight-bold mb-6">signup</h1>
                <v-text-field label="Enter username" outlined prepend-icon="mdi-account"
                    v-model="username"></v-text-field>

                <v-text-field label="Enter your mail" outlined prepend-icon="mdi-email" v-model="mail"></v-text-field>
                <v-text-field label="Enter phoneno" outlined prepend-icon="mdi-cellphone"
                    v-model="phoneno"></v-text-field>
                <v-text-field label="Enter password" :type="showpass ? 'text' : 'password'" outlined
                    :prepend-icon="showpass ? 'mdi-eye' : 'mdi-eye-off'" @click:prepend="eyehideshow"
                    v-model="password"></v-text-field>

                <v-card-actions class="justify-center">
                    <v-btn color="primary" elevation="2" large block @click="signup">
                        signup
                    </v-btn>
                </v-card-actions>
            </v-card>
        </v-container>
    </div>


</template>



<script>
import eventservices from "../../../../../bus_booking_app/src/Services/eventservices";

export default {
    data() {
        return {
            emptyinput: false,
            alerterror: false,
            showAlert: false,
            anyerror: false,
            mail: "",
            password: "",
            error: "",
            phoneno: "",
            username: "",
            showpass: false
        }
    },
    methods: {
        fetchvalue() {
            const mailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
            if (this.username === "" || this.mail === "" || this.phoneno === "" || this.password === "") {
                this.anyerror = true;
                this.error = "Please fill in all boxes.";
                setTimeout(() => {
                    this.anyerror = false;
                }, 2000);
                return;
            }
            if (mailPattern.test(this.mail) === false) {
                this.anyerror = true;
                this.error = "Please enter a valid email address.";
                setTimeout(() => {
                    this.anyerror = false;
                }, 2000);
                return;
            }
         
            eventservices.FetchApiDetail(this.username, this.mail, this.phoneno, this.password).then((pResponse) => {
                // console.log(pResponse);
                console.log(pResponse.data)
                if (pResponse.data.status === "s") {
                    this.showAlert = true
                    this.$store.commit('themail', this.mail)
                    setTimeout(() => {
                        this.showAlert = false
                        this.$router.push("/login")
                    }, 3000);
                } else {

                    this.anyerror = true
                    this.error = "enter the values first"
                    setTimeout(() => {
                        this.anyerror = false
                    }, 2000)

                }
            })
                .catch((error) => {
                    console.error(error);
                });
        },

        signup() {
            this.fetchvalue();
        },

        eyehideshow() {
            this.showpass = !this.showpass;
        }
    }
}
</script>
































<!-- <script>

import eventservices from "../../../../../bus_booking_app/src/Services/eventservices";


export default {
    data() {
        return {
            emptyinput: false,
            alerterror: false,
            showAlert: false,
            anyerror: false,
            mail: "",
            password: "",
            error: "",
            phoneno: "",
            username: "",
            showpass: false
        }
    },
    methods: {
        fetchvalue() {
            eventservices
                .FetchApiDetail(this.username,this.mail,this.phoneno,this.password)
                .then((pResponse) => {
                    console.log(pResponse);
                    console.log(pResponse.data)
                    this.Responsedata = pResponse.data;
                    // this.names=pResponse.data[0]
                })
                .catch((error) => {
                    console.error(error);
                });
        },

        signup() {
            this.anyerror = false
            this.showAlert = false

            if (this.mail.trim() === "" || this.password.trim() === "" || this.phoneno.trim() === "" || this.username.trim() === "") {
                this.error = "input value cant empty"
                this.anyerror = true
                setTimeout(() => {
                    this.anyerror = false
                }, 1000)

            } else if (!/^[a-z0-9]+@[a-z]+\.[a-z]+$/.test(this.mail.trim())) {
                this.error = "enter valid email"
                this.anyerror = true
                setTimeout(() => {
                    this.anyerror = false
                }, 1000)
            } else if (this.mail.trim().length < 6 || this.password.trim().length < 6 || this.username.trim().length < 6 || this.phoneno.trim().length < 6) {
                this.error = "at least 6 characterss reuquired"
                this.anyerror = true
                setTimeout(() => {
                    this.anyerror = false
                }, 1000)

            } else {
                this.showAlert = true
                this.$store.commit('signal', { id: Date.now(), name: this.mail, password: this.password, username: this.username, phoneno: this.phoneno })
                setTimeout(() => {
                    this.showAlert = false
                    this.$router.push("/login")
                }, 3000)
            }
        },
        eyehideshow() {
            this.showpass = !this.showpass
        }


    }
}
</script> -->
