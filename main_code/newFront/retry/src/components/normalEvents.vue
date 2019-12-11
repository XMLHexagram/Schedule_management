<template>
    <div>
        <div v-for="affair in affairs" :key="affair.ID">
            <!--            <SwipeCell :title="affair.Title" :desc="affair.Extra" status="" @click="showPopup">-->
            <swipe-cell>
                <template slot="left">
                    <Button square type="primary" text="日期" v-on:click="showPopup"/>
                </template>

                <cell :border="false" :title="affair.Title" :value="affair.Extra"/>

                <template slot="right">
                    <Button square type="danger" text="删除"/>
                    <Button square type="primary" text="修改"/>
                </template>

            </swipe-cell>
            <popup v-model="show" position="top" :style="{height:'20%'}">
                <div>创建于:{{affair.CreatedAt}}</div>
                <div>Deadline:{{affair.Deadline}}</div>
            </popup>
        </div>
    </div>
</template>

<script>
    import axios from "axios";
    import {SwipeCell, Cell, Popup, Button} from "vant";

    export default {
        components: {SwipeCell, Cell, Popup, Button},
        name: "normalEvents",
        data() {
            let Title;
            let Extra;
            let Deadline;
            let ID_;
            return {
                show: false,
                affairs: [],
                Title,
                Extra,
                Deadline,
                ID_,
            }
        },
        methods: {
            showPopup() {
                this.show = true;
            },
            deleteAffairs: function (ID) {
                // alert(ID);
                // // axios.delete('http://121.199.40.243:1221/opera/' + ID).then(() => {
                axios.delete('http://localhost:1221/opera/' + ID).then(() => {
                    this.getAllAffairs()
                })
                // }
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
                    // axios.get('http://localhost:1221/allAffairs').then(res => {
                    this.affairs = res.data.data
                    // console.log(this.affairs)
                }).catch(err => {
                    this.$Message.error(err)
                })
            },
            pushModifyAffair: function () {
                let str, temp;
                temp = this.Deadline + "";
                str = temp.split(" ");
                temp = str[0] + "T" + str[1] + "Z";

                axios({
                    method: 'put',
                    url: 'http://121.199.40.243:1221/opera/' + this.ID_,
                    // url: 'http://localhost:1221/opera/' + this.ID_,
                    data: {
                        Title: this.Title,
                        Deadline: temp,
                        Extra: this.Extra
                    }
                }).then(() => {
                    this.getAllAffairs()
                })
            },
        },
        mounted() {
            this.getAllAffairs();
        },
    }
</script>

<style scoped>

</style>
