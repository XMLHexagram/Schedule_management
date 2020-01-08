<template>
    <div>
        <van-tabs v-model="active" swipeable>
            <van-tab title="每日任务">
                <slot>
                    <DailyAffairs/>
                </slot>
            </van-tab>
            <van-tab title="长期事务">
                <slot>
                    <Affairs/>
                </slot>
            </van-tab>
        </van-tabs>
        <!--    <HelloWorld msg="Welcome to Your Vue.js App"/>-->
    </div>
</template>

<script>
    // @ is an alias to /src
    import DailyAffairs from '@/components/DailyAffairs.vue';
    import Affairs from "@/components/Affairs";

    export default {
        name: 'home',
        components: {
            DailyAffairs,
            Affairs
        },
        data() {
            return {
                active: 2
            };
        }
    }
</script>
