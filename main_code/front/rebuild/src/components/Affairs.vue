<template>
    <div>
        <div v-for="(affair,index) in Affairs" v-bind:key="affair.ID">
            <van-swipe-cell>

                <van-cell :border="false" :title="affair.title" :value="affair.extra"
                          v-on:click="controlShowDetail(index)"/>

                <template slot="right">
                    <van-button square type="primary" text="完成" v-on:click="deleteAffair(affair.id)"/>
                </template>
            </van-swipe-cell>

            <!--            <van-popup v-model="showDetail" position="bottom" :style="{height:'20%'}">内容</van-popup>-->
        </div>
        <van-popup
                v-model="showDetail"
                position="top"
                :style="{ height: '50%' }"
        >
            <div>
                <van-field
                        :value="tempAffair.title"
                        label="事务"
                        rows="2"
                        autosize
                        disabled
                        type="textarea"
                />
                <van-field
                        :value="tempAffair.extra"
                        label="详细说明"
                        rows="2"
                        autosize
                        disabled
                        type="textarea"
                />
                <van-field
                        :value="tempAffair.deadline"
                        label="截止时间"
                        rows="2"
                        autosize
                        disabled
                        type="textarea"
                />
                <van-field
                        :value="tempAffair.created_at"
                        label="创建时间"
                        rows="2"
                        autosize
                        disabled
                        type="textarea"
                />

                <van-row>
                    <van-col offset="16">
                        <van-button type="info" v-on:click="showModify=true">
                            修改
                        </van-button>
                    </van-col>
                </van-row>


            </div>
        </van-popup>

        <transition name="van-fade">
            <van-popup
                    v-show="showModify"
                    v-model="showModify"
                    position="top"
                    :style="{ height: '50%' }"
                    :duration="0"
            >
                <div>
                    <form @submit="modifyAffair(tempAffair.id)">
                        <van-field
                                v-model="tempAffair.title"
                                label="事务"
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
                        <van-field
                                :value="tempAffair.created_at"
                                label="创建时间"
                                rows="2"
                                autosize
                                disabled
                                type="textarea"
                        />

                        <van-row>
                            <van-col offset="16">
                                <van-button plain type="info submit">
                                    提交
                                </van-button>
                            </van-col>
                        </van-row>
                    </form>

                </div>
            </van-popup>
        </transition>

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
                    v-on:change="changeDeadline"
            />
        </van-popup>

    </div>
</template>

<script>
    import axios from 'axios'
    // import {apiGet} from "../api";

    export default {
        name: 'Affairs',
        props: {
            // msg: String
        },
        data() {
            return {
                Affairs: [],
                showDetail: false,
                showModify: false,
                showTimePacker: false,
                tempAffair: {},
                minDate: new Date(),
                maxDate: new Date(2023, 0, 1),
                currentDate: new Date()
            }
        },
        created() {
            this.getAffairs()
        },
        methods: {
            controlShowDetail: function (index) {
                this.showDetail = true;
                this.tempAffair = this.Affairs[index]
            },
            getAffairs: function () {
                axios.get("http://www.sweetbeecr.com:1221/all/affairs").then((res) => {
                    this.Affairs = res.data.data
                }).catch(err => {
                    this.Affairs = err
                    alert("我们遇到了未知错误，这有可能导致程序无法正常运行");
                })
            },
            deleteAffair: function (id) {
                let url = 'http://www.sweetbeecr.com:1221/opera?id=' + id;
                axios.delete(url).then(() => {
                    // axios.delete('http://localhost:1221/opera/' + ID).then(() => {
                    this.getAffairs()
                })
            },
            modifyAffair:function(id){
                this.showModify = false;
                axios({
                        method: 'put',
                        url: 'http://www.sweetbeecr.com:1221/opera?id=' + id,
                        data: {
                            "title":this.tempAffair.title,
                            "deadline":this.currentDate,
                            //这里如果传入this.tempAffair.deadline会返回json格式错误
                            //可能因为这个格式无法被解析成时间
                            "extra":this.tempAffair.extra
                        }
                    }
                ).then(() => {
                    this.getAffairs();
                })
            },
            changeDeadline: function () {
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
            }
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
