<template>
    <div>
        <!-- <h1>i am main data{{ maindata }}</h1> -->
        <v-app-bar color="primary" dark class="px-4">
              <v-icon x-large>mdi-bus</v-icon>
            <v-toolbar-title class="font-weight-bold text-h6">
                BusBooking
            </v-toolbar-title>
            <v-spacer></v-spacer>

            <v-btn text class="mx-1" @click="home">Home</v-btn>
            <v-btn text class="mx-1" @click=myticket>My Tickets</v-btn>
            <v-btn text class="mx-1" @click="setting">Settings</v-btn>
        </v-app-bar>
        <!-- <allbooks @value="changes"></allbooks> -->
          <allbooks :seats-string="$route.query.values" @value="changes"></allbooks>
        <!-- <gpay :childval="maindata"></gpay> -->
    </div>
</template>

<script>

import allbooks from './component/allbooks.vue';
// import gpay from './component/gpay.vue';
export default {
    data() {
        return {
            maindata: [],

        }
    },
    components: {
        allbooks,
        // gpay
    },
    methods: {
        changes(values) {
            this.maindata = values
        },
        home() {
            this.$router.push("/dash")
        },
        myticket() {
            this.$router.push("/ticket")
        },
        setting() {
            this.$router.push("/setting")
        }
    },

}

</script>