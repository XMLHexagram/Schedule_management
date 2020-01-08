<template>
    <div>
        <div v-for="affair in Affairs" v-bind:key="affair.ID">
            <van-swipe-cell>

                <van-cell :border="false" :title="affair.title" :value="affair.extra" v-on:click="showDetail=true"/>

                <template slot="right">
                    <van-button square type="primary" text="选择"/>
                </template>
            </van-swipe-cell>

<!--            <van-popup v-model="showDetail" position="bottom" :style="{height:'20%'}">内容</van-popup>-->
            <van-overlay :show="showDetail" @click="showDetail = false" />
        </div>
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
            }
        },
        created() {
            this.getDailyEvents()
        },
        methods: {
            getDailyEvents: function () {
                axios.get("http://www.sweetbeecr.com:1221/all/affairs").then((res) => {
                    this.Affairs = res.data.data
                }).catch(err => {
                    this.Affairs = err
                    alert("我们遇到了未知错误，这有可能导致程序无法正常运行");
                })
            },
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
