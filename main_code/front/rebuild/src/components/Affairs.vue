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

        <!--        以下是以后要抽象成components的部分-->
        <van-row>
            <van-col offset="15">
                <van-button type="info" v-on:click="showAddAffair">
                    添加事务
                </van-button>
            </van-col>
        </van-row>

        <!--        以下是用于展示的弹出层中的内容-->
        <van-popup
                v-model="showDetail"
                position="top"
                :style="{ height: '55%' }"
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

        <!--以下是用于修改的弹出层-->
        <transition name="van-fade">
            <van-popup
                    v-show="showModify"
                    v-model="showModify"
                    position="top"
                    :style="{ height: '55%' }"
                    :duration="0"
            >
                <div>
                    <form>
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
                            <van-col span="6" offset="3">
                                <van-button plain type="info" v-on:click="showModify=false">
                                    返回
                                </van-button>
                            </van-col>
                            <van-col span="6" offset="7">
                                <van-button plain type="primary" v-on:click="modifyAffair(tempAffair.id)">
                                    提交
                                </van-button>
                            </van-col>
                        </van-row>
                    </form>

                </div>
            </van-popup>
        </transition>

<!--添加新事物的弹出层-->
            <van-popup
                    v-model="showAdd"
                    position="top"
                    :style="{ height: '55%' }"
            >
                <div>
                    <form>
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

                        <van-row>
                            <van-col span="6" offset="3">
                                <van-button plain type="info" v-on:click="showAdd=false">
                                    返回
                                </van-button>
                            </van-col>
                            <van-col span="6" offset="7">
                                <van-button plain type="primary" v-on:click="addAffair(tempAffair.id)">
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
</template>

<script>
    import axios from 'axios';
    import {baseURL} from "../main";
    import {Dialog} from "vant";
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
                showAdd:false,
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
                axios.get(baseURL+"/all/affairs").then((res) => {
                    this.Affairs = res.data.data
                }).catch(err => {
                    this.Affairs = err
                    alert("我们遇到了未知错误，这有可能导致程序无法正常运行");
                })
            },
            deleteAffair: function (id) {
                Dialog.confirm({
                    // title: '标题',
                    message: '确定完成了吗'
                }).then(() => {
                    let url = baseURL + '/opera?id=' + id;
                    axios.delete(url).then(() => {
                        // axios.delete('http://localhost:1221/opera/' + ID).then(() => {
                        this.getAffairs()
                    })
                }).catch(() => {
                    // on cancel
                });


            },
            modifyAffair: function (id) {
                this.showModify = false;
                axios({
                        method: 'put',
                        url: baseURL+'/opera?id=' + id,
                        data: {
                            "title": this.tempAffair.title,
                            "deadline": this.currentDate,
                            //这里如果传入this.tempAffair.deadline会返回json格式错误
                            //可能因为这个格式无法被解析成时间
                            "extra": this.tempAffair.extra
                        }
                    }
                ).then(() => {
                    this.getAffairs();
                })
            },
            addAffair: function () {
                this.showAdd = false;
                axios({
                    method: 'post',
                    url: baseURL+'/opera/add',
                    data: {
                        "title": this.tempAffair.title,
                        "extra": this.tempAffair.extra,
                        "deadline": this.currentDate,
                    }
                }).then(() => {
                    this.getAffairs();
                })
            },
            showAddAffair: function () {
                for (let key in this.tempAffair) {
                    delete this.tempAffair[key];
                }
                this.showAdd = true;
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
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
