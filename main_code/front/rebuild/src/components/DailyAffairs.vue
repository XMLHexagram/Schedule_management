<template>
    <div>
        <div v-for="affair in dailyAffairs" v-bind:key="affair.ID">
            <van-swipe-cell>
                <template slot="left">
                    <van-button square type="primary" text="选择"/>
                </template>

                <van-cell :border="false" :title="affair.Title" :value="affair.Extra"/>

            </van-swipe-cell>
        </div>
    </div>
</template>

<script>
    import axios from 'axios'
    // import {apiGet} from "../api";

    export default {
        name: 'HelloWorld',
        props: {
            // msg: String
        },
        data() {
            return {
                dailyAffairs: []
            }
        },
        created() {
            this.getDailyEvents()
        },
        methods: {
            getDailyEvents: function () {
                axios.get("http://www.sweetbeecr.com:1221/all/dailyEvents").then((res) => {
                    this.dailyAffairs = res.data.data
                }).catch(err => {
                    this.dailyAffairs = err
                    alert("我们遇到了未知错误，这有可能导致程序无法正常运行");
                })
            }
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
