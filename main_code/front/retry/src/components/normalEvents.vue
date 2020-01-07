<template>
    <div>
        <div v-for="affair in affairs" :key="affair.ID">
            <!--            <SwipeCell :title="affair.Title" :desc="affair.Extra" status="" @click="showPopup">-->
            <swipe-cell>
                <template slot="left">
                    <Button square type="primary" text="日期" v-on:click="showPopup(affair.created_at,affair.deadline)"/>
                </template>

                <cell :border="false" :title="affair.title" :value="affair.extra"/>

                <template slot="right">
                    <Button square type="danger" text="删除" v-on:click="deleteAffairs(affair.ID)"/>
                    <Button square type="primary" text="修改" v-on:click="modifyAffairs(affair.ID)"/>
                </template>

            </swipe-cell>
        </div>


        <popup v-model="showEventTime" position="top">
            <CellGroup>
                <cell>创建于:{{this.tempCreatTime}}</cell>
                <Cell>Deadline:{{this.tempDeadline}}</Cell>
            </CellGroup>
        </popup>

        <Overlay :show="showModify">
            <div class="wrapper">
                <div class="block">
                    <div>
                        <CellGroup>
                            <van-field
                                    v-model="Title"
                                    required
                                    clearable
                                    label="事务"
                                    placeholder="Title"
                            />

                            <van-field
                                    v-model="Extra"
                                    label="附加"
                                    placeholder="Extra"
                            />
                            <van-datetime-picker v-model="tempTime"
                                                 :max-date="maxDate"
                                                 v-on:cancel="cancelModifyEvent"
                                                 v-on:confirm="pushModifyAffair"/>
                        </CellGroup>
                    </div>
                </div>
            </div>
        </Overlay>
        <Button plain type="info" v-on:click="showAddEvent=true">
            添加事务
        </Button>
        <Overlay :show="showAddEvent">
            <div class="wrapper">
                <div class="block">
                    <div>
                        <CellGroup>
                            <van-field
                                    v-model="Title"
                                    required
                                    clearable
                                    label="事务"
                                    placeholder="Title"
                            />

                            <van-field
                                    v-model="Extra"
                                    label="附加"
                                    placeholder="Extra"
                            />
                            <van-datetime-picker
                                    v-model="tempTime"
                                    :min-date="minDate"
                                    :max-date="maxDate"
                                    v-on:cancel="showAddEvent=false"
                                    v-on:confirm="addEvents"/>
                        </CellGroup>
                    </div>
                    <!--                    <Button plain type="info" v-on:click="show1=false">-->
                    <!--                        退出-->
                    <!--                    </Button>-->
                    <!--                    <Button plain type="primary" v-on:click="addEvents">-->
                    <!--                        提交-->
                    <!--                    </Button>-->
                </div>
            </div>
        </Overlay>
    </div>
</template>

<script>
    import axios from "axios";
    import {SwipeCell, CellGroup, Cell, Popup, Button, Overlay} from "vant";

    export default {
        components: {SwipeCell, CellGroup, Cell, Popup, Button, Overlay},
        name: "normalEvents",
        data() {
            let Title;
            let Extra;
            let tempID;
            let tempCreatTime;
            let tempDeadline;
            return {
                showEventTime: false,
                showAddEvent: false,
                showModify: false,
                affairs: [],
                tempCreatTime,
                tempDeadline,
                Title,
                Extra,
                tempTime: new Date(),
                minDate: new Date(),
                maxDate: new Date(2020, 11, 31),
                tempID,
            }
        },
        methods: {
            showPopup: function (creatTime, deadline) {
                this.showEventTime = true;
                this.tempCreatTime = creatTime;
                this.tempDeadline = deadline;

            },
            deleteAffairs: function (ID) {
                // alert(ID);
                axios.delete('http://121.199.40.243:1221/opera/' + ID).then(() => {
                    // axios.delete('http://localhost:1221/opera/' + ID).then(() => {
                    this.getAllAffairs()
                })
                // }
            },
            modifyAffairs: function (ID) {
                //回调
                // alert(ID);
                this.tempID = ID;
                let i;
                i = 0;
                while (i < 1000) {
                    // alert(this.affairs[i].ID);
                    if (this.affairs[i].ID == ID) {
                        let tempToChange = this.affairs[i].Deadline;
                        let tempData = new Date(tempToChange);
                        this.tempTime = tempData;
                        // alert(this.tempTime);
                        this.Title = this.affairs[i].Title;
                        this.Extra = this.affairs[i].Extra;
                        break;
                    }
                    i++;
                }
                this.showModify = true;
                // this.pushModifyAffair();
                // this.getAllAffairs();
            },
            getAllAffairs: function () {
                axios.get('http://121.199.40.243:1221/all/affairs').then(res => {
                    // axios.get('http://localhost:1221/allAffairs').then(res => {

                    // eslint-disable-next-line no-console
                    console.log(res)
                    this.affairs = res.data.data
                    // eslint-disable-next-line no-console
                    console.log(this.affairs)
                    // console.log(this.affairs)
                })
                // .catch(err => {
                // this.$Message.error(err)
                // })
            },
            pushModifyAffair: function () {
                this.showModify = false;
                axios({
                    method: 'put',
                    url: 'http://121.199.40.243:1221/opera/' + this.tempID,
                    // url: 'http://localhost:1221/opera/' + this.ID_,
                    data: {
                        Title: this.Title,
                        Deadline: this.tempTime,
                        Extra: this.Extra
                    }
                }).then(() => {
                    this.getAllAffairs();
                    this.Title="";
                    this.Extra="";
                })
            },
            addEvents: function () {
                axios({
                    method: 'post',
                    url: 'http://121.199.40.243:1221/opera/add',
                    data: {
                        Title: this.Title,
                        Deadline: this.tempTime,
                        Extra: this.Extra,
                    }
                }).then(() => {
                    this.getAllAffairs();
                    this.showAddEvent = false;
                })
            },
            cancelModifyEvent: function () {
                this.showModify=false;
                this.Title="";
                this.Extra="";
            }
        },
        mounted() {
            this.getAllAffairs();
        },
    }
</script>

<style scoped>
    .wrapper {
        display: flex;
        align-items: center;
        justify-content: center;
        height: 100%;
    }


    .block {
        width: 500px;
        height: 500px;
        background-color: #fff;
    }

</style>
