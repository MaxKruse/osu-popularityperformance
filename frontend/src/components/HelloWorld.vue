<script setup>
// Import needed elements from quasar
import { QLayout, QHeader, QToolbar, QAvatar, QTabs, QRouteTab, QPageContainer, QFooter, QToolbarTitle } from 'quasar'

const TITLE = "Osu! Performance v2.5";
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

    <q-page-container>
      <router-view />
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
