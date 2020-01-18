<template>
    <div class="my_div">
        <div v-for="(affair,index) in Affairs" v-bind:key="affair.ID" class="my_div">
            <van-swipe-cell class="my_div">
                <van-cell
                        :border="false"
                        :title="affair.title"
                        :value="affair.extra"
                        v-on:click="controlShowDetail(index)"
                        class="my_div"
                />
                <template slot="right">
                    <van-button square type="primary" text="完成" v-on:click="tempt(affair.id)"/>
                </template>
            </van-swipe-cell>

            <!--            <van-popup v-model="showDetail" position="bottom" :style="{height:'20%'}">内容</van-popup>-->
        </div>

        <!--        以下是以后要抽象成components的部分-->
        <van-row style="display: flex; justify-content: center">
            <!--            <van-col span="6" offset="5">-->
            <!--                <van-button type="warning">-->
            <!--                    删除事务-->
            <!--                </van-button>-->
            <!--            </van-col>-->
            <!-- <van-col span="6" offset="15"> -->
            <van-button type="info" @click="() => showPopup = true" style="margin: 10px">添加任务</van-button>
            <!-- </van-col> -->
            <!-- <van-col span="6" offset="15"> -->
            <!-- </van-col> -->
        </van-row>

        <!--        以下是用于展示的弹出层中的内容-->
        <van-popup v-model="showDetail" position="top" :style="{ height: '55%' }">
            <div>
                <van-field :value="tempAffair.title" label="事务" rows="2" autosize disabled type="textarea"/>
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
                        <van-button type="info" v-on:click="showModify=true">修改</van-button>
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
                        <van-field v-model="tempAffair.title" label="事务" rows="2" autosize type="textarea"/>
                        <van-field v-model="tempAffair.extra" label="详细说明" rows="2" autosize type="textarea"/>
                        <van-field
                                v-model="tempAffair.deadline"
                                label="截止时间"
                                rows="2"
                                autosize
                                disabled
                                type="textarea"
                                v-on:click="showTimePacker=true"
                                v-on:change="changeToDeadline"
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
                                <van-button plain type="info" v-on:click="showModify=false">返回</van-button>
                            </van-col>
                            <van-col span="6" offset="7">
                                <van-button plain type="primary" v-on:click="modifyAffair(tempAffair.id)">提交
                                </van-button>
                            </van-col>
                        </van-row>
                    </form>
                </div>
            </van-popup>
        </transition>
        <!--添加新事物的弹出层-->
        <AddLTA v-model="showPopup"/>
        <!-- 时间选择 -->
        <van-popup v-model="showTimePacker" position="bottom" :style="{ height: '30%' }">
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
    import axios from "axios";
    import {baseURL} from "../main";
    import {Dialog} from "vant";
    import AddLTA from "./AddLongTermAffair";
    import {ApiInstance} from "../instances/index";

    export default {
        name: "Affairs",
        props: {},
        components: {
            AddLTA
        },
        data() {
            return {
                Affairs: [],
                showDetail: false,
                showModify: false,
                showPopup: false,
                tempAffair: [],
                showTimePacker: false,
                currentDate: new Date(),
                minDate: new Date(),
                maxDate: new Date(2023, 0, 1),
            };
        },
        created() {
            this.getAffairs();
        },
        methods: {
            async tempt(id) {
                this.deleteAffair(id)
                setTimeout(this.getAffairs,1000)
            },
            controlShowDetail: function (index) {
                this.showDetail = true;
                this.tempAffair = this.Affairs[index];
            },
            changeToDeadline: function () {
                let y = this.currentDate.getFullYear();
                let m = this.currentDate.getMonth() + 1;
                m = m < 10 ? "0" + m : m;
                let d = this.currentDate.getDate();
                d = d < 10 ? "0" + d : d;
                let h = this.currentDate.getHours();
                let minute = this.currentDate.getMinutes();
                minute = minute < 10 ? "0" + minute : minute;
                this.tempAffair.deadline =
                    y + "-" + m + "-" + d + "   " + h + ":" + minute;
            },
            async getAffairs() {
                try {
                    const token = localStorage.getItem("token");
                    const {data: res} = await ApiInstance.get(baseURL + "/all/affairs", {
                        headers: {
                            Authorization: token
                        }
                    });
                    this.Affairs = res.data;
                } catch (e) {
                }
            },
            async deleteAffair(id) {
                Dialog.confirm({
                    // title: '标题',
                    message: "确定完成了吗?"
                }).then(() => {
                    let url = baseURL + "/opera?id=" + id;
                    try {
                        const token = localStorage.getItem("token");
                        const {data: res} = ApiInstance.delete(url, {
                            headers: {
                                Authorization: token
                            }
                        });
                    } catch (e) {
                    }
                });
            },
            async modifyAffair(id) {
                this.showModify = false;
                this.tempAffair.deadline = this.currentDate;
                try {
                    const token = localStorage.getItem("token");
                    const {data: res} = ApiInstance.put(
                        baseURL + "/opera?id=" + id,
                        {
                            title: this.tempAffair.title,
                            deadline: this.tempAffair.deadline,
                            extra: this.tempAffair.extra
                        },
                        {
                            headers: {
                                Authorization: token
                            }
                        }
                    );
                    window.location.reload();
                } catch (e) {
                }
            }
        }
    };
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    .input-text {
        border: 1px;
        padding: 30%;
        background-color: #66ccff;
        border-radius: 5px;
    }

    .body {
        margin: 0px;
        padding: 0px;
        height: 100%;
        width: 100%;
    }

    .my_div {
        background: transparent;
    }
</style>
