<template>
    <div id="app" style="">
        <!-- <div id="app2"> -->
        <div>
            <Daily></Daily>
        </div>
        <!-- </div> -->
        <div class="row at-row flex-center" style="position: relative;top: 20px">
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
            <at-button type="primary" style="float: left;">delete
            </at-button>
            <at-button type="primary" style="float: left;">modify
            </at-button>
        </div>

        <div v-for="affair in affairs" :key="affair.ID">
            <div class="row at-row flex-center">
                <div class="col-md-4">
                    <div class="at-box-row bg-c-brand-dark common">{{affair.ID}}</div>
                </div>
                <div class="col-md-4">
                    <div class="at-box-row bg-c-brand-light common">{{affair.Title}}</div>
                </div>
                <div class="col-md-4">
                    <div class="at-box-row bg-c-brand-light common">{{affair.CreatedAt}}</div>
                </div>
                <div class="col-md-4">
                    <div class="at-box-row bg-c-brand-dark common">{{affair.Deadline}}</div>
                </div>
                <div class="col-md-4">
                    <div class="at-box-row bg-c-brand-light common">{{affair.Extra}}</div>
                </div>
                <at-button v-on:click="deleteAffairs(affair.ID)" type="primary" style="float: left;">delete
                </at-button>
                <at-button v-on:click="modifyAffairs(affair.ID)" type="primary" style="float: left;">modify
                </at-button>
            </div>
        </div>
        <div class="row at-row flex-center" style="position:relative;top: 30px">
            <div class="col-md-12">
                <InputForm></InputForm>
            </div>


            <div class="col-md-12">
                <at-card style="width: 600px">
                    <h4>修改事务</h4>
                    <div>
                        <!--                <p>{{this.nowTime}}</p>-->
                        <form>
                            <at-input v-model="Title" placeholder="Title" style="width: 500px"></at-input>
                            <at-input v-model="Extra" placeholder="Extra" style="width: 500px"></at-input>
                            <!--                    <at-input-number :min="1" :max="12"></at-input-number>-->
                            <at-input v-model="Deadline" placeholder="Deadline" style="width: 500px"></at-input>
                            <at-button v-on:click="pushModifyAffair()" style="float: left"></at-button>
                            <at-button v-on:click="getAllAffairs"></at-button>
                        </form>
                    </div>
                </at-card>
            </div>
        </div>
    </div>
</template>

<script>
    import axios from 'axios';
    import InputForm from "./components/InputForm";
    import Daily from "./components/daily"

    export default {
        name: 'app',
        components: {Daily, InputForm},
        data() {
            let Title;
            let Extra;
            let Deadline;
            let ID_;
            return {
                affairs: [],
                Title,
                Extra,
                Deadline,
                ID_,
            }
        },
        methods: {
            deleteAffairs: function (ID) {
                // alert(ID);
                axios.delete('http://121.199.40.243:1221/opera/' + ID).then(() => {
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
                        this.Deadline = this.affairs[i].Deadline;
                        this.Title = this.affairs[i].Title;
                        this.Extra = this.affairs[i].Extra;
                        break;
                    }
                    i++;
                }
                // this.getAllAffairs();
            },
            getAllAffairs: function () {
                axios.get('http://121.199.40.243:1221/allAffairs').then(res => {
                    this.affairs = res.data.data
                    // console.log(this.affairs)
                }).catch(err => {
                    this.$Message.error(err)
                })
            },
            pushModifyAffair: function () {
                axios({
                    method: 'put',
                    url: 'http://121.199.40.243:1221/opera/' + this.ID_,
                    data: {
                        Title: this.Title,
                        Deadline: this.Deadline,
                        Extra: this.Extra
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
