<template>
    <div class="row at-row flex-center">
        <!--        <input v-model="nowYear">-->
        <!--        <p>{{nowYear}}</p>-->
        <at-card style="width: 600px">
            <h4>添加新事务</h4>
            <div>
                <!--                <p>{{this.nowTime}}</p>-->
                <form>
                    <at-input v-model="eventTitle" placeholder="events_title" style="width: 500px"></at-input>
                    <at-input v-model="extraSth" placeholder="extra_sth" style="width: 500px"></at-input>
                    <!--                    <at-input-number :min="1" :max="12"></at-input-number>-->
                    <div class="row at-row">
                        <div class="col-sm-12 col-md-4">
                            <at-input-number v-model="nowYear" style="width: 35px"></at-input-number>
                            <!--                            <span>{{nowYear}}</span>-->
                        </div>
                        <div class="col-sm-12 col-md-4">
                            <at-input-number v-model="nowMouth" style="width: 35px"></at-input-number>
                        </div>
                        <div class="col-sm-12 col-md-4">
                            <at-input-number v-model="nowDay" style="width: 35px"></at-input-number>
                        </div>
                    </div>
                    <div class="row at-row">
                        <div class="col-sm-12 col-md-4">
                            <at-input-number v-model="nowHour" style="width: 35px"></at-input-number>
                        </div>
                        <div class="col-sm-12 col-md-4">
                            <at-input-number v-model="nowMinute" style="width: 35px"></at-input-number>
                        </div>
                        <div class="col-sm-12 col-md-4">
                            <at-input-number v-model="nowSecond" style="width: 35px"></at-input-number>
                            <!--                            <span>{{nowSecond}}</span>-->
                        </div>
                    </div>

                    <at-button v-on:click="addAffair()" style="float: left"></at-button>
                </form>
            </div>
        </at-card>
    </div>
</template>

<script>
    import axios from 'axios';

    let temp = new Date();
    export default {
        name: "InputForm",
        data() {
            let eventTitle;
            let extraSth;
            return {
                // nowTime:{},
                eventTitle,
                extraSth,
                nowYear: {},
                nowMouth: {},
                nowDay: {},
                nowHour: {},
                nowMinute: {},
                nowSecond: {},
            }
        },
        methods: {
            getDate: function () {
                // this.nowTime.year =temp.getFullYear();
                // this.nowTime.mouth=temp.getMonth()+1;
                // this.nowTime.day =temp.getDay();
                // this.nowTime.hour = temp.getHours();
                // this.nowTime.minute = temp.getMinutes();
                // this.nowTime.second =temp.getSeconds();
                // console.log(this.nowTime);
                this.nowYear = temp.getFullYear();
                this.nowMouth = temp.getMonth() + 1;
                this.nowDay = temp.getDay();
                this.nowHour = temp.getHours();
                this.nowMinute = temp.getMinutes();
                this.nowSecond = temp.getSeconds();
            },
            addAffair: function () {
                axios({
                    method: 'post',
                    url: 'http://localhost:1221/opera/add',
                    data: {
                        events_title: this.eventTitle,
                        events_deadline: this.nowYear + "/" + this.nowMouth + "/" + this.nowDay + " " + this.nowHour + "/" + this.nowMinute + "/" + this.nowSecond,
                        extra_tips: this.extraSth
                    }
                });
            },
            setDate: function (year, month, day, hour, minute, second, title, sth) {
                this.nowYear = year;
                alert(year);
                alert(this.nowYear);
                this.nowMouth = month;
                this.nowDay = day;
                this.nowHour = hour;
                this.nowMinute = minute;
                this.nowSecond = second;
                this.eventTitle = title;
                this.extraSth = sth;
            }
        },
        mounted() {
            this.getDate()
        },
        created() {
            this.getDate()
        }
    }
</script>

<style scoped>

</style>