<template>
  <div class="my_div">
    <div v-for="(affair,index) in Affairs" v-bind:key="affair.ID" class="my_div">
      <van-swipe-cell class="my_div">
        <van-cell
          :border="false"
          :title="affair.title"
          :value="affair.extra"
          v-on:click="controlShowDetail(index)"
          class="my_div"
        />

        <template slot="right">
          <van-button square type="primary" text="完成" v-on:click="deleteAffair(affair.id)" />
        </template>
      </van-swipe-cell>

      <!--            <van-popup v-model="showDetail" position="bottom" :style="{height:'20%'}">内容</van-popup>-->
    </div>

    <!--        以下是以后要抽象成components的部分-->
    <van-row style="display: flex; justify-content: center">
      <!--            <van-col span="6" offset="5">-->
      <!--                <van-button type="warning">-->
      <!--                    删除事务-->
      <!--                </van-button>-->
      <!--            </van-col>-->
      <!-- <van-col span="6" offset="15"> -->
      <van-button type="info" @click="() => showPopup = true" style="margin: 10px">添加任务</van-button>
      <!-- </van-col> -->
      <!-- <van-col span="6" offset="15"> -->
      <!-- </van-col> -->
    </van-row>

    <!--        以下是用于展示的弹出层中的内容-->
    <van-popup v-model="showDetail" position="top" :style="{ height: '55%' }">
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
            <van-button type="info" v-on:click="showModify=true">修改</van-button>
          </van-col>
        </van-row>
      </div>
    </van-popup>

    <!--以下是用于修改的弹出层-->
    <transition name="van-fade">
      <van-popup
        v-show="showModify"
        v-model="showModify"
        position="top"
        :style="{ height: '55%' }"
        :duration="0"
      >
        <div>
          <form>
            <van-field v-model="tempAffair.title" label="事务" rows="2" autosize type="textarea" />
            <van-field v-model="tempAffair.extra" label="详细说明" rows="2" autosize type="textarea" />
            <van-field
              v-model="tempAffair.deadline"
              label="截止时间"
              rows="2"
              autosize
              disabled
              type="textarea"
              v-on:click="showTimePacker=true"
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
              <van-col span="6" offset="3">
                <van-button plain type="info" v-on:click="showModify=false">返回</van-button>
              </van-col>
              <van-col span="6" offset="7">
                <van-button plain type="primary" v-on:click="modifyAffair(tempAffair.id)">提交</van-button>
              </van-col>
            </van-row>
          </form>
        </div>
      </van-popup>
    </transition>

    <!--添加新事物的弹出层-->
    <AddLTA v-model="showPopup" />
  </div>
</template>

<script>
import axios from "axios";
import { baseURL } from "../main";
import { Dialog } from "vant";
import AddLTA from "./AddLongTermAffair";
import { ApiInstance } from "../instances/index";
// import changeBG from "./ChangeBG";
// import {apiGet} from "../api";

export default {
  name: "Affairs",
  props: {
    // msg: String
  },
  components: {
    AddLTA
    // changeBG
  },
  data() {
    return {
      Affairs: [],
      showDetail: false,
      showModify: false,
      showPopup: false,
      showAdd: false,
      tempAffair: []
    };
  },
  created() {
    this.getAffairs();
  },
  methods: {
    controlShowDetail: function(index) {
      this.showDetail = true;
      this.tempAffair = this.Affairs[index];
    },
    async getAffairs() {
      //   const token = localStorage.getItem("token")
      //   axios
      //     .get(baseURL + "/all/affairs",
      //         {
      //             headers: {
      //                 Authorization: token
      //             }
      //         }
      //     )
      //     .then(res => {
      //       this.Affairs = res.data.data;
      //     })
      //     .catch(err => {
      //       this.Affairs = err;
      //       alert("我们遇到了未知错误，这有可能导致程序无法正常运行");
      //     });
      try {
        const token = localStorage.getItem("token");
        const { data: res } = await ApiInstance.get(baseURL + "/all/affairs", {
          headers: {
            Authorization: token
          }
        });
      } catch (e) {}
    },

    deleteAffair: function(id) {
      Dialog.confirm({
        // title: '标题',
        message: "确定完成了吗"
      })
        .then(() => {
          let url = baseURL + "/opera?id=" + id;
          axios.delete(url).then(() => {
            // axios.delete('http://localhost:1221/opera/' + ID).then(() => {
            this.getAffairs();
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
        this.getAffairs();
      });
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
