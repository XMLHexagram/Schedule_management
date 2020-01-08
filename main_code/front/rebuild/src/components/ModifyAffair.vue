<template>
    <div>
        <van-popup
                v-model="showDetail"
                position="top"
                :style="{ height: '50%' }"
        >
            <div>
                <van-field
                        :value="tempAffair.title"
                        label="事务"
                        rows="2"
                        autosize
                        disabled
                        type="textarea"
                />
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
                        <van-button type="info">
                            修改
                        </van-button>
                    </van-col>
                </van-row>
            </div>
        </van-popup>
    </div>
</template>

<script>
    export default {
        name: 'ModifyAffairs',
        props: {
            // msg: String
            showDetail:Boolean,
            tempAffair:{},
        },
        data() {
        },
        created() {
        },
        methods: {
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
