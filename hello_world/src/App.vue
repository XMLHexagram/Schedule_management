<template>
    <div id="app" style="">
        <!-- <div id="app2"> -->
        <div class="row at-row flex-around" style="position:relative;top: -50px">
            <div class="col-md-12">
                <InputForm></InputForm>
            </div>


            <div class="col-md-12">
                <at-card style="width: 600px">
                    <h4>修改事务</h4>
                    <div>
                        <!--                <p>{{this.nowTime}}</p>-->
                        <form>
                            <at-input v-model="eventTitle" placeholder="events_title" style="width: 500px"></at-input>
                            <at-input v-model="extraSth" placeholder="extra_sth" style="width: 500px"></at-input>
                            <!--                    <at-input-number :min="1" :max="12"></at-input-number>-->
                            <at-input v-model="deadline" placeholder="deadline" style="width: 500px"></at-input>
                            <at-button v-on:click="pushModifyAffair()" style="float: left"></at-button>
                            <at-button v-on:click="getAllAffairs"></at-button>
                        </form>
                    </div>
                </at-card>
            </div>
        </div>

        <!-- </div> -->
        <div class="row at-row" style="position: relative;left: 39px;top:-20px">
            <div class="col-md-4">
                <div class="at-box-row bg-c-brand-dark common">ID</div>
            </div>
            <div class="col-md-4">
                <div class="at-box-row bg-c-brand-light common">内容</div>
            </div>
            <div class="col-md-4">
                <div class="at-box-row bg-c-brand-light common">创建时间</div>
            </div>
            <div class="col-md-4">
                <div class="at-box-row bg-c-brand-dark common">截止时间</div>
            </div>
            <div class="col-md-4">
                <div class="at-box-row bg-c-brand-light common">附加内容</div>
            </div>
        </div>

        <div v-for="affair in affairs" :key="affair.ID">
            <div class="row at-row flex-center">
                <div class="col-md-4">
                    <div class="at-box-row bg-c-brand-dark common">{{affair.ID}}</div>
                </div>
                <div class="col-md-4">
                    <div class="at-box-row bg-c-brand-light common">{{affair.events_title}}</div>
                </div>
                <div class="col-md-4">
                    <div class="at-box-row bg-c-brand-light common">{{affair.CreatedAt}}</div>
                </div>
                <div class="col-md-4">
                    <div class="at-box-row bg-c-brand-dark common">{{affair.events_deadline}}</div>
                </div>
                <div class="col-md-4">
                    <div class="at-box-row bg-c-brand-light common">{{affair.extre_sth}}</div>
                </div>
                <at-button v-on:click="deleteAffairs(affair.ID)" type="primary" style="float: left;">delete
                </at-button>
                <at-button v-on:click="modifyAffairs(affair.ID)" type="primary" style="float: left;">modify
                </at-button>
            </div>
        </div>
    </div>
</template>

<script>
    import axios from 'axios';
    import InputForm from "./components/InputForm";

    export default {
        name: 'app',
        components: {InputForm},
        data() {
            let eventTitle;
            let extraSth;
            let deadline;
            let ID_;
            return {
                affairs: [],
                eventTitle,
                extraSth,
                deadline,
                ID_,
            }
        },
        methods: {
            deleteAffairs: function (ID) {
                // alert(ID);
                axios.delete('http://localhost:1221/opera/' + ID).then(() => {
                    this.getAllAffairs()
                })
            },
            modifyAffairs: function (ID) {
                //回调
                // alert(ID);
                this.ID_ = ID;
                let i;
                i = 0;
                while (i < 100) {
                    // alert(this.affairs[i].ID);
                    if (this.affairs[i].ID == ID) {
                        this.deadline = this.affairs[i].events_deadline;
                        this.eventTitle = this.affairs[i].events_title;
                        this.extraSth = this.affairs[i].extre_sth;
                        break;
                    }
                    i++;
                }
                // this.getAllAffairs();
            },
            getAllAffairs: function () {
                axios.get('http://localhost:1221/allAffairs').then(res => {
                    this.affairs = res.data.data
                    // console.log(this.affairs)
                }).catch(err => {
                    this.$Message.error(err)
                })
            },
            pushModifyAffair: function () {
                axios({
                    method: 'put',
                    url: 'http://localhost:1221/opera/' + this.ID_,
                    data: {
                        events_title: this.eventTitle,
                        events_deadline: this.deadline,
                        extra_tips: this.extraSth
                    }
                }).then(() => {
                    this.getAllAffairs()
                })
            }
        },
        mounted() {
            this.getAllAffairs();

        },
    }
</script>

<style>
    #app {
        font-family: 'Avenir', Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: center;
        color: #2c3e50;
        margin-top: 60px;
    }

    div.common {
        border-radius: 5px;
        background-color: #66B3F3;
        height: 30px;
    }
</style>
