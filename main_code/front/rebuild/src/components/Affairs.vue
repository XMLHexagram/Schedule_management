<template>
    <div>
        <div v-for="affair in Affairs" v-bind:key="affair.ID">
            <van-swipe-cell>

                <van-cell :border="false" :title="affair.title" :value="affair.extra" v-on:click="showDetail=true"/>

                <template slot="right">
                    <van-button square type="primary" text="完成" v-on:click="deleteAffair(affair.id)"/>
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
            this.getAffairs()
        },
        methods: {
            getAffairs: function () {
                axios.get("http://www.sweetbeecr.com:1221/all/affairs").then((res) => {
                    this.Affairs = res.data.data
                }).catch(err => {
                    this.Affairs = err
                    alert("我们遇到了未知错误，这有可能导致程序无法正常运行");
                })
            },
            deleteAffair:function (ID) {
                let url='http://www.sweetbeecr.com:1221/opera/?id='+ID;
                axios.delete(url).then(() => {
                    // axios.delete('http://localhost:1221/opera/' + ID).then(() => {
                    this.getAffairs()
                })
            },
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
