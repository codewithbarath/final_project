<template>
    <div>
        <!-- <h1>values:{{ mainarr }}</h1> -->
        <v-app-bar color="primary" dark class="px-4">
            <v-icon x-large>mdi-bus</v-icon>
            <v-toolbar-title class="font-weight-bold text-h6">
                BusBooking
            </v-toolbar-title>
            <v-spacer></v-spacer>

            <v-btn text class="mx-1" @click="home">Home</v-btn>
            <!-- <v-btn text class="mx-1" @click="myticket">My Tickets</v-btn> -->
            <v-btn text class="mx-1" @click="setting">Settings</v-btn>
        </v-app-bar>
        <!-- <h4>val is{{ mainarr }}</h4> -->
        <!-- <h1>{{ allvalues }}</h1> -->
        <v-btn clas="mt-4 ml-4" @click="back">back</v-btn>
        <div class="d-flex justify-center">
            <h1>YOUR TICKETS</h1>
        </div>
        <div class="d-flex flex-column align-center justify-center">
            <v-card color="red" v-for="(value, index) in mainarr" :key="index" style="width: 900px;" class="pt-3 mt-10">
                <div class="d-flex justify-space-around">
                    <h1 class="pa-4">{{ value.from }}</h1>
                    <v-divider class="mt-10"></v-divider>
                    <h1 class="pt-4">to</h1>
                    <v-divider class="mt-10"></v-divider>
                    <h1 class="pa-4">{{ value.to }}</h1>
                </div>

                <v-card class="d-flex justify-center">
                    <h3>{{ value.departureTime }}--to--</h3>
                    <h3>{{ value.arrivalTime }}</h3>
                </v-card>

                <v-card class="pa-4">
                    <h1 color="blue" class="text-center">{{ value.busName }}</h1>
                    <div>
                        <h3 class="pa-4 text-center">Passenger Name: {{ value.gender }} (Age: {{ value.name }})</h3>
                        <h3 class="pa-4 text-center">Seat No: {{ value.seatNo }} | Gender: {{ value.age }}</h3>
                    </div>
                </v-card>

                <v-card class="d-flex justify-space-around pa-5">
                    <h3 class="green--text">₹{{ value.totalPaidAmount }}</h3>
                    <h3 class="green--text">MOBILE: {{ value.mobileNo }}</h3>
                    <h3 class="green--text">BUS ID: {{ value.busId }}</h3>
                </v-card>
            </v-card>
        </div>



    </div>
</template>

<script>
import eventservices from '../../../Services/eventservices';



export default {
    data() {
        return {
            allvalues: this.$store.state.allbooking,
            names: this.$store.state.totalnofbookedseats,
            mainarr: []
        }
    },
    mounted() {
        var mail = this.$store.state.theloginmail
        if (!mail) {
            alert("login first")
            this.$router.push("/login")
            return
        }
        eventservices.getTicketmethod(mail).then((pResponse) => {
            console.log(pResponse.data)
            alert("success")
            this.mainarr = pResponse.data.data
        })
            .catch((error) => {
                console.log(error)
            })
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













































<!--         
        <div class="d-flex flex-column align-center justify-center">
            <v-card color="red"  v-for="(value, index) in mainarr" :key="index" style="width: 900px;"
                class="pt-3 mt-10">
                <div class="d-flex justify-space-around">
                    <h1 class="pa-4">{{ value.from }}</h1>
                    <v-divider class="mt-10"></v-divider>
                    <h1 class="pt-4">to</h1>
                    <v-divider class="mt-10"></v-divider>
                    <h1 class="pa-4">{{ value.to }}</h1>

                </div>
                <v-card>

                </v-card>
                <v-card class="d-flex justify-center">
                    <h3>{{ value.departureTime }}--to--</h3>
                    <h3>{{ value.arrivalTime }}</h3>
                </v-card>
                <v-card class="pa-4">
                    <h1 color="blue" class="text-center">{{ value.busName}}</h1>
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
        </div> -->

















<template>
    <div class="tickets-page">
        <!-- App bar -->
        <v-app-bar color="primary" dark elevation="4" class="px-4">
            <v-icon x-large class="mr-2">mdi-bus</v-icon>
            <v-toolbar-title class="font-weight-bold text-h6">BusBooking</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-btn text class="mx-1" @click="home">
                <v-icon left>mdi-home</v-icon> Home
            </v-btn>
            <v-btn text class="mx-1" @click="setting">
                <v-icon left>mdi-cog</v-icon> Settings
            </v-btn>
        </v-app-bar>

        <v-container class="py-8">
            <!-- Back + Title -->
            <div class="d-flex align-center mb-6">
                <v-btn icon color="primary" @click="back">
                    <v-icon>mdi-arrow-left</v-icon>
                </v-btn>
                <v-icon x-large color="primary" class="ml-3 mr-3">mdi-ticket-confirmation</v-icon>
                <h1 class="text-h4 font-weight-bold primary--text">YOUR TICKETS</h1>
            </div>

            <!-- Loading skeleton -->
            <v-row v-if="loading" justify="center">
                <v-col cols="12" md="10" lg="8">
                    <v-skeleton-loader type="article, actions" class="rounded-xl"></v-skeleton-loader>
                </v-col>
            </v-row>

            <!-- Empty state -->
            <v-card v-else-if="mainarr.length === 0" class="mx-auto pa-10 text-center rounded-xl" max-width="600"
                elevation="2">
                <v-icon size="80" color="grey lighten-1">mdi-ticket-outline</v-icon>
                <h2 class="text-h5 mt-4 grey--text">No tickets yet</h2>
                <p class="text-subtitle-1 grey--text mt-2">
                    Book your first journey to see tickets here
                </p>
                <v-btn color="primary" rounded class="mt-4" @click="home">
                    <v-icon left>mdi-bus</v-icon> Book a Ticket
                </v-btn>
            </v-card>

            <!-- Tickets -->
            <v-row v-else justify="center">
                <v-col v-for="(value, index) in mainarr" :key="index" cols="12" md="10" lg="8">
                    <v-card class="ticket-card rounded-xl elevation-6">
                        <!-- Gradient header -->
                        <div class="ticket-header pa-5 d-flex justify-space-between align-center">
                            <div class="d-flex align-center">
                                <v-icon color="white" large class="mr-3">mdi-bus-side</v-icon>
                                <div>
                                    <div class="white--text text-h6 font-weight-bold">
                                        {{ value.busName }}
                                    </div>
                                    <div class="white--text text-caption">Bus ID: {{ value.busId }}</div>
                                </div>
                            </div>
                            <v-chip small color="white" text-color="primary" label>
                                <v-icon small left>mdi-check-decagram</v-icon> Confirmed
                            </v-chip>
                        </div>

                        <!-- Route -->
                        <v-card-text class="pa-6">
                            <v-row align="center" no-gutters>
                                <v-col cols="5" class="text-center">
                                    <div class="text-caption grey--text text-uppercase">From</div>
                                    <div class="text-h5 font-weight-bold">{{ value.from }}</div>
                                    <div class="text-body-2 primary--text mt-1">
                                        {{ value.departureTime }}
                                    </div>
                                </v-col>
                                <v-col cols="2" class="text-center">
                                    <v-icon color="primary" size="32">mdi-arrow-right-bold</v-icon>
                                </v-col>
                                <v-col cols="5" class="text-center">
                                    <div class="text-caption grey--text text-uppercase">To</div>
                                    <div class="text-h5 font-weight-bold">{{ value.to }}</div>
                                    <div class="text-body-2 primary--text mt-1">
                                        {{ value.arrivalTime }}
                                    </div>
                                </v-col>
                            </v-row>
                        </v-card-text>

                        <!-- Perforation (real ticket look) -->
                        <div class="perforation">
                            <span class="notch notch-left"></span>
                            <span class="notch notch-right"></span>
                        </div>

                        <!-- Passenger -->
                        <v-card-text class="pa-6">
                            <div class="text-subtitle-2 grey--text text-uppercase mb-3">
                                <v-icon small class="mr-1">mdi-account</v-icon> Passenger
                            </div>
                            <v-card outlined class="pa-4 rounded-lg">
                                <div class="d-flex align-center flex-wrap">
                                    <v-avatar color="primary" size="48" class="mr-4">
                                        <span class="white--text text-h6 font-weight-bold">
                                            {{ (value.name || '?').charAt(0).toUpperCase() }}
                                        </span>
                                    </v-avatar>
                                    <div class="flex-grow-1">
                                        <div class="text-h6 font-weight-medium">{{ value.name }}</div>
                                        <div class="text-caption grey--text">Age {{ value.age }}</div>
                                    </div>
                                    <v-chip small :color="genderColor(value.gender)" text-color="white" class="mr-2">
                                        <v-icon small left>{{ genderIcon(value.gender) }}</v-icon>
                                        {{ value.gender }}
                                    </v-chip>
                                    <v-chip small color="primary" text-color="white">
                                        <v-icon small left>mdi-seat</v-icon> {{ value.seatNo }}
                                    </v-chip>
                                </div>
                            </v-card>
                        </v-card-text>

                        <!-- Footer -->
                        <v-card-actions class="pa-5 ticket-footer">
                            <v-icon color="success" class="mr-2">mdi-currency-inr</v-icon>
                            <div>
                                <div class="text-caption grey--text">Total Paid</div>
                                <div class="text-h5 font-weight-bold success--text">
                                    ₹{{ value.totalPaidAmount }}
                                </div>
                            </div>
                            <v-spacer></v-spacer>
                            <div class="text-right mr-4">
                                <div class="text-caption grey--text">Mobile</div>
                                <div class="text-body-1 font-weight-medium">{{ value.mobileNo }}</div>
                            </div>
                            <v-btn color="primary" rounded outlined>
                                <v-icon left>mdi-download</v-icon> Download
                            </v-btn>
                        </v-card-actions>
                    </v-card>
                </v-col>
            </v-row>
        </v-container>
    </div>
</template>

<script>
import eventservices from '../../../Services/eventservices';

export default {
  data() {
    return {
      mainarr: [],
      loading: false,
    };
  },
  mounted() {
    const mail = this.$store.state.theloginmail;
    if (!mail) {
      alert('Please login first');
      this.$router.push('/login');
      return;
    }
    this.loading = true;
    eventservices
      .getTicketmethod(mail)
      .then((pResponse) => {
        this.mainarr = pResponse.data.data || [];
      })
      .catch((error) => {
        console.error(error);
      })
      .finally(() => {
        this.loading = false;
      });
  },
  methods: {
    back() { this.$router.push('/phone'); },
    home() { this.$router.push('/dash'); },
    myticket() { this.$router.push('/allticket'); },
    setting() { this.$router.push('/setting'); },
    genderColor(gender) {
      return gender && gender.toLowerCase().startsWith('m') ? 'blue' : 'pink';
    },
    genderIcon(gender) {
      return gender && gender.toLowerCase().startsWith('m')
        ? 'mdi-gender-male'
        : 'mdi-gender-female';
    },
  },
};
</script>

<style scoped>
.tickets-page {
    background: #f5f7fa;
    min-height: 100vh;
}

/* Card hover lift */
.ticket-card {
    position: relative;
    transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.ticket-card:hover {
    transform: translateY(-4px);
    box-shadow: 0 12px 24px rgba(0, 0, 0, 0.15) !important;
}

/* Gradient header */
.ticket-header {
    background: linear-gradient(135deg, #1976d2 0%, #42a5f5 100%);
}

/* Perforation strip with side notches */
.perforation {
    position: relative;
    height: 0;
    border-top: 2px dashed #e0e0e0;
    margin: 0 24px;
}

.notch {
    position: absolute;
    top: -14px;
    width: 28px;
    height: 28px;
    background: #f5f7fa;
    border-radius: 50%;
    box-shadow: inset 0 2px 4px rgba(0, 0, 0, 0.08);
}

.notch-left {
    left: -38px;
}

.notch-right {
    right: -38px;
}

/* Footer */
.ticket-footer {
    background: #fafafa;
    border-top: 1px solid #eee;
}
</style>
