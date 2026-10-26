<template>
    <div>
        <!-- <h1>{{ mainarr }}</h1> -->
        <v-btn color="secondary" elevation="14" large small @click="back" class="ml-5 mt-5"> back</v-btn>
        <!-- <h1>All Book</h1> -->
        <!-- <h1>{{ passengers }}</h1> -->
        <!-- <h1>{{ values }}</h1> -->

        <div class="d-flex justify-center">
            <div class="d-flex flex-column aling-center">
                <v-card v-for="(seat, index) in mainarr" :key="index" class="pa-5 mt-10" elevation="4"
                    style="width:1100px">
                    <h2 style="background-color: orange; color:white " class="d-flex justify-center">Passenger :{{ index
                        + 1 }}</h2>
                    <h3 class="pt-3">*required information</h3>
                    <v-text-field label="Passenger1 name" outlined v-model="seat.name" class="pt-2"
                        append-icon="mdi-account"></v-text-field>
                    <v-select :items="items(index)" label="select gender" v-model="seat.gender"></v-select>

                    <v-text-field label="passenger age" outlined v-model="seat.age"
                        append-icon="mdi-cake-variant"></v-text-field>

                </v-card>
                <div style="width: 500px; margin-top: 40px; margin-left: 300px;" class="d-flex justify-center">

                    <v-card-actions class="d-flex justify-center">
                        <v-btn style="width: 300px;" color="success" @click="updates" :loading="navi">proceed</v-btn>
                    </v-card-actions>
                </div>
            </div>
        </div>
    </div>
</template>


<script>
export default {
    data() {
        return {
            passengers: [],
            updatedpassengers: [],
            mobileno: "",
            emailid: "",
            navi: false,
            mainarr: []
            // age:""
        }
    },
    computed: {
        values() {
            return this.$store.state.bookvalue
        }
    },
    created() {
    
            this.mainarr = JSON.parse(this.seatsString);
        
        for (let i = 0; i < this.values.length; i++) {
            this.passengers.push({
                busid: this.values[i].busid,
                seatno: this.values[i].seatno,
                gender: this.values[i].prefer == 'female' ? "female" : "",
                name: ""

            })
        }
        // this.mainval=JSON.parse(this.$route.query.buses)

    },
    methods: {
        back() {
            this.$router.push("/seatbook")
            localStorage.clear()
        },
        updates() {
            for (let i = 0; i < this.mainarr.length; i++) {
                if (this.mainarr[i].name.trim() === "" || this.mainarr[i].gender === "") {
                    alert("please fill valid information");
                    return;
                }
            }

            this.updatedpassengers = []
            for (let i = 0; i < this.passengers.length; i++) {
                this.updatedpassengers.push({
                    busid: this.passengers[i].busid,
                    seatno: this.passengers[i].seatno,
                    gender: this.passengers[i].gender,
                    name: this.passengers[i].name,
                    age: this.passengers[i].age
                })

            }
            console.log(this.updatedpassengers)
            this.$emit("value", this.updatedpassengers)
            this.$store.commit("seatdetail", this.updatedpassengers)
            this.navi = true
            setTimeout(() => {
                this.navi = false
                this.$router.push({path:"/gpay",query:{values:JSON.stringify(this.mainarr)}})
            }, 2000)
            // this.$router.push("/gpay")



        },
        items(index) {
            for (let i = 0; i < this.values.length; i++) {
                if (this.values[index].gender === 'female') {
                    return ['female']
                }
                else {
                    return ["male", "female"]
                }
            }
        }

    },
    props: {
        seatsString: {
            type: String,
            default: ""
        }
    },

}
</script>
