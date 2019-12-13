<template>
    <div>
        <daily-events></daily-events>
        <normal-events></normal-events>
    </div>

</template>

<script>
    import DailyEvents from "./components/dailyEvents";
    import normalEvents from "./components/normalEvents";

    export default {
        name: 'app',
        components: {
            normalEvents,
            DailyEvents,
        },
    }
</script>

<style>
    /*#app {*/
    /*    font-family: 'Avenir', Helvetica, Arial, sans-serif;*/
    /*    -webkit-font-smoothing: antialiased;*/
    /*    -moz-osx-font-smoothing: grayscale;*/
    /*    text-align: center;*/
    /*    color: #2c3e50;*/
    /*    margin-top: 60px;*/
    /*}*/
</style>
