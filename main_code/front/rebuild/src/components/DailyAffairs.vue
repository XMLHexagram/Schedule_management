<template>
    <div>
        <button v-on:click="getData"></button>
        <p>{{dailyEvents}}</p>
    </div>
</template>

<script>
    import {apiGet} from "../api";

    export default {
        name: 'DailyAffairs',
        props: {},
        data (){
            return{
                dailyEvents:[]
            }
        },
        methods: {
            getData:function () {
                this.dailyEvents = apiGet("http://www.sweetbeecr.com:1221/all/dailyEvents");
                // console.log(this.dailyEvents)
            }
        },
    }
</script>

<style scoped>
</style>
