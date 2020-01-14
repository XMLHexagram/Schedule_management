<template>
  <div class="outter" >
<!--    <div id="nav">-->
<!--      <router-link to="/">Home</router-link> |-->
<!--      <router-link to="/about">About</router-link>-->
<!--    </div>-->
    <!-- <changeBG /> -->
    <router-view/>
  </div>
</template>

<style>
  .outter {
    margin-top: 0px;
    width: 100%;
    height: 100%;
    padding: 0px;

  }
</style>

<script>
//   import changeBG from './components/ChangeBG'
//   export default {
//     data() {
//       return {
//         backGround: 'url(https://s0.xinger.ink/acgimg/acgurl.php)',
//         ckey: 1,
//       }
//     },
//     components:{
//       changeBG,
//     },
//     methods: {
//         customParentFunction(msg){
//             console.log('子组件传过来的值',msg)
//             this.ckey += 1
//         }
//     }
//   }
</script>