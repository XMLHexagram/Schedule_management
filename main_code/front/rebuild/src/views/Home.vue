<template>
  <div class="inner" :style="{ background: backGround }" :key="ckey" id="home">
    <van-tabs
      v-model="active"
      swipeable
      style="
                width: 100%;
                height: 100%;
                padding: 0px;"
    >
      <van-tab title="每日任务">
        <slot>
          <DailyAffairs />
        </slot>
      </van-tab>
      <van-tab
        title="长期事务"
        style="
                width: 100%;
                height: 100%;
                padding: 0px;"
      >
        <slot>
          <Affairs />
        </slot>
      </van-tab>
    </van-tabs>
    <changeBG />
  </div>
</template>

<script>
// @ is an alias to /src
import DailyAffairs from "@/components/DailyAffairs.vue";
import Affairs from "@/components/Affairs";
import changeBG from "@/components/ChangeBG";

export default {
  name: "home",
  props: ["parentFunction"],
  components: {
    DailyAffairs,
    Affairs,
    changeBG
  },
  data() {
    return {
      active: 2,
      backGround: "url(https://s0.xinger.ink/acgimg/acgurl.php)",
      //https://s0.xinger.ink/fj/fjurl.php
      //https://s0.xinger.ink/acgimg/acgurl.php
      ckey: 1
    };
  },
  methods: {
    // ParentFunction: function(msg) {
    //   console.log("fa");
    //   console.log("子组件传过来的值", msg);
    //   this.ckey += msg;
    // }
  }
};
</script>
<style scoped>
.inner {
  width: 100%;
  height: 100%;
  padding: 0px;
  display: flex;
  align-content: center;
  align-items: center;
  flex-direction: column;
}
</style>
