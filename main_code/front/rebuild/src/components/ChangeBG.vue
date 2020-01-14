<template>
  <div>
    <van-button type="info" v-on:click="sendDataToParent" style="margin: 10px" >更换背景</van-button>
  </div>
</template>

<script>
  export default {
    name: 'changeBG',
    data(){
        return{
            childmsg: 1
        }
    },
    methods:{
        sendDataToParent: function () {
            this.$emit('parentFunction', this.childmsg)
            location.reload();
      }
    }
  }
</script>