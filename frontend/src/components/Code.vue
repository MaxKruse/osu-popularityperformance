<script setup>
import { inject, ref, toRef } from 'vue';
import { useRoute, useRouter } from "vue-router";

const router = useRouter();
const route = useRoute();

let session = route.query?.session || "";

console.log(session)
// get the session from the http params
// save session as cookie
inject('$cookies').set('session', session);

// redirect to the main page
router.push({ path: '/' });

</script>
