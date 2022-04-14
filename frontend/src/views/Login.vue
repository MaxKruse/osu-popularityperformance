<script setup>
import { QBtn } from "quasar"
</script>

<template>

<q-btn label="Login with osu!" href="https://localhost:5000/api/oauth/start" padding="lg" color="pink-4 is-rounded" class="absolute-center" size="xl" @click="redirectToLogin" />

</template>
