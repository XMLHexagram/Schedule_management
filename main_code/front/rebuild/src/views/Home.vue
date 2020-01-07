<template>
    <div>
        <van-tabs v-model="active" animated swipeable>
            <van-tab title="每日任务">
                <div>
                    <DailyAffairs/>
                </div>
            </van-tab>
            <van-tab title="长期事务">内容 2</van-tab>
        </van-tabs>
        <!--    <HelloWorld msg="Welcome to Your Vue.js App"/>-->
    </div>
</template>

<script>
    // @ is an alias to /src
    import DailyAffairs from '@/components/DailyAffairs.vue'

    export default {
        name: 'home',
        components: {
            DailyAffairs
        },
        data() {
            return {
                active: 2
            };
        }
    }
</script>
