<script setup>
import { QLayout, QHeader, QToolbar, QAvatar, QPageContainer, QFooter, QToolbarTitle, useQuasar } from 'quasar'
import { useStore } from './store';
import { onMounted, ref, watch } from 'vue';
import { GetSelf } from "./backend";
import Login from './views/Login.vue';
import { useRoute, useRouter } from "vue-router";

const TITLE = "Osu! Performance v2.5";

const q = useQuasar();
const store = useStore();

const route = useRoute();

// set darkmode
q.dark.set(true);

const isLoggedIn = ref(true);

// try to get self from api
onMounted(async () => {

  if(route.path === "/code") {
    return;
  }

  try {
    store.user = await GetSelf();
  } catch (e) {
    isLoggedIn.value = false;
    console.error(e);
    q.notify({
            type: "negative",
            message: "Error getting user info. Try to login again.",
            timeout: 10000,
            progress: true,
            position: 'bottom'
        })
  }
})

</script>

<template>
  <q-layout view="hHh Lpr fFf">

    <q-header elevated class="bg-primary text-white" height-hint="98">
      <q-toolbar>
        <q-toolbar-title>
          <q-avatar>
            <img src="https://cdn.quasar.dev/logo-v2/svg/logo-mono-white.svg">
          </q-avatar>
          {{ TITLE }}
        </q-toolbar-title>
      </q-toolbar>
    </q-header>

    <q-page-container v-if="!isLoggedIn">
      <Login />
    </q-page-container>

    <q-page-container v-else class="fit row wrap justify-center items-center content-center">
      <router-view class="col-auto" style="overflow: auto; min-width: 60%; max-width: 60%;"/>
    </q-page-container>

    <q-footer elevated class="bg-grey-8 text-white">
      <q-toolbar>
        <q-toolbar-title>
          <div>{{ TITLE }}</div>
        </q-toolbar-title>
      </q-toolbar>
    </q-footer>

  </q-layout>
</template>

<style>
</style>
