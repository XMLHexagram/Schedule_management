<template>
    <div class="row at-row" style="float: left">
        <!--        <input v-model="nowYear">-->
        <!--        <p>{{nowYear}}</p>-->
        <at-card style="width: 600px">
            <h4>添加新事务</h4>
            <div>
                <!--                <p>{{this.nowTime}}</p>-->
                <form>
                    <at-input v-model="Title" placeholder="Title" style="width: 500px"></at-input>
                    <at-input v-model="Extra" placeholder="Extra" style="width: 500px"></at-input>
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

                        <at-button v-on:click="addAffair()">提交事务</at-button>
                    </div>
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
            let Title;
            let Extra;
            return {
                // nowTime:{},
                Title,
                Extra,
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
                if (this.nowMouth < 10) {
                    this.nowMouth = "0" + this.nowMouth;
                }
                if (this.nowDay < 10) {
                    this.nowDay = "0" + this.nowDay;
                }
                if (this.nowHour < 10) {
                    this.nowHour = "0" + this.nowHour;
                }
                if (this.nowMinute < 10) {
                    this.nowMinute = "0" + this.nowMinute;
                }
                if (this.nowSecond < 10) {
                    this.nowSecond = "0" + this.nowSecond;
                }

                axios({
                    method: 'post',
                    url: 'http://121.199.40.243:1221/opera/add',
                    // url: 'http://localhost:1221/opera/add',
                    data: {
                        Title: this.Title,
                        Deadline: this.nowYear + "-" + this.nowMouth + "-" + this.nowDay + "T" + this.nowHour + ":" + this.nowMinute + ":" + this.nowSecond + "Z",
                        Extra: this.Extra
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