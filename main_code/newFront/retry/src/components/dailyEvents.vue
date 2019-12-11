<template>
    <div>
        <!--        <div v-for="event in dailyEvents" :key="event.Title">-->
        <cell v-on:click="show = true" title="显示每日事务" is-link arrow-direction="down" style="background-color: aquamarine"></cell>
        <Overlay :show="show" v-on:click="show = false">
            <van-row>
                <div v-for="event in dailyEvents" :key="event.Title">
                    <cell-group>
                        <cell :title="event.Title" :value="event.Extra"/>
                    </cell-group>
                </div>
            </van-row>
        </Overlay>
        <!--        </div>-->
    </div>
</template>

<script>
    import axios from "axios";
    // import Vue from 'vue';
    import {Cell, CellGroup, Overlay} from 'vant';

    // Vue.use(Cell).use(CellGroup);

    export default {
        components: {Cell, CellGroup, Overlay},
        name: "dailyEvents",
        data() {
            return {
                show: false,
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
                    // eslint-disable-next-line no-console
                    // console.log(this.dailyEvents)
                }).catch(err => {
                    this.$Message.error(err)
                })
            },
        },
        created() {
            this.getDailyEvents();
        }
    }
</script>

<style scoped>

</style>
