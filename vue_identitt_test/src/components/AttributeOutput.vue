<template>
    <div class="results-wrapper">
        <h1>Attributes:</h1>
        <pre>{{result | pretty}}</pre>
    </div>
</template>
<script>
export default {
    props: ['result'],
    filters: {
        pretty(value) {
            return JSON.stringify(JSON.parse(value), null, 2);
        }
    }
}
</script>
<style lang="scss">
    .results-wrapper {
        margin-bottom: 200px;
        margin-left: 50px;
        display: flex;
        flex-flow: column nowrap;
        text-align: left;
        align-items: center;
    }
</style>
