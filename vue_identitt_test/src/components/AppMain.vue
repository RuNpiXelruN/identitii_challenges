<template>
    <transition name="fade" appear>
    <div class="identitii-wrapper">
        <div class="inner-wrapper">
            <h1>Identitii Frontend Challenges</h1>
            <h3>by Justin :)</h3>
            <div class="challenges-wrapper">
                <router-link tag="v-btn" class="info" :to="{name: 'AttributeForm'}">Attribute Selection Form</router-link>
                <router-link tag="v-btn" class="info" :to="{name: 'PingPong'}">Ping Pong</router-link>                
            </div>
        </div>
    </div>    
    </transition>
</template>
<script>
import FormComponent from '@/components/FormComponent.vue'
export default {
    components: {
        FormComponent,
    }
}
</script>
<style lang="scss">
    .identitii-wrapper {
        min-height: 100%;
        display: flex;
        flex-flow: row nowrap;
        align-items: center;
        justify-content: center;
        width: 100%;

        .inner-wrapper {
            margin-bottom: 60px;
        }

        .challenges-wrapper {
            margin-top: 60px;
        }
    }

    // animation / transition
    .fade-enter {
        opacity: 0
    }

    .fade-enter-active {
        transition: opacity 0.5s ease;
    }
    
    // .fade-leave {}

    .fade-leave-active {
        opacity: 0;
    }
</style>


