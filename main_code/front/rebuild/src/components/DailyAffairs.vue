<template>
  <div class="my_div">
    <div v-for="affair in dailyAffairs" v-bind:key="affair.ID" class="my_div">
      <van-swipe-cell class="my_div">
        <template slot="left">
          <van-button square type="primary" text="删除" v-on:click="showCheck(affair.ID)" />
        </template>
        <van-cell :border="false" :title="affair.Title" :value="affair.Extra" class="my_div" />
      </van-swipe-cell>
    </div>

    <van-row style="display: flex; justify-content: center">
      <van-button type="info" @click="() => showPopup = true" style="margin: 10px">添加任务</van-button>
    </van-row>
    <AddAffair v-model="showPopup" />
  </div>
</template>

<script>
import axios from "axios";
import { Dialog } from "vant";
import { baseURL } from "../main";
import AddAffair from "./AddAffair";
import {ApiInstance} from "../instances/index"

export default {
  name: "DailyAffairs",
  props: {
    // msg: String
  },
  components: {
    AddAffair
  },
  data() {
    return {
      dailyAffairs: [],
      showPopup: false
    };
  },
  created() {
    console.log(2)
    this.getDailyAffairs();
  },
  methods: {
    async getDailyAffairs() {
      try {
        const token = localStorage.getItem("token")
        const { data: res } = await ApiInstance.get(
          baseURL + "/all/dailyAffairs",
          {
            headers: {
              Authorization: token,
            }
          }
        );
        this.dailyAffairs = res.data
      } catch (e) {}
    },
    showCheck: function(id) {
      Dialog.confirm({
        // title: '标题',
        message: "确定要删除吗(删除后只能通过数据库恢复)"
      })
        .then(() => {
          this.deleteDailyAffair(id);
        })
        .catch(() => {
          // on cancel
        });
    },
    async deleteDailyAffair(id) {
      try {
        const token = localStorage.getItem("token")
        const { data: res } = await ApiInstance.delete(
          baseURL + "/operaDaily?id=" + id,
          {
            headers: {
              Authorization: token,
            }
          }
        );
        this.getDailyAffairs();
      } catch (e) {}
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
