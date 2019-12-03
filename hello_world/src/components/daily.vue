<template>
    <div class="row at-row no-gutter">
        <div v-for="event in dailyEvents" :key="event.Title">
            <at-card class="" style="width: 300px">
                <h4 slot="title">{{event.Title}}</h4>
                <div>{{event.Extra}}</div>
            </at-card>
        </div>
    </div>
</template>

<script>
    import axios from 'axios';

    export default {
        name: "daily",
        data() {
            return {
                dailyEvents: []
            }
        },
        methods: {
            getDailyEvents: function () {
                axios.get('http://localhost:1221/dailyEvents').then(res => {
                    // alert(res)
                    // alert("1")
                    this.dailyEvents = res.data.data
                    // console.log(this.dailyEvents)
                }).catch(err => {
                    this.$Message.error(err)
                })
            },
        },
        mounted() {
            this.getDailyEvents();
        },
        created() {
            this.getDailyEvents();
        }
    }
</script>

<style scoped>

</style>