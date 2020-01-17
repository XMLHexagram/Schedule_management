<template>
  <!--        时间输入控件-->
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
</template>
<script>
export default {
    data() {
        return {
            currentDate: new Date(),
            minDate: new Date(),
            maxDate: new Date(2023, 0, 1),
        }
    },
}
</script>