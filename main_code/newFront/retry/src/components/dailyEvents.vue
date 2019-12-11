<template>
    <div>
<!--        <div v-for="event in dailyEvents" :key="event.Title">-->
            <van-cell-group>
                <van-cell title="单元格" value="内容"/>
                <van-cell title="单元格" value="内容" :label="dailyEvents" :border="'{{ dailyEvents }}'"/>
            </van-cell-group>
<!--        </div>-->
    </div>
</template>

<script>
    import {AxiosInstance as axios} from "axios";
    import Vue from 'vue';
    import {Cell, CellGroup} from 'vant';

    Vue.use(Cell).use(CellGroup);

    export default {
        name: "dailyEvents",
        data() {
            return {
                dailyEvents: []
            }
        },
        methods: {
            getDailyEvents: function () {
                axios.get('http://121.199.40.243:1221/dailyEvents').then(res => {
                    // axios.get('http://localhost:1221/dailyEvents').then(res => {
                    // alert(res)
                    // alert("1")
                    this.dailyEvents = res.data.data
                    // console.log(this.dailyEvents)
                }).catch(err => {
                    this.$Message.error(err)
                })
            },
        }
    }
</script>

<style scoped>

</style>
