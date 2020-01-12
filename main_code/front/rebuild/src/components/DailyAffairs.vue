<template>
    <div :style="{ background: backGround }">
        <div style="padding-left: 10% , padding-right: 10%">
            <div v-for="affair in dailyAffairs" v-bind:key="affair.ID">

                <van-swipe-cell>
                    <template slot="left">
                        <van-button square type="primary" text="删除" v-on:click="showCheck(affair.ID)"/>
                    </template>

                    <van-cell :border="false" :title="affair.Title" :value="affair.Extra"/>

                </van-swipe-cell>
            </div>

            <van-row  style="display: flex; justify-content: center">
                <!--            <van-col span="6" offset="5">-->
                <!--                <van-button type="warning">-->
                <!--                    删除事务-->
                <!--                </van-button>-->
                <!--            </van-col>-->
                <!-- <van-col span="6" offset="15"> -->
                    <van-button type="info" v-on:click="showAdd" style="margin: 10px">
                        添加任务
                    </van-button>
                <!-- </van-col> -->
                <!-- <van-col span="6" offset="15"> -->
                    <van-button type="info" v-on:click="changeBackGround" style="margin: 10px" >
                        更换背景
                    </van-button>
                <!-- </van-col> -->
            </van-row>

    <!--        用于添加新事务的弹出层-->
            <van-popup
                    v-model="showAddPopup"
                    position="top"
                    style="border-radius: 0 0 15px 15px"
                    :style="{ height: '35%' }"
            >
                <div>
                    <form>
                        <van-field
                                v-model="tempAffair.title"
                                label="每日任务"
                                rows="2"
                                autosize
                                type="textarea"
                                
                        />
                        <van-field
                                v-model="tempAffair.extra"
                                label="详细说明"
                                rows="2"
                                autosize
                                type="textarea"
                        />
                        <van-field
                                v-model="tempAffair.deadline"
                                label="截止时间"
                                rows="2"
                                autosize
                                disabled
                                type="textarea"
                                v-on:click="showTimePacker=true"
                        />

                        <van-row>
                            <van-col span="6" offset="3">
                                <van-button plain type="info" v-on:click="showAddPopup=false">
                                    返回
                                </van-button>
                            </van-col>
                            <van-col span="6" offset="7">
                                <van-button plain type="primary" v-on:click="addDailyAffair(tempAffair.id)">
                                    提交
                                </van-button>
                            </van-col>
                        </van-row>
                    </form>

                </div>
            </van-popup>

            <!--        时间输入控件-->
            <van-popup
                    v-model="showTimePacker"
                    position="bottom"
                    :style="{ height: '30%' }"
            >
                <van-datetime-picker
                        :show-toolbar="false"
                        v-model="currentDate"
                        type="datetime"
                        :min-date="minDate"
                        :max-date="maxDate"
                        v-on:change="changeToDeadline"
                />
            </van-popup>
        </div>
    </div>


</template>

<script>
    import axios from 'axios'
    import {Dialog} from "vant";
    import {baseURL} from "../main";
    // import {apiGet} from "../api";

    export default {
        name: 'DailyAffairs',
        props: {
            // msg: String
        },
        data() {
            return {
                dailyAffairs: [],
                showAddPopup: false,
                tempAffair: {},
                minDate: new Date(),
                maxDate: new Date(2023, 0, 1),
                currentDate: new Date(),
                showTimePacker:false,
                backGround: null
            }
        },
        created() {
            this.getDailyAffairs()
        },
        methods: {
            getDailyAffairs: function () {
                axios.get(baseURL + "/all/dailyAffairs").then((res) => {
                    this.dailyAffairs = res.data.data
                }).catch(err => {
                    this.dailyAffairs = err
                    alert("我们遇到了未知错误，这有可能导致程序无法正常运行");
                })
            },
            showAdd: function () {
                this.showAddPopup = true;
            },
            showCheck: function (id) {
                Dialog.confirm({
                    // title: '标题',
                    message: '确定要删除吗(删除后只能通过数据库恢复)'
                }).then(() => {
                    this.deleteDailyAffair(id)
                }).catch(() => {
                    // on cancel
                });
            },
            deleteDailyAffair: function (id) {
                axios.delete(baseURL + "/operaDaily?id=" + id).then(() => {
                    this.getDailyAffairs()
                }).catch(err => {
                    this.dailyAffairs = err
                    alert("我们遇到了未知错误，这有可能导致程序无法正常运行");
                })
            },
            changeToDeadline: function () {
                let y = this.currentDate.getFullYear();
                let m = this.currentDate.getMonth() + 1;
                m = m < 10 ? ('0' + m) : m;
                let d = this.currentDate.getDate();
                d = d < 10 ? ('0' + d) : d;
                let h = this.currentDate.getHours();
                let minute = this.currentDate.getMinutes();
                minute = minute < 10 ? ('0' + minute) : minute;
                let second = this.currentDate.getSeconds();
                second = minute < 10 ? ('0' + second) : second;

                this.tempAffair.deadline = y + '-' + m + '-' + d + ' ' + h + ':' + minute + ':' + second;
            },
            addDailyAffair: function () {
                try {
                    this.showAddPopup = false;
                axios({
                    method: 'post',
                    url: baseURL+'/operaDaily/add',
                    data: {
                        "title": this.tempAffair.title,
                        "extra": this.tempAffair.extra,
                        "deadline": this.currentDate,
                    }
                }).then(() => {
                    this.getDailyAffairs();
                })
                } catch (error) {
                    alert("添加失败")
                }
                
            },
            changeBackGround: function() {
                axios.get("https://s0.xinger.ink/acgimg/acgurl.php").then((res) => {
                    this.backGround = res
                    console.log(111)
                }).catch(err => {
                    this.dailyAffairs = err
                    alert("背景加载失败");
                })
                // this.backGround = this.backGround + "https://s0.xinger.ink/acgimg/acgurl.php"
            }
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    .input-text {
        border: 1px;
        padding: 30%;
        background-color: #66ccff;
        border-radius: 5px;
    }
</style>
