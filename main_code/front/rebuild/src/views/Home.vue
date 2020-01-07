<template>
    <div class="home">
        <van-tabs v-model="active" swipeable animated>
            <van-tab title="每日事务" name="1">
                <DailyAffairs/>
            </van-tab>
            <van-tab title="长期任务" name="2">
                <Affairs/>
            </van-tab>
            <!--      <van-tab title="标签 3" name="c">内容 3</van-tab>-->
        </van-tabs>
    </div>
</template>

<script>
    import DailyAffairs from "../components/DailyAffairs";
    import Affairs from "../components/Affairs";

    export default {
        components: {
            DailyAffairs,
            Affairs,
        },
        data() {
            return {
                active: '1'
            };
        }
    }
</script>
