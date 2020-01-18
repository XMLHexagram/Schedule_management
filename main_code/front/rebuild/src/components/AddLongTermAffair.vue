<template>
  <div>
    <van-popup
      :value="show"
      @input="handleInput"
      position="top"
      style="border-radius: 0 0 15px 15px; background-color: rgba(0.7,0.7,0.7,0.3)"
      :style="{ height: '45%' }"
    >
      <div style="background-color: rgba(0,0,0,0)">
        <form>
          <van-field v-model="tempAffair.title" label="每日任务" rows="2" autosize type="textarea" />
          <van-field v-model="tempAffair.extra" label="详细说明" rows="2" autosize type="textarea" />
          <van-field
            v-model="tempAffair.deadline"
            label="截止时间"
            rows="2"
            autosize
            disabled
            type="textarea"
            v-on:click="() => showTimePacker=true"
          />

          <van-row>
            <van-col span="6" offset="3">
              <van-button plain type="info" v-on:click="() => show = false">返回</van-button>
            </van-col>
            <van-col span="6" offset="7">
              <van-button plain type="primary" v-on:click="addDailyAffair(tempAffair.id)" v-bind:disabled="dis">提交</van-button>
            </van-col>
          </van-row>
        </form>
      </div>
    </van-popup>
    <!-- 时间选择 -->
    <van-popup v-model="showTimePacker" position="bottom" :style="{ height: '30%' }">
      <van-datetime-picker
        :show-toolbar="false"
        v-model="currentDate"
        type="datetime"
        :min-date="minDate"
        :max-date="maxDate"
        v-on:change="changeToDeadline"
      />
    </van-popup>
  </div>
</template>

<script>
import axios from "axios";
import { baseURL } from "../main";
import { ApiInstance } from "../instances/index";
export default {
  name: "AddAffair",
  props: ["value"],
  data() {
    return {
      tempAffair: {},
      show: this.value,
      showTimePacker: false,
      currentDate: new Date(),
      minDate: new Date(),
      maxDate: new Date(2023, 0, 1),
      dis: false,
    };
  },
  watch: {
    value(v) {
      this.show = v;
    }
  },
  methods: {
    async addDailyAffair() {
      console.log(1);
      try {
        this.showAdd = false;
        this.dis = true;
        const token = localStorage.getItem("token");
        const { data: res } = await ApiInstance.post(
          baseURL + "/opera/add",
          {
            title: this.tempAffair.title,
            extra: this.tempAffair.extra,
            deadline: this.currentDate
          },
          {
            headers: {
              Authorization: token
            }
          }
        );
        this.dis = false;
        window.location.reload();
      } catch (e) {}
    },
    handleInput(value) {
      this.$emit("input", value);
    },
    changeToDeadline: function() {
      let y = this.currentDate.getFullYear();
      let m = this.currentDate.getMonth() + 1;
      m = m < 10 ? "0" + m : m;
      let d = this.currentDate.getDate();
      d = d < 10 ? "0" + d : d;
      let h = this.currentDate.getHours();
      let minute = this.currentDate.getMinutes();
      minute = minute < 10 ? "0" + minute : minute;
      this.tempAffair.deadline =
        y + "-" + m + "-" + d + "   " + h + ":" + minute;
    }
  }
};
</script>