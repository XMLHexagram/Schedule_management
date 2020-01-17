<template>
  <div>
    <van-popup v-model="showDetail" position="top" :style="{ height: '50%' }">
      <div>
        <van-field :value="tempAffair.title" label="事务" rows="2" autosize disabled type="textarea" />
        <van-field
          :value="tempAffair.extra"
          label="详细说明"
          rows="2"
          autosize
          disabled
          type="textarea"
        />
        <van-field
          :value="tempAffair.deadline"
          label="截止时间"
          rows="2"
          autosize
          disabled
          type="textarea"
        />
        <van-field
          :value="tempAffair.created_at"
          label="创建时间"
          rows="2"
          autosize
          disabled
          type="textarea"
        />

        <van-row>
          <van-col offset="16">
            <van-button type="info">修改</van-button>
          </van-col>
        </van-row>
      </div>
    </van-popup>
    <!-- 时间输入 -->
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
import { Dialog } from "vant";
export default {
  name: "ModifyAffairs",
  props: {
    // msg: String
    showDetail: Boolean,
    tempAffair: {}
  },
  data() {
    return {
      showModify: false,
      showTimePacker: false,
      showAdd: false,
      minDate: new Date(),
      maxDate: new Date(2023, 0, 1),
      currentDate: new Date()
    };
  },
  created() {},
  methods: {
    deleteAffair: function(id) {
      Dialog.confirm({
        // title: '标题',
        message: "确定完成了吗"
      })
        .then(() => {
          let url = baseURL + "/opera?id=" + id;
          axios.delete(url).then(() => {
            // axios.delete('http://localhost:1221/opera/' + ID).then(() => {
            location.reload()
          });
        })
        .catch(() => {
          // on cancel
        });
    },
    modifyAffair: function(id) {
      this.showModify = false;
      axios({
        method: "put",
        url: baseURL + "/opera?id=" + id,
        data: {
          title: this.tempAffair.title,
          deadline: this.currentDate,
          //这里如果传入this.tempAffair.deadline会返回json格式错误
          //可能因为这个格式无法被解析成时间
          extra: this.tempAffair.extra
        }
      }).then(() => {
        location.reload()
      });
    },
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
