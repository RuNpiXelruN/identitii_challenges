<template>
    <transition name="fade" appear>
    <div class="ping-pong-wrapper">
        <router-link tag="v-btn" class="blue accent-3 back-btn v-btn--outline v-btn--depressed blue--text text--accent-3" to="/">Home</router-link>
        <h1 class="ping-pong-heading">Identitii - Ping Pong</h1>
        <v-btn color="blue accent-1 white--text" @click="updatePingPong">Ping Pong</v-btn>
        <h2 class="pingpong">{{ pingPong | capitalize }}</h2>
    </div>
    </transition>
</template>
<script>
export default {
    computed: {
        pingPong() {
            return this.$store.getters.getPingPong
        }
    },
    methods: {
        updatePingPong() {
            this.$store.dispatch('setPingPong')
        }
    },
    filters: {
        capitalize(value) {
            return value.charAt(0).toUpperCase() + value.slice(1)
        }
    }
}
</script>
<style lang="scss">
    .ping-pong-wrapper {
        min-height: 100%;
        display: flex;
        flex-flow: column nowrap;
        align-items: center;
        justify-content: center;
        width: 100%;

        button {
            margin-bottom: 30px;
        }

        .ping-pong-heading {
            margin-bottom: 30px;
            border-bottom: 1px solid;
        }

        .pingpong {
            font-size: 42px;
        }
    }

    // animation / transition
    .fade-enter {
        opacity: 0
    }

    .fade-enter-active {
        transition: opacity 1s ease;
    }
    
    // .fade-leave {}

    .fade-leave-active {
        opacity: 0;
    }
</style>


